package services

import (
	"go-go-go/src/services/test"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupEngine() *gin.Engine {
	engine := gin.New()
	engine.RedirectTrailingSlash = true
	if false {
		gin.SetMode(gin.ReleaseMode)
	}
	if true {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowCredentials = true
		corsConfig.AddAllowMethods(http.MethodPatch)
		corsConfig.AddAllowMethods(http.MethodDelete)
		corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Cache-Control", "Pragma")
		corsConfig.AllowOriginFunc = func(origin string) bool {
			return true
		}
		engine.Use(cors.New(corsConfig))
	}
	engine.Use(QueryMonitorMiddleware)
	registerRouters(engine)
	return engine
}

func registerRouters(engine *gin.Engine) {
	apiGroupLv1 := engine.Group("/lv1")
	apiGroupLv2 := apiGroupLv1.Group("/lv2")
	includeRoutes(apiGroupLv2, test.MonitorRoutes)

	engine.Static("/assets", "./src/assets")
	engine.StaticFS("/assets_list", http.Dir("src/assets"))
}

func includeRoutes(group *gin.RouterGroup, routes map[string]map[string]gin.HandlersChain) {
	for url, methodHandlerChain := range routes {
		for method, handlerChain := range methodHandlerChain {
			group.Handle(method, url, handlerChain...)
		}
	}
}
