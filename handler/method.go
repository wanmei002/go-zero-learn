package handler

import (
    "fmt"
    "github.com/wanmei002/go-zero-learn/server"
    "net/http"
    "time"
)

func Home(ctx *server.Context) http.HandlerFunc {
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, err := w.Write([]byte(`{"hello":"zzh"}`))
        if err != nil {
            fmt.Println("home write failed, err :", err)
        }
    })
    
}

func ZZH(ctx *server.Context) http.HandlerFunc {
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(10e9)
      
        w.WriteHeader(200)
        
        _, err := w.Write([]byte(`{"hello":"zyn"}`))
        
        if err != nil {
            fmt.Println("zzh method write failed; err :", err)
        }
        
    })
    
}
