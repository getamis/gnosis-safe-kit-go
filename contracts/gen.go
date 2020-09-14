package contracts

//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/GnosisSafe.sol
//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/proxies/GnosisSafeProxyFactory.sol

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=GnosisSafe --lang=go --pkg=contracts --abi $PWD/build/contracts/GnosisSafe.abi --bin $PWD/build/contracts/GnosisSafe.bin --out $PWD/contracts/gnosis_safe.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=GnosisSafeProxyFactory --lang=go --pkg=contracts --abi $PWD/build/contracts/GnosisSafeProxyFactory.abi --bin $PWD/build/contracts/GnosisSafeProxyFactory.bin --out $PWD/contracts/gnosis_safe_proxy_factory.go
