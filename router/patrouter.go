package router

import (
    "errors"
    "fmt"
    "net/http"
    "path"
)

type (
    node struct {
        item     interface{}
        children [2]map[string]*node
    }
    
    Tree struct {
        root *node
    }

    PatRouter struct {
    trees       map[string]*Tree
    }
)

// NewRouter returns a httpx.Router.
func NewRouter() *PatRouter {
    return &PatRouter{
        trees: make(map[string]*Tree),
    }
}

func (pr *PatRouter) Handle(method, reqPath string, handler http.Handler) error {
    
    if !(method == "GET" || method == "POST") {
        return errors.New("must GET|POST")
    }
    
    if len(reqPath) == 0 || reqPath[0] != '/' {
        return errors.New("path must be begin with '/'")
    }
    
    fmt.Println("reqPath before:", reqPath)
    
    cleanPath := path.Clean(reqPath)
    fmt.Println("cleanPath:", cleanPath)
    tree, ok := pr.trees[method]
    if ok {
        err := tree.Add(cleanPath, handler)
        return err
    }
    
    tree = NewTree()
    pr.trees[method] = tree
    fmt.Printf("pr: %+v\n", pr)
    err := tree.Add(cleanPath, handler)
    fmt.Printf("tree : %+v\n", tree)
    return err
}



func (pr *PatRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    
    fmt.Println(r.URL.Path)
    
    // 开始匹配字典树
    reqPath := path.Clean(r.URL.Path)
    if tree, ok := pr.trees[r.Method]; ok {
        if result, ok := tree.Search(reqPath); ok {
            if len(result.Params) > 0 {
                // 把数据加入 r 中
                
                for k, v := range result.Params {
                    r.Form.Set(k, v)
                }
            }
            result.Item.(http.HandlerFunc).ServeHTTP(w, r)
            return
        }
    }
    
    // 没有匹配到router
    w.WriteHeader(404)
    _, err := w.Write([]byte(`{"code":404,"msg":""}`))
    if err != nil {
        fmt.Println("404 page err:", err)
    }
}


