package views

import (
	"errors"
	"github.com/jodealter/go_blog/common"
	"github.com/jodealter/go_blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 首页
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错:", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！！！"))
	}

	index.WriteData(w, hr)
}
