package controller

import (
	"github.com/gin-gonic/gin"
	"wine/model"
	"strconv"
	"time"
	"log"
	"net/http"
)

func AddFans(c *gin.Context)  {
	var p model.Fans
	//get请求 c.Query post请求 c.PostForm 传的是一个数组的话c.QueryMap
	//var firstname = c.Query("name")
	//var lastname = c.Query("lastname")
	name := c.PostForm("name")
	address := c.PostForm("address")
	//var ctime = c.PostForm("ctime")
	//将得到的值传给结构体
	p.Name = name
	p.Address = address

	id, _ := p.AddFans()
	p.Id = int(id)
	c.JSON(200,gin.H{
		"id":p.Id,
		"status":true,
		"code":"000",
		"message":"新增粉丝成功",
	})
}

func GetFans(c *gin.Context)  {
	var p model.Fans
	//id := c.Param("id")  //路由中的id
	id := c.Query("id")  //路由中的id
	p.Id, _ = strconv.Atoi(id) //字符串转数字
	fan := p.GetFans()

	//将要显示的结果存入map
	fans := make(map[string]interface{})
	fans["ctime"] = time.Unix(int64(fan.Ctime), 0).Format("2006-01-02 15:04:05")  //将int时间转换为字符串时间
	fans["Id"] = p.Id
	fans["name"] = fan.Name
	fans["address"] = fan.Address

	c.JSON(200,gin.H{
		"code":"000",
		"status":true,
		"message":"获取粉丝成功",
		"data":fans,
	})
}

func GetAllFans(c *gin.Context)  {
	var p model.Fans
	fans,_ := p.GetAllFan()
	datas := make([]map[string]interface{},0) //创建切片，要指定长度

	//遍历结果集
	for _,v := range fans {
		data := make(map[string]interface{})
		data["name"] = v.Name
		data["ctime"] = time.Unix(int64(v.Ctime), 0).Format("2006-01-02 15:04:05")
		datas = append(datas,data)
	}

	c.JSON(200,gin.H{
		"code":"000",
		"status":true,
		"message":"获取粉丝成功",
		"data":datas,
	})
}

//删除粉丝
func DelFan(c *gin.Context)  {
	var p model.Fans
	var id = c.Query("id")
	p.Id,_ = strconv.Atoi(id)
	num, _ := p.DelFans()      //判断下 也可以不判断
	if num == 0 {
		c.JSON(200,"请不要重复删除")
	} else {
		c.JSON(200,"ok")
	}
}

//更新粉丝
func UpdateFan(c *gin.Context)  {
	var p model.Fans
	id := c.Query("id")
	name := c.Query("name")
	address := c.Query("address")
	p.Id,_ = strconv.Atoi(id)
	p.Name = name
	p.Address = address
	_, err := p.UpdateFan()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":"000",
		"status":true,
		"message":"更新成功",
	})
}
