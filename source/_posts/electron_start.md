

# Electron

> Electron = Chromium   + Node



## 安装

```bash
 $ sudo npm install -g electron
 
 #  如果报错: 
 # Error: EACCES: permission denied, mkdir '/usr/lib/node_modules/electron/.electron'
 $ sudo npm install -g electron --unsafe-perm=true
```



## Hello World

配置webstorm, 建立工程, 目录结构如下, 

```
.
├── app
│   └── main.js
├── node_modules
└── package.json
```

package.json:

```json
{
    "name": "ihello",
    "version": "1.0.0",
    "description": "hello",
    "main": "./app/main.js",
    "scripts": {
        "start": "electron .",
        "test": "echo \"Error: no test specified\" && exit 1"
    },
    "author": "amas",
    "dependencies": {
        "electron": "^4.1.1"
    }
}
```



1.  在webstrom中选择Open, 打开这个目录
2. 新建一个npm的task, 命令选择start



main.js:

```js
const {app} = require('electron');
app.on('ready', () => {
 console.log('Hello from Electron');
});
```

```shell
$ npm start
> electron .
Hello from Electron
```



## 进程模型

![img](https://cdn-images-1.medium.com/max/900/1*EJETq7XOPz5RVY5IfF6NIg@2x.png)





### 进程间通讯

简单的方法, 在Renderer进程中使用remote对象.

![](/home/amas/2019-03-28-102749_875x413_scrot.png)



ipcMain和ipcRenderer的命名方式是以你在哪里使用来命名的, 所以ipcMain只能在主进程中出现, 而ipcRenderer在renderer进程中出现.

```js
ipcMain.on('ping', () => {
  console.log('pong');
});

# 发送消息到主进程
ipcRenderer.send('ping', 'ping from renderer');
```

![](/home/amas/2019-03-28-142315_797x662_scrot.png)



## 使用require

我们可以在html中使用require函数来加载javascript, 加载的路径是当前HTML文件的所在路径

```html
<script>
  require('./render.js');
</script>
```



## 如何打开外部浏览器?

```
const {shell} = require('electron');
shell.openExternal('http://www.google.com');
```

## 如何调用操作系统原生的对话框?

```js
# 在renderer进程中
const { dialog } = require('electron').remote

# 在主进程中
const { dialog } = require('electron')


# 彈出對話框
const files = dialog.showOpenDialog({
    title: "爸爸, 请选择一个文件",
    properties: ['openFile'],
    filters: [
        { name: 'Images', extensions: ['jpg', 'png', 'gif'] },
        { name: 'Movies', extensions: ['mkv', 'avi', 'mp4'] },
        { name: 'Custom File Type', extensions: ['as'] },
        { name: 'All Files', extensions: ['*'] }
    ]
});
```



## BrowserWindow

![](/home/amas/2019-03-28-103400_788x317_scrot.png)

Each BrowserWindow instance has a property called webContents , which stores an
object responsible for the web browser window that we create when we call new Browser-
Window() . webContents is similar to app because it emits events based on the lifecycle
of the web page in the renderer process.
The following is an incomplete list of some of the events that you can listen for on
the webContents object:

- did-start-loading
- did-stop-loading
- dom-ready
- blur
- focus
- resize
- enter-full-screen
- leave-full-screen

## electron-compile

無論前端技術如何變化, 瀏覽器始終都只在處理三種東西:

- CSS
- HTML
- JavaScript

![](/home/amas/2019-03-28-152620_970x237_scrot.png)



> 经过深思熟虑, 我准备采用
>
> Stylus + Jade



## CSS Grid

- https://css-tricks.com/snippets/css/complete-guide-grid/

### grid-template + grid-column + grid-row



```css
.container {
  display: grid | inline-grid;
}

.item {
  grid-column-start: <number> | <name> | span <number> | span <name> | auto
  grid-column-end: <number> | <name> | span <number> | span <name> | auto
  grid-row-start: <number> | <name> | span <number> | span <name> | auto
  grid-row-end: <number> | <name> | span <number> | span <name> | auto
}
<line> - can be a number to refer to a numbered grid line, or a name to refer to a named grid line
span <number> - the item will span across the provided number of grid tracks
span <name> - the item will span across until it hits the next line with the provided name
auto - indicates auto-placement, an automatic span, or a default span of one

.container {
  grid-template-columns: <track-size> ... | <line-name> <track-size> ...;
  grid-template-rows: <track-size> ... | <line-name> <track-size> ...;
}
<track-size> - can be a length, a percentage, or a fraction of the free space in the grid (using the fr unit)
<line-name> - an arbitrary name of your choosing
```



![](/home/amas/2019-03-29-111643_650x783_scrot.png)





![](/home/amas/2019-03-29-111827_629x754_scrot.png)



![](/home/amas/2019-03-29-111827_629x754_scrot.png)



![](/home/amas/2019-03-29-171438_628x754_scrot.png)

### fr: 一个特别的单位

### grid-template-areas + grid-area

````css
.item {
  grid-area: <name> | <row-start> / <column-start> / <row-end> / <column-end>;
}
````

>start / end





```css
.item-a {
  grid-area: header;
}
.item-b {
  grid-area: main;
}
.item-c {
  grid-area: sidebar;
}
.item-d {
  grid-area: footer;
}

.container {
  display: grid;
  grid-template-columns: 50px 50px 50px 50px;
  grid-template-rows: auto;
  grid-template-areas: 
    "header header header header"
    "main main . sidebar"
    "footer footer footer footer";
}
```

![](/home/amas/2019-03-29-181333_612x423_scrot.png)



### grid-gap | grid-column-gap | grid-row-gap

调整单元格之间的间距

### place-self | justify-self | align-self

- align: 控制单元格中的元素上下移动
- justify: 控制单元格中的元素左右移动
- place: 控制二者

### place-items | justify-items | align-items

用来控制所有item的对齐方式

### place-content | justify-content | align-content

将所有items作为一个整体来控制, 调整其上下左右对齐方式

```css
.container {
	justify-content: start | end | center | stretch | space-around | space-between | space-evenly;	
}
```

### grid-auto-flow

### grid-auto-columns | grid-auto-rows

```css
.container {
	grid-auto-columns: <track-size> ...;
	grid-auto-rows: <track-size> ...;
}
```



## CSS 伪类

### 链接相关

|          |      |      |
| -------- | ---- | ---- |
| :link    |      |      |
| :hover   |      |      |
| :active  |      |      |
| :visited |      |      |

### 输入相关

- :target
- :enabled
- :disabled
- :checked
- :indeterminate 初始化未操作
- :requited
- :optional
- :read-only
- :read-write

### 选择子元素相关

- :root  通常就是html标签
- :first-child 
- :last-child
- :nth-child(n)
- :nth-last-child(n)
- :nth-of-type(n)
- :first-of-type
- :last-of-type
- :nth-last-of-type(n)
- :only-of-type



![](/home/amas/2019-04-02-141444_689x573_scrot.png)



### 特殊选择

#### :not()

```css
div:not(.music)
input:not([disabled])
```

#### :empty

选择没有text, 且没有子元素的标签

```
<p></p>
```



#### ::before

在标签之前可添加内容, 可以是文字图片, 但是这些内容不在DOM中

#### ::after

在标签之后可添加内容, 可以是文字图片, 但是这些内容不在DOM中



### 文本相关

#### ::first-letter

#### ::last-letter

#### :lang()

这个只有IE8支持???



## CSS渐变

制造一个光线渐变, 我们需要指定三个东西:

1. 圆点, 可以理解为光源
2. 半径, 可以理解外光照的范围
3. 渐变颜色, 至少是两个

```
<radial-gradient> = radial-gradient(
  [ [ <shape> || <size> ] [ at <position> ]? , |
    at <position>, 
  ]?
  <color-stop> [ , <color-stop> ]+
)
```

- shape
  - circle
  - ellipse
- size: 可以是具体的长度单位比如20px, 20em, 20%等, 也可以是下面几个特殊点
  - closest-side / fathest-side (coloset-side指距离光源点最近的边,)
  - closest-corner / fathest-coner
- position
- color-stop

```css
#g1 {
  width: 100px;
  height: 100px;
  background: radial-gradient(#000, #fff);
}
#g2 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle closest-side, #000, #fff);
}
#g3 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle farthest-side, #000, #fff);
}
#g4 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle 20px, #000, #fff);
}
#g5 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle at top right, #000, #fff);
}
#g6 {
  width: 100px;
  height: 200px;
  background: -webkit-repeating-radial-gradient(center, circle, #f00, #00f 20%);
}
#g7 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle farthest-corner at center center, #000, #fff);
}
#g8 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle farthest-corner at left, #000, #fff);
}
#g9 {
  width: 100px;
  height: 200px;
  background: radial-gradient(circle closest-corner at left, #000, #fff);
}
```

## 颜色

### rgb

- red
- green
- blue

### rgba

我们可以通过三原色+透明度来定义一种颜色

- red
- green
- blue
- alpha

### hlsa

![Image result for hue](data:image/png;base64,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)

色相/饱和度/亮度/透明度 四元组也可以定义颜色

- **Hue** - Think of a color wheel. Around 0o and 360o are reds 120o are greens, 240o are blues. Use anything in between 0-360. Values above and below will be modulus 360.
- **Saturation** - 0% is completely denatured (grayscale). 100% is fully saturated (full color).
- **Lightness** - 0% is completely dark (black). 100% is completely light (white). 50% is average lightness.
- **alpha** 



## inherit

```css
input, select, button { font: inherit; }
a { color: inherit; }
a { background: inherit; }
```

## semi-transparent colors

```css
hsla(0, 0%, 100%, .5)
rgba(255,255,255,.5)
```

## 内圆角

```stylus
// 较为复杂的方法
#inner-round
  width 100px
  margin 2em
  background  #655
  padding .8em
  border-radius .8em
  &>div
    padding 1em
    background tan
    border-radius .6em

// 简单方法, 只需要一个div
#inner-round2
  width 100px
  margin 2em
  background tan
  padding 1em
  border-radius .8em
  box-shadow: 0 0 0 .6em #655;
```



## 调整背景图片的位置

```css
background: url(code-pirate.svg) no-repeat bottom right #58a;
background-position: right 20px bottom 10px;


// 假如我们调整了容器的padding, 那么上面的css还需要随之修改, 而通过background-origin: 可以指定背景图片在哪个盒子里进行偏移
background: url("code-pirate.svg") no-repeat #58a bottom right; /* or 100% 100% */
background-origin: content-box;
```

## 背景栅格

```stylus
div.rect
  width 200px
  height 100px
  margin .5em
  background linear-gradient(red, blue)

div.rect.g2
  background linear-gradient(red 20%, blue 80%)

div.rect.g3
  background linear-gradient(red 50%, blue 50%)

div.rect.g4
  background linear-gradient(40deg,red 50%, blue 50%)

div.rect.g5
  background linear-gradient(red 50%, blue 50%)
  background-size 100% 10px

// 红占25%, blue占75%
div.rect.g6
  background linear-gradient(red 25%, blue 0)
  background-size 100% 50%

div.rect.g7
  background linear-gradient(90deg, red 50%, blue 50%)
  background-size 10px 100%

div.rect.g8
  background linear-gradient(90deg, red 50%, transparent 50%)
  background-size 5px 100%

div.rect.g9
  background repeating-linear-gradient(red, red 10px, yellow 0, yellow 20px)

```



## 背景模式

基于inear-gradient/repeating-linear-gradient 我们可以制造出非常多的花样, 下面这个网站告诉你可以有多作:

- https://leaverou.github.io/css3patterns/



### 图片应用于border

```css
	border: 1em solid transparent;
	background: linear-gradient(white, white) padding-box,
	            url(http://csssecrets.io/images/stone-art.jpg) border-box  0 / cove
```





## 形状

### border-radius: 100px / 75px;

## 組件庫

- http://photonkit.com/
- https://github.com/sindresorhus/awesome-electron
- https://pixinvent.com/materialize-material-design-admin-template/landing/
- https://materializecss.com/showcase.html