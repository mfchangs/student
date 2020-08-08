package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


type Rename struct {
	help bool
	detail bool
	force bool
	recursion bool
	link bool
	hardLink bool
	keep bool
	backup string
}

var (
	BACKUP    bool
	HELP      bool
	KEEP      bool
	HARDLINK  bool
	RECURSION bool
	FORCE     bool
	DETAIL    bool
	LINK      bool
	SUFFIX    string
)

func GetCurPath() string {
	getwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return getwd
}

func FileAbs(filePath string) string {
	if ! filepath.IsAbs(filePath) {
		filePath, _ = filepath.Abs(filePath)
	}
	return filePath
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func CopyFile(src, dst string) {

}


func Start(files []string) {
	fileNumber := len(files)
	dstFile := FileAbs(files[fileNumber - 1])
	srcFiles := files[:fileNumber - 1]

	for i := 0; i < len(srcFiles); i++ {
		srcFiles[i] = FileAbs(srcFiles[i])
	}

	Judge(srcFiles, dstFile)

	for i := 0; i < len(srcFiles); i++ {
		CopyFile(srcFiles[i], dstFile)
	}
}


func Judge(src []string, dst string) {
	// 判断动作

	for i := 0; i < len(srcFiles); i++ {
		if ! fileExists(srcFiles[i]) {
			err := fmt.Errorf("not exists file %s", srcFiles[i])
			fmt.Println(err)
			os.Exit(100)
		}
	}

	if strings.HasPrefix(dst, "/") {
		dst = dst[:len(dst) - 1]
		_, err := os.Stat(dst)
		if err != nil {
			err := os.Mkdir(dst, 766)
			if err != nil {
				panic(err)
			}
		}
	}

	dstType, err := os.Stat(dst)
	if err != nil {
		if len(src) != 1 {
			fmt.Printf("目录路径不存在")
			os.Exit(10)
		}
	}
	if len(src) != 1 && ! dstType.IsDir() {
		fmt.Printf("源文件或目录有多个，但目录路径不是目录")
		os.Exit(10)
	}

}


func Help()  {
	flag.Usage()
}

func main() {

	flag.BoolVar(&DETAIL, "d", false, "显示复制源文件和目标文件")
	flag.BoolVar(&FORCE, "f", false, "强制覆盖目标文件")
	flag.BoolVar(&HELP, "h", false, "查看帮助")
	flag.BoolVar(&RECURSION, "r", false, "递归复制目录")
	flag.BoolVar(&LINK, "l", false, "复制为软链接")
	flag.BoolVar(&HARDLINK, "hl", false, "复制为硬链接")
	flag.BoolVar(&KEEP, "k", false, "保留目标文件, -f时无效")
	flag.BoolVar(&BACKUP, "b", false, "备份目标文件, 备份后的文件名是dst_filename_20200805")
	flag.StringVar(&SUFFIX, "s", "", "指定备份目标文件的后缀, 与-b配合使用")

	flag.Parse()

	if flag.NArg() < 2 {
		Help()
		os.Exit(1)
	} else if HELP {
		Help()
		os.Exit(0)
	}

	Start(flag.Args())
}
