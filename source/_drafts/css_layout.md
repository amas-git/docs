## CSS GRID 布局

> GRID是二维的

### 什么是GRID LINE

### 什么是GRID CELL

### 什么是GRID TRAC

### 什么是GRID AREA

>```
>grid-area: <name> | <row-start> / <column-start> / <row-end> / <column-end>;
>```

### GRID CONTAINER 的属性

> 1. grid-template-columns 和  grid-template-rows来定义如何分配容器的空间
> 2. rows/colums都可以指定名字, 方便使用

#### display:  grid | inline-grid

```css
display: grid
```

#### grid-template-columns: track-size ... |  line-name track-size, ...

```jade
body
    .container.demo1
        .item
        .item
        .item
        .item
        .item
        .item
        .item
        .item
        .item
    .container.demo2
        .item
        .item
        .item
        .item
        .item
        .item
        .item
        .item
        .item
```



```stylus
div.container
  display grid
  height 100px
  width 100px
  margin 2px 0 2px
  background grey

div.item
  background chocolate
  border white 1px solid


div.container.demo1
  grid-template-columns  10px auto 10px
  grid-template-rows 10px
```

> auto-fit

#### grid-template-rows

#### grid-gap: grid-row-gap  grid-column-gap;



#### 对齐

```
justify-items: start|end|center|strech
align-items: start|end|center|strech
place-items: <align-items> <justify-items>
```



### GRID ITEM 的属性

> **注意:**
> `float`, `display: inline-block`, `display: table-cell`, `vertical-align` and `column-*` 在GRID中不起作用

#### justify-self: 水平对齐方式

```
justify-self: start | end | center | stretch
```



#### align-self: 垂直对齐方式

```
align-self: start | end | center | stretch
```





## CSS FLEX布局

> FLEX是一维的