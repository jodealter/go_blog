package router

import (
	"github.com/jodealter/go_blog/api"
	"github.com/jodealter/go_blog/views"
	"net/http"
)

func Router() {
	//这两部分是路由
	//api 与 路由router的区别是，api属于路由的一部分，router负责指派任务，而api是负责处理数据的那一部分
	//路由分三种，1返回页面viewa,2api 返回数据 3静态资源

	//1.处理页面
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)

	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
