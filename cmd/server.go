package main

import (
	"flag"
	"net/http"
	"os"
	"github.com/phpor/go_simple_fileserver/config"
	"github.com/phpor/go_simple_fileserver/controller"
)

func main() {
	// 参数处理
	port := flag.String("port", "8082", "port to listen")
	data_dir := flag.String("data_dir", "/tmp/data", "dir to store file")

	_,err := os.Stat(*data_dir)
	if err != nil {
		os.MkdirAll(*data_dir, 0777)
	}
	config.Set("data_dir", *data_dir)

	// 配置
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controller.Index(w, r)
	})

	// 启动
	http.ListenAndServe(":"+ *port, nil)



}