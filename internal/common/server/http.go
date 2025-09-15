package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//启动HTTP服务
func RunHTTPServer(serviceName string , wapper func(router *gin.Engine)){
	addr := viper.Sub(serviceName).GetString("http-addr")
	RunHTTPServerOnAddr(addr,wapper)
}

func RunHTTPServerOnAddr(addr string , wapper func(router *gin.Engine)){
	apiRouter := gin.New()
	wapper(apiRouter)
	apiRouter.Group("/api")
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}