package server

import (
	"github.com/jodealter/go_blog/router"
	"log"
	"net/http"
)

var App = &MsServer{}

type MsServer struct {
}

func (s *MsServer) Start(ip, port string) {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port

	//路由功能
	router.Router()
	server := http.Server{
		Addr: ip + ":" + port,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
