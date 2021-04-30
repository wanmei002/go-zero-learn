package main

import (
    "fmt"
    "github.com/wanmei002/go-zero-learn/handler"
    "github.com/wanmei002/go-zero-learn/middleware"
    "github.com/wanmei002/go-zero-learn/server"
)

func main(){
    
    svr := server.NewServer()
    ctx := &server.Context{}
    svr.Use(middleware.RequestTime)
    
    handler.RegisterHandler(svr, ctx)
    
    if err := svr.Start("127.0.0.1", "8093"); err != nil {
        fmt.Println("server start failed; err :", err)
    }
}
