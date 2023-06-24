// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "forge-std/Script.sol";
import "../src/KeyValue.sol";

contract KeyValueScript is Script {

    function run() external  {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // KeyValue keyValue = new KeyValue();

        vm.stopBroadcast();
    }
}
