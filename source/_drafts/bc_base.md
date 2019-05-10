

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

当你打开网页很可能就已经正在使用椭圆曲线加密了, https协议



> $y^2 = x^3+ax+b$

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



使用openssl生成secp256k1

```zsh
$ openssl ecparam -name secp256k1 -out secp256k1.pem
$ openssl ecparam -in secp256k1.pem -text -param_enc  explicit -noout
# openssl ecparam -name secp256k1 | openssl ecparam -text -param_enc  explicit -noout
Field Type: prime-field
Prime:
    00:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:
    ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:fe:ff:
    ff:fc:2f
A:    0
B:    7 (0x7)
Generator (uncompressed):
    04:79:be:66:7e:f9:dc:bb:ac:55:a0:62:95:ce:87:
    0b:07:02:9b:fc:db:2d:ce:28:d9:59:f2:81:5b:16:
    f8:17:98:48:3a:da:77:26:a3:c4:65:5d:a4:fb:fc:
    0e:11:08:a8:fd:17:b4:48:a6:85:54:19:9c:47:d0:
    8f:fb:10:d4:b8
Order: 
    00:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:ff:
    ff:fe:ba:ae:dc:e6:af:48:a0:3b:bf:d2:5e:8c:d0:
    36:41:41
Cofactor:  1 (0x1)
```



