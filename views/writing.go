package views

import (
	"github.com/jodealter/go_blog/common"
	"github.com/jodealter/go_blog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
