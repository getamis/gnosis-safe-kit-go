package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"

	"github.com/getamis/gnosis-safe-kit-go"
)

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=Token --lang=go --pkg=main --abi $PWD/examples/safe-transaction/IERC20.abi --out $PWD/examples/safe-transaction/ierc20.go

func signHash(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}

	if sig[64] < 2 {
		sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	return sig, nil
}

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

	recipient := common.HexToAddress(os.Getenv("BOB_ADDRESS"))

	alicePrivateKeyHex := os.Getenv("ALICE_PRIVATE_KEY")
	if len(alicePrivateKeyHex) == 0 {
		log.Fatalln("Please specify ALICE_PRIVATE_KEY in the environment variables")
	}

	alicePrivateKey, err := crypto.ToECDSA(common.FromHex(alicePrivateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load the private key: %v", err)
	}

	alice := bind.NewKeyedTransactor(alicePrivateKey)

	fmt.Println("Alice address:", alice.From.Hex())

	// If the gasPrice is 0 we assume that nearly all available gas can be used (it is always more than safeTxGas)
	gasPrice := big.NewInt(0)
	safeTxGas := big.NewInt(0)

	relayerPrivateKeyHex := os.Getenv("RELAYER_PRIVATE_KEY")
	if len(relayerPrivateKeyHex) == 0 {
		log.Fatalln("Please specify RELAYER_PRIVATE_KEY in the environment variables")
	}

	relayerPrivateKey, err := crypto.ToECDSA(common.FromHex(relayerPrivateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load the relayer's private key: %v", err)
	}

	relayer := bind.NewKeyedTransactor(relayerPrivateKey)

	fmt.Println("relayer address:", relayer.From.Hex())

	proxyNonce, err := safeContract.Nonce(&bind.CallOpts{
		Pending: true,
	})
	if err != nil {
		log.Fatalf("Failed to get the proxy's nonce: %v", err)
	}

	safeTx := gnosis.SafeTransaction{
		To:             recipient,                         // to
		Value:          new(big.Int).SetInt64(1000000000), // value
		Data:           nil,                               // data, the contract call user wants
		Operation:      gnosis.Call,                       // operation
		SafeTxGas:      safeTxGas,                         // safeTxGas
		BaseGas:        new(big.Int),                      // baseGas
		GasPrice:       gasPrice,                          // gasPrice
		GasToken:       common.Address{},                  // gasToken
		RefundReceiver: common.Address{},                  // refundReceiver
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
	safeTxSignature, err := signHash(safeTxHashBytes, alicePrivateKey)
	if err != nil {
		log.Fatalf("Failed to sign the safe transaction: %v", err)
	}

	// Relay the signed transaction
	relayTx, err := safeContract.ExecTransaction(
		&bind.TransactOpts{
			From:     relayer.From,
			Signer:   relayer.Signer,
			GasLimit: 200000,
		},
		safeTx.To,             // to
		safeTx.Value,          // value
		safeTx.Data,           // data
		gnosis.Call,           // operation
		safeTx.SafeTxGas,      // safeTxGas
		safeTx.BaseGas,        // baseGas
		safeTx.GasPrice,       // gasPrice
		safeTx.GasToken,       // gasToken
		safeTx.RefundReceiver, // refundReceiver
		safeTxSignature,       // signature
	)
	if err != nil {
		log.Fatalf("Failed to relay the transaction: %v", err)
	}

	fmt.Println(relayTx.Hash().Hex())
}
