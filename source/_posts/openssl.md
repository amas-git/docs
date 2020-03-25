### OpenSSL



```bash
$ openssl version
OpenSSL 1.1.1d  10 Sep 2019

# CHECK COMPLETE VERSION
$ openssl version -a
OpenSSL 1.1.1d  10 Sep 2019
built on: Wed Nov 13 16:09:29 2019 UTC
platform: linux-x86_64
options:  bn(64,64) rc4(16x,int) des(int) idea(int) blowfish(ptr) 
compiler: gcc -fPIC -pthread -m64 -Wa,--noexecstack -march=x86-64 -mtune=generic -O2 -pipe -fno-plt -Wa,--noexecstack -D_FORTIFY_SOURCE=2 -march=x86-64 -mtune=generic -O2 -pipe -fno-plt -Wl,-O1,--sort-common,--as-needed,-z,relro,-z,now -DOPENSSL_USE_NODELETE -DL_ENDIAN -DOPENSSL_PIC -DOPENSSL_CPUID_OBJ -DOPENSSL_IA32_SSE2 -DOPENSSL_BN_ASM_MONT -DOPENSSL_BN_ASM_MONT5 -DOPENSSL_BN_ASM_GF2m -DSHA1_ASM -DSHA256_ASM -DSHA512_ASM -DKECCAK1600_ASM -DRC4_ASM -DMD5_ASM -DVPAES_ASM -DGHASH_ASM -DECP_NISTZ256_ASM -DX25519_ASM -DPOLY1305_ASM -DNDEBUG -D_FORTIFY_SOURCE=2
OPENSSLDIR: "/etc/ssl" # 注意这个目录
ENGINESDIR: "/usr/lib/engines-1.1"
Seeding source: os-specific

$ tree -d /etc/ssl
/etc/ssl
├── certs     # 保存证书
│   ├── edk2
│   └── java
├── misc      # 一些辅助脚本，比如自定义证书的脚本
└── private
```



openssl的功能

```bash
$ openssl help
Standard commands
asn1parse         ca                ciphers           cms               
crl               crl2pkcs7         dgst              dhparam           
dsa               dsaparam          ec                ecparam           
enc               engine            errstr            gendsa            
genpkey           genrsa            help              list              
nseq              ocsp              passwd            pkcs12            
pkcs7             pkcs8             pkey              pkeyparam         
pkeyutl           prime             rand              rehash            
req               rsa               rsautl            s_client          
s_server          s_time            sess_id           smime             
speed             spkac             srp               storeutl          
ts                verify            version           x509              

Message Digest commands (see the `dgst' command for more details)
blake2b512        blake2s256        gost              md4               
md5               mdc2              rmd160            sha1              
sha224            sha256            sha3-224          sha3-256          
sha3-384          sha3-512          sha384            sha512            
sha512-224        sha512-256        shake128          shake256          
sm3               

