package tiny

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"tiu/config"
	"tiu/dao"
)

func Backup() {
	var iamgeList []dao.ImageTable
	db, err := gorm.Open(sqlite.Open(filepath.Join(config.ExecutePath, "data", "image.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Find(&iamgeList)
	os.Mkdir(filepath.Join(config.ExecutePath, "data", "image"), 0666)
	for _, v := range iamgeList {
		// 将其写入到文件中
		os.WriteFile(filepath.Join(config.ExecutePath, "data", "image", v.Filename), v.Data, 0666)
	}
}

//func DownloadImageDB() {
//	url := viper.GetString("upload.host") + "pg/backup"
//	key := config.Key
//	// 发送POST请求
//	// 这里是一个post请求
//	http.NewRequest("POST", url, nil)
//
//}
