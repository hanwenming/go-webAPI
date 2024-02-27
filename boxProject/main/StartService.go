package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-webAPI/tools"
	"log"
	"net/http"
	"os"
)

func main() {

	//err := tools.InitMysql()
	//if err != nil {
	//	os.Exit(1)
	//}

	ip, err := tools.GetHostIp()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("当前主机IP : ", ip)

	// 创建 Gin 的默认引擎
	r := gin.Default()

	// 定义一个 GET 请求的路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Welcome to the API!",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 运行 HTTP 服务器，监听在 8080 端口
	err = r.Run(":8080")
	if err != nil {
		os.Exit(1)
	}
}
