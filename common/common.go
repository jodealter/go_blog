package common

import (
	"encoding/json"
	"github.com/jodealter/go_blog/config"
	"github.com/jodealter/go_blog/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

//通用文件

var Template models.HTMLTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var param map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &param)
	return param
}
func Success(w http.ResponseWriter, data interface{}) {

	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultjson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write(resultjson)
	if err != nil {
		log.Println(err)
	}
}
func Error(w http.ResponseWriter, err error) {

	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultjson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(resultjson)
	if err != nil {
		log.Println(err)
	}
}
