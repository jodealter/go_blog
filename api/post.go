package api

import (
	"errors"
	"github.com/jodealter/go_blog/common"
	"github.com/jodealter/go_blog/dao"
	"github.com/jodealter/go_blog/models"
	"github.com/jodealter/go_blog/service"
	"github.com/jodealter/go_blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户id
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)

		categoryId, _ := strconv.Atoi(cId)

		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		service.SavePost(post)
		common.Success(w, post)

		//put代表更新
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)

		categoryId, _ := strconv.Atoi(cId)

		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidfloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidfloat)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		service.UpdatePost(post)
		common.Success(w, post)
		//put代表更新
	}
}
func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")

	//去除后缀
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此路径"))
		return
	}
	post, err := dao.GetPostByid(pId)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}
func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchRsp := service.SearchPost(condition)
	common.Success(w, searchRsp)
}
