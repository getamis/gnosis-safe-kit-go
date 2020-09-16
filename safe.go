package gnosis

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/getamis/gnosis-safe-kit-go/contracts"
)

// NewSafe creates a new instance of Safe, bound to a specific deployed contract.
var NewSafe = contracts.NewGnosisSafe

// Safe is an auto generated Go binding around an Ethereum contract.
type Safe = contracts.GnosisSafe

type SafeTransaction struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signature      []byte
}
