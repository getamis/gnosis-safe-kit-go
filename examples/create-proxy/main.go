package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

	relayerPrivateKeyHex := os.Getenv("RELAYER_PRIVATE_KEY")
	if len(relayerPrivateKeyHex) == 0 {
		log.Fatalln("Please specify RELAYER_PRIVATE_KEY in the environment variables")
	}

	relayerPrivateKey, err := crypto.ToECDSA(common.FromHex(relayerPrivateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load relayer's private key: %v", err)
	}

	relayer := bind.NewKeyedTransactor(relayerPrivateKey)

	proxyFactory, err := contracts.NewGnosisSafeProxyFactory(common.HexToAddress(os.Getenv("PROXY_FACTORY_ADDRESS")), conn)
	if err != nil {
		log.Fatalf("Failed to create proxy factory: %v", err)
	}

	proxyFactorySession := &contracts.GnosisSafeProxyFactorySession{
		Contract: proxyFactory,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     relayer.From,
			Signer:   relayer.Signer,
			GasLimit: big.NewInt(1000000).Uint64(),
		},
	}
	_ = proxyFactorySession

	masterCopy := os.Getenv("MASTER_COPY")
	if len(masterCopy) == 0 {
		log.Fatalln("Please specify MASTER_COPY in the environment variables")
	}

	masterCopyAddress := common.HexToAddress(masterCopy)
	_ = masterCopyAddress

	gnosisSafeABIpacker := func(owners []common.Address, threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) ([]byte, error) {
		gnosisSafeABI, err := abi.JSON(strings.NewReader(contracts.GnosisSafeABI))
		if err != nil {
			return nil, err
		}
		return gnosisSafeABI.Pack("setup", owners, threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
	}

	alicePrivateKeyHex := os.Getenv("ALICE_PRIVATE_KEY")
	if len(alicePrivateKeyHex) == 0 {
		log.Fatalln("Please specify ALICE_PRIVATE_KEY in the environment variables")
	}

	alicePrivateKey, err := crypto.ToECDSA(common.FromHex(alicePrivateKeyHex))
	if err != nil {
		log.Fatalf("Failed to load Alice's private key: %v", err)
	}

	aliceAddress := crypto.PubkeyToAddress(alicePrivateKey.PublicKey)

	initializer, err := gnosisSafeABIpacker(
		[]common.Address{aliceAddress}, // owners
		big.NewInt(1),                  // threshold, we have only one owner
		common.Address{},               // to common.Address, keep it empty
		[]byte{},                       // data []byte
		common.Address{},               // fallbackHandler common.Address
		common.Address{},               // paymentToken common.Address
		big.NewInt(0),                  // payment *big.Int
		common.Address{},               // paymentReceiver common.Address
	)
	if err != nil {
		log.Fatalf("Failed to pack initializer ABI: %v", err)
	}

	nonce := big.NewInt(1)
	tx, err := proxyFactorySession.CreateProxyWithNonce(masterCopyAddress, initializer, nonce)
	if err != nil {
		log.Fatalf("Failed to create proxy with nonce %v: %v", nonce, err)
	}

	fmt.Println("Tx Hash:", tx.Hash().Hex())
	fmt.Println("Cost:", tx.Cost())

	eventCh := make(chan *contracts.GnosisSafeProxyFactoryProxyCreation, 1)
	subscription, err := proxyFactory.WatchProxyCreation(&bind.WatchOpts{}, eventCh)
	if err != nil {
		log.Fatalf("Failed to watch ProxyCreation: %v", err)
	}

	proxyCreationEvent := <-eventCh
	subscription.Unsubscribe()

	fmt.Println("Proxy Address:", proxyCreationEvent.Proxy.Hex())
}
