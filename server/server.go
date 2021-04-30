package server

import (
    "fmt"
    "github.com/wanmei002/go-zero-learn/router"
    "net/http"
)

// Context 上下文
type (
    
    Middleware func(http.HandlerFunc) http.HandlerFunc
    
    Context struct {
    
    }
    
    Conf struct {
        Host string
        Port string
    }
    
    
    Engine struct {
        trees []router.Route
        middlewares []Middleware
    }
    
    
    // Server 主要配置参数
    // 和处理路由中间件  路由什么的
    Server struct {
        Conf
        Engine
    }
)

func NewServer() *Server {
    return &Server{}
}


func (s *Server) AddRoutes(routes []router.Route) {
    s.Engine.trees = append(s.Engine.trees, routes...)
    fmt.Println(s.Engine.trees)
}
// Use 添加中间件
func (s *Server) Use(handler Middleware) {
    s.Engine.Use(handler)
}

func (s *Server) Start(args ...string) error {
    return s.Engine.Start(args...)
}
