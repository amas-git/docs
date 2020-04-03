





# CSS

## 选择器

## 伪元素

## 颜色

## 渐变

## 背景

### background

## 盒子模型

![](/src/amas/docs/source/_drafts/assets/2019-04-17-115909_249x216_scrot.png)

- margin-box (这个并没有)
- border-box
- padding-box
- content-box

### box-sizing: content-box



### inline / inline-block

- 对于inline和inline-block可以设置line-height

### line-height

![img](https://pic2.zhimg.com/80/v2-200f0335ec68bf245c6507afef744ae8_hd.jpg)

### inline元素的幽灵空白问题

问题是这样的

```html
<nav>
  <a href="#">AAA</a>
  <a href="#">BBB</a>
  <a href="#">CCC</a>
</nav>
```

```css
nav a {
  display: inline-block;
  padding: 5px;
  background: #008000;
}
```



结果显示出来:

![](/src/amas/docs/source/_drafts/assets/2019-04-18-193607_143x42_scrot.png)

> Q: 那么中间的空隙是啥啊?
>
> A: 这并不是bug, 如果把a标签后面有换行字符, 这就相当于单词之间放置空白一样是正常行为



那么怎么消除呢? 去掉换行符就可以了,  这时候的脑洞就比较大了:

```html
<nav>
  <a href="#">AAA</a><a href="#">BBB</a><a href="#">CCC</a>
</nav>

<nav>
  <a href="#">AAA</a><!--
  --><a href="#">BBB</a><!--
  --><a href="#">CCC</a>
</nav>

<nav>
  <a href="#">AAA</a
    >a href="#">BBB</a
    ><a href="#">CCC</a>
</nav>
```

另一个思路是, 既然这是因为空白字符造成的, 那么将字号设置为0不久可以了

```css
nav {
  font-size: 0;
}
nav a {
  font-size: 16px;
}
```

> 这个方法慎用, 一方面Android某些版本上有bug, 另一方面当字号使用em单位的时候

![](/src/amas/docs/source/_drafts/assets/2019-04-18-211820_165x43_scrot.png)





## 边框

## 文本

![](/src/amas/docs/source/_drafts/assets/2019-04-19-170652_918x881_scrot.png)







## 动画

## 布局

### 基础布局

### Flex布局

### Grid布局

### Table布局

### 如何实现水平居中?

#### inline/inline-*元素的水平居中
```
text-align: center
```
#### 块元素的水平居中
```
margin: 0 auto
```

### 多个块元素水平居中?

可以使用flex布局

### 如何实现垂直居中?

CSS实现垂直居中较为曲折

#### inline/inline-*元素的垂直居中

> 单行垂直居中
>
> 方法一: 上下padding一致法
>
> ```
> padding-top: 30px;
>  padding-bottom: 30px;
> ```
>
> 方法二: line-height与height一致法
>
> ```
>   height 100px
>   line-height 100px
>   white-space: nowrap // 防止文字过长冒出去, 干脆禁止文字换行
> ```



> 多行垂直居中
>
> 方法一: 上下padding一致法, 不罗嗦了
>
> 方法二: table + table-cell法
>
> ```stylus
> #parent
> display table
> #child
> display table-cell
> vertical-align middle
> ```
>
> 方法三: flex法
>
> ```stylus
> #parent
> height 100px // 必须设置高度
> display flex
> justify-content: center
> align-items: center;
> flex-direction: column
> ```
>
> 方法四: 幽灵法
>
> ```stylus
> #parent
> position relative
> height 100px
> width 100px
> overflow auto
> background grey
>   &::before // 伪元素占了1%的宽度, 占满整个行高
>    content " "
>    display inline-block
>    height 100%
>    width 1%
>    vertical-align middle
> 
> #child
>   margin 0
>   width 94% // 宽度的设置十分重要, 因为伪元素实际占据了1%, 但为啥98%不行呢????  inline-*的幽灵空白
>   display inline-block
>   vertical-align middle
> ```
>
> 方法五: absolute+transform法
>
> ```stylus
> #parent
> 	position relative
> #child
>   position absolute
>   top 50%
>   left 50%
>   transform translate(-50%, -50%)
> ```
>
> 方法六: margin auto 大法
>
> ```stylus
> #parent
>   position relative
>   width 100px
>   height 100px
>   border 1px solid
> #child
>   background-color yellow
>   position: absolute // 1. 使用绝对定位
>   top 0
>   left 0
>   right 0
>   bottom 0              // 2. 这个定位法是不可能实现的, 因此浏览器感到很为难
>   height 50px
>   width 50px
>   margin auto         // 3. auto导致浏览器居中显示, 否则child会出现在0,0点, 还记得水平居中块元素的方法么??
> ```
>
> 

万能方法: 使用table包裹

> 这个方法用div+table layout也是一样的效果, 就不罗嗦了

```
    table(style="width: 100%;")
        tr
            td(style="text-align: center; vertical-align: middle;")
                p This is the element need to be center
```



幽灵元素法大概是这个意思:

![](/src/amas/docs/source/_drafts/assets/2019-04-18-191910_496x500_scrot.png)

1. 一般我们可以用:before伪元素, 设置为inline-block, 然后就可以设置height或line-height充满整个容器
2. 这个时候再设置vertical-align, 就可以让元素垂直居中

#### block元素的垂直居中







#### 方法一: margin大法

> 这个方法兼容性非常好, 但比较罗嗦, 有两个变种:
>
> ```
> #parent
> position relative
> width 100px
> height 100px
> border 1px solid
> #child
> background-color yellow
> position: absolute
> top 50%
> left 50%
> height 50%
> width 50%
> margin-top -25%
> margin-left -25%
> ```
>
> 第一种的缺点是只能使用百分比控制子元素的尺寸, 下面这种方法则可以指定子元素的尺寸, 但是IE7是不行的.
>
> ```
> 
> ```
>
> #### 

#### 方法二: vertical-align?

这玩意看上去好像很管用, 其实不然.

1. 它只在table-cell中好使, 这时候怎么居中取决于table的row高度, 这个row高度具体又受控于 [table-height-algorithm](http://www.w3.org/TR/CSS2/tables.html#height-layout)
2. 有部分inline元素也会起作用, 取决于line-height

举个例子:

```html
<div id="foo">Hello</div>
```

#### 方法三:  line-height大法

> 如果只需要垂直居中一行文字

这种方法及其简单, 就是将line-height设置为和height一样即可,只适用于一行文字的垂直居中, 比如你想搞个按钮的话这种方法最合适

``` stylus
#foo
  height 100px
  line-height 100px
  white-space: nowrap // 防止文字过长冒出去, 干脆禁止文字换行
```

通过这种方法居中一个img也是可以的

```
<div id="parent">
    <img src="image.png" alt="" />
</div>



#parent {
    line-height: 200px;
}
#parent img {
    vertical-align: middle;
}
```



#### 方法四: table + table-cell

```jade
    div#parent
        div#child this is cool this is cool
```



```stylus

```

#### 方法五: margin大法

#### 方法六: padding大法

这个思路比较简介, 适合一行文字的垂直居中, 只需要将上下内边距设置为相同即可

```
 
```







## 参考

 - https://css-tricks.com/centering-css-complete-guide/
 - 