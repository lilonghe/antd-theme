# Antd Theme
某些场景下需要改变整体的 theme，因为涉及到第三方框架，所以这里利用预生成的方式生成对应的 theme 文件。

# Intro
```
npm install
npm run theme
npm start
```

# FAQ

## 闪烁
因为示例 theme 是后加载的，所以在第一屏可能会出现闪烁，所以推荐有个 loading 态，或者界面动态生成。

## 自定义的组件也想使用 theme 里面的颜色
修改 common.less, 在里面用 CSS 定义使用到的全局变量，比如：
```
:root {
    --primary-color: @primary-color;
}
```
因此示例中 `Hello` 文字的颜色和 `Button` 组件的颜色保持了一致。

# Demo
![demo](./demo/1.png)