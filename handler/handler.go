package handler

import (
    "github.com/wanmei002/go-zero-learn/router"
    "github.com/wanmei002/go-zero-learn/server"
    "net/http"
)

// 在这里添加路由

func RegisterHandler(s *server.Server, c *server.Context) {
    s.AddRoutes([]router.Route{
        {
            Path:"/",
            Method:http.MethodGet,
            Handler:Home(c),
        },
        {
            Path:"/zzh",
            Method:http.MethodGet,
            Handler:ZZH(c),
        },
    })
}