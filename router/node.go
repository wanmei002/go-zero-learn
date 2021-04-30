package router

import "fmt"

func (nd *node) getChildren(route string) map[string]*node {
    if len(route) > 0 && route[0] == ':' {
        return nd.children[1]
    }
    
    return nd.children[0]
}

func newNode(item interface{}) *node {
    return &node{
        item: item,
        children: [2]map[string]*node{
            make(map[string]*node),
            make(map[string]*node),
        },
    }
}

func (nd *node) forEach(f func(string, *node) bool) bool {
    for _, children := range nd.children {
        for k, v := range children {
            fmt.Printf("foreach :%v, %v\n", k, v)
            return f(k, v)
        }
    }
    return false
}