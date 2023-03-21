package service

import (
	"github.com/jodealter/go_blog/config"
	"github.com/jodealter/go_blog/dao"
	"github.com/jodealter/go_blog/models"
	"html/template"
)

func GetAllIndexInfo(slug string, page, pagesize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	var total int
	if slug == "" {

		total = dao.CountGetAllPost()
		posts, err = dao.GetPostPage(page, pagesize)
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pagesize)
		total = dao.CountGetAllPostBySlug(slug)
	}

	var postMores []models.PostMore
	for _, post := range posts {

		categoryName := dao.GetCategoryNameById(post.CategoryId)
		UserName := dao.GetUserNameById(post.UserId)

		//这个rune是转换中文格式，每个中文对应一个1然后从0到100的额时候，才是切100个字
		context := []rune(post.Content)
		if len(context) > 100 {
			context = context[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(context),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     UserName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	return hr, nil
}
