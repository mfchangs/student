package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//var 变量
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

//GetCurPath 获取当前路径
func GetCurPath() string {
	getwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return getwd
}

//FileAbs 把相对路径变成绝对路径
func FileAbs(filePath string) string {
	if !filepath.IsAbs(filePath) {
		filePath, _ = filepath.Abs(filePath)
	}
	return filePath
}

//FileExists 判断路径是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//FileType 判断路径类型
func FileType(filePath string) string {
	_type, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}
	if _type.IsDir() {
		return "dir"
	}
	return "file"
}

//ReadInput 接受输入
func ReadInput(filepath string) string {
	for {
		fmt.Printf("%s exists override? y/n   ", filepath)
		r := bufio.NewReader(os.Stdin)
		input, _, _ := r.ReadLine()
		inputContext := strings.TrimSpace(string(input))
		fmt.Printf("input: %s", inputContext)
		if strings.ToLower(inputContext) == "y" || strings.ToLower(inputContext) == "n" {
			return inputContext
		}
	}
}

//MobileAction 执行文件移动动作
func MobileAction(src, dst string) error {
	srcContext, err := os.Open(src)

	if err != nil {
		errContext := fmt.Errorf("error: read %s failed, %s", src, err)
		return errContext
	}
	defer srcContext.Close()

	dstContext, err := os.Create(dst)
	if err != nil {
		errContext := fmt.Errorf("error: read %s failed, %s", dst, err)
		return errContext
	}
	defer dstContext.Close()

	_, err = io.Copy(dstContext, srcContext)
	if err != nil {
		errContext := fmt.Errorf("error: copy src %s to %s failed, %s", src, dst, err)
		return errContext
	}
	return nil

}

//CopyFile 复制
func CopyFile(src, dst string) {

	if FileExists(dst) {
		if !FORCE {
			if ReadInput(dst) == "n" {
				return
			}
		}
		if BACKUP {
			backupDst := dst + "_" + SUFFIX
			err := MobileAction(dst, backupDst)
			if err != nil {
				panic(err)
			}
		}
	}

	if DETAIL {
		fmt.Printf("%s -> %s", src, dst)
	}

	err := MobileAction(src, dst)
	if err != nil {
		panic(err)
	}
}

//Judge 判断
func Judge(src []string, dst string) {

	// 判断源文件或目录是否存在
	for i := 0; i < len(src); i++ {
		if !FileExists(src[i]) {
			err := fmt.Errorf("not exists src file %s", src[i])
			fmt.Println(err)
			os.Exit(100)
		}
	}

	if len(src) != 1 {
		if FileType(dst) != "dir" {
			fmt.Printf("%s not is direcoty", dst)
		}
	} else if len(src) == 1 {
		if strings.HasSuffix(dst, "/") {
			dst = dst[:len(dst)-1]
			if !FileExists(dst) {
				fmt.Printf("not exists file %s", dst)
				os.Exit(100)
			}
			if FileType(dst) != "dir" {
				fmt.Printf("%s not is direcoty", dst)
				os.Exit(100)
			}
		} else {
			dstDir := filepath.Dir(dst)
			if !FileExists(dstDir) {
				fmt.Printf("%s not exists", dstDir)
			}
			if FileType(dstDir) != "dir" {
				fmt.Printf("%s not is direcoty", dstDir)
				os.Exit(100)
			}
		}
	}
}

//Help 显示帮助
func Help() {
	flag.Usage()
}

//Start 开始支持
func Start(files []string) {
	fileNumber := len(files)
	dstFile := FileAbs(files[fileNumber-1])
	srcFiles := files[:fileNumber-1]

	for i := 0; i < len(srcFiles); i++ {
		srcFiles[i] = FileAbs(srcFiles[i])
	}

	Judge(srcFiles, dstFile)

	for i := 0; i < len(srcFiles); i++ {
		if FileType(srcFiles[i]) == "file" {
			CopyFile(srcFiles[i], dstFile)
		} else {

		}
	}
}

//Init 参数初始化
func Init() {
	if FORCE && KEEP {
		KEEP = false
	}
	if BACKUP && SUFFIX == "" {
		SUFFIX = time.Now().Format("200601021504")
	}

	if !BACKUP && SUFFIX != "" {
		SUFFIX = ""
	}
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

	Init()
	Start(flag.Args())
}
