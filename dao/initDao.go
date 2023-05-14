// 这里是我们使用sqlite的部分

package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

type ImageTable struct {
	gorm.Model
	ID       uint64 `gorm:"primarykey"`
	Filename string `gorm:"type:varchar(20)"`
	Data     []byte `gorm:"type:blob"`
	Length   int
}

// 备份功能的实现就是为了解决这样的情况的
// 数据库中没有了
// 我们可以字节

var (
	db  *gorm.DB
	err error
)

func InitDao() {
	db, err = gorm.Open(sqlite.Open(filepath.Join("data", "image.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&ImageTable{})
	db.Raw("create index index_length on image_tables (id,length);")
	db.Raw("create index index_filename on image_tables (id,filename);")
}
