# Category Theory

## Category: 合成的本质

> Category由两个概念构成：
>
> - Objects
> - Arrows



> Arrows = Functions = Mophisms



> g . f = g | f



## 合成的性质

1. 结合律：h◦(g◦f) = (h◦g)◦f = h◦g◦f
2. 存在UNIT, 一个叫做ID的函数，可以使 f ◦ id = f，当然 id ◦ f = f



```c++
template<class T> T id(T x) { return x; }
```

```haskell
-- Haskell的标准库中有id函数
id :: a -> a 
id x = x
```



## 程序的本质是合成

> 人脑的短期记忆一般能处理5-9个
>
> 我们常说这段代码优雅美丽，它的意思往往是更容易被人脑所理解
>
> 另一种说法是Surface Area的增长小于Volume的增长，如果是面向对象Classs和Interface就是SurfaceArea, 如果是函数式编程，函数的声明即使SurfaceArea
>
> CategoryTheory并不鼓励你打开Object内部，而仅仅鼓励你关心Object之间的关系，



### 类型与函数

类型的Battle:

- 动态 Vs 静态
  - 区别： 动态类型在运行时检查，静态类型再编译时检查
- 强类型 Vs 弱类型

> 无论语言的系统类型如何，通常都提供绕过类型系统的机制



> Q: 什么是类型推导？
>
> A: 借助编译器程序中的类型信息
>
> Q: 单元测试和强类型哪个更好？



## 类型

> 类型 = 集合  + Bottom = Hask



> Q: 计算机的函数与数学中的函数有何不同？
>
> A: 计算机中的函数需要进行运算，有可能终结也有可能无法终结，如计算中所谓的运行时错误导致无法完成计算，我们把这种情况下函数返回值得类型记为: \_|\_ , Bottom, 
>
> ```hs
> f::Bool -> Bool
> 返回值可能是true,false或者_|_，在haskell中undefine表示Bottom
> ```
>
> 可能返回Bottom的函数叫partial函数



>
>
>Q: 为什么我们再软件开发中需要数学模型？
>
>A: 很难证明操作语义（operational semantics）下的计算是否正确，所以我们会用单元测试的方法检测一个函数执行结果和预期一致， 另外一种方法absurd :: Void -> a是denotational semantics ， 就像数学公式的推导过程类似



> 如果类型是集合，那么类型中的空集是什么类型？是c中的void么？
>
> absurd :: Void -> a
>
> 一个类型，只有一个元素的集合，那是什么类型？ 这个是c中的void
>
> void, () 叫做unit, 
>
> f2() { return 2; } // void -> int
>
> f2 () = 2               //





如何解释:

```
void f() {}
f x = ()
```







## 纯函数和脏函数





##  Curry-Howard isomorphism



## ZERO MORPHISMS



## MONOID

- 一个集合于一个二元操作符，记为+, 该运算是封闭的 
- 结合律：(a+b)+c = a + (b +c)
- 存在zero, 使得zero + a = a



## Kleisli



## FUNCTORS



## PRODUCTS/COPRODUCTS



## LIMITS/COLIMITS