```zsh
$ openssl ecparam -list_curves
secp256k1 : SECG curve over a 256 bit prime field
# 1. 首先我们需要生成ecparm, 然而这个东西其实没有什么用
$ openssl ecparam -name secp256k1 -out ecparam.pem
# 2. 下面我们来正式生成secp256k1
$ openssl ecparam -name secp256k1 -genkey 
-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIPXtSAxoBvKnZkZVT7OUn4qMEoo0zv26sf27C8hukJIkoAcGBSuBBAAK
oUQDQgAE+FcrUh50oPr/m2sOy1kiwM8U/4UttGqEwcmaQfLZeXs20SzRF/g9D4WK
J2MTrjtwtujnC3n77Nys/4MwcZ/Szg==
-----END EC PRIVATE KEY-----
d9:84:47:92:ae:7b:e7:58:31:e3:c3:b9:c5:cb:e2:88:3a:dc:8b:7f
# 上面输出的是.pem文件格式, 有两部分, 上面是ecparam, 目前没用, 加上-noout把它干掉, 下面是私钥
$ openssl ecparam -name secp256k1 -genkey -noout -out private.pem
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIOf4xq0sgD0JtVO90Cck1gu/H5/zjGrj9D7duApYPmLaoAcGBSuBBAAK
oUQDQgAEYDMQnSU9ZKDTs1tGMtr3AkWhKtAAbOCZnCZDiin4l+BJw1goeHmv0K61
QQoivkovD61zQDnmMtQWs6YXxJJqtA==
-----END EC PRIVATE KEY-----

# 我们来解读一下这个pem文件
$ openssl ec -in private.pem  -inform PEM -text -noout  
read EC key
Private-Key: (256 bit)
priv:
    e7:f8:c6:ad:2c:80:3d:09:b5:53:bd:d0:27:24:d6:
    0b:bf:1f:9f:f3:8c:6a:e3:f4:3e:dd:b8:0a:58:3e:
    62:da
pub:
    04:60:33:10:9d:25:3d:64:a0:d3:b3:5b:46:32:da:
    f7:02:45:a1:2a:d0:00:6c:e0:99:9c:26:43:8a:29:
    f8:97:e0:49:c3:58:28:78:79:af:d0:ae:b5:41:0a:
    22:be:4a:2f:0f:ad:73:40:39:e6:32:d4:16:b3:a6:
    17:c4:92:6a:b4
ASN1 OID: secp256k1
# 1. 
#   priv: 这个是私钥
#   pub: 是公钥, 04开头可知是非压缩SEC格式

# 2. private.pem 文件里保存了什么?
# -----BEGIN EC PRIVATE KEY-----
# Base64编码之后的DER格式
# ----END EC PRIVATE KEY----- 
# 来验证一下:
# 注意: 我们在dump的时候一定要用big-endian的方式, 要不然你就对不上了
$ echo MHQCAQEEIOf4xq0sgD0JtVO90Cck1gu/H5/zjGrj9D7duApYPmLaoAcGBSuBBAAKoUQDQgAEYDMQnSU9ZKDTs1tGMtr3AkWhKtAAbOCZnCZDiin4l+BJw1goeHmv0K61QQoivkovD61zQDnmMtQWs6YXxJJqtA== | base64 -d | od -An -x --endian big 
 3074 0201 0104 20e7 f8c6 ad2c 803d 09b5
 53bd d027 24d6 0bbf 1f9f f38c 6ae3 f43e
 ddb8 0a58 3e62 daa0 0706 052b 8104 000a
 a144 0342 0004 6033 109d 253d 64a0 d3b3
 5b46 32da f702 45a1 2ad0 006c e099 9c26
 438a 29f8 97e0 49c3 5828 7879 afd0 aeb5
 410a 22be 4a2f 0fad 7340 39e6 32d4 16b3
 a617 c492 6ab4


$ openssl ec -in private.pem -outform DER > private.der
$ cat private.der | od -An -x --endian big 
 3074 0201 0104 20e7 f8c6 ad2c 803d 09b5
 53bd d027 24d6 0bbf 1f9f f38c 6ae3 f43e
 ddb8 0a58 3e62 daa0 0706 052b 8104 000a
 a144 0342 0004 6033 109d 253d 64a0 d3b3
 5b46 32da f702 45a1 2ad0 006c e099 9c26
 438a 29f8 97e0 49c3 5828 7879 afd0 aeb5
 410a 22be 4a2f 0fad 7340 39e6 32d4 16b3
 a617 c492 6ab4

 
# 我们来研究下der里面有什么信息, 下面这俩命令都是一个效果
$ openssl asn1parse -in private.pem -inform pem
$ openssl asn1parse -in private.der -inform der
    0:d=0  hl=2 l= 116 cons: SEQUENCE          
    2:d=1  hl=2 l=   1 prim: INTEGER           :01
    5:d=1  hl=2 l=  32 prim: OCTET STRING      [HEX DUMP]:E7F8C6AD2C803D09B553BDD02724D60BBF1F9FF38C6AE3F43EDDB80A583E62DA
   39:d=1  hl=2 l=   7 cons: cont [ 0 ]        
   41:d=2  hl=2 l=   5 prim: OBJECT            :secp256k1
   48:d=1  hl=2 l=  68 cons: cont [ 1 ]        
   50:d=2  hl=2 l=  66 prim: BIT STRING     
# OCTEC STRING是啥? 看起来像是私钥, 就是256bit的随机数, 来验证一下
$ echo -n E7F8C6AD2C803D09B553BDD02724D60BBF1F9FF38C6AE3F43EDDB80A583E62DA | wc -c
64
# 64 x 4 = 256 没错
# 解释一下这几个信息, 数据长度单位都是字节(16bit或2bytes)
#  d  = 数据的层级
#  hl = header length
# 	l  = 数据长度
# 下标从0开始, 0 + 2 + 5 = 7 数据从第8个字节开始, 往后取32字节
$ cat private.der | tail -c +8 | head -c 32  > r
$ cat r | od -An -x --endian big 
 e7f8 c6ad 2c80 3d09 b553 bdd0 2724 d60b
 bf1f 9ff3 8c6a e3f4 3edd b80a 583e 62da

# 用私钥算公钥, 保存成DER格式, 默认公钥采用非压缩格式存储
$ openssl ec -in private.pem -pubout -outform DER | od -An -x --endian big
 3056 3010 0607 2a86 48ce 3d02 0106 052b
 8104 000a 0342 0004 6033 109d 253d 64a0
 d3b3 5b46 32da f702 45a1 2ad0 006c e099
 9c26 438a 29f8 97e0 49c3 5828 7879 afd0
 aeb5 410a 22be 4a2f 0fad 7340 39e6 32d4
 16b3 a617 c492 6ab4
 
# 非压缩格式的SEC和压缩格式的SEC都是可以的, 后面会区分处理生成不同的地址
$ openssl ec -in private.pem -pubout -conv_form compressed -outform DER > public_key.der
$ cat public_key.der | od -An -x --endian big
 3036 3010 0607 2a86 48ce 3d02 0106 052b
 8104 000a 0322 0002 6033 109d 253d 64a0
 d3b3 5b46 32da f702 45a1 2ad0 006c e099
 9c26 438a 29f8 97e0
# 用openssl来进一步确认
$ openssl pkey -pubin -in public_key.der -inform der -text -noout
Public-Key: (256 bit)
pub:
    02:60:33:10:9d:25:3d:64:a0:d3:b3:5b:46:32:da:
    f7:02:45:a1:2a:d0:00:6c:e0:99:9c:26:43:8a:29:
    f8:97:e0
ASN1 OID: secp256k1


# 把SEC格式的公钥从DER文件中取出来
$ cat public_key.der | tail -c +24 > public_key.sec
$ cat public_key.sec | od -An -x --endian big
 0260 3310 9d25 3d64 a0d3 b35b 4632 daf7
 0245 a12a d000 6ce0 999c 2643 8a29 f897
 e000
$ cat public_key.sec | wc -c
33 # 33bytes正是压缩SEC的长度, 如果产生的点过小也可能是32bytes


# 解决了椭圆曲线的问题, 我们可以正式开搞
# https://en.bitcoin.it/wiki/Address
# 1. 计算SHA256(PublicKey)
$ sha256sum public_key.sec 
f376a6732f1d5d433301183fd03a4f09a1d65bb4059530dab7ea1b8c2f455475  public_key.sec

# 2. ripemd160
# 注意: echo -n禁止加换行字符, 否则算出来的hash就不对了
$ echo -n f376a6732f1d5d433301183fd03a4f09a1d65bb4059530dab7ea1b8c2f455475  | openssl rmd160 
(stdin)= 9f2d1416d8169cb356cfa6bdda2e626b88bab1d2

# 3. 


# 压缩格式的地址
# 1. sha256
$ sha256sum  public_key_uncompressed.sec
f376a6732f1d5d433301183fd03a4f09a1d65bb4059530dab7ea1b8c2f455475  public_key_uncompressed.sec
# 2. ripemd160
# 注意openssl需要接收的是raw数据, 所以我们需要用xxd把hex string逆成原始字节串
$ echo -n f376a6732f1d5d433301183fd03a4f09a1d65bb4059530dab7ea1b8c2f455475 | xxd -r -p | openssl rmd160
(stdin)= be01a788e9ff48b514095bad08e86957e7a7f3bd
# 3. 计算checksum, 
# 0x00 be01a788e9ff48b514095bad08e86957e7a7f3bd
$ echo - n 00be01a788e9ff48b514095bad08e86957e7a7f3bd | xxd -r -p | openssl sha256
(stdin)= ad23aea0ea8b7355b60def91859b41e992d032d5969d37a640a5f14e55b38688


```
```
 3074 0201 0104 20[e7 f8c6 ad2c 803d 09b5
 53bd d027 24d6 0bbf 1f9f f38c 6ae3 f43e
 ddb8 0a58 3e62 da]a0 0706 052b 8104 000a
 a144 0342 0004 6033 109d 253d 64a0 d3b3
 5b46 32da f702 45a1 2ad0 006c e099 9c26
 438a 29f8 97e0 49c3 5828 7879 afd0 aeb5
 410a 22be 4a2f 0fad 7340 39e6 32d4 16b3
 a617 c492 6ab4
```



