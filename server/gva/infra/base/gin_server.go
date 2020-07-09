package base

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gva/infra"
	"gva/middleware"
	"net/http"
	"time"
)

var ginEngine *gin.Engine

var httpSrvHandler *http.Server

func Gin() *gin.Engine {
	return ginEngine
}

func HttpSrvHandler() *http.Server {
	return httpSrvHandler
}

type GinServerStarter struct {
	infra.BaseStarter
}

func (p *GinServerStarter) Init(ctx infra.StarterContext) {
	//gin application实例
	ginEngine = initGin()
	//日志组件配置和扩展
	//主要中间件配置：recover,日志输出中间件的自定义
	ginEngine.Use(middleware.RequestLog())
	ginEngine.Use(middleware.RecoveryMiddleware())
}


func (p *GinServerStarter) Start(ctx infra.StarterContext) {
	port := ctx.Props().GetDefault("app.server.port", "8082")
	address := fmt.Sprintf(":%s", port)
	httpSrvHandler = &http.Server{
		Addr:           address,
		Handler:        Gin(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("HttpServerRun:", address)
	if err := httpSrvHandler.ListenAndServe(); err != nil {
		log.Fatalf(" HttpServerRun:%s err:%v\n", address, err)
	}
}

func (p *GinServerStarter) StartBlocking() bool {
	return true
}

func (p *GinServerStarter) Stop(starterCtx infra.StarterContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler().Shutdown(ctx); err != nil {
		log.Fatalf("HttpServerStop err:%v\n", err)
	}
	log.Info("HttpServerStop stopped")
}

func initGin() *gin.Engine {
	app := gin.New()

	return app
}
