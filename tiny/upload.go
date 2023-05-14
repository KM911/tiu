package tiny

import (
	"bytes"
	"fmt"
	"github.com/KM911/oslib"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"tiu/config"
)

func Upload(file_ string) {
	file_ = DownloadFile(file_)
	file_ = ImageConvent(file_)
	Upload2Server(file_)
}

func DownloadFile(url string) string {
	if strings.Contains(url, viper.GetString("upload.host")) {
		println("图片已经上传")
		os.Exit(1)
	}
	// 如果是本地文件的话 我们直接返回
	if url[:4] != "http" {
		return url
	}
	// 根据上一次的经验 我们还需要写一个就是更加具有普世性的获取图片类型/名字的方法
	fileName := path.Base(url)
	filepath := path.Join(config.ExecutePath, "image", "download", fileName)
	return filepath
}

// 转换图片格式 这里考虑到就是其实未必是一个好的事情
func ImageConvent(file_ string) string {
	if viper.GetBool("upload.convent") {
		// 压缩命令
		ext := filepath.Ext(file_)
		switch ext {
		case ".png", ".jpg":
			output := filepath.Join(config.ExecutePath, "data", "webp", oslib.FileName(file_)+".webp")
			oslib.Run(path.Join(config.ExecutePath, "data", "cwebp.exe") + " -q 80 " + file_ + " -o " + output)
			os.Remove(file_)
			return output
		default:
			return file_
		}
	}
	return file_
}

/*
将图片上传到服务器
可以是任意格式的图片包括gif吗?
*/
func Upload2Server(file_ string) {
	file, err := os.Open(file_)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 设置body的buff大小
	var body bytes.Buffer
	//body.Grow(1024 * 16)
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", file_)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(part, file)
	writer.Close()
	req, err := http.NewRequest("POST", viper.GetString("upload.host")+"upload", &body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	data := make([]byte, 1024)
	n, err := resp.Body.Read(data)
	fmt.Println(string(data[:n]))
	if viper.GetBool("upload.clip") {
		oslib.Run("echo " + string(data[:n]) + " | clip")
	}
}
