

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

### place-content | justify-content | align-content

```css
.container {
	justify-content: start | end | center | stretch | space-around | space-between | space-evenly;	
}
```



### grid-auto-columns | grid-auto-rows

```css
.container {
	grid-auto-columns: <track-size> ...;
	grid-auto-rows: <track-size> ...;
}
```



## 組件庫

- http://photonkit.com/
- https://github.com/sindresorhus/awesome-electron