package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 设置路由及处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go Web Server!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	// 启动服务器
	port := "8080"
	fmt.Printf("Server is running at http://localhost:%s/\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
