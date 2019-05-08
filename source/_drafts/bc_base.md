

## 整数除余数运算
## 有限域
## 质因数分解

破解质因数分解:

 - http://mathworld.wolfram.com/QuadraticSieve.html
 - https://en.wikipedia.org/wiki/General_number_field_sieve

## 陷门函数

> Q: 新中国成立是几月几日? 
>
> A: 10月1日
>
> Q: 有件事情发生于10月1日,你知道是什么事么?
>
> A: 鬼才知道

## 椭圆曲线(Elliptic Curve)

> $y^2 = ax^3+ax+b$

EC在整数,复数,实数集合上有不同的表现.

### F~p~

> TODO:
>
> 1. 绘制F~101~ 在整数集合上的散点图
> 2. 计算某个点的内积
> 3. 为什么积是离散对数问题?

### 椭圆曲线的加法

### 椭圆曲线的乘法

> $A+A = 2A$

椭圆曲线乘法比较容易计算, 可是它的逆运算也就是椭圆曲线除法非常困难, 利用这个性质就可以构造陷门函数从而实现非对称加密.





## 椭圆曲线加密算法

离散对数问题:

Problems Related to DLP
I Given an abelian group (G, ·) and g ∈ G of order n.
I Discrete Logarithm Problem (DLP) :
Given h ∈ G such that h = g
x find x. (DLP(g, h) → x)
I Computational Diffie-Hellman Problem (CDH) :
Given a = g
x and b = g
y find c = g
xy (CDH(g, a, b) → c).
I Decisional Diffie-Hellman Problem (DDH) :
Given a = g
x
, b = g
y and c = g
z
, determine if
g
xy = g
z or equivalently xy ≡ z mod n
(DDH(g, a, b, c) → true/false)
Frederik Vercauteren ESAT/COSIC - K.U. Leuven ECRYPT Summer School 2008 Disc



## 地址

> $K = kG$

1. k是一个随机数($0< k <2^256$)
2. K是公钥, 是一个坐标(x,y) (x和y均为256bit)

知道k和G很容易算出K, 但是通过K很难算出k, 这个就是离散对数问题决定的.

## 椭圆曲线签名算法(ECDSA)

> Elliptic Curve Digital Signature Algorithm
>
> 假如:
>
> 你有k, 并且计算出K = kG, 其他人知道K和G, 怎么向其他人证明你确实知道k呢?这个就是签名算法的核心问题

### 如何计算签名

思路大概是这样的:

首先我们随机选取一个点, 这个点记为R, 坐标为(r,y),  我们需要做的是从K点出发命中R点

现在我们就有三个为众人所知的点, 分别是G, K, R

接下来我们要要算出两个值u, v, 使得这三个点之间满足

> $uG + vK = R$



为此可以推到出:

> $uG + vkG = mR$
>
> $k = (m-u) / v$



接下来计算s:

1. $u = digest / s$
2. $v = r /s$
3. $k = (m - u) /v$



> $k = (m - (digest/s)) / (r/s)$
>
> $k = (m - (digest/s)) s/r$
>
> $k = (sm - digist)/r$
>
> $kr = sm - digest​$
>
> $s = (kr + digest)/m​$

m: 保密, 但对应的R点是公开的

digest: 公开

r: 公开

s: 公开, 代表了r,k,m,digest之间的关系, 因为r,digest是公开的,  其实由k,m可以唯一确定s?

最后得到签名:

> (r,s)即为签名



那么问题来了, 当你知道了r,s也就知道了u,v, 那么你是否能通过r,s,u,v计算出k? 我们知道:

> $s = (kr + digest)/m$

其中m和k你都是不知道的, 因此想要通过s,r,digest逆推出k是很难的.

这里其实也是一个陷门函数, 当你知道k和m, 计算s很容易, 但是想从s得到k, m是不太可能的,虽然找到一个(k,m)满足等式并不难, 但恰好是真正的(k,m)的概率极低,几乎不可能.

> 不信我们来玩一个游戏
>
> 我随便想到两个数m和n把它记在纸上不让你看
>
> 我告诉你$m/n=6$
>
> 我给你一百年时间你来告诉我m和n分别是多少? 猜中给你一个亿



### 如何验证签名

签名验证则需要验证以下公式相等即可

> $uG + vK = mG = R$
>
> 1. 计算消息的摘要,记为digest'
> 2. $ u = digest' / s$
> 3. $v = r/s$

最后计算出R, 只要R的x坐标等于r, 就可以证明你能够通过K找到一点射中R, 本质上是到合适(u,v)值, 只有知道k的人才能轻易做到, 于是也就证明了签名的有效性.



## secp256k1
> $y^2=x^3+7​$


这个曲线长这样: https://www.wolframalpha.com/input/?i=y%5E2+%3D+x%5E3+%2B+7

