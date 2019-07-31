## 比特币交易

四要素:

- 版本
- Input
- Output
- Locktime

### 二进制形式

> 注意: 交易体二进制采用小端存储

![](/src/amas/docs/source/_drafts/assets/2019-05-16-150954_1212x249_scrot.png)

variable-length field

### 版本号

通常是1, 有一些特别场景会出现2, BIP0112中使用OP_CHECKSEQUENCEVERIFY.



## Inputs

包括记录两个东西:

1. 一个output引用, 你所花的钱必有来源 
2. 证明这个交易确实是属于你(用ECDSA签名)


Input中包含的字段:
  - Previous transaction ID
  - Previous transaction index
  - ScriptSig
  - Sequence: Replace-By-Fee (RBF) and OP_CHECKSEQUENCEVERIFY




> Sequence and Locktime
>
> high-frequency trade transaction
>
> 这个玩意设计的目的是在于高频交易,啥是高频交易, 就是A和 B来来回回频繁支付的场景下, 比如A给B 2BTC, B给A 0.5BTC, 其实只要用一笔交易记录A给B 1.5 BTC. 此种场景下优化链上的交易数量.
>
> 最先这笔交易需要
>
> sequence为0, 同时指定一个far-away locktime, 比如100, 就是说往后100个块有效,
>
> 这个功能有漏洞,矿工可能作弊



### Outputs

outputs里面记录了BTC的去向, 至少有一个, 可以有多个. 

> 比如交易所可以利用批量支付, 把很多人的交易放在一个tx中

每个output中包含

- amount: 比特币数量, 聪为单位,  0.001,000,000,000
- ScriptPubKey.(variable-length field)



### UTXOs

>  The entire set of unspent transaction
> outputs at any given moment is called the UTXO set
>
> UTXO SET: 
>
> 没有被使用的output就构成了UTXO SET
>
> 这东西之所以重要是因为这个其实才是你真正能使用的钱,你的地址里有多少比特币, 其实就是UTXO set

如果一个交易使用了UTXO set以外的output, 那明显就是出现了双花.



### Locktime

Locktime主要是为了应对高频交易,  取值是整整数. 是一种延迟的交易,locktime > 500,000,000时是Unix时间戳, 小于的话是块高度, 当时间或者块的高度没有大于locktime的时候, 这个交易可以被签名, 但是不能被花费, 直到超过locktime.



> locktime相当于银行的post dated check, 现实生活中的用途包括:
>
> ### [Post dated check](https://www.accountingtools.com/articles/what-is-a-post-dated-check.html)
>
> - *Deliberate payment delay*. The issuer does this in order to delay payment to the recipient, while the recipient may accept it simply because the check represents a firm date on which it will be able to deposit the check. This situation represents a risk to the check recipient, since the passage of time may result in there being no cash left in the issuer's [bank account](https://www.accountingtools.com/articles/2017/5/15/bank-account-types) to be used to pay the amount listed on the check when it is eventually presented to the bank for payment.
> - *Collection method*. The recipient may require the issuer to hand over a set of post dated checks to cover a series of future payments, which the recipient agrees to cash on the specified dates. This approach is used to improve the odds of being paid, especially when the issuer has little credit.



使用locktime, 收款方会存在一定的风险, 当到达locktime的时候,很可能付款方已经余额不足.

> 另外, 当input里面的secquence为FFFFFFFF时候忽略locktime



### 矿工费

> inputs = outputs  + tx_fee

每笔交易需要支付矿工费,所以outputs通常小于等于inputs



如何计算矿工费?

> 我们知道inputs - outputs就是这个交易的矿工费, 但是outputs里面每一笔是有amount字段的, 但是inputs确没有, 因此需要计算inputs的总量. 



