package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// 接收客户端 request，并将 request 中带的 header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 当访问 localhost/healthz 时，应返回 200

func healthz(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {

		w.Header().Set(k, strings.Join(v, ","))

	}
	w.Header().Set("version", os.Getenv("version"))
	fmt.Printf("%s\n", req.RemoteAddr)
	fmt.Fprintf(w, "200")

}

func main() {

	http.HandleFunc("/healthz", healthz)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
