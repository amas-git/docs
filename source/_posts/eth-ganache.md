---
title: Ganache:以太坊私有链开发
date: 2018-01-15 16:08:31
tags:
---
# Ganache
> https://github.com/trufflesuite/ganache
## 使用solc编译以太坊合约
Hola.sol:
```
pragma solidity ^0.4.0;

// just store a number
contract Hola{
    uint _n;

    function set(uint n) public {
        _n = n;
    }

    function get() public constant returns (uint) {
        return _n;
    }
}
```

```
$ npm install solc
$ node
> code = fs.readFileSync('Hola.sol').toString()
...
> solc.compile(code)
{ contracts: 
   { ':Hola': 
      { assembly: [Object],
        bytecode: '6060604052341561000f57600080fd5b60d38061001d6000396000f3006060604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806360fe47b114604e5780636d4ce63c14606e575b600080fd5b3415605857600080fd5b606c60048080359060200190919050506094565b005b3415607857600080fd5b607e609e565b6040518082815260200191505060405180910390f35b8060008190555050565b600080549050905600a165627a7a7230582074f04aadab7a9c5e94eed3085df4e2479330b586a25c56f042c06fe92979fe470029',
        functionHashes: [Object],
        gasEstimates: [Object],
        interface: '[{"constant":false,"inputs":[{"name":"n","type":"uint256"}],"name":"set","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]',
        metadata: '{"compiler":{"version":"0.4.19+commit.c4cbbb05"},"language":"Solidity","output":{"abi":[{"constant":false,"inputs":[{"name":"n","type":"uint256"}],"name":"set","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}],"devdoc":{"methods":{}},"userdoc":{"methods":{}}},"settings":{"compilationTarget":{"":"Hola"},"libraries":{},"optimizer":{"enabled":false,"runs":200},"remappings":[]},"sources":{"":{"keccak256":"0x6b33e855d2ece02355cacc4d3f225774b8e37be189fe188092b2118bc8ec5a0f","urls":["bzzr://54182ae4934b15b3ed71063a6b314a7f1453f3ea6073bddf19efdbd13675286f"]}},"version":1}',
        opcodes: 'PUSH1 0x60 PUSH1 0x40 MSTORE CALLVALUE ISZERO PUSH2 0xF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0xD3 DUP1 PUSH2 0x1D PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x60 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH1 0x49 JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x60FE47B1 EQ PUSH1 0x4E JUMPI DUP1 PUSH4 0x6D4CE63C EQ PUSH1 0x6E JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST CALLVALUE ISZERO PUSH1 0x58 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x6C PUSH1 0x4 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH1 0x94 JUMP JUMPDEST STOP JUMPDEST CALLVALUE ISZERO PUSH1 0x78 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x7E PUSH1 0x9E JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST DUP1 PUSH1 0x0 DUP2 SWAP1 SSTORE POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD SWAP1 POP SWAP1 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 PUSH21 0xF04AADAB7A9C5E94EED3085DF4E2479330B586A25C JUMP CREATE TIMESTAMP 0xc0 PUSH16 0xE92979FE470029000000000000000000 ',
        runtimeBytecode: '6060604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806360fe47b114604e5780636d4ce63c14606e575b600080fd5b3415605857600080fd5b606c60048080359060200190919050506094565b005b3415607857600080fd5b607e609e565b6040518082815260200191505060405180910390f35b8060008190555050565b600080549050905600a165627a7a7230582074f04aadab7a9c5e94eed3085df4e2479330b586a25c56f042c06fe92979fe470029',
        srcmap: '48:164:0:-;;;;;;;;;;;;;;;;;',
        srcmapRuntime: '48:164:0:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;81:51;;;;;;;;;;;;;;;;;;;;;;;;;;138:72;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;81:51;124:1;119:2;:6;;;;81:51;:::o;138:72::-;178:4;201:2;;194:9;;138:72;:::o' } },
  sourceList: [ '' ],
  sources: { '': { AST: [Object] } } }

```

类型:
    - address: 160bit

Coin.sol: 造币合约
```
pragma solidity ^0.4.0;

contract Coin {
    // The keyword "public" makes those variables
    // readable from outside.
    address public minter;
    mapping (address => uint) public balances;

    // Events allow light clients to react on
    // changes efficiently.
    event Sent(address from, address to, uint amount);

    // This is the constructor whose code is
    // run only when the contract is created.
    function Coin() public {
        minter = msg.sender;
    }

    function mint(address receiver, uint amount) public {
        if (msg.sender != minter) return;
        balances[receiver] += amount;
    }

    function send(address receiver, uint amount) public {
        if (balances[msg.sender] < amount) return;
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        Sent(msg.sender, receiver, amount);
    }
}
```
    - mapping (address => uint) : map
    - address public minter:
        - address: 160bit
        - public: 这个属性可以被访问，solidity会自动生成一个对应名字的get方法: minter()

> function minter() returns (address) { return minter; }
> function balances(address _account) public view returns (uint) {
>    return balances[_account];
> }

## 安装
```
$ npm install -g ganache-cli
```

安装正经的ganache,
ganache的安装非常简单，因为它用AppImage打包的，所以就是下载一个文件，然后直接运行就可以了，以我的ArchLinux为例:
```
$ wget https://github.com/trufflesuite/ganache/releases/download/v1.0.2/ganache-1.0.2-x86_64.AppImage
$ chmod +x ganache-1.0.2-x86_64.AppImage
$ ./ganache-1.0.2-x86_64.AppImage
```



