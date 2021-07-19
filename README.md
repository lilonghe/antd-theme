# Antd Theme
某些场景下需要改变整体的 theme，这里利用预生成的方式生成对应的 theme 文件。

# 注意
动态创建标签加载 theme 样式文件会有一定延迟，可以考虑适当加个 loading 状态或者动态渲染 html。

# FAQ
## 自定义的组件也想使用 theme 里面的颜色
修改 common.less, 在里面用 CSS 定义使用到的全局变量，比如：
```
:root {
    --primary-color: @primary-color;
}
```