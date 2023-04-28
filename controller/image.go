package controller

import (
	"github.com/KM911/oslib"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strconv"
	"tiu/dao"
)

type Image struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func SaveImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	filedata, err := file.Open()
	if err != nil {
		return
	}
	defer filedata.Close()
	// 从header中读取信息 Content-Length
	length, _ := strconv.Atoi(c.GetHeader("Content-Length"))
	var buffer = make([]byte, length)
	filedata.Read(buffer)
	file.Filename = dao.SaveImage(file.Filename, buffer, length)
	c.String(200, viper.GetString("upload.host")+"image/"+file.Filename)
}

func GetImage(c *gin.Context) {
	filename := c.Param("filename")
	data := dao.FindImage(filename)
	filetype := path.Ext(filename)[1:]
	if data == nil {
		c.String(200, "not found")
		return
	}
	// 这里需要维护一个就是mine 不然无法获取到正确的文件类型
	c.Data(200, "image/"+filetype, data)
}

func DelImage(c *gin.Context) {
	filename := c.Param("filename")
	dao.DeleteImage(filename)
}

func BackUp(c *gin.Context) {
	// 这里是一个post请求
	key := c.PostForm("key")
	if key == viper.GetString("backup.key") {
		// 返回我们的image.db文件
		c.File(filepath.Join(oslib.ExecutePath(), "dao", "image.db"))
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "密钥错误",
		})
	}
}