我们用bitcoin-tool验证一下计算过程是不是正确:

```
$ bitcoin-tool --input-type private-key  --input-format raw  --input-file r --public-key-compression uncompressed  --network bitcoin --output-type all 
address.hex:00be01a788e9ff48b514095bad08e86957e7a7f3bd
address.base58:13eXsokAbxaQ4eisEgjzYm7DfUZxp
address.base58check:1JKfNQBuHb39SjQW1GXHv9AXixYxNM5jxv
address-checksum.hex:00be01a788e9ff48b514095bad08e86957e7a7f3bd19ec46d3
address-checksum.base58:1JKfNQBuHb39SjQW1GXHv9AXixYxNM5jxv
public-key-ripemd160.hex:be01a788e9ff48b514095bad08e86957e7a7f3bd
public-key-ripemd160.base58:3eXsokAbxaQ4eisEgjzYm7DfUZxp
public-key-ripemd160.base58check:JKfNQBuHb39SjQW1GXHv9AXixYxP3a8pH
public-key-sha256.hex:f376a6732f1d5d433301183fd03a4f09a1d65bb4059530dab7ea1b8c2f455475
public-key-sha256.base58:HPP1K8ekJRSrAw8XfMXawcfyENbVNzH69JgsvdhYBn1a
public-key-sha256.base58check:2rDwyHTqjPXuMFurKyvKXMAKR1hn7WZrPCD9kbY5HZxDRVbZqw
public-key.hex:046033109d253d64a0d3b35b4632daf70245a12ad0006ce0999c26438a29f897e049c358287879afd0aeb5410a22be4a2f0fad734039e632d416b3a617c4926ab4
public-key.base58:PPzQbeScDAXdJF9fdgmuvs4kdZZi8cNezM2zptbupzzH7wxsfyL9eTgUWmqe9QTuT1UvrxtWoXMRdGheLj8TFa4s
public-key.base58check:3XZ8c7TNTbcZ3JUGg2ax8j7ygWM8tieco7Gm5s5cyXCWv9NzJuc3vNd4DqTyMgnVSTuBvrhc1PYazEfvkp2LBHCzAi7P22
private-key-wif.hex:80e7f8c6ad2c803d09b553bdd02724d60bbf1f9ff38c6ae3f43eddb80a583e62da
private-key-wif.base58:fJ14dxi4AptS14RMExFQTGtNiTGYLfdAmW6DE9FGJJEws
private-key-wif.base58check:5KaSzLMrC9hejqCNZbHHkwD2si4XcGdtCrXqFVoYUMWWQhFafaV
private-key.hex:e7f8c6ad2c803d09b553bdd02724d60bbf1f9ff38c6ae3f43eddb80a583e62da
private-key.base58:GcXCRjqPCr3e3s7msngWWdhxU4AyJz8fba7byLc4CKSZ
private-key.base58check:2mAQeum7h8fUpp6XPqPftTuf7EcQ8moV3U42KkoEUesF2XKi6e

```



