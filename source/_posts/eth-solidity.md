---
title: 以太坊Solidity
date: 2018-01-18 10:28:20
tags:
---
# 以太坊区块链基础

## 交易:Transactions
 - 区块链是一个全局共享的交易数据库
 - 只要加入以太坊网络就可以获得这些数据
 - 如果你要改变数据库，你必须发送一个交易，并让其他人接受这笔交易
 - 交易总是被发送者签名的

> 交易的本质是消息，这个消息从一个账户发出到目标账户
> 交易中可携带payload和Ether
> 如果目标账户包含代码，那么payload将作为代码的输入数据
> 创建合约交易: Zero-Aaccount, 如果目标地址是0，则交易会创建新合约, 此时payload中包含的是合约代码, EVM发现是合约创建交易，会执行一段固定的代码返回交易中的payload,这个payload也就是code将被永久存储到区块链中???


## 区块:Blocks
 - 出现矛盾的交易总是会有一个先被打包进入区块，其他的交易将变得不合法从而被拒绝
 - 区块以一定的时间间隔出现

## 以太币:Ether

## 账户: Accounts
 - External Accounts: 被公钥私钥所控制
 - Contract Accounts: 被存储在账户中的代码所控制

## EVM: The Ethereum Virtual Machine
 - 智能合约本质上就是一段用来履行合约的代码
 - EVM是以太坊运行合约代码的虚拟机
 - 合约代码不能访问网络，文件系统，其他进程，可以有限的访问其他智能合约
 - EVM是stack machine
 - 指令集
 
### 账户地址: Address
 - External Accounts: 地址由公钥计算得出
 - Contract Accounts: 地址在合约创建的时候被生成，由创建者的地址+Noce计算得出

### Storage
每个账户内部有一个持久存储键值对的空间，叫做storage

### Balance

### Message Calls
合约可以调用其他合约 或者 向其他非合约账户发送Ether, Message
Calls类似于Transactions
  - souce
  - target
  - payload
  - Ether
  - gas
  - return data
  - Memory: 每MessageCalls会有一段内存空间叫memory, memory按照256bit的大小一块一块的，写的时候可以使8bits或者256bits,内存可以扩展，扩展要花费gas???


## 开发环境
  - ethereumjs-testrpc
  - truffle
  
```sh
$ npm install -g ethereumjs-testrpc
$ npm install -g truffle
$ mkdir hello-dapp
$ truffle init
Downloading...
Unpacking...
Setting up...
Unbox successful. Sweet!

Commands:

  Compile:        truffle compile
  Migrate:        truffle migrate
  Test contracts: truffle test

# 查看工程中的结构  
$ tree
.
├── contracts
│   └── Migrations.sol
├── migrations
│   └── 1_initial_migration.js
├── test
├── truffle-config.js
└── truffle.js

# 建立一个新合约
$ truffle create contract HelloDapp
$ tree
.
├── contracts
│   ├── HelloDapp.sol
│   └── Migrations.sol
├── migrations
│   └── 1_initial_migration.js
├── test
├── truffle-config.js
└── truffle.js
$

# 编译
$ 
```

## 定义一个合约
```
pragma solidity ^0.4.15;
contract <contract-name> {

}

# 合约也可以继承其他的合约
contract <contract-name1> is <contact-name2> {

}
```
### 状态变量: State Variable
``` 
contract C {
    uint storedData; // State variable
}
``` 

#### Constant State Variables
``` 
contract C {
    uint    constant x = 32**22 + 8;
    string  constant text = "abc";
    bytes32 constant myHash = keccak256("abc");
}
```

  - keccak256
  - sha256
  - ripemd160
  - ecrecover
  - addmod 
  - mulmod

### 合约函数
有两种类型的合约函数
 - ReadOnly或Constant函数     : 不改变合约状态, 不需要消耗gas
 - Trasaction函数数 ： 会改变合状态, 需要消耗gas才能执行

函数调用可分为两种:
 - 内部调用: 非message call
 - 外部调用: 又叫做message call
 
对于合约函数和状态变量，一般有四种可见度属性:
 - external 
   - 不适用于状态变量(StateVariable)
   - external函数是合约的调用接口， 表示其他合约可以调用这个方法
   - 如果是在合约内部则需要用this.f()的形式来调用
 - public
   - 合约函数默认是public的
   - public函数是合约的调用接口，同external
   - public的StateVariable会自带getter和setter方法
 - internal
   - StateVariable默认是internal的
   - internal的StateVariable和函数仅能在合约内部使用， 使用的时候不需要加this
 - private
   - 仅在合约内部使用， 不可被继承
 

