// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Dummy {
    uint256 public x;

    constructor() {
        x = 10;
    }

    function setX(uint256 _x) public {
        x = _x;
    }
}
