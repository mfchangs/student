package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func test()  {
	filename := "/tmp/fstab"
	open, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(open)
	for {
		readString, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		if strings.Index(readString, "UUID") != -1 {
			fmt.Printf(readString)
		}
	}
}

func main() {
	test()
}
