pragma solidity >=0.5.0 <0.7.0;

/// @title TestContract calls a Gnosis Safe contract and check return values.
contract TestContract {
    /// Call revert
    function callRevert() public payable {
        revert();
    }

    /// Call require
    function callRequire(bool ok) public payable {
        require(ok);
    }

    /// Call assert
    function callAssert(bool ok) public payable {
        assert(ok);
    }
}
