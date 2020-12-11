// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contracts

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
)

var (
	testContractABI, _ = abi.JSON(strings.NewReader(TestContractABI))
)

var _ = Describe("AMIS trust contract test", func() {
	var (
		ctx = context.Background()

		backend                     *backends.SimulatedBackend
		deployer, owner             *ecdsa.PrivateKey
		deployerAuth, relayerAuth   *bind.TransactOpts
		proxyContract               *GnosisSafe
		testContractAddr, proxyAddr common.Address
		amisTrust                   *AMISTrust
	)

	BeforeEach(func() {
		var err error
		ctx = context.Background()
		// init deployer/owner
		deployer, err = crypto.GenerateKey()
		Expect(err).Should(BeNil())
		owner, err = crypto.GenerateKey()
		Expect(err).Should(BeNil())
		ownerAddr := crypto.PubkeyToAddress(owner.PublicKey)
		relayer, err := crypto.GenerateKey()
		Expect(err).Should(BeNil())
		backend = backends.NewSimulatedBackend(
			core.GenesisAlloc{
				crypto.PubkeyToAddress(deployer.PublicKey): {Balance: big.NewInt(1000000000000000000)},
				crypto.PubkeyToAddress(relayer.PublicKey):  {Balance: big.NewInt(1000000000000000000)},
			},
			10000000,
		)

		deployerAuth = bind.NewKeyedTransactor(deployer)
		relayerAuth = bind.NewKeyedTransactor(relayer)

		// Deploy gnosis master/gnosis proxy/amis trsut/test contract
		safeAddr, _, _, err := DeployGnosisSafe(deployerAuth, backend)
		Expect(err).Should(BeNil())
		proxyAddr, _, _, err = DeployGnosisSafeProxy(deployerAuth, backend, safeAddr)
		Expect(err).Should(BeNil())
		_, _, amisTrust, err = DeployAMISTrust(deployerAuth, backend)
		Expect(err).Should(BeNil())
		testContractAddr, _, _, err = DeployTestContract(deployerAuth, backend)
		Expect(err).Should(BeNil())
		backend.Commit()

		// Setup owner
		proxyContract, err = NewGnosisSafe(proxyAddr, backend)
		Expect(err).Should(BeNil())
		owners := []common.Address{ownerAddr}
		threshold := big.NewInt(1)
		data := []byte{}
		fallbackHandler := common.Address{}
		paymentToken := common.Address{}
		payment := big.NewInt(0)
		paymentReceiver := common.Address{}
		tx, err := proxyContract.Setup(deployerAuth, owners, threshold, common.Address{}, data, fallbackHandler, paymentToken, payment, paymentReceiver)
		Expect(err).Should(BeNil())
		backend.Commit()
		receipt, err := backend.TransactionReceipt(ctx, tx.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	It("callRevert", func() {
		value := big.NewInt(1)
		relayerAuth.Value = value
		callRevertTxData, err := testContractABI.Pack("callRevert")
		Expect(err).Should(BeNil())
		callOpts := &bind.CallOpts{}
		relayCallRevertTxData, err := proxyContract.EncodeTransactionData(callOpts, testContractAddr, value, callRevertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, big.NewInt(0))
		Expect(err).Should(BeNil())
		sig, err := signHash(relayCallRevertTxData, owner)
		Expect(err).Should(BeNil())

		// Call amis trust contract, the eth will not keep in proxy contract
		relayCallRevertTx, err := amisTrust.ExecTransaction(relayerAuth, proxyAddr, testContractAddr, value, callRevertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
		Expect(err).ShouldNot(BeNil())
		Expect(relayCallRevertTx).Should(BeNil())
		backend.Commit()
		got, err := backend.BalanceAt(ctx, proxyAddr, nil)
		Expect(err).Should(BeNil())
		Expect(got).Should(Equal(big.NewInt(0)))

		//  Call gnosis safe contract, the eth will keep in proxy contract
		relaySafeCallRevertTx, err := proxyContract.ExecTransaction(relayerAuth, testContractAddr, value, callRevertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
		Expect(err).Should(BeNil())
		backend.Commit()
		relaySafeCallRevertTxReceipt, err := backend.TransactionReceipt(ctx, relaySafeCallRevertTx.Hash())
		Expect(err).Should(BeNil())
		Expect(relaySafeCallRevertTxReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
		got, err = backend.BalanceAt(ctx, proxyAddr, nil)
		Expect(err).Should(BeNil())
		Expect(got).Should(Equal(big.NewInt(1)))
	})

	Context("callRequire", func() {
		It("true", func() {
			value := big.NewInt(1)
			relayerAuth.Value = value
			callRequireTxData, err := testContractABI.Pack("callRequire", true)
			Expect(err).Should(BeNil())
			callOpts := &bind.CallOpts{}
			relayCallRequireTxData, err := proxyContract.EncodeTransactionData(callOpts, testContractAddr, value, callRequireTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, big.NewInt(0))
			Expect(err).Should(BeNil())
			sig, err := signHash(relayCallRequireTxData, owner)
			Expect(err).Should(BeNil())

			// Call amis trust contract, the eth will transfer to test contract
			relayCallRequireTx, err := amisTrust.ExecTransaction(relayerAuth, proxyAddr, testContractAddr, value, callRequireTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
			Expect(err).Should(BeNil())
			backend.Commit()
			relayCallRequireTxReceipt, err := backend.TransactionReceipt(ctx, relayCallRequireTx.Hash())
			Expect(err).Should(BeNil())
			Expect(relayCallRequireTxReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
			got, err := backend.BalanceAt(ctx, testContractAddr, nil)
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(big.NewInt(1)))
		})

		It("false", func() {
			value := big.NewInt(1)
			relayerAuth.Value = value
			callRequireTxData, err := testContractABI.Pack("callRequire", false)
			Expect(err).Should(BeNil())
			callOpts := &bind.CallOpts{}
			relayCallRequireTxData, err := proxyContract.EncodeTransactionData(callOpts, testContractAddr, value, callRequireTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, big.NewInt(0))
			Expect(err).Should(BeNil())
			sig, err := signHash(relayCallRequireTxData, owner)
			Expect(err).Should(BeNil())

			// Call amis trust contract, the eth will not keep in proxy contract
			relayCallRequireTx, err := amisTrust.ExecTransaction(relayerAuth, proxyAddr, testContractAddr, value, callRequireTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
			Expect(err).ShouldNot(BeNil())
			Expect(relayCallRequireTx).Should(BeNil())
			backend.Commit()
			got, err := backend.BalanceAt(ctx, proxyAddr, nil)
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(big.NewInt(0)))

			//  Call gnosis safe contract, the eth will keep in proxy contract
			relaySafeCallRequireTx, err := proxyContract.ExecTransaction(relayerAuth, testContractAddr, value, callRequireTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
			Expect(err).Should(BeNil())
			backend.Commit()
			relaySafeCallRequireTxReceipt, err := backend.TransactionReceipt(ctx, relaySafeCallRequireTx.Hash())
			Expect(err).Should(BeNil())
			Expect(relaySafeCallRequireTxReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
			got, err = backend.BalanceAt(ctx, proxyAddr, nil)
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(big.NewInt(1)))
		})
	})

	Context("callAssert", func() {
		It("true", func() {
			value := big.NewInt(1)
			relayerAuth.Value = value
			callAssertTxData, err := testContractABI.Pack("callAssert", true)
			Expect(err).Should(BeNil())
			callOpts := &bind.CallOpts{}
			relayCallAssertTxData, err := proxyContract.EncodeTransactionData(callOpts, testContractAddr, value, callAssertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, big.NewInt(0))
			Expect(err).Should(BeNil())
			sig, err := signHash(relayCallAssertTxData, owner)
			Expect(err).Should(BeNil())

			// Call amis trust contract, the eth will transfer to test contract
			relayCallAssertTx, err := amisTrust.ExecTransaction(relayerAuth, proxyAddr, testContractAddr, value, callAssertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
			Expect(err).Should(BeNil())
			backend.Commit()
			relayCallAssertTxReceipt, err := backend.TransactionReceipt(ctx, relayCallAssertTx.Hash())
			Expect(err).Should(BeNil())
			Expect(relayCallAssertTxReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
			got, err := backend.BalanceAt(ctx, testContractAddr, nil)
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(big.NewInt(1)))
		})

		It("false", func() {
			value := big.NewInt(1)
			relayerAuth.Value = value
			callAssertTxData, err := testContractABI.Pack("callAssert", false)
			Expect(err).Should(BeNil())
			callOpts := &bind.CallOpts{}
			relayCallAssertTxData, err := proxyContract.EncodeTransactionData(callOpts, testContractAddr, value, callAssertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, big.NewInt(0))
			Expect(err).Should(BeNil())
			sig, err := signHash(relayCallAssertTxData, owner)
			Expect(err).Should(BeNil())

			// Call amis trust contract, the eth will not keep in proxy contract
			relayCallAssertTx, err := amisTrust.ExecTransaction(relayerAuth, proxyAddr, testContractAddr, value, callAssertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
			Expect(err).ShouldNot(BeNil())
			Expect(relayCallAssertTx).Should(BeNil())
			backend.Commit()
			got, err := backend.BalanceAt(ctx, proxyAddr, nil)
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(big.NewInt(0)))

			//  Call gnosis safe contract, the eth will keep in proxy contract
			relaySafeCallAssertTx, err := proxyContract.ExecTransaction(relayerAuth, testContractAddr, value, callAssertTxData, 0, common.Big0, common.Big0, common.Big0, common.Address{}, common.Address{}, sig)
			Expect(err).Should(BeNil())
			backend.Commit()
			relaySafeCallAssertTxReceipt, err := backend.TransactionReceipt(ctx, relaySafeCallAssertTx.Hash())
			Expect(err).Should(BeNil())
			Expect(relaySafeCallAssertTxReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
			got, err = backend.BalanceAt(ctx, proxyAddr, nil)
			Expect(err).Should(BeNil())
			Expect(got).Should(Equal(big.NewInt(1)))
		})
	})
})

func TestContracts(t *testing.T) {
	if _, noColor := os.LookupEnv("GINKGO_NO_COLOR"); noColor {
		config.DefaultReporterConfig.NoColor = true
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contracts package suite")
}

func signHash(rawSafeTx []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	safeTxHashBytes := crypto.Keccak256(rawSafeTx)
	sig, err := crypto.Sign(safeTxHashBytes, privateKey)
	if err != nil {
		return nil, err
	}

	if sig[64] < 2 {
		sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	return sig, nil
}