```sol

function <function-name> [internal|external] [pure|constant|view|payable|] [returns()] (<params>) {
    <body>
}
```
 - function-name: 函数名
 - internal: 函数默认就是internal的，所以一般不用写
 - external:
 - pure: 确保此函数不会修改state
 - view: 确保此函数不会修改state， 通常getter函数都是view类型的
 - payable:
 - constant: 是否为constant函数
 - returns: 函数返回值
 
 > selector: this.<function-name>.selector
 
#### Fallback Function
合约还有一个没有名字的函数，这个函数也不能传入参数。所有找不到调用方法的最后都会调用到这个函数。

这个函数大概是这个样子:
``` 
function() public payable {...}
```
 - payable是可选的， 如果设置了， 则可以接受以太币
 
```
// 这个合约会接受所有发送给它的以太币，然后保存起来，而且没办法找回来 
contract Sink {
    function() public payable { }
}
```
#### 函数的重载

#### 构造函数
``` 
contract A {
    uint public a;

    constructor(uint _a) internal {
        a = _a;
    }
}
```
  - 其中被internal修饰了构造函数的合约只能用于继承

 
### 合约状态的变更
以下行为均被视为修改合约状态
 - 写入StateVariables
 - 提交事件
 - 建立合约
 - 自我销毁
 - 调用非view或pure函数
 - low-level调用
 - 执行包含特定opcode的内连汇编

 
### Data Location: [memory|storage]
数组和struct等复杂的数据类型，还可以制定一个存储位置， 可以是memory或者storage
 - 在函数参数中， 则默认是memory
 - 如果是本地变量或者状态变量， 则默认是storage
 
> calldata: 除了memory和storage之外，其实还有一个calldata区， 只读不可持久化的存储区域， 通常调用函数

``` 
pragma solidity ^0.4.0;

contract C {
    uint[] x; /*：storage */

    
    function f(uint[] memoryArray /* :memory */) public {
        x = memoryArray; /* x:storage (memoryArray背拷贝到storage中)  */
        var y = x;       /* y:storage (并没有复制操作，y指向x的storage) */
        y[7];            
        y.length = 2;    /* 实际上修改了x */
        delete x;        /* 实际上修改了x, y也随之受到影响 */
        g(x);            /* 调用g, 传递的是x的引用 */
        h(x);            /* 调用h, 将x引用的数组复制一份到memory中，然后传递给h处理 */
    }

    function g(uint[] storage storageArray) internal {}
    function h(uint[] memoryArray) public {}
}
```
### 合约中可以使用的数据类型
可以使用的数据类型:
 - bool
 - int, int8, int16, ... int256(int)
 - uint, uint8, uint16, ... uint256(uint)
 - fixedMxN: M:8-256, N:0-80
 - ufixedMxN: M:8-256, N:0-80
 - Address: 20byte
   - 0.5版本之前，合约本身集成了Address, 之后不是继承关系，但是合约仍然可以转换为Address类型
   - 针对一个地址类型可以使用以下操作(以下单位通通都是Wei)
     - <address>.balance: uint256, 账户余额
     - <address>.transfer(uint256 amount):
     - <address>.send(uint256 amount) returns (bool):
     - <address>.call(...) returns (bool):
     - <address>.callcode(...) returns (bool):
     - <address>.delegatecall(...) returns (bool):
 - bytes1(byte), bytes2, ..., bytes32
 - bytes: 动态长度的数组
 - string: 字符串
 - address: 地址字面量，通常39-41
 - 十六进制: hex"001122FF"
 - Enum: 

```
enum <enum-name> { value1, value2, ..., valueN }
```
> 注意: send和transfer的区别， send是transfer的底层版本， send发送失败不会导致合约终止， 所以使用send要检测其返回值。如非必须，尽量使用transfer


### 数组
 - bytes和string是特殊的数组
 - bytes类似于byte[]
 - strings等价与bytes, 但是禁止通过索引访问其中的元素

> bytes首选， 因为它的开销更低

在memory中创建数组需要使用new:
```
pragma solidity ^0.4.16;

contract C {
    function f(uint len) public pure {
        uint[] memory a = new uint[](7);
        bytes  memory b = new bytes(len);
    }
}
```

``` 
[1, 2, 3]     // uint8[3] memory
[uint(1),2,3] // uint[3] memory
```
#### push(x) : 向数组中追加元素
#### length : 数组的长度
> 由于EVM的限制， 函数的返回值不可以是动态长度的数组

## 结构体: struct
``` 
    struct <struct-name> {
        ...
    }
```

### 映射表: mapping
``` 
    mapping( <key-type> => <value-type> ) <variable-name>;
```
 - mapping只能在storage中
 - <key-type>不能是mapping类型, <value-type>可以是mapping类型
 - 无论key是什么，最终都是使用keccak256散列方法
 - mapping是无法遍历的， 但是如果我们将key保存在数组里，还是可以实现遍历的功能
 
