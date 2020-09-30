package options

import (
	"flag"
	"os"
)

var (
	LIST     bool
	RELOAD   bool
	RESTART  bool
	STOP     bool
	MSTOP    string
	MRESTART string
	HELP1    bool
)

// H 变量
var H = HelpOptions{}

// Options 选项
func Options() {
	flag.BoolVar(&HELP, "h", false, "查看帮助")
	flag.BoolVar(&LIST, "list", false, "显示进程监控状态")
	flag.BoolVar(&RELOAD, "reload", false, "重载mprocess进程配置")
	flag.BoolVar(&RESTART, "restart", false, "重启mprocess进程")
	flag.BoolVar(&STOP, "stop", false, "停止mprocess进程,默认pid文件是/var/run/mprocess.pid")
	flag.StringVar(&CONFIG, "config", "/etc/mprocess/mprocess.ini", "指定mprocess进程配置文件")
	flag.StringVar(&MSTOP, "mstop", "", "停止指定监控进程")
	flag.StringVar(&MRESTART, "mrestart", "", "重启指定监控进程")

	flag.Parse()
	if HELP {
		H.Help()
		os.Exit(0)
	}
}
