package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Test 中间件demo实现
type TestServer struct {

}

func (t *TestServer)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

func main() {
	t := &TestServer{}
	http.Handle("/", timeMiddleware(t))
	log.Fatal(http.ListenAndServe(":8080",nil))
}