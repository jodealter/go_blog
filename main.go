package main

import (
	"github.com/jodealter/go_blog/common"
	"github.com/jodealter/go_blog/server"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func init() {

	//模版加载
	common.LoadTemplate()
}
func main() {
	server.App.Start("127.0.0.1", "8080")
}
