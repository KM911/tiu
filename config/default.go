package config

import (
	"github.com/KM911/oslib"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	ExecutePath = ""
)

func InitConifg() {
	// 对于共同的配置进行设置
	ExecutePath = oslib.ExecutePath()
	configFile := filepath.Join(ExecutePath, "data", "default.json")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		println("配置文件不存在 生成配置文件中 ...  ")
		os.Mkdir(filepath.Join(ExecutePath, "data"), os.ModePerm)
		os.Mkdir(filepath.Join(ExecutePath, "data", "webp"), os.ModePerm)
		os.Create(filepath.Join(ExecutePath, "data", "default.json"))
		viper.Set("model", 0)
		viper.Set("server.port", 3000)
		viper.Set("upload.host", "http://81.68.91.70/pg/")
		viper.Set("upload.clean", false)
		viper.Set("upload.convent", false)
		viper.Set("upload.clip", true)
		viper.Set("backup.key", "")
		viper.WriteConfig()
		viper.WatchConfig()
	}
	Model(viper.GetInt("model"))
}

// 开发者设置
func Model(mod int) {
	switch mod {
	case 1:
		// 开发者一
		// 设置端口转发
		viper.Set("upload.host", "http://localhost:3000/pg/")
	case 2:
		// 开发者二
	case 3:
		// 开发者三
	default:
		// 默认设置 也就是用户设置
		// 在我看来是很好的一种设计方式

	}
}
