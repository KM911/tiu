package dao

import (
	"github.com/KM911/oslib"
	"path"
	"strings"
)

/*
保存图片 如果长度相同 则更新 如果长度不同 则插入
大小不同则一定是不同的文件
大小相同不一定是相同的文件 还需要去比对原始的文件名 这里我还是推荐你使用原始的文件名而不是时间戳
当然了你返回时间肯定是可以的
*/
func SaveImage(filename string, data []byte, length int) string {
	var Image ImageTable
	db.Find(&Image, "length = ?", length)
	if Image.ID != 0 {
		db.Model(&Image).Update("data", data)
		return Image.Filename
	}
	if !IsTIUFormat(filename) {
		filename = "KM" + oslib.RandomStringName(10) + path.Ext(filename)
	}
	db.Create(&ImageTable{Filename: filename, Data: data, Length: length})
	return filename
}

func DeleteImage(filename string) {
	db.Delete(&ImageTable{}, "filename = ?", filename)
}

func FindImage(filename string) []byte {
	var image ImageTable
	db.Find(&image, "filename = ?", filename)
	if image.ID != 0 {
		return image.Data
	} else {
		return nil
	}
}

/*
lens and
*/
func IsTIUFormat(file_ string) bool {
	filename := oslib.FileName(file_)
	if len(filename) == 12 && strings.HasPrefix(filename, "KM") {
		return true
	}
	return false
}
