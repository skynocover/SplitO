package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	filename := Getfile()
	file := Readfile(filename)
	creatFile(file)
}

func creatFile(program []string) {
	var j = 0
	for j < len(program) {
		if strings.IndexAny(program[j], "O") == 0 { //找到開頭為O的行數
			file, err := os.Create(program[j]) //依照當前文字建立檔案
			if err != nil {
				return
			}
			defer file.Close()
			file.WriteString(program[j] + "\n") //寫入當前文字
			j++
			for strings.IndexAny(program[j], "O") != 0 { //不斷的寫入直到遇到下一次的O開頭
				file.WriteString(program[j] + "\n")
				j++
				if j == len(program) { //如果已經到最後一行則return
					return
				}
			}
		} else { //若不是則繼續找下一行
			j++
		}
	}
}

func Getfile() (filename string) {
	fmt.Println("Please enter filename: ")
	fmt.Scanln(&filename)
	return
}

func Readfile(filename string) (lines []string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines = strings.Split(string(content), "\r\n")
	if len(lines) == 1 { //如果不是以windowns的方式分行則只會有一段
		lines = strings.Split(string(content), "\n")
	}
	return
}
