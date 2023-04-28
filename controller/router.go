package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// 这里我们全部将图片文件放到sqlite数据库中，所以不需要进行静态文件的处理

	apiRouter := r.Group("/pg")
	apiRouter.GET("/test", Test)
	apiRouter.GET("/ip", IP)
	apiRouter.POST("/upload", SaveImage)
	// 这里假如是 这样的形势  /api/1.png 那么我们就可以使用这个方法
	apiRouter.GET("/image/:filename", GetImage)
	apiRouter.GET("/del/:filename", DelImage)
	// 理论上来说 这里肯定是需要加密的 不然任何人都可以自由得将你得文件内容进行一个上传了就是不是吗?用户鉴权 我们可以放到后面去做
	apiRouter.POST("/backup", BackUp)
}
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func IP(c *gin.Context) {
	c.JSON(200, gin.H{
		"IP": c.RemoteIP(),
	})
}
