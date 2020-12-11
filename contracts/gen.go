package contracts

//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/GnosisSafe.sol
//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/proxies/GnosisSafeProxyFactory.sol
//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/proxies/GnosisSafeProxy.sol
//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/AMISTrust.sol
//go:generate docker run -v $PWD/contracts:/contracts -v $PWD/build/contracts:/build ethereum/solc:0.5.17-alpine -o /build --overwrite --abi --bin /contracts/TestContract.sol

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=GnosisSafe --lang=go --pkg=contracts --abi $PWD/build/contracts/GnosisSafe.abi --bin $PWD/build/contracts/GnosisSafe.bin --out $PWD/contracts/gnosis_safe.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=GnosisSafeProxyFactory --lang=go --pkg=contracts --abi $PWD/build/contracts/GnosisSafeProxyFactory.abi --bin $PWD/build/contracts/GnosisSafeProxyFactory.bin --out $PWD/contracts/gnosis_safe_proxy_factory.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=GnosisSafeProxy --lang=go --pkg=contracts --abi $PWD/build/contracts/GnosisSafeProxy.abi --bin $PWD/build/contracts/GnosisSafeProxy.bin --out $PWD/contracts/gnosis_safe_proxy.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=AMISTrust --lang=go --pkg=contracts --abi $PWD/build/contracts/AMISTrust.abi --bin $PWD/build/contracts/AMISTrust.bin --out $PWD/contracts/amis_trust.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --type=TestContract --lang=go --pkg=contracts --abi $PWD/build/contracts/TestContract.abi --bin $PWD/build/contracts/TestContract.bin --out $PWD/contracts/test_contract.go
