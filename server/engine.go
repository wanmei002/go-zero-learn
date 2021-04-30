package server

import (
    "errors"
    "fmt"
    "github.com/wanmei002/go-zero-learn/router"
    "net/http"
    "strings"
)

// engine Middleware
func (e *Engine) Use(middleware Middleware) {
    e.middlewares = append(e.middlewares, middleware)
}

// 开始处理路由 生成字典树
// 1. 中间件处理函数
// 2. 处理过后的函数 生成字典树
// 3. 在 ServeHTTP 中匹配

// ChainMiddleware 把方法用中间件处理下
func (e *Engine) ChainMiddleware(handler http.HandlerFunc) http.HandlerFunc {
    midLen := len(e.middlewares)
    for i := range e.middlewares {// 得倒着，从后面往前面来
        handler = e.middlewares[midLen-1-i](handler)
    }
    
    return handler
}

// 开始 循环调用
func (e *Engine) Start(args ...string) error {
    r := router.NewRouter()
    
    if len(e.trees) < 1 {
        return errors.New("没有处理请求的方法")
    }
    
    for _, route := range e.trees {
        // 1. 需要先对路由进行中间件处理
        route.Handler = e.ChainMiddleware(route.Handler)
        // 2. 再进行字典树处理
        fmt.Println("route handler : ", route.Handler)
        err := r.Handle(route.Method, route.Path, route.Handler)
        if err != nil {
            return err
        }
    }
    addr := "127.0.0.1:8092"
    if len(args) == 2 {
        if args[0] != "" && args[1] != "" {
            addr = strings.Trim(args[0], " ") + ":" + strings.Trim(args[1], " ")
        }
    }
    
    fmt.Println("addr :", addr)
    svr := http.Server{
        Addr:       addr,
        Handler:    r,
    }
    
    // 还需要在字典树里
    
    return svr.ListenAndServe()
    
}