```
## 代码

​```js
const F = function (p) {
  return o = {
    add(x, y) {
      return (x + y) % p;
    },

    mui(x, y) {
      return (x * y) % p;
    },

    // 检测运算是否封闭
    isClosed() {

    },

    // x/y = x*y^(p-2) (mod p)
    div(x, y) {
      //return o.mui(x, y**(p-2n));
      return o.mui(x, o.pow(y, (p-2n)));
    },

    pow(x, n) {
      // 可以优化
      n = (n % (p - 1n));
      // -----------
      return (x**n) % p;
    }
  };
};

const Z = function(g, p) {
  return o = {
    permutation(max) {
      let rs = [];
      for(i=1n; i<max; ++i) {
        let r = (g**i) % p;
        rs.push(r);
      }
      return rs;
    }
  };
};


// 费马小定理
// p是质数
// a**(p-1) = 1 (mod p)
console.log((10**(3-1)  % 3));
console.log((66666666666666666666666n**(3n-1n)  % 3n)); // a 不能是p的倍数
console.log((66666666666666666666661n**(3n-1n)  % 3n)); // a 不能是p的倍数
console.log((10**(7-1)  % 7));
console.log((10**(13-1) % 13));
console.log((8**(13-1) % 13));
console.log((2**(13-1) % 13));

// x/y = x * (1/y) = x * ((y**(p-1) / y)) = x * (y**(p-2))

// 如果y是p的倍数
// x/y = x/p = x/0  // 这个是没有意义的


const F19 = F(19n);
console.log(F19.mui(3n,7n));  // 2
console.log(F19.div(2n,7n));  // 3
console.log(F19.mui(3n,7n));  // 2

console.log(F19.mui(4n, 5n)); // 1
console.log(F19.div(1n, 5n)); // 4


console.log(F19.mui(10n, 13n)); // 16

// y % 19 = 0
console.log(F19.div(16n, 19n)); // 0?

console.log(F19.pow(2n,13n));
console.log(F19.pow(2n,19n*13n));
console.log(F19.pow(2n,19n*19n*13n));

// Computational Diffie-Hellman Problem
// a = g**x , b= g**y, c = g**(xy)
// CDH(g,a,b) -> c ?


const Z5 = Z(5n, 277n);

//console.log(Z5.permutation(64n));

const Zm107 = Z(5n, 2n**107n - 1n);
console.log(Zm107.permutation(164n));
```



## SEC Format

为了便于使用和计算, 有人提出了SEC格式, 其实就是如何表示椭圆曲线上的坐标. 首先我们来看非压缩格式:

> 0x04 <x><y> (总计65字节)

有必要同时保存x和 y么?

我们知道椭圆曲线上任意一点x可能对应两个y, 有没有可能只记录x, 那么给定x, 我们就可以计算出y, 或者p-y

最多只有这两种取值.

那么我们如何表达是哪个呢? 我们知道p是质数, 必然是奇数, 如果y是奇数, p-y就是偶数, 如果y是偶数那么p-y就是奇数. 于是:

>0x02: 偶数y
>
>0x03: 奇数y
>
><x>

于是我们可以只记录x坐标, 节省出一个坐标的空间(32bytes).



## DER Signatures (Distinguished Encoding Rules)

还记得我们计算出来的签名(r,s)么? 



## 为什么使用Base58?

因为Base64中的一些字符看起来很像,比如0和o, I和l等等, 为了增强可读性

> Base58 = 10个数字 + 26个小写英文字母 + 26个大写英文字母 - 大写O - 数字0 - 大写I - 小写l



### ripemd160 hash



## 比特币地址如何组成

> 1. a = SEC(P)
> 2. Hash160(X) = RIPENMD160(SHA256(X))
> 3. b = Hash160(a)
> 4. c = [0x00 | 0x06] + b
> 5. checksum = Hash256(c)[0,4]  # 4bytes的checksum
> 6. address = Base58(c+checksum)



## WIF Format (Wallet Import Format)

对于256bit的私钥我们并不需要经常传输, 但有时会做私钥迁移.为此设计了WIF格式.





>1. prefix = [0x80 | 0xef] # 主网 | 测试网
>2. a = prefix + bigendian(k) + [0x01] 如果使用SEC格式的的公钥添加0x01后缀
>3. checksum = SHA256(a)[0,4]
>4. WIF = Base58(b + checksum)

## 参考

- https://www.youtube.com/watch?v=XmygBPb7DPM
- https://www.youtube.com/watch?v=F3zzNa42-tQ
- https://trustica.cz/en/2018/04/26/elliptic-curves-prime-order-curves/
- https://www.esat.kuleuven.be/cosic/publications/talk-78.pdf 
- https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/