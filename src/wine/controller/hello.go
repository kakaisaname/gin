package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Hello(c *gin.Context)  {
	c.String(http.StatusOK,"hello")
}

func Test(c *gin.Context)  {
	c.String(http.StatusOK,"test")
}

func Async(c *gin.Context)  {
	var c_copy = c.Copy() //goroute 中请使用只读的Context,用Context.Copy可以返回一个只读的Context
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c_copy.Request.URL.Path)  //5秒后打印
	}()
	c.String(http.StatusOK,"ok")
}

func Pong(c *gin.Context)  {
	c.JSON(200,map[string]interface{}{"code":1,"msg":"pong"})
}