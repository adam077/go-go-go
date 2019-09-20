package main

import (
	"go-go-go/src/services"
)

func main() {
	//scheduler.Run()
	engine := services.SetupEngine()
	engine.Run("0.0.0.0:8080")
	//engine.RunTLS("0.0.0.0:8080", certPath, keyPath)
}
