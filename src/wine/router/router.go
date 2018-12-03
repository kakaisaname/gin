package router

import (
	"github.com/gin-gonic/gin"
	. "wine/controller"   //Notice: ".", 点操作 导入包后 调用这个包的函数时 可以省略包名
)

func InitRouter() *gin.Engine  {
	router := gin.Default()

	router.GET("async",Async)
	router.GET("pong",Pong)

	//添加fans
	router.POST("/fan",AddFans)
	router.GET("/fans",GetFans)
	router.GET("/allFans",GetAllFans)
	router.DELETE("/fan",DelFan)
	router.PUT("/fan",UpdateFan)

	return router
}