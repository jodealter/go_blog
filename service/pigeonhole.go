package service

import (
	"github.com/jodealter/go_blog/config"
	"github.com/jodealter/go_blog/dao"
	"github.com/jodealter/go_blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}
}
