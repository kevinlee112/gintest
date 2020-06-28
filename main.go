package main

import (
	"context"
	"fmt"
	"gintest/app/models"
	"gintest/config"
	"gintest/routes"
	"gintest/util/gredis"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @license.name MIT
func main()  {
	// Init Config
	config.InitConfig()

	// db 初始化
	models.InitDB()

	// redis 初始化
	gredis.InitRedis()


	g := gin.New()

	debugMode(g)

	// router register
	routes.Register(g)

	// 启动
	fmt.Println("|-----------------------------------|")
	fmt.Println("|               gin-api             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port" + config.AppConfig.Addr + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	server := &http.Server{
		Addr         : config.AppConfig.Addr,
		Handler      : g,
		ReadTimeout  : config.AppConfig.ReadTimeout * time.Second,
		WriteTimeout :  config.AppConfig.WriteTimeout * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}


// 只有非 release 时才可用该函数
func debugMode(g *gin.Engine) {
	if config.AppConfig.RunMode == "release" {
		return
	}
	// 性能分析 - 正式环境不要使用！！！
	pprof.Register(g)
}
