// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/**
 * @title Interview Assignment for Backend/Blockchain Engineer
 * @author Arash Alaei <arashalaei22@gmail.com>
 * @notice a smart contract that can store key-value pairs and retrieve values using keys.
 */
contract KeyValue {
    // State variables
    mapping (uint256 => uint256) private store;

    // Functions
    /**
     * 
     * @param _key The key by which we want to retrieve the data from the map
     * @return The value retrieved
     */
    function get(uint256 _key) public view returns(uint256){
        return store[_key];
    }

    /**
     * 
     * @param _key The key with which we want to store a value
     * @param _value The value retrieved by the key
     */
    function set(uint256 _key, uint256 _value) public {
        store[_key] = _value;
    }
    
}