package main

import (
	"go-go-go/src/data"
	"go-go-go/src/scheduler"
	"go-go-go/src/services"
)

func main() {
	if data.GetConfig(data.Scheduler) != "" {
		// 多实例的话感觉推荐用环境变量
		scheduler.Run()
	}
	engine := services.SetupEngine()
	engine.Run("0.0.0.0:8080")
}