>
>
>### Why We Minimize Trusting Third Parties
>As Nick Szabo eloquently wrote in his seminal essay “Trusted
>Third Parties are Security Holes”, trusting third parties to provide
>correct data is not a good security practice. The third party may be
>behaving well now, but you never know when it may get hacked,
>have an employee go rogue, or start implementing policies that are
>against your interests. Part of what makes Bitcoin secure is not
>trusting, but verifying the data that we’re given.





## Script

OP CODE:

|              |      |      |
| ------------ | ---- | ---- |
| OP_DUP       | 0x76 |      |
| OP_CHECKSIG  |      |      |
| OP_0         | 0x00 |      |
| OP_1         | 0x51 |      |
| OP_16        | 0x60 |      |
| OP_ADD       | 0x93 |      |
| OP_HASH160   | 0xa9 |      |
| OP_CHECKSIG  | 0xac |      |
| OP_PUSHDATA1 |      |      |
| OP_PUSHDATA2 |      |      |
| OP_PUSHDATA4 |      |      |
|              |      |      |

标准脚本

|        |                            |      |
| ------ | -------------------------- | ---- |
| p2pk   | pay to pubkey              |      |
| p2pkh  | pay to pubkey hash         |      |
| p2sh   | pay to script hash         |      |
| p2wpkh | pay to witness pubkey hash |      |
| p2wsh  | pay to witness-script hash |      |
|        |                            |      |
|        |                            |      |
|        |                            |      |

完成一个支付, 首先要有两个东西

1. UTXOs里的ScriptPubkey, 这个证明了UTXOs是谁的
2. 当前交易里的ScriptSig

## p2pk

早期交易中大量使用的脚本.也叫p2pk UTXOs, 

p2pk ScriptPubKey:

```
xx  - 公钥长度
... - 公钥
ac  - OP_CHECKSIG 
```



p2pk ScriptSig (解锁脚本):

```
xx  - 签名长度
... - 签名
```





p2pk ScriptSig和p2pk ScriptPubKey拼在一起:

script:

```
-----------------------------[ p2pk ScriptSig    ]
xx  - 签名长度
... - 签名
-----------------------------[ p2pk ScriptPubKey ]
xx  - 公钥长度
... - 公钥
ac  - OP_CHECKSIG 
```

执行过程:
```
#-----------------------------------| 1.初始状态 
SCRIPT       STACK
signature
pubkey
OP_CHECKSIG
#-----------------------------------| 2.参数进栈   
SCRIPT       STACK
OP_CHECKSIG
             signature
             pubkey
#-----------------------------------| 3.执行OP_CHECKSIG            
SCRIPT       STACK
             0或1                        
```


> p2pk的问题
>
> 1. 公钥太长了, SEC格式的公钥或33或65 bytes, 为了传输用hex编码, 结果最长要66或130个字符
> 2. p2pk用于Ip-to-Ip的支付, 机器之间无所谓长短
> 3. 导致UTXO set变大
> 4. pubkey暴露在外, 万一那天ECDSA被破解了, 那么这些p2pk的UTXO就都可以被偷走了



为了解决以上问题, 产生了p2pkh

## p2pkh

>p2pkh的优点
>
>1. 地址更短
>2. 不再长期暴露公钥, 安全性更好(注意: p2pkh把公钥移动到ScriptSig中)

p2pkh ScriptPubKey:

```
76  - OP_DUP
a9  - OP_HASH160
xx  - Hash长度
... - Hash
88  - OP_EQUALVERIFY
ac  - OP_CHECKSIG
```



p2pkh ScriptSig

```
xx  - 签名长度 
... - 签名(DER格式)
xx  - 公钥长度
... - 公钥
```



我们来看下脚本是怎么执行的:

```
#-----------------------------------| 1.初始状态 
SCRIPT           STACK
signature
pubkey
OP_DUP
OP_HASH160
hash
OP_EQUALVERIFY
OP_CHECKSIG
#-----------------------------------| 2.参数进栈   
SCRIPT           STACK
OP_DUP
OP_HASH160
hash
OP_EQUALVERIFY
OP_CHECKSIG
                 signature
                 pubkey
#-----------------------------------| 3.执行OP_DUP         
SCRIPT           STACK
OP_HASH160
hash
OP_EQUALVERIFY
OP_CHECKSIG
                 signature
                 signature
                 pubkey   
#-----------------------------------| 4.执行OP_HASH160         
SCRIPT           STACK
hash
OP_EQUALVERIFY
OP_CHECKSIG
                 hash'
                 signature
                 pubkey       
#-----------------------------------| 5.执行OP_HASH160         
SCRIPT           STACK
OP_EQUALVERIFY
OP_CHECKSIG      hash
                 hash'
                 signature
                 pubkey   
#-----------------------------------| 6.OP_EQUALVERIFY      
SCRIPT           STACK
OP_CHECKSIG     
                 signature
                 pubkey 
#-----------------------------------| 6.OP_EQUALVERIFY      
SCRIPT           STACK
                 0或1
```



> 注意: ,如果input中的ScriptSig可以让指定的output中的ScriptPubKey计算结果为1, 就好象钥匙打开了锁一样, 这个过程也叫做unlock. 



了解了签名脚本, 你就可以实现一个自己的签名脚本,当这个脚本执行成功就可以解锁比特币出来.

## 交易的创建和验证

![](/src/amas/docs/source/_drafts/assets/2019-05-20-135006_583x196_scrot.png)

有三方面:

1. inputs里面的utxos是不是有效的
2. inputs比特币 >= outputs比特币
3. ScriptSig可以解锁ScriptPubKey

>inputs >= outputs, 且inputs > 0
>
>### The Value Overflow Incident
>Back in 2010, there was a transaction that created 184 billion new
>bitcoins. This was due to the fact that in C++, the amount field is a
>signed integer and not an unsigned integer. That is, the value could
>be negative!
>The clever transaction passed all the checks, including the one for
>not creating new bitcoins, but only because the output amounts
>overflowed past the maximum number. 2 64 is ~1.84 × 10 19 satoshis,
>which is 184 billion bitcoins. The fee was negative by enough that
>the C++ code was tricked into believing that the fee was actually
>positive by 0.1 BTC!
>The vulnerability is detailed in CVE-2010-5139 and was patched
>via a soft fork in Bitcoin Core 0.3.11. The transaction and the extra
>bitcoins it created were invalidated retroactively by a block reor‐
>ganization, which is another way of saying that the block including
>the value overflow transaction and all the blocks built on top of it
>were replaced



## 创建交易

三要素

1. 钱从哪来
2. 钱到哪去
3. 你需要多快的交易速度(给多少手续费)

> ## 为什么重复使用地址是不安全的?
>
> 迄今为止,比特币的资产安全根本上是由DLP问题求解难度保证的.
>
> 1. p2pk的资产安全性全靠ECDSA
> 2. p2pkh依靠ECDSA和SHA256+RMP160
>
> p2pkh
>
> 尽管与p2pk相比不会把公钥保存到区块里(outputs的ScriptPubKey中), 但是当发起一笔交易的时候, 仍然需要暴露公钥, 所以重复使用一个地址, 就增加了暴露的机会, 这使得SHA256+RMP160不再有用.
>
> ## 怎么做?
>
> 支付的时候使用一个地址, 剩下的钱招零到新地址里. 这样就只需要暴露一次公钥, 而且最好这个地址不要再使用, 没有攻击者会对一个没钱的地址感兴趣.
>
> ## 如何计算手续费?
>
> 交易字节数 x 单价
>
> 注意: Segwit中有点不同
>
> ## 如何估计单价?
>
> 1. 根据历史成交
> 2. 根据当前mempool中的出价
> 3. 固定手续费

## p2sh

>  p2pk和p2pkh是单签名交易, 每个input只需要用同一个私钥签名即可.
>
>  p2sh是一个多签名交易,