### delete x: 相当于赋初始值
如果x是int, 则delete x等价于x=0

## 事件: events
### 定义事件
``` 
    event <event-name> (
        ...
    );
```

### 发送事件
``` 
    event Hello (
        address _from,
        uint _value
    );
    
    function hello(int value) {
        emit Hello(msg.sender, value); 
    }
```

## 继承: A is B, C
```
contract X {}
contract A is X {}
contract C is A, X {}
```

## 库: library
```
library L {
    function hello() public {
        ...
    }
}

// 调用
L.hello();
```

有关libary还有一个特殊的用法, using for
using for 可以将libary和某个数据对象绑定在一起，在库的内部可以通过self引用到数据对象， 从而可以扩展数据对象的能力。
``` 

library Search {
    function indexOf(uint[] storage self, uint value)
        public
        view
        returns (uint)
    {
        for (uint i = 0; i < self.length; i++)
            if (self[i] == value) return i;
        return uint(-1);
    }
}

contract C {
    using Search for uint[];
    uint[] data;

    function append(uint value) public {
        data.push(value);
    }

    function replace(uint _old, uint _new) public {
        // This performs the library function call
        uint index = data.indexOf(_old);
        if (index == uint(-1))
            data.push(_new);
        else
            data[index] = _new;
    }
}
```

## 接口: interface
诸如ERC20等合约均为接口定义.
``` 
interface Token {
    function transfer(address recipient, uint amount) public;
}
``` 

## 投票合约代码分析
```
pragma solidity ^0.4.16; // solidity语言版本

// 声名一个合约对象
contract Ballot {
  
    // voter对象，代表一个投票的人
    struct Voter {
        uint weight;      // Voter拥有的投票权，如果被主席授予投票权，则weight=1, 如果被人委托投票，则给weight加1
        bool voted;       // 是否已投票
        address delegate; // 委托投票人的账户地址
        uint vote;        // 投票结果(Proposal数组中某个索引)
    }

    // 提议对象
    struct Proposal {
        bytes32 name;    // 名字
        uint voteCount;  // 总票数
    }

    address public chairperson; // 主席

    mapping(address => Voter) public voters; // 参与投票的人

    Proposal[] public proposals;             // 提议对象集合


    // 发起一轮投票，需要给出提议对象的集合作为输入参数
    function Ballot(bytes32[] proposalNames) public {
        chairperson = msg.sender;  //发起人即主席
        voters[chairperson].weight = 1; // 主席拥有一票

        // 构建提议对象
        for (uint i = 0; i < proposalNames.length; i++) {
            proposals.push(Proposal({
                name: proposalNames[i],
                voteCount: 0
            }));
        }
    }

    // 授予某个voter投票权
    function giveRightToVote(address voter) public {
        
        require((msg.sender == chairperson)        // 必须是主席身份
                && !voters[voter].voted            // voter没有投票
                && (voters[voter].weight == 0));   // voter没有任何投票权
        voters[voter].weight = 1;                  // 分配一个投票权
    }


    // 将投票权委托给某个voter
    function delegate(address to) public {
        Voter storage sender = voters[msg.sender];
        require(!sender.voted); // 发起委托的voter必须还没有投票

        require(to != msg.sender); // 不能委托给自己

        // 有一种情况可能存在， 即被委托人已经将自己的投票权委托给另外的人,
        // 形成一个委托链，如果最终委托链指向了当前委托的人，那么这种情况是不允许的
        // 如: A想委托给B
        // B -> C      : 这种情况是OK的
        // B -> C -> A : 因为A要把自己的投票委托出去，但A又被委托了，因此不能接受这种委托， 因为一旦允许，A,B,C相当于没有人投票了
        while (voters[to].delegate != address(0)) {
            to = voters[to].delegate;
            require(to != msg.sender);
        }

        sender.voted = true;  // 委托者在委托成功后相当于完成了投票
        sender.delegate = to; // 记录委托人

        Voter storage delegate = voters[to];
        if (delegate.voted) {
            // 如果被委托人已经投票，那就给被委托人选择的提议增加票数
            proposals[delegate.vote].voteCount += sender.weight;
        } else {
            // 否则就把投票权转移给被委托人
            delegate.weight += sender.weight;
        }
    }

    // 投票
    function vote(uint proposal) public {
        Voter storage sender = voters[msg.sender];
        require(!sender.voted);
        sender.voted = true;
        sender.vote = proposal;

        proposals[proposal].voteCount += sender.weight; //给对应的提议增加票数， 委托人通常至少有2票，委托人只能将全部的票投给一个提议
    }

    // 获胜提议的索引
    function winningProposal() public view
            returns (uint winningProposal)
    {
        // 选出票数最多的即可
        uint winningVoteCount = 0;
        for (uint p = 0; p < proposals.length; p++) {
            if (proposals[p].voteCount > winningVoteCount) {
                winningVoteCount = proposals[p].voteCount;
                winningProposal = p;
   ER         }
        }
    }

    // 获胜提议的名字
    function winnerName() public view
            returns (bytes32 winnerName)
    {
        winnerName = proposals[winningProposal()].name;
    }
}
```