```zsh
# 如果你没有od命令, 用openssl enc -base64 -d -A


#---------------------------------------[ 私钥 | (7*8 + 3) * 16 = ]
 7430 0102 0401 e720 c6f8 2cad 3d80 b509
 bd53 27d0 d624 bf0b 9f1f 8cf3 e36a 3ef4
 b8dd 580a 623e a0da 0607 2b05 0481 0a00
 44a1 4203 0400 3360 9d10 3d25 a064 b3d3
 465b da32 02f7 a145 d02a 6c00 99e0 269c
 8a43 f829 e097 c349 2858 7978 d0af b5ae
 0a41 be22 2f4a ad0f 4073 e639 d432 b316
 17a6 92c4 b46a

#--------------------------------------[ 公钥  ] 
 5630 1030 0706 862a ce48 023d 0601 2b05
 0481 0a00 4203 0400 3360 9d10 3d25 a064
 b3d3 465b da32 02f7 a145 d02a 6c00 99e0
 269c 8a43 f829 e097 c349 2858 7978 d0af
 b5ae 0a41 be22 2f4a ad0f 4073 e639 d432
 b316 17a6 92c4 b46a
```







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

Base58算法把输入看成一个大数字, 然后不断的用这个数字去对58求模, 每次计算得到余数从字典中替换. 直到除尽. Base58需要58个字符的字典.

在将Base58之前, 我们先来看一个问题, 假如有一个数字100, 我们可以写成以下几种形式:

```
100     十进制
 64     十六进制
144     八进制
10201   三进制
1100100 二进制
```

Base58有一定的压缩功能, 可以使数据变得更短,

常见的字典包括:

```
BTC = 123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
XRP = rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz
Flickr = 123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
```

```
#!
```



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

```
$ ./bitcoin-tool \
  --input-file <(echo -n 'Hi guys!' | openssl dgst -sha256 -binary) \
  --input-format raw \
  --input-type private-key \
  --network bitcoin \
  --output-type private-key-wif \
  --output-format base58check \
  --public-key-compression uncompressed
```



## 参考



- https://www.youtube.com/watch?v=XmygBPb7DPM
- https://www.youtube.com/watch?v=F3zzNa42-tQ
- https://trustica.cz/en/2018/04/26/elliptic-curves-prime-order-curves/
- https://www.esat.kuleuven.be/cosic/publications/talk-78.pdf 
- https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/
- https://en.bitcoin.it/wiki/Address
- 地址生成工具: https://gobittest.appspot.com/Address
- Base58: https://learnmeabitcoin.com/glossary/base58 