Cipher commands (see the `enc' command for more details)
aes-128-cbc       aes-128-ecb       aes-192-cbc       aes-192-ecb       
aes-256-cbc       aes-256-ecb       aria-128-cbc      aria-128-cfb      
aria-128-cfb1     aria-128-cfb8     aria-128-ctr      aria-128-ecb      
aria-128-ofb      aria-192-cbc      aria-192-cfb      aria-192-cfb1     
aria-192-cfb8     aria-192-ctr      aria-192-ecb      aria-192-ofb      
aria-256-cbc      aria-256-cfb      aria-256-cfb1     aria-256-cfb8     
aria-256-ctr      aria-256-ecb      aria-256-ofb      base64            
bf                bf-cbc            bf-cfb            bf-ecb            
bf-ofb            camellia-128-cbc  camellia-128-ecb  camellia-192-cbc  
camellia-192-ecb  camellia-256-cbc  camellia-256-ecb  cast              
cast-cbc          cast5-cbc         cast5-cfb         cast5-ecb         
cast5-ofb         des               des-cbc           des-cfb           
des-ecb           des-ede           des-ede-cbc       des-ede-cfb       
des-ede-ofb       des-ede3          des-ede3-cbc      des-ede3-cfb      
des-ede3-ofb      des-ofb           des3              desx              
idea              idea-cbc          idea-cfb          idea-ecb          
idea-ofb          rc2               rc2-40-cbc        rc2-64-cbc        
rc2-cbc           rc2-cfb           rc2-ecb           rc2-ofb           
rc4               rc4-40            seed              seed-cbc          
seed-cfb          seed-ecb          seed-ofb          sm4-cbc           
sm4-cfb           sm4-ctr           sm4-ecb           sm4-ofb    
```



## 构建TrustStore

TrustStore也就是RootCA, OpenSSL本身并不携带任何RootCA, 通常使用操作系统内置的。缺点是有些操作系统的RootCA更新比较慢。你可以使用Mozilla提供的最新RootCA:

- https://curl.haxx.se/docs/caextract.html (PEM格式)
- https://hg.mozilla.org/mozilla-central/raw-file/tip/security/nss/lib/ckfw/builtins/certdata.txt (Mozilla专用格式)
  - 有两个项目可以将Moziila的RootCA转成PEM格式
    - https://github.com/agl/extract-nss-root-certs
    - https://raw.github.com/bagder/curl/master/lib/mk-ca-bundle.pl

## 私钥(Key)和证书管理

使用SSL/TLS,需要做三件事情:

1. 生成一个足够强壮的钥匙(Key)
2. 生成CSR(Certifacate Signing Request), 发给CA
3. 安装CA提供的证书

### 1. 生成钥匙

|                      |                                                |      |
| -------------------- | ---------------------------------------------- | ---- |
| 算法(Key algorithm)  | RSA,DSA,ECDSA                                  | 必须 |
| 长度(Key size)       | RSA大于2048, ECDSA大于224                      | 必须 |
| 密码短语(Passphrase) | 设置钥匙的密码，如果被别人拷走了也没法直接使用 | 可选 |

```bash
# 1. 制作不带密码短语的RSA钥匙
$ openssl genrsa -out server.key 2048

# 2. 制作用AES256加密算法保护的RSA钥匙
$ openssl genrsa -aes256 -out server.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
....................................................+++++
.......................+++++
e is 65537 (0x010001) # 注意: RSA的e值默认是65537,之所以使用这个数字是因为安全和性能上的平衡，你也可以使用3获得更快的速度
Enter pass phrase for server.key: ******
Enter pass phrase for server.key: ******

$ cat server.key     # PEM格式
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,FC56592FFA758F8954794A338D2A86D0

yc4dKg7aU5lhGFBRKg0Uqcc1OuLdMhB7fOXF5iYfHjFov3HjltnQGReQvciVuPC4
qaV5h24/GMpyiIM8Q+nhGpj/fvJMzhGHrYoBIA0Uku2H+WIRvkSS4zKGxV6lsyTh
ak3Z4yVns3t3/PN5pHrHompZMPPBiLQUMv9mEGBaXIFvAW77bvAOBxA205hwlfxU
m3bRXjnd7cWrkYWSeic3OUjWgAxp2T//QY1g3kzRFfriVyF26owohGBbFqz3eDMt
R5F5ZRsJzCGjcIalHUaL/Yypq94BG+FSSySDWdcYwssdLJSptcCgM+ez+liNaTPw
MchdJy3Fv6FH2I/dKq67Teht1UQj5TkXwl4xVCPpFWkIAhMoKxoDLyACdEQXc4/I
/ePgd4QLimmi0yEMzZ/N80n7r3EVQOdBL7kD6tZHuxj1qiUAztPFf1G5oKTd1Ubj
joFGeSCaX4sdanhC6vFnhOmxbZoy5wsvkcQ6ZlvJJMtDpFEjHSsRWlplyRzPdxK9
dh0rS+/+UICiQE+6IQDpZfpJMsXLTLW4g5RUcXBle9+DEeJTjedeHIkqC1BrGeXN
Qu7WpqP14AAfvXUlWrwvn96SK1aIe+GHiVLzK/dxUGoQqqOuAPgnX4D90EKKhzQV
dXk5yDpuPTJIuCVdKE7/pQEupDtTq9YFnJLbKKynGvZPJuWoLYgH9PyIFKiUhGro
RLpmQ1IKTOPtkBD+vdzjUJ/9AH+19AIoEdx8IuV5gHM8tzf9EGIXsBBlpt30leHo
A1l374DjQDX7b+iIxv+wP8s6//p8cCfkOvC4838H/MyOcvPQEUbzkhg5kx5rxUOj
FjauFl4zyfTVzF9K20hp6LXp7g8GYqbEg1HEtYNEHFRMwmcwxZ2wzmStrY5YzXwx
F3N7Tkw8KshhFRotIuyQLJroe/GYe6hnhyhehacXIrus3sLX5ogUZWKemS8P/Ynz
QGiEZFChQqYP50/vONjPDIDHblP08PWz1dXBFWaiQdvmnjUMiZlX1fmosQZU3bR4
MortyMXL/JJId82/dHYznBuS6B90m5c+4I0xUw/ElABa7e/aVy+RsuZSLFuwcDCD
AthF0Gm0vu0/sQVLDUt+ggOhFy8eGYRnRYUioaKiqRfyibgT+GQ9Tm+9huTRNdrp
Pc7Z1bRsz99krJ7/wjxHI/LuIvtNcxfSF5O+nODHBuKNteiFZjZ+GE3juhXdwcBL
F1lFb7VN2iicqj3VK8CVxwIJgFh/PzXrG4SyQx0AzExkLLrRH3IyUHP/xb1x6ym/
IrBJrdT9bcoemy5l5GlkYfhMIf4JmlO9DMlclmyoNMyIZ97hmIYwX47EsloDU9Q0
fXtEYyybqaGGYIg8k98jveAXr6m5nD7FpwoavkjIyhbxMa8nbDqexNicmLlNBd2d
K6/nB6ndyH0YsGSUxFT3mUwWBmvZsCWpQVj32U4D8ux476C9AWq25OU8pZQzkp8g
MNMpa1NrldcKPOa1Bu0keuHGnCSZOK8yOW5k4nHkCpitdlbOaXNH4bwEjBeu/gxD
PaEYpdVclilzk1BePQNrdk2V0NCKrLIdwjNE1Xg5lx6L2owtQXPRm+O0HeQZKf9g
-----END RSA PRIVATE KEY----

# openssl可以解读
$ openssl rsa -text -in server.key
Enter pass phrase for server.key: ****** # 输入你的密码短语
RSA Private-Key: (2048 bit, 2 primes)
modulus: # RSA模数
    00:ad:7b:d0:4c:48:6c:af:33:69:d9:da:32:2f:68:
    2a:b6:9b:9f:a5:4a:54:5d:29:2a:55:fa:ea:b6:d8:
    50:47:8b:d9:22:b0:dc:67:ce:49:a0:56:f1:98:f3:
    46:6d:8c:eb:59:37:ab:d4:0e:66:30:87:de:c4:fc:
    84:4c:d2:ac:ee:fc:73:4e:80:1e:4c:3b:77:30:42:
    48:76:24:39:44:f8:72:e7:ad:0e:cb:2e:e9:6d:9e:
    4f:2a:06:03:5a:fd:c3:9c:13:57:ff:a2:a6:df:2c:
    88:e1:6e:4d:b6:5a:98:24:f1:e6:77:a2:29:69:73:
    b1:69:5b:48:55:b9:c4:ff:b7:c1:c8:92:b4:75:19:
    dd:c2:80:82:05:c3:e6:48:ee:15:ab:eb:f1:53:d2:
    9b:b1:33:a2:73:36:ca:18:cb:5f:64:f7:ec:a8:52:
    e3:99:68:59:d4:43:4a:c1:71:26:28:82:35:2f:ff:
    87:6c:73:83:77:d3:49:f3:e5:a2:c2:ba:1a:80:f2:
    18:22:ef:e2:ef:c1:ab:94:06:fa:f2:64:01:1a:3a:
    32:94:91:ca:09:5b:59:58:05:bf:de:0a:ba:bd:53:
    3a:99:6d:ac:62:e4:dc:32:d0:38:46:a8:41:75:be:
    7c:1d:c8:5b:73:b4:a2:ae:12:dc:b8:ef:75:92:e3:
    ba:59
publicExponent: 65537 (0x10001)
privateExponent:
    00:90:8f:b2:e4:4e:19:9a:e8:f9:d4:9a:58:e5:5e:
    24:f1:a3:be:a5:8a:c9:c0:13:b4:7e:8f:27:15:14:
    2d:d9:60:b4:a0:8e:c6:2b:0e:20:16:27:3d:0d:59:
    f3:1f:08:a5:78:e3:c6:20:5f:9b:51:e7:76:7f:a9:
    78:49:57:e8:d9:00:ae:c9:04:43:ba:fc:76:76:55:
    55:72:74:fe:ef:f5:24:32:df:e7:8a:f2:5d:7b:85:
    c6:ab:da:f4:e2:d6:c1:30:86:81:f6:b8:3f:db:8a:
    8d:c9:64:14:07:d7:78:1e:99:20:96:22:1f:e6:0a:
    8d:14:a1:07:26:c0:35:63:6c:0a:07:99:de:39:c9:
    cf:77:8e:ae:93:97:f9:dc:47:0b:ca:72:9b:17:24:
    a2:87:ab:43:9d:c4:3b:0a:b0:2a:1c:70:ec:c8:fd:
    63:86:19:4f:e4:46:a9:88:2e:56:7d:40:d2:45:7f:
    9d:fe:b1:69:c8:dd:5f:e7:fa:d7:b2:c4:eb:89:9e:
    bf:c0:54:e1:79:74:12:27:54:7f:b2:a4:08:ce:84:
    fb:55:ba:e1:46:9b:18:c8:8c:e3:c6:7d:41:dc:64:
    bd:83:8d:de:a2:9b:73:c8:58:30:ae:c6:8e:b2:3e:
    1a:8a:f2:8b:4c:9d:99:6c:69:38:a9:5d:09:14:d7:
    b2:71
prime1: # RSA大质数1
    00:d8:c3:c2:29:a7:cc:ef:f6:b5:64:9c:99:25:eb:
    6a:dc:6b:26:6e:46:9c:c3:2e:ab:ed:f4:3f:65:95:
    30:76:f9:38:f4:13:cc:bb:cd:de:db:99:50:b3:e5:
    63:42:d1:c2:2e:bf:26:da:eb:06:89:20:7d:bf:76:
    cf:5c:99:37:2d:52:b3:19:08:39:12:ca:0a:f3:ff:
    0c:36:87:96:03:56:c6:f3:2a:b5:1d:71:92:da:33:
    45:ca:75:1d:ab:1b:fc:f6:6b:d4:22:39:c6:be:51:
    ff:44:10:ab:89:dd:1c:cc:f5:24:24:61:62:10:ba:
    e8:44:a2:b8:92:e8:ad:5b:0d
prime2: # RSA大质数2
    00:cc:e2:89:e3:d6:49:62:40:39:cf:42:c8:11:8a:
    7f:07:20:07:e0:bb:56:14:74:6b:cc:3f:96:e2:16:
    ad:11:a9:11:ff:9d:e7:3f:d7:37:04:8a:1c:63:f2:
    7c:4a:24:5d:80:55:65:4e:a6:d3:b2:7d:d0:92:1b:
    f1:e9:c0:2a:a1:d8:c6:1e:22:f4:f1:dc:14:db:7b:
    b0:e8:60:3a:ff:38:00:60:54:f7:c9:ae:b7:e6:7e:
    b4:22:7d:5d:3e:d6:45:b9:4e:ed:71:a6:ad:62:93:
    cf:be:90:a4:38:76:47:f2:d5:b8:67:65:75:99:e2:
    d7:3d:e3:53:7e:0b:42:19:7d
exponent1:
    49:64:ce:e3:27:cb:be:1c:3c:82:da:7b:08:59:d3:
    8e:da:40:e2:e9:c9:be:54:99:26:32:a3:1c:94:0d:
    1a:db:7b:ab:38:e1:03:5d:cb:6d:73:55:dd:f3:77:
    4e:72:93:5b:1c:a1:dd:51:e2:9e:9f:7f:b4:4a:58:
    1f:b4:48:f8:71:9d:ee:85:d8:3d:42:67:bf:01:c6:
    72:d8:29:b4:eb:b6:e5:32:ba:ac:43:7e:43:9f:44:
    ac:2e:47:63:5e:50:a2:67:14:26:9a:85:6e:7f:78:
    fc:e4:e0:10:07:eb:ee:81:ba:41:0f:30:13:16:15:
    fa:d7:55:c3:78:40:ea:29
exponent2:
    20:6d:b6:cf:86:0a:45:6f:ce:f6:9c:26:58:88:68:
    44:b3:70:2b:c4:db:02:0f:cf:44:1b:c8:80:ee:7f:
    e4:2c:b3:79:96:ff:94:1e:37:4a:13:a6:1c:b7:b3:
    ae:74:85:0c:1b:f8:15:f3:d7:cc:07:97:ec:98:59:
    b8:da:be:a4:b1:4f:e2:53:3c:1b:cf:ee:c9:32:91:
    b4:a6:0e:90:78:c7:ae:77:a4:64:9a:af:e6:de:a5:
    1a:54:67:5e:db:c1:5b:6c:3f:ae:de:67:d1:13:7e:
    2d:36:6d:97:b6:38:fb:19:92:bf:62:d2:b3:51:b1:
    29:cd:82:58:8f:e5:9d:6d
coefficient:
    00:95:29:44:0b:94:e0:aa:0e:42:59:61:e1:b0:ad:
    b5:f9:ee:0b:ad:8e:70:e3:7f:d0:55:43:cd:ae:cf:
    4d:f8:1e:ad:c4:3c:fd:53:4c:0d:ce:f5:4f:70:37:
    20:80:5d:eb:c4:2b:1f:88:ab:49:9f:8b:9e:84:f5:
    67:16:42:77:44:40:66:0b:85:0e:d5:a9:05:82:5b:
    45:fd:d9:80:2f:b1:02:7c:1f:bc:b2:c2:53:48:cb:
    45:e7:f2:fa:33:c5:50:47:d8:ea:c9:0d:d3:f0:c4:
    4c:7e:ac:98:48:d7:9e:39:97:7a:d6:25:fa:3c:67:
    24:2b:25:76:b1:cd:b6:23:23
writing RSA key
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEArXvQTEhsrzNp2doyL2gqtpufpUpUXSkqVfrqtthQR4vZIrDc
Z85JoFbxmPNGbYzrWTer1A5mMIfexPyETNKs7vxzToAeTDt3MEJIdiQ5RPhy560O
yy7pbZ5PKgYDWv3DnBNX/6Km3yyI4W5NtlqYJPHmd6IpaXOxaVtIVbnE/7fByJK0
dRndwoCCBcPmSO4Vq+vxU9KbsTOiczbKGMtfZPfsqFLjmWhZ1ENKwXEmKII1L/+H
bHODd9NJ8+WiwroagPIYIu/i78GrlAb68mQBGjoylJHKCVtZWAW/3gq6vVM6mW2s
YuTcMtA4RqhBdb58Hchbc7SirhLcuO91kuO6WQIDAQABAoIBAQCQj7LkThma6PnU
mljlXiTxo76lisnAE7R+jycVFC3ZYLSgjsYrDiAWJz0NWfMfCKV448YgX5tR53Z/
qXhJV+jZAK7JBEO6/HZ2VVVydP7v9SQy3+eK8l17hcar2vTi1sEwhoH2uD/bio3J
ZBQH13gemSCWIh/mCo0UoQcmwDVjbAoHmd45yc93jq6Tl/ncRwvKcpsXJKKHq0Od
xDsKsCoccOzI/WOGGU/kRqmILlZ9QNJFf53+sWnI3V/n+teyxOuJnr/AVOF5dBIn
VH+ypAjOhPtVuuFGmxjIjOPGfUHcZL2Djd6im3PIWDCuxo6yPhqK8otMnZlsaTip
XQkU17JxAoGBANjDwimnzO/2tWScmSXratxrJm5GnMMuq+30P2WVMHb5OPQTzLvN
3tuZULPlY0LRwi6/JtrrBokgfb92z1yZNy1SsxkIORLKCvP/DDaHlgNWxvMqtR1x
ktozRcp1Hasb/PZr1CI5xr5R/0QQq4ndHMz1JCRhYhC66ESiuJLorVsNAoGBAMzi
iePWSWJAOc9CyBGKfwcgB+C7VhR0a8w/luIWrRGpEf+d5z/XNwSKHGPyfEokXYBV
ZU6m07J90JIb8enAKqHYxh4i9PHcFNt7sOhgOv84AGBU98mut+Z+tCJ9XT7WRblO
7XGmrWKTz76QpDh2R/LVuGdldZni1z3jU34LQhl9AoGASWTO4yfLvhw8gtp7CFnT
jtpA4unJvlSZJjKjHJQNGtt7qzjhA13LbXNV3fN3TnKTWxyh3VHinp9/tEpYH7RI
+HGd7oXYPUJnvwHGctgptOu25TK6rEN+Q59ErC5HY15QomcUJpqFbn94/OTgEAfr
7oG6QQ8wExYV+tdVw3hA6ikCgYAgbbbPhgpFb872nCZYiGhEs3ArxNsCD89EG8iA
7n/kLLN5lv+UHjdKE6Yct7OudIUMG/gV89fMB5fsmFm42r6ksU/iUzwbz+7JMpG0
pg6QeMeud6Rkmq/m3qUaVGde28FbbD+u3mfRE34tNm2Xtjj7GZK/YtKzUbEpzYJY
j+WdbQKBgQCVKUQLlOCqDkJZYeGwrbX57gutjnDjf9BVQ82uz034Hq3EPP1TTA3O
9U9wNyCAXevEKx+Iq0mfi56E9WcWQndEQGYLhQ7VqQWCW0X92YAvsQJ8H7yywlNI
y0Xn8vozxVBH2OrJDdPwxEx+rJhI1545l3rWJfo8ZyQrJXaxzbYjIw==
-----END RSA PRIVATE KEY-----

# 如果需要公钥
$ openssl rsa -in server.key -pubout -out server.pub
# 注意 ： 验证输出文件是一种美德，如果你忘记加-pubout,结果将导出私钥
$ cat server.pub
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArXvQTEhsrzNp2doyL2gq
tpufpUpUXSkqVfrqtthQR4vZIrDcZ85JoFbxmPNGbYzrWTer1A5mMIfexPyETNKs
7vxzToAeTDt3MEJIdiQ5RPhy560Oyy7pbZ5PKgYDWv3DnBNX/6Km3yyI4W5NtlqY
JPHmd6IpaXOxaVtIVbnE/7fByJK0dRndwoCCBcPmSO4Vq+vxU9KbsTOiczbKGMtf
ZPfsqFLjmWhZ1ENKwXEmKII1L/+HbHODd9NJ8+WiwroagPIYIu/i78GrlAb68mQB
GjoylJHKCVtZWAW/3gq6vVM6mW2sYuTcMtA4RqhBdb58Hchbc7SirhLcuO91kuO6
WQIDAQAB
-----END PUBLIC KEY-----


# 制作DSA钥匙，包含两步:
# 1. 生成私钥(随机数)
# 2. 制作DSA钥匙
$ openssl dsaparam -genkey 2048 | openssl dsa -out dsa.key -aes128

# 制作ECDSA要是，包含两步:
# 1. 选择椭圆曲线(EC)算法，生成对应的私钥
# 2. 制作ECDSA钥匙
$ openssl ecparam -genkey -name secp256r1 | openssl ec -out ec.key -aes128
```

### 2. 创建CSR

 ```bash
# 如果有些字段你想设置为空，输入'.'
# 如果你按回车，当前字段会使用默认值
# 
# 常用的信息
# - Country Name           :
# - Locality Name          :
# - Organization Name      :
# - Organization Unit Name :
# - Comman Name            : FQDN, 通常是服务器的全域名
# - Email Address          :
# Extra Attributes
# - Challenge Password     : (可选) 撤销证书时输入的密码，这个并不能增强证书的安全性，只是一个功能
# - Optional Company Name  :
$ openssl req -new -key server.key -out server.csr
Enter pass phrase for server.key: ****** # 输入密码短语
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:Beijing
Locality Name (eg, city) []:Beijing
Organization Name (eg, company) [Internet Widgits Pty Ltd]:amas
Organizational Unit Name (eg, section) []:.
Common Name (e.g. server FQDN or YOUR name) []:server.amas.org
Email Address []:amas@gmail.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:.
An optional company name []:.

# 解读，做好CSR之后再次确认信息是否正确是好习惯
$ openssl req -text -in service.csr -noout   
Certificate Request:
    Data:
        Version: 1 (0x0)
        Subject: C = CN, ST = Beijing, L = Beijing, O = amas, CN = server.amas.org, emailAddress = amas@gmail.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:ad:7b:d0:4c:48:6c:af:33:69:d9:da:32:2f:68:
                    2a:b6:9b:9f:a5:4a:54:5d:29:2a:55:fa:ea:b6:d8:
                    50:47:8b:d9:22:b0:dc:67:ce:49:a0:56:f1:98:f3:
                    46:6d:8c:eb:59:37:ab:d4:0e:66:30:87:de:c4:fc:
                    84:4c:d2:ac:ee:fc:73:4e:80:1e:4c:3b:77:30:42:
                    48:76:24:39:44:f8:72:e7:ad:0e:cb:2e:e9:6d:9e:
                    4f:2a:06:03:5a:fd:c3:9c:13:57:ff:a2:a6:df:2c:
                    88:e1:6e:4d:b6:5a:98:24:f1:e6:77:a2:29:69:73:
                    b1:69:5b:48:55:b9:c4:ff:b7:c1:c8:92:b4:75:19:
                    dd:c2:80:82:05:c3:e6:48:ee:15:ab:eb:f1:53:d2:
                    9b:b1:33:a2:73:36:ca:18:cb:5f:64:f7:ec:a8:52:
                    e3:99:68:59:d4:43:4a:c1:71:26:28:82:35:2f:ff:
                    87:6c:73:83:77:d3:49:f3:e5:a2:c2:ba:1a:80:f2:
                    18:22:ef:e2:ef:c1:ab:94:06:fa:f2:64:01:1a:3a:
                    32:94:91:ca:09:5b:59:58:05:bf:de:0a:ba:bd:53:
                    3a:99:6d:ac:62:e4:dc:32:d0:38:46:a8:41:75:be:
                    7c:1d:c8:5b:73:b4:a2:ae:12:dc:b8:ef:75:92:e3:
                    ba:59
                Exponent: 65537 (0x10001)
        Attributes:
            a0:00
    Signature Algorithm: sha256WithRSAEncryption
         2d:fb:83:5f:d7:e0:10:28:de:94:97:b9:c1:9b:b2:a5:df:89:
         71:6e:56:da:4f:05:77:fa:8e:72:9d:18:6a:43:8e:57:ec:c7:
         27:7b:22:91:48:3d:73:c7:ad:14:b6:98:6c:a3:7a:ef:4c:99:
         be:d7:69:12:39:22:54:ec:65:cf:9d:de:cc:36:f5:eb:26:81:
         0d:45:96:d2:5a:d3:0c:70:b8:91:8a:41:c2:c5:09:a7:70:45:
         83:95:24:bc:b4:ad:be:8c:a4:2f:e7:00:e0:66:ff:a8:7a:cb:
         e8:5c:ed:f0:18:28:99:df:6d:e5:32:1a:b5:bd:66:fd:a6:d3:
         15:45:c6:57:c1:30:3d:6c:66:03:bc:01:1c:38:5d:cc:d3:c2:
         75:8f:f1:22:ee:c7:a4:fc:9f:12:42:48:f8:2d:79:8c:50:c0:
         17:71:b4:ea:ce:b5:40:60:db:87:80:04:44:33:5b:00:b3:4c:
         d0:c4:c2:d1:ef:3c:0f:60:61:18:bc:5d:2a:b3:c7:54:7a:56:
         05:07:b5:87:5a:18:f8:34:6d:61:e5:c8:26:4c:30:79:fc:48:
         5e:76:1d:f4:5c:a9:bc:80:5a:ee:ff:43:80:fb:cf:62:18:b4:
         e2:74:53:c4:0f:d3:58:a8:7e:72:a1:13:a7:f1:eb:35:eb:5e:
         9c:df:de:a2

# 也可以通过已有证书生成CSR
$ openssl x509 -x509toreq -in existed.crt -out new.csr -signkey existed.key
 ```

以上过程也可以支持非交互方式，这样可以实现证书自动化, 建立一个配置文件server.cnf

```ini
[req]
prompt = no
distinguished_name = distinguished_name
[distinguished_name]
CN = www.amas.org
emailAddress = amas@gmail.com
O = amas
L = Beijing
C = CN
```

```bash
$ openssl req -new -config server.cnf -key server.key -out server.csr
```



做好CSR之后，你可以干三件事情

1. 发给其他机构，比如PublicCA, 让他们签发证书(.crt)
2. 使用已有证书(.crt)给CSR颁发证书
3. 自己给CSR签名，生成自签名证书



我们先来操作3, 自签名证书:

```bash
$ openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
$ cat server.crt
-----BEGIN CERTIFICATE-----
MIIDeTCCAmECFBtP7j7syGImz149qYFXMQY5YmrXMA0GCSqGSIb3DQEBCwUAMHkx
CzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5n
MQ0wCwYDVQQKDARhbWFzMRgwFgYDVQQDDA9zZXJ2ZXIuYW1hcy5vcmcxHTAbBgkq
hkiG9w0BCQEWDmFtYXNAZ21haWwuY29tMB4XDTIwMDMyNTAzMDQ1MloXDTIxMDMy
NTAzMDQ1MloweTELMAkGA1UEBhMCQ04xEDAOBgNVBAgMB0JlaWppbmcxEDAOBgNV
BAcMB0JlaWppbmcxDTALBgNVBAoMBGFtYXMxGDAWBgNVBAMMD3NlcnZlci5hbWFz
Lm9yZzEdMBsGCSqGSIb3DQEJARYOYW1hc0BnbWFpbC5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQCte9BMSGyvM2nZ2jIvaCq2m5+lSlRdKSpV+uq2
2FBHi9kisNxnzkmgVvGY80ZtjOtZN6vUDmYwh97E/IRM0qzu/HNOgB5MO3cwQkh2
JDlE+HLnrQ7LLultnk8qBgNa/cOcE1f/oqbfLIjhbk22Wpgk8eZ3oilpc7FpW0hV
ucT/t8HIkrR1Gd3CgIIFw+ZI7hWr6/FT0puxM6JzNsoYy19k9+yoUuOZaFnUQ0rB
cSYogjUv/4dsc4N300nz5aLCuhqA8hgi7+LvwauUBvryZAEaOjKUkcoJW1lYBb/e
Crq9UzqZbaxi5Nwy0DhGqEF1vnwdyFtztKKuEty473WS47pZAgMBAAEwDQYJKoZI
hvcNAQELBQADggEBADkId1+VTvT5Vz3bR028wZJqVtfqUrRx9TJy9SPHJV0hH2GH
S8Q8aZ2hNijtwDIcD85J99nt15mpqTgeN7MOHWbB5IixAIE3hQ9SfjKNBYM+N7RC
xMmQtW7U7HnU7QhslN4LBvhuJBOephb1XmgHbzifFIU8miIfNa/ty030W4a268wa
umofnJ7i1WHjQiL3WtSYxFebea5s20na+cohxZnfcWCkGXPBDgJOVd15yVo3746z
pMxHA8HWRMbOqQEacFZ7Nhkc6Cbo/HAJbZCkMfifot0JQDG1b1/M+gOHxf4U0oYX
1gjKndm9SsKaBDawa+SZamb+MYzZyiXbqSQ3T2M=
-----END CERTIFICATE-----


# 解释下server.crt
$ openssl x509 -in server.crt -text -noout
Certificate: 
    Data:
        Version: 1 (0x0)
        Serial Number:
            1b:4f:ee:3e:ec:c8:62:26:cf:5e:3d:a9:81:57:31:06:39:62:6a:d7
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = CN, ST = Beijing, L = Beijing, O = amas, CN = server.amas.org, emailAddress = amas@gmail.com
        Validity
            Not Before: Mar 25 03:04:52 2020 GMT
            Not After : Mar 25 03:04:52 2021 GMT
        Subject: C = CN, ST = Beijing, L = Beijing, O = amas, CN = server.amas.org, emailAddress = amas@gmail.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:ad:7b:d0:4c:48:6c:af:33:69:d9:da:32:2f:68:
                    2a:b6:9b:9f:a5:4a:54:5d:29:2a:55:fa:ea:b6:d8:
                    50:47:8b:d9:22:b0:dc:67:ce:49:a0:56:f1:98:f3:
                    46:6d:8c:eb:59:37:ab:d4:0e:66:30:87:de:c4:fc:
                    84:4c:d2:ac:ee:fc:73:4e:80:1e:4c:3b:77:30:42:
                    48:76:24:39:44:f8:72:e7:ad:0e:cb:2e:e9:6d:9e:
                    4f:2a:06:03:5a:fd:c3:9c:13:57:ff:a2:a6:df:2c:
                    88:e1:6e:4d:b6:5a:98:24:f1:e6:77:a2:29:69:73:
                    b1:69:5b:48:55:b9:c4:ff:b7:c1:c8:92:b4:75:19:
                    dd:c2:80:82:05:c3:e6:48:ee:15:ab:eb:f1:53:d2:
                    9b:b1:33:a2:73:36:ca:18:cb:5f:64:f7:ec:a8:52:
                    e3:99:68:59:d4:43:4a:c1:71:26:28:82:35:2f:ff:
                    87:6c:73:83:77:d3:49:f3:e5:a2:c2:ba:1a:80:f2:
                    18:22:ef:e2:ef:c1:ab:94:06:fa:f2:64:01:1a:3a:
                    32:94:91:ca:09:5b:59:58:05:bf:de:0a:ba:bd:53:
                    3a:99:6d:ac:62:e4:dc:32:d0:38:46:a8:41:75:be:
                    7c:1d:c8:5b:73:b4:a2:ae:12:dc:b8:ef:75:92:e3:
                    ba:59
                Exponent: 65537 (0x10001)
    Signature Algorithm: sha256WithRSAEncryption
         39:08:77:5f:95:4e:f4:f9:57:3d:db:47:4d:bc:c1:92:6a:56:
         d7:ea:52:b4:71:f5:32:72:f5:23:c7:25:5d:21:1f:61:87:4b:
         c4:3c:69:9d:a1:36:28:ed:c0:32:1c:0f:ce:49:f7:d9:ed:d7:
         99:a9:a9:38:1e:37:b3:0e:1d:66:c1:e4:88:b1:00:81:37:85:
         0f:52:7e:32:8d:05:83:3e:37:b4:42:c4:c9:90:b5:6e:d4:ec:
         79:d4:ed:08:6c:94:de:0b:06:f8:6e:24:13:9e:a6:16:f5:5e:
         68:07:6f:38:9f:14:85:3c:9a:22:1f:35:af:ed:cb:4d:f4:5b:
         86:b6:eb:cc:1a:ba:6a:1f:9c:9e:e2:d5:61:e3:42:22:f7:5a:
         d4:98:c4:57:9b:79:ae:6c:db:49:da:f9:ca:21:c5:99:df:71:
         60:a4:19:73:c1:0e:02:4e:55:dd:79:c9:5a:37:ef:8e:b3:a4:
         cc:47:03:c1:d6:44:c6:ce:a9:01:1a:70:56:7b:36:19:1c:e8:
         26:e8:fc:70:09:6d:90:a4:31:f8:9f:a2:dd:09:40:31:b5:6f:
         5f:cc:fa:03:87:c5:fe:14:d2:86:17:d6:08:ca:9d:d9:bd:4a:
         c2:9a:04:36:b0:6b:e4:99:6a:66:fe:31:8c:d9:ca:25:db:a9:
         24:37:4f:63
```





再来看2,使用其他证书给CSR签名颁发证书:

```bash
# 1. 先来生成一个root.crt, 用来给server.crt签名
$ openssl req -x509 -newkey rsa:4096 -keyout root.key -out root.crt -days 365
writing new private key to 'root.key'
Enter PEM pass phrase:             # 密码: 123456
Verifying - Enter PEM pass phrase: # 密码: 123456
-----
...
-----
Country Name (2 letter code) [AU]:.
State or Province Name (full name) [Some-State]:.
Locality Name (eg, city) []:.
Organization Name (eg, company) [Internet Widgits Pty Ltd]:.
Organizational Unit Name (eg, section) []:.
Common Name (e.g. server FQDN or YOUR name) []:.
Email Address []:root@gmail.com

$ ls -l root*
root.crt
root.key

# 2. 用root.key给server.csr签名
$ openssl x509 -req -days 365 -in server.csr -signkey root.key -out server.crt
Getting request Private Key
Enter pass phrase for root.key: #输入: 123456
Generating certificate request
# ----------- DONE ------------
```



### 创建可以给多个主机名使用的证书

> SNA: Subject Alternative Name

openssl默认创建的证书只包含CommanName，而CommanName只能设置一个Host, 这样便引发一个问题，这个证书不能应用于多个域名，增加了管理证书的成本，因此X.509证书引入SAN机制来解决这个问题。简单来说就是在证书中加入一个新的字段，名为`subjectAltName`，这个字段中可以设置多个Hostname, 并支持i通配符(*)

1. 设置了SAN, CommonName将被忽略
2. 新的证书很可能都不会包含CommonName



先准sna.ext

```
subjectAltName = DNS:*.amas.org, DNS:amas.org, DNS:xxx.org
```

签发证书的时候`-extfile`引用这个文件

```bash
$ openssl x509 -req -days 365 -in server.csr -signkey root.key -out server.crt -extfile sna.ext  

$ openssl x509 -text -in server.crt -text -nnout    
...
        X509v3 extensions:
            X509v3 Subject Alternative Name: 
                DNS:*.amas.org, DNS:amas.org, DNS:xxx.org # SNA
...
```



### CA颁发的证书

自签名证书和CA颁发的证书有些区别，我们来看一下区别

```bash
# 从/etc/ssl/certs目录下随便找一个证书,留意X509v3 extensions这个字段
$ openssl x509 -text -noout -in /etc/ssl/certs/00673b5b.0
...
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage: critical
            	Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
            	TLS Web Server Authentication, TLS Web Client Authentication
...
```

#### BC

```
X509v3 Basic Constraints: critical
	CA:TRUE
```

自签名证书要么没有`Basic Constraint`字段，要么是CA:FALSE

#### KU和EKU

```
X509v3 Key Usage: critical
	Digital Signature, Key Encipherment
```

X509v3 Key Usage: 限制了证书使用的范围

#### CRL

```
X509v3 CRL Distribution Points:
    Full Name:
    URI:http://crl.starfieldtech.com/sfs3-20.crl
```

X509v3 CRL Distribution Points: Revocation LIst的URL, 通常作废证书的CRLs和CA-singed列表每7天更新一次

#### CPS

```
X509v3 Certificate Policies:
    Policy: 2.16.840.1.114414.1.7.23.3
    CPS: http://certificates.starfieldtech.com/repository/
```

Certificate Policy Statement (CPS)

#### AIA

```
Authority Information Access:
    OCSP - URI:http://ocsp.starfieldtech.com/
    CA Issuers - URI:http://certificates.starfieldtech.com/repository/sf...
    _intermediate.crt
```

包含两个重要信息:

1. OCSP(CA's Online Certificate Status Protocol), 可以实时检测证书是否已被回收
2. ?

#### SKI

```
X509v3 Subject Key Identifier: 
	7B:5B:45:CF:AF:CE:CB:7A:FD:31:92:1A:6A:B6:F3:46:EB:57:48:50
```

Subject Key Indentifier , 

#### SAN

```
X509v3 extensions:
    X509v3 Subject Altern ative Name: 
    DNS:*.amas.org, DNS:amas.org, DNS:xxx.org # SNA
```



### 常见的证书文件格式

| 格式         | 说明                                                 |
| ------------ | ---------------------------------------------------- |
| DER          | DER ASN.1二进制编码                                  |
| PEM          | ASCII, Base64编码DER, 可以有多个，用BEGIN和END来区分 |
| PKCS#7       | RFC2315                                              |
| PKCS#12(PFX) | 微软产品中比较常见                                   |



```bash
# PEM -> DER
$ openssl x509 -outform PEM -in server.crt -out server.pem 

# PEM -> PFX
$ openssl pkcs12 -export -out server.p12 -inkey sever.key -in server.crt -certfile

# PEM -> PKCS#7
$ openssl crl2pkcs7 -nocrl -out server.p7b -certfile server.crt

# PKCS#7 -> PEM
$ openssl pkcs7 -in server.p7b -print_certs -out server.crt
```

## SSL的配置

SSL最需要关注的配置就是加密套件，加密套件用来配置整个SSL过程中所使用的加密算法，总的来说就是一个算法列表，下面所有的工作不过是通过一些配置得到一个算法组合列表

### 加密套件

```sh
# 查看所有支持的加密套件
$ openssl ciphers -v 'ALL:COMPLEMENTOFALL' 
TLS_AES_256_GCM_SHA384  TLSv1.3 Kx=any      Au=any  Enc=AESGCM(256) Mac=AEAD
...
# TLS_AES_256_GCM_SHA384 : 套件名
# TLSv1.3                : 支持的最低协议版本
# Kx                     : 密钥交换算法
# Au                     : 认证算法, Authentication Algorithm
# Enc                    : 加密算法和强度
# Mac                    : MAC算法
```

### 加密套件关键字

| KEYWORD             | 说明                                             |
| ------------------- | ------------------------------------------------ |
| DEFAULT             |                                                  |
| COMPLEMENTOFDEFAULT | 目前是ADN                                        |
| ALL                 | 除了eNULL之外的所有套件                          |
| COMPLEMENTOFALL     |                                                  |
| HIGH                | 高加密强度，通常是说密钥长度大于128bit           |
| MEDIUM              | 通常是128bit密钥                                 |
| LOW                 | 地强度密钥，通常是低于128bit的40或56bit (不安全) |
| EXP, EXPORT         | (不安全)                                         |
| EXPORT40            | (不安全)                                         |
| EXPORT56            | (不安全)                                         |
| TLSv1, SSLv3, SSLv2 |                                                  |

### 摘要算法关键字

| KEYWORD       | 说明 |
| ------------- | ---- |
| MD5           | 过时，不安全     |
| SHA,SHA1      |      |
| SHA256(v1.0+) |      |
| SHA384(v1.0+) |      |

### 认证关键字

| KEYWORD       | 说明                             |
| ------------- | -------------------------------- |
| aDH           | DH authentication (没有实现)     |
| aDSS, DSS     | DSS 认证， 证书里包含DSS钥匙     |
| aECDH(v1.0+)  | ECDH认证                         |
| aECDSA(v1.0+) | ECDSA认证                        |
| aNULL         | 不认证，目前是匿名DH算法(不安全) |
| aRSA          | RSA认证，证书中包含RSA钥匙       |
| PSK           | PSK(Pre-Shared Key)认证          |
| SPR           | Secure Remote Password认证       |

### 密钥交换算法关键字

| KEYWORD       | 说明           |
| ------------- | -------------- |
| ADH           | 匿名DH(不安全) |
| AECDH (v1.0+) | 匿名ECDH(不安全)   |
| DH            | DH |
| ECDH (v1.0+)  | ECDH |
| EDH (v1.0+)   | Ephemeral DH |
| EECDH (v1.0+) | Ephemeral ECDH |
| kECDH( v1.0+) | ECDH key agreement |
| kEDH          | Ephemeral DH + key agreement |
| kEECDH(v1.0+) | Ephemeral ECDH + key agreement |
| kRSA,RSA      | RSA密钥交换 |

### 加密机器关键字

| KEYWORD       | 说明           |
| ------------- | -------------- |
| 3DES          |                |
| AES           |                |
| AESGCM(v1.0+) |                |
| CAMELLIA      |                |
| DES           | 不安全         |
| eNULL,NULL    | 不加密，不安全 |
| IDEA          |                |
| RC2           | 不安全         |
| RC4           |                |
| SEED          |                |

### 关键字修饰符
| 修饰符       | 说明           |
| ------------- | -------------- |
| Append          | 向套件列表末尾添加匹配的套件，已存在的不改变在列表中的位置 |
| -          | 移除匹配套件，但是后面若添加，则仍然可以使用 |
| ! | 永久的移除匹配的套件，即便后面添加了也没有用 |
| +     | 将匹配的套件移动到列表末尾(只对已经存在的套件生效，不添加) |

总的来说:

1. 强度要大于128
2. 使用带认证功能的
3. 不要使用MD5
4. 用ECDH密钥交换，比DHE要快
5. 用ECDSA
6. TLS 1.2使用AES GCM
7. 不要用RC4

先过滤掉不安全的套件, 一句话

```
!aNULL !eNULL !LOW !3DES !MD5 !EXP !DSS !PSK !SRP !kECDH !CAMELLIA
```

我们要尽可能使用ECDH和ECDSA

```
kEECDH+ECDSA kEECDH kEDH !aNULL !eNULL !LOW !3DES !MD5 !EXP !DSS !PSK !SRP !kECDH !CAMELLIA
```

最后我们将HIGH追加到列表尾部，然后+SHA将所有SHA套件移动到列表尾部，+RC4将所有RC4移动到列表尾部，最后我们把比较弱的RC4添加到列表尾部，以防有比较古老的客户端也可以使用

```
kEECDH+ECDSA kEECDH kEDH HIGH +SHA +RC4 RC4 !aNULL !eNULL !LOW !3DES !MD5 !EXP !DSS !PSK !SRP !kECDH !CAMELLIA
```

```bash
$ openssl ciphers -v 'kEECDH+ECDSA kEECDH kEDH HIGH +SHA +RC4 RC4 !aNULL !eNULL !LOW !3DES !MD5 !EXP !DSS !PSK !SRP !kECDH !CAMELLIA'
```

## 性能测试

openssl内置了性能比较

```bash
# 单线程测试
$ openssl speed rsa ecdh 

# 多线程测试
$ openssl speed -multi 4 rsa ecdh 
```

## 最佳实践

### 1. 私钥管理

1. 使用RSA2048的私钥或ECC224
2. 保护好你的私钥
   1. 在可信设备上创建CSR，不要使用CA提供的服务
   2. 每年使用新的私钥RENEW证书
   3. 保护好私钥，做好备份
3. SNA中设置好足够用的Hostname
4. 从可信CA获取证书

### 2. SSL配置

1. 部署完整的信任链
2. 使用安全的协议(TLSv1.2以上)
3. 只使用安全的加密套件
4. 了解服务选择密码套件的方法，使之可控

### 3. 性能

​	1. 不使用过长的密钥

### 4. 应用程序设计

	1. 100%的加密通讯(HTTPS)
 	2. 避免混合内容
 	3. 从第三方获得的资源也要加密(HTTPS)
 	4. 安全Cookies

### 5. Validation
### 6. 高级主题
1. Extended Validation certificate
2. Public key pinning
3. ECDSA private keys
4. OCSP Stapling





## 参考

- 《OPENSSL COOKBOOK》: A Guide to the Most Frequently Used OpenSSL Features and Commands