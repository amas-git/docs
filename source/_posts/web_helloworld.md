# Web开发入门

## html/head

```
html(lang="zh")
  head
    meta(charset="utf-8")
    meta(name="author" content="amas")
    meta(name="description" content="something need to say")
    meta(name="keywords" content="blog,tech,web,html")
    title Hello
    link
    base
    meta
    noscript
    script
    style
```

html/body

```
html
	head
	body
```

## HTML 和 HTML5的区别

|      | HTML5             | OLD HTML |
| ---- | ----------------- | -------- |
| 属性 | <name=a disabled> | <name="a" disabled="true">    |
| / | <br>                  | </br>          |
| End Tag | <p>... | <p>...</p> |

HTML5去掉的标签
	- font
	- center

HTML新加的i标签
	- footer
	- nav
	- sidebar
	- section
	- audio, video
	- canvas

HTML5新特性
	- Drag&drop
	- Web storage
	- GEO

## 如何阅读HTML5规范
> ## Content Model Cagegories
> CMC定义了容器里允许放什么样的内容
> Typical default display
>  定义了元素的默认display属性i以及其他显示细节


## Block Element
> 虽然HTML5中没有这个说法

	- blockquote
	- cite
	- pre

## Editing Elements
	- ins
	- del

## dfn, abbr,time
```
	abbr(title="World Wide Web Consortium") W3C
```
## Code Related Elements: kbd, code, samp, var
 - kbd: keyboard
 - samp: sample

## wbr, br
	- br: break
	- wbr: word break

## sub, sup,s,mark,small
	- sup: 上标
	- sub: 下i标
	- s: 被划掉的文字 
	- mark: 高亮

## strong, em, b, u, i
## span
如果i一段文字的样子山面怎么 都不能满足，那么用span+css

## Character Reference
有很多字符不好打，可以参考： https://www.w3.org/TR/html51/syntax.html#named-character-references

|  Character    |       Character Reference       |
| ---- | ----------------- |
| < | lt |
| > | gt |
|space| nbsp |
|&| amp |
| ✓ | check |
| &copy; |copy|
| &quot; |quot|
|&apos;|apos|
|&#110;|#110|
|&pi;|pi|