## address.transfer
使用address.transfer可以给其他账户转以太币，也可以给合约转以太币， 有几点需要注意一下:
 - 如果地址是合约地址，则会调用合约的FallbackFunction
 - 如果调用层级超过1024, 则会转账失败
 
## 从合约中提现(Withdrawal from Contract)

### pragma solidity <version>;
### import <filename>;
### import * as <alias> from <filename>;
### import {<s1> as <alias>,s2,...} from <filename>;
### import <filename> as <alias>;
### State Variables
永久存储在合约中的变量
```
contract Hola {
    uint data; // State Variables
}
```

    - bool
    - int
    - uint
    - uint9
    - uint256
    - fixed
    - ufixed
    - fixedMxN
    - ufixedMxN
    - address

> M: 8 16 24 32 40 48 56 64 ... 256
> N: 0-80

### 以太币单位
```
wei
finney
szabo
ether
```

### 时间单位
```
1         : 1 seconds
1 minutes : 60 seconds
1 hours   : 60 minutes
1 days    : 24 hours
1 weeks   : 7 days
1 years   : 365 days
```
> 注意，这种单位不可以给变量赋值，仅可用于比较

### 枚举: enum <name> {value1,value2,...,valueN}
### 函数 
> function <name> (<parameters>) {interal|external|private|public} [pure|constant|view|payable] [return (<types>)]

#### external
    - 是合约接口
    - 可以被其他合约调用
    - 在内部必须用this调用
#### public
    - 是合约接口
    - 即可被外部合约调用也可以直接内部调用
    - public的变量会自动生成gette方法, getter方法是external的

#### internal
    - 只能在合约内部访问
    - 派生合约也可以访问

#### private
    - 智能在合约内部访问
    - 派生合约不可访问

> 注意: 虽然private/internal这类修饰可以阻止外部合约使用，但是这些变量方法对于整个世界都是可见的
### 全局对象和全局函数
#### address是个对象
```
<address>.balance (uint256):                   balance of the Address in Wei
<address>.transfer(uint256 amount):            send given amount of Wei to Address, throws on failure
<address>.send(uint256 amount) returns (bool): send given amount of Wei to Address, returns false on failure
<address>.call(...) returns (bool):            issue low-level CALL, returns false on failure
<address>.callcode(...) returns (bool):        issue low-level CALLCODE, returns false on failure
<address>.delegatecall(...) returns (bool):    issue low-level DELEGATECALL, returns false on failure
```
> 注意: 因为contract也是继承于address,所以在合约中使用this.balance可以查看当前账户的balance
#### array

#### 区块和交易相关的属性
```
block.blockhash(uint blockNumber) returns (bytes32): 给定区块的Hash, 仅支持最近的256个区块
block.coinbase (address):                            当前矿工的地址
block.difficulty (uint):                             当前区块的难度
block.gaslimit (uint):                               当前区块的gaslimit
block.number (uint):                                 当前区块号
block.timestamp (uint):                              当前区块时间戳(unix epoch)
msg.data (bytes):                                    Payload
msg.gas (uint):                                      剩余的gas
msg.sender (address):                                消息/交易发送者的地址
msg.sig (bytes4):                                    msg.data的前4bytes
msg.value (uint):                                    发送者为这条消息附带的Ether, (单位是wei)
now (uint):                                          相当于block.timestamp
tx.gasprice (uint):                                  gas价格
tx.origin (address):                                 sender of the transaction (full call chain)
```
#### 加密函数
```
addmod(uint x, uint y, uint k) returns (uint): (x+y) % k
mulmod(uint x, uint y, uint k) returns (uint): (x*y) % k
keccak256(...) returns (bytes32): Ethereum-SHA-3 (Keccak-256) hash
sha256(...) returns (bytes32):    SHA256 Hash
sha3(...) returns (bytes32):      SHA3 Hash
ripemd160(...) returns (bytes20): RIPEMD-160 Hash
ecrecover(bytes32 hash, uint8 v, bytes32 r, bytes32 s) returns (address): 

```

## 参考
 - [vim-solidity](https://github.com/tomlion/vim-solidity/blob/master/README.md)