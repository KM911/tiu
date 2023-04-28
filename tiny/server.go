package tiny

import (
	"github.com/spf13/viper"
	"strconv"
	"tiu/controller"
	"tiu/dao"

	"github.com/gin-gonic/gin"
)

func Server() {
	println("server start")
	// 关闭debug模式
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	controller.InitRouter(r)
	dao.InitDao()
	port := ":" + strconv.Itoa(viper.GetInt("server.port"))
	r.Run(port)
}
