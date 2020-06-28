package routes

import (
	"gintest/app/Controllers/api/blog_v1"
	"gintest/app/Controllers/test"
	_ "gintest/docs"
	"gintest/util/response"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func registerApi(g *gin.Engine) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // API 注释
	//测试路由
	g.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1,"pong", nil)
	})
	// 业务接口测试
	apiV1 := g.Group("/api/v1")
	{
		apiV1.GET("/blog/:id", blog_v1.Detail)
		apiV1.POST("/blog/add", blog_v1.Add)
	}



	// 测试加密性能
	TestRouter := g.Group("/test")
	{
		// 测试 MD5 组合 的性能
		TestRouter.GET("/sign", test.Md5Test)

		// 测试 AES 的性能
		TestRouter.GET("/aes", test.AesTest)

		// 测试 RSA 的性能
		TestRouter.GET("/rsa", test.RsaTest)
	}
}
