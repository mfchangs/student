package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("config.ini")
	getErr("load config", err)

	// 遍历所有的section
	for _, v := range cfg.Sections() {
		fmt.Println("v: ", v.Name())
		fmt.Println(v.KeyStrings())
	}

	fmt.Println(cfg.Section("name:ping").Key("bin").String())
}

func getErr(msg string, err error) {
	if err != nil {
		log.Printf("%v err->%v\n", msg, err)
	}
}
