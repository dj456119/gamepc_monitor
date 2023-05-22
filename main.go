package main

import (
	http_controller "gamepc_monitor/controller/http"
	"gamepc_monitor/dlog"
)

func main() {
	dlog.Init(dlog.LOG_LEVEL_DEBUG, dlog.OUT_TYPE_STDOUT)
	dlog.Info("Starting GamePC Monitor...")
	go http_controller.Init()
	select {}
}
