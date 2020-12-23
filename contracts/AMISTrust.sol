pragma solidity >=0.5.0 <0.7.0;
import "./GnosisSafe.sol";

/// @title AMISTrust calls a Gnosis Safe contract and check return values.
contract AMISTrust {
    /// @dev Call a Gnosis Safe contract and revert if the return vaule of execTransaction is false.
    function execTransaction(
        address payable contractAddr,
        address to,
        uint256 value,
        bytes calldata data,
        Enum.Operation operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address payable refundReceiver,
        bytes calldata signatures
    )
        external
        payable
        returns (bool success)
    {
        require(contractAddr!=address(0), "Invalid contract address");
        if (msg.value > 0) {
            (success, ) = contractAddr.call.value(msg.value)("");
            require(success, "Failed to send value");
        }
        GnosisSafe gnosisSafeContract = GnosisSafe(contractAddr);
        success = gnosisSafeContract.execTransaction(to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures);
        require(success, "Failed to exec tranasction");
    }
}
