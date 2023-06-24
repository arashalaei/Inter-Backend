// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "forge-std/Test.sol";
import "../src/KeyValue.sol";
/**
 * @title KeyValue smart test
 * @author Arash Alaei <arashalaei22@gmail.com>
 * @notice It tests all functions in the KeyValue smart contract
 */
contract KeyValueTest is Test {
    // State variables
    KeyValue public keyValue; // an instance of KeyValue smart contract


    function setUp() public {
        keyValue = new KeyValue();   
        keyValue.set(1, 1);     
        keyValue.set(3, 9);     
    }

    /**
     * @notice test get function
     */
    function testGet() public {
        assertEq(keyValue.get(1), 1);
        assertEq(keyValue.get(3), 9);
    }

    /**
     * @notice test set function
     */
    function testSet(uint256 _key, uint256 _value) public{
        keyValue.set(_key, _value);
        assertEq(keyValue.get(_key), _value);
    }
}
