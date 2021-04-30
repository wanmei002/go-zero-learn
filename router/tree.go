package router

import (
    "errors"
    "fmt"
    "reflect"
)

type (
    Result struct {
        Item interface{}
        Params map[string]string
        Found bool
    }
    
    innerResult struct {
        key string
        value string
        named bool
        found bool
    }
    
)

var NotFond Result

func NewTree() *Tree {
    return &Tree{
        root: newNode(nil),
    }
}

// Add adds item to associate with route.
func (t *Tree) Add(route string, item interface{}) error {
    if len(route) == 0 || route[0] != '/' {
        return errors.New("path should start with '/'")
    }
    
    if item == nil {
        return errors.New("empty item")
    }
    
    return add(t.root, route[1:], item)
}

func add(nd *node, route string, item interface{}) error {
    if len(route) == 0 {
        if nd.item != nil {
            return errors.New("duplicated item")
        }
        
        nd.item = item
        return nil
    }
    
    if route[0] == '/' {
        return errors.New("duplicated slash")
    }
    
    for i := range route {
        if route[i] == '/' {
            token := route[:i]
            children := nd.getChildren(token)
            if child, ok := children[token]; ok {
                if child != nil {
                    return add(child, route[i+1:], item)
                }
                
                return errors.New("search tree is in an invalid state")
            }
            
            child := newNode(nil)
            children[token] = child
            return add(child, route[i+1:], item)
        }
    }
    
    children := nd.getChildren(route)
    if child, ok := children[route]; ok {
        if child.item != nil {
            return errors.New("duplicated item")
        }
        
        child.item = item
    } else {
        children[route] = newNode(item)
    }
    fmt.Printf("route:%v;", route)
    getMethod(nd)
    
    return nil
}

// 获取 node 绑定的方法
func getMethod(nd *node) {
    if nd == nil {
        return
    }
    fmt.Println(nd.item)
    valNode := reflect.ValueOf(nd.item)
    fmt.Println("kind:", valNode.Kind(), "; func:", reflect.Func)
    valNode.Method(0)
    fmt.Printf("=== item :%+v\n", valNode)
    for _, k := range nd.children {
        for kk, vv := range k {
            fmt.Println("==", kk)
            getMethod(vv)
        }
    }
}


// Search 方法查找匹配的path
func (t *Tree) Search(path string) (Result, bool) {
    if len(path) <= 0 || path[0] != '/' {
        return NotFond, false
    }
    var result Result
    // 递归查找字典树匹配
    
    ok := t.next(t.root, path[1:], &result)
    
    return result, ok
}


// next 递归匹配路由
func (t *Tree) next(n *node, path string, result *Result) bool {
    // 如果 path 最后是 `/`, 会走到这
    if len(path) <= 0 && n.item != nil {
        result.Item = n.item
        return true
    }
    
    for i := range path {
        if path[i] == '/' {
            token := path[:i]
            // 开始匹配路由
            // 思路分析
            n.forEach(func(s string, nd *node) bool {
                if r := match(s, token); r.found {
                    if t.next(nd, path[i+1:], result) {
                        if r.named {
                            addParam(result, r.key, r.value)
                        }
                        return true
                    }
                }
                return false
            })
            
        }
    }
    
    // path 最后不是 `/` 根据 `/` 切分path的时候，
    
    return n.forEach(func(s string, v *node) bool {
        if r := match(s, path); r.found && v.item != nil {
            if r.named {
                addParam(result, r.key, r.value)
            }
            result.Item = v.item
            return true
        }
        return false
    })
}


func addParam(result *Result, k, v string) {
    if result.Params == nil {
        result.Params = make(map[string]string)
    }
    
    result.Params[k] = v
    
}


func match(pat, token string) innerResult {
    if pat[0] == ':' {
        return innerResult{
            key:    pat[1:],
            value:  token,
            named:  true,
            found:  true,
        }
    }
    
    return innerResult{
        found:  pat == token,
    }
}
