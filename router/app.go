package router

import (
	"OJ_sandbox/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 执行路由
	r.GET("/test") // 后期可能用于集群判断
	r.POST("/code_execute", service.RunCodeByDocker)

	return r
}
