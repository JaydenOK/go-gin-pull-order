package main

import (
	"fmt"
	"gin-pull-order/config"
	"gin-pull-order/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	engine.Run(":" + config.PORT)
	fmt.Println("Pull Order Task Running ...")
}
