package views

import (
	"github.com/jodealter/go_blog/common"
	"github.com/jodealter/go_blog/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
