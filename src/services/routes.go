package services

import (
	"fmt"
	"go-go-go/src/data"
	"go-go-go/src/services/test"
	"go-go-go/src/taobao"
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

	engine.GET("/tb/:itemId", func(c *gin.Context) {
		itemId := c.Param("itemId")
		code := taobao.GetItemCode(itemId, data.GetConfig("taobao_cookie"))
		//c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")

		c.Header("Content-Type", "text/html; charset=utf-8")
		html := fmt.Sprintf("<html><body><a href=%s target=_blank>点击跳转至淘宝商品页</a>  </br> <p>或复制代码 %s 至淘宝APP</p></body> </html>",
			code.Data.ShortLinkInfo.CouponUrl, code.Data.TaoTokenInfo.CouponUrl)
		c.String(200, html)
	})
}

func includeRoutes(group *gin.RouterGroup, routes map[string]map[string]gin.HandlersChain) {
	for url, methodHandlerChain := range routes {
		for method, handlerChain := range methodHandlerChain {
			group.Handle(method, url, handlerChain...)
		}
	}
}
