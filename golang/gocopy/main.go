package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}


func copyFileAction(src, dst string, showProgress, force bool) {
	if force {
		if fileExists(dst) {
			fmt.Printf("%s exists override ? y/n   ", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			if strings.TrimSpace(string(data)) != "y" {
				return
			}
		}
 	}
 	copyFile(src, dst)

	if showProgress {
		fmt.Printf("%s -> %s\n", src, dst)
	}
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	destFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()
	return io.Copy(destFile, srcFile)
}

func main() {
	var help, detail, force, recursion, link, hardLink, keep, backup bool
	var suffix string
	flag.BoolVar(&detail, "d", false, "显示复制源文件和目标文件")
	flag.BoolVar(&force, "f", false, "强制覆盖目标文件")
	flag.BoolVar(&help, "h", false, "查看帮助")
	flag.BoolVar(&recursion, "r", false, "递归复制目录")
	flag.BoolVar(&link, "l", false, "复制为软链接")
	flag.BoolVar(&hardLink, "hl", false, "复制为硬链接")
	flag.BoolVar(&keep, "k", false, "保留目标文件, -f时无效")
	flag.BoolVar(&backup, "b", false, "备份目标文件, 备份后的文件名是dst_filename_20200805")
	flag.StringVar(&suffix, "s", "", "指定备份目标文件的后缀, 与-b配合使用")

	flag.Parse()
	flag.Visit(test())
	//if flag.NArg() < 2 {
	//	flag.Usage()
	//	return
	//}
	flag.Usage()
	fmt.Println("args: ", flag.Args())

	//copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)
}

func test(testArgs flag.Flag)  {
	fmt.Println(testArgs)
}