package Util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Hi() {
	fmt.Println("hi")
}

func ReadFile() ([]string){
	fi, _ := os.Open("./example/my.log")

	defer fi.Close()

	var lines []string
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines,string(a))
		//fmt.Println("M:"+string(a))
	}

	return lines
}