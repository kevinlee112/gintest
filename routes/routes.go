package routes

import (
	"gintest/middleware"
	"gintest/middleware/exception"
	"gintest/middleware/logger"
	"gintest/util/response"
	"github.com/gin-gonic/gin"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	// ---------------------------------- 注册全局中间件 ----------------------------------
	//自定义全局中间件
	g.Use(logger.SetUp(), exception.SetUp())
	g.Use(middleware.Cors())     // cors跨域
	// ---------------------------------- 注册路由 ----------------------------------
	//404
	g.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404,"请求方法不存在", nil)
	})
	// web
	registerWeb(g)
	// api
	registerApi(g)

	return g
}
