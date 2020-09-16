package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	_ "github.com/joho/godotenv/autoload"

	"github.com/getamis/gnosis-safe-kit-go"
)

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=Token --lang=go --pkg=main --abi $PWD/examples/safe-transaction/IERC20.abi --out $PWD/examples/safe-transaction/ierc20.go

func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(os.Getenv("ETHEREUM_RPC_ENDPOINT"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	aliceProxyAddress := common.HexToAddress(os.Getenv("ALICE_PROXY_ADDRESS"))
	safeContract, err := gnosis.NewSafe(aliceProxyAddress, conn)
	if err != nil {
		log.Fatalf("Failed to initialize the safe account: %v", err)
	}

	tokenAddress := common.HexToAddress(os.Getenv("ERC20_TOKEN_ADDRESS"))

	tokenTransferAbiPacker := func(recipient common.Address, amount *big.Int) ([]byte, error) {
		packer, err := abi.JSON(strings.NewReader(TokenABI))
		if err != nil {
			return nil, err
		}
		return packer.Pack("transfer", recipient, amount)
	}

	recipient := common.HexToAddress(os.Getenv("BOB_ADDRESS"))
	contractCall, err := tokenTransferAbiPacker(recipient, big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to pack the transfer tx: %v", err)
	}

	alicePrivateKeyHex := os.Getenv("ALICE_PRIVATE_KEY")
	if len(alicePrivateKeyHex) == 0 {
		log.Fatalln("Please specify ALICE_PRIVATE_KEY in the environment variables")
	}

	alicePrivateKey, err := crypto.ToECDSA(common.FromHex(alicePrivateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load the private key: %v", err)
	}

	alice := bind.NewKeyedTransactor(alicePrivateKey)
	nonce, err := conn.PendingNonceAt(context.Background(), alice.From)
	if err != nil {
		log.Fatalf("Failed to get pending nonce: %v", err)
	}

	gasLimit := uint64(2000000)
	gasPrice := big.NewInt(200000000000)

	// Create the transaction, sign it and schedule it for execution
	aliceTx := types.NewTransaction(nonce, tokenAddress, new(big.Int), gasLimit, gasPrice, contractCall)

	signedAliceTx, err := alice.Signer(types.HomesteadSigner{}, alice.From, aliceTx)
	if err != nil {
		log.Fatalf("Failed to sign the transaction: %v", err)
	}
	rawSignedAliceTx, err := rlp.EncodeToBytes(signedAliceTx)
	if err != nil {
		log.Fatalf("Failed to encode the signed transaction: %v", err)
	}

	relayerPrivateKeyHex := os.Getenv("RELAYER_PRIVATE_KEY")
	if len(relayerPrivateKeyHex) == 0 {
		log.Fatalln("Please specify RELAYER_PRIVATE_KEY in the environment variables")
	}

	relayerPrivateKey, err := crypto.ToECDSA(common.FromHex(relayerPrivateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load the relayer's private key: %v", err)
	}

	relayer := bind.NewKeyedTransactor(relayerPrivateKey)

	proxyNonce, err := safeContract.Nonce(&bind.CallOpts{
		Pending: true,
	})
	if err != nil {
		log.Fatalf("Failed to get the proxy's nonce: %v", err)
	}

	safeTx := gnosis.SafeTransaction{
		To:             aliceProxyAddress,   // to
		Value:          new(big.Int),        // value
		Data:           rawSignedAliceTx,    // data, the contract call user wants
		Operation:      gnosis.DelegateCall, // operation
		SafeTxGas:      new(big.Int),        // safeTxGas
		BaseGas:        new(big.Int),        // baseGas
		GasPrice:       new(big.Int),        // gasPrice
		GasToken:       common.Address{},    // gasToken
		RefundReceiver: common.Address{},    // refundReceiver
	}

	rawSafeTx, err := safeContract.EncodeTransactionData(
		&bind.CallOpts{
			Pending: true,
		},
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		proxyNonce,
	)
	if err != nil {
		log.Fatalf("Failed to encode the safe transaction: %v", err)
	}

	safeTxHashBytes := crypto.Keccak256(rawSafeTx)

	// alice signs the safe transaction
	safeTxSignature, err := crypto.Sign(safeTxHashBytes, alicePrivateKey)
	if err != nil {
		log.Fatalf("Failed to sign the safe transaction: %v", err)
	}

	// Relay the signed transaction
	relayTx, err := safeContract.ExecTransaction(
		&bind.TransactOpts{
			From:   relayer.From,
			Signer: relayer.Signer,
		},
		aliceProxyAddress,   // to
		new(big.Int),        // value
		rawSignedAliceTx,    // data
		gnosis.DelegateCall, // operation
		new(big.Int),        // safeTxGas
		new(big.Int),        // baseGas
		new(big.Int),        // gasPrice
		common.Address{},    // gasToken
		common.Address{},    // refundReceiver
		safeTxSignature,     // signature
	)
	if err != nil {
		log.Fatalf("Failed to relay the transaction: %v", err)
	}

	fmt.Println(relayTx.Hash())
}
