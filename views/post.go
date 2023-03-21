package views

import (
	"errors"
	"github.com/jodealter/go_blog/common"
	"github.com/jodealter/go_blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//获取路径参数
	path := r.URL.Path

	//去除前缀
	pIdStr := strings.TrimPrefix(path, "/p/")

	//去除后缀
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此路径"))
		return
	}

	postResL, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询不出"))
		return
	}
	detail.WriteData(w, postResL)
}
