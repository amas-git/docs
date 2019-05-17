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
  - ScriptSig: 签名脚本
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