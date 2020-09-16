package gnosis

import (
	"github.com/getamis/gnosis-safe-kit-go/contracts"
)

// NewProxyFactory creates a new instance of ProxyFactory, bound to a specific deployed contract.
var NewProxyFactory = contracts.NewGnosisSafeProxyFactory

// ProxyFactory is an auto generated Go binding around an Ethereum contract.
type ProxyFactory = contracts.GnosisSafeProxyFactory
