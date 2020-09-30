package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func test() {
	logfile, _ := os.Open("/tmp/test.log")
	shell := exec.Command("ping", "www.baidu.com")
	shell.Stdout = logfile
	err := shell.Start()
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	test()
	fmt.Print("pid: ", os.Getpid())
	time.Sleep(10000000000)
}
