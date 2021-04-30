package middleware

import (
    "fmt"
    "net/http"
    "time"
)

func RequestTime(next http.HandlerFunc) http.HandlerFunc {
    
    return func(w http.ResponseWriter, r *http.Request) {
        t1 := time.Now().UnixNano()/1000
        
        next.ServeHTTP(w, r)
        
        sub := time.Now().UnixNano()/1000
        
        fmt.Println(r.URL.Path, " used ", sub-t1)
    }
}
