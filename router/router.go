package router

import "net/http"

// path包 用来生成字典树

// router包用来把事件跟路由绑定在一起

// 1. 先把方法  路径 Handler 放到一个 struct 里
// 2. 把 struct 集合放到 engine 里
type (
    Route struct {
        Path string // 路径
        Method string // request 方法类型,比如 GET POST 等
        Handler http.HandlerFunc
    }
)


