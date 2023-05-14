package main

import (
	"os"
	"oslib"
	"tiu/config"
	"tiu/tiny"
)

func main() {
	// 用于性能分析
	// f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	defer oslib.TimerStart().End()
	config.InitConifg()
	argvLens := len(os.Args)
	switch argvLens {
	case 1:
		// 需要进行测试就把代码放到这里来
		// tiny.Server()
		println("need a valid argv")
	case 2:
		switch os.Args[1] {
		case "server", "s":
			tiny.Server()
			println("start server")
		case "load", "l":
			tiny.Load(os.Args[2])
		case "backup", "b":
			tiny.Backup()
			println("backup your iamge")
		default:
			println("invalid  argv")
		}
	default:
		if os.Args[1] == "upload" || os.Args[1] == "u" {
			for i := 2; i < argvLens; i++ {
				tiny.Upload(os.Args[i])
			}
		} else {
			println("invalid  argv")
		}
	}
}
