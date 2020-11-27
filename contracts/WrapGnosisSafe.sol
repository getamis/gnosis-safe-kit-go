pragma solidity >=0.5.0 <0.7.0;
import "./GnosisSafe.sol";

/// @title Wrap Gnosis Safe calls a Gnosis Safe contract and check return values.
contract WrapGnosisSafe {
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
        GnosisSafe gnosisSafeContract = GnosisSafe(contractAddr);
        contractAddr.transfer(msg.value);
        success = gnosisSafeContract.execTransaction(to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures);
        require(success, "Failed to exec tranasction");
    }
}
