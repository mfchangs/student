package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	isDir()
}

func isDir()  {
	fmt.Println(filepath.Abs("../tmp"))
}