package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"

	"github.com/getamis/gnosis-safe-kit-go/contracts"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(os.Getenv("ETHEREUM_RPC_ENDPOINT"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKeyHex := os.Getenv("RELAYER_PRIVATE_KEY")
	if len(privateKeyHex) == 0 {
		log.Fatalln("Please specify PRIVATE_KEY in the environment variables")
	}

	privateKey, err := crypto.ToECDSA(common.FromHex(privateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load the private key: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = 7000000

	address, tx, _, err := contracts.DeployGnosisSafe(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy GnosisSafe Proxy Factory: %v", err)
	}

	fmt.Println("Address:", address.Hex())
	fmt.Println("Tx Hash:", tx.Hash().Hex())
}
