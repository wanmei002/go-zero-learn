## 描述
本框架是本人浏览 go-zero 单体框架原理，自己学习制作了一个简易框架
go-zero 地址 [https://github.com/tal-tech/go-zero](https://github.com/tal-tech/go-zero)

## 目录介绍
### handler
handler 里面存放了路由注册和处理路由对应的方法
`RegisterHandler` 注册路由的方法
handler/method.go 文件里有两个示例的方法

### middleware
存放中间件的文件夹
`中间件的使用 Server -> Use` Server 结构体的 `Use` 方法

### router
这个文件夹里面的源码包含了 路由字典树的生成，路由的匹配等功能

### server
包含了中间件的注册，路由的注册等逻辑
