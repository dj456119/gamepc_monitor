package http_controller

import (
	"gamepc_monitor/dlog"
	"gamepc_monitor/service"
	"net/http"
	"text/template"
)

var httpController HttpController

type HttpController struct {
	Service service.Service
}

func GetSensorModel(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/html/model.html")
	if err != nil {
		dlog.Error(err)
		return
	}

	data, err := httpController.Service.GetSensorModel()
	if err != nil {
		dlog.Error(err)
		return
	}
	dlog.Debug(data)
	err = tmpl.Execute(w, data)
	if err != nil {
		dlog.Error(err)
	}
}

func Init() {
	httpController = HttpController{
		Service: service.NewService(),
	}
	dlog.Info("Listening on :7777...")
	http.HandleFunc("/model", GetSensorModel)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		dlog.Fatal(err)
	}
}