>Satoshi probably didn’t test multisig, as it has an off-by-one error (see
>OP_CHECKMULTISIG Off-by-One Bug). The bug has had to stay in the protocol
>because fixing it would require a hard fork.



> Schnorr签名可能未来会替代 ECDSA, 其对多签名具有更好的支持. 相同点是二者本质安全都靠DLP.



## 什么是M-of-N Multisig, Multisig Output

m-of-n多签名, 其实还有一个条件m <=n, m就是minimum, 有n个签名, 至少需要用m个签名通过才能解锁比特别. 

现实生活中有非常多的支付场景需要m-of-n多签名. 

|      |       |
| ---- | ---- |
|    1-of-2  | 夫妻两人的共同账户, 任意一人的签名都支付           |
|    2-of-2  | 夫妻的大额共同存款, 需要两个人的签名才可以支付            |
|    2-of-3  | 三个合伙人拥有的共同基金, 需要至少两个人的签名才能支付            |

https://bitcoin.org/en/developer-examples#offline-signing

多签名脚本:



### 什么是Bare Multisig?

## BIP0016:RedeemScript



## 交易的延展性:Transaction Malleability

延展性是个装逼十足的说法,其实就是可变性.



交易延展性指的是指改变交易ID而不改变交易的含义.

交易的ID具有可变性, 这听起来有点不可思议, 但的确如此. 因为在计算交易ID的时候, 

是从交易体计算而来的,  绝大多数交易体中的字段都是不能随便更改的, 需要验证签名, 但是有一个例外, 每一个input单色ScriptSig可以不经过签名校验就更改. 这样就带来一个问题,如果有人修改了ScriptSig, 那么交易ID也随之改变. 这个有什么用呢? 首先如果修改了ScriptSig, 这肯定是不能解锁ouput的. 如果基于这个修改过的交易再构建出一串交易, 这一堆交易也都是最终不会进入到区块链的. 它的真正影响在于对支付通道的影响, 比如闪电网络, 有自己的侧链,如果同一个交易ID变来变去就会比较危险了. 另一方面, 虽然这种可变的交易最终不会确认, 但是我们可以制造大量这样的交易链, 然后以此DoS比特币网络.





我们怎么消灭可变性?

解决方案要兼容老的版本, 采用增加一个witness的字段来存储ScriptSig

### BIP0143: Segwit是什么?

> 隔离见证就是将签名信息保存到新的字段里, 为了兼容老客户端, 需要搞一个软分叉, 这个软分叉从477,120.开始.



> ### 解决什么问题?
>
> 1. 解决了交易ID不可变
> 2. 闪电网络的基础

Segwit多了三个字段

- marker

- flag

- witness
  - signature

一部分矿工, 绝大多是中国矿工在软分叉的时候发生了分期, 后来成为了BCH.

### Bech32, defined in BIP0173

## p2sh-p2wpkh

p2wpkh虽然挺好,但是是一种新的支付脚本, 老钱包不支持,咋整? 工程师来解决, 将

p2wpkh搞到p2sh里

p2sh-p2wpkh地址是一个正常的p2sh地址, RedeemScript是[OP_0, <20 bytes hash>]







### Contributing

```

```

A large part of the Bitcoin ethic is contributing back to the community. The main way
you can do that is through open source projects. There are almost too many to list,
but here’s a sample:
Bitcoin Core
The reference client
Libbitcoin
An alternate implementation of Bitcoin in C++
btcd
A Golang-based implementation of Bitcoin
Bcoin
A JavaScript-based implementation of Bitcoin, maintained by purse.io
248
|
Chapter 14: Advanced Topics and Next Stepspycoin
A Python library for Bitcoin
BitcoinJ
A Java library for Bitcoin
BitcoinJS
A JavaScript library for Bitcoin
BTCPay
A Bitcoin payment processing engine written in C#

```

```