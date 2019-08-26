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
	program := Getprogram(file)
	Creatfile(program)
}
func Creatfile(program [][]string) {
	for i := 1; i < len(program); i++ { //i=檔案數 j=程式的行
		file, err := os.Create(program[i][0]) //建立第一個陣列的檔案
		if err != nil {
			return
		}
		defer file.Close()
		//逐漸將每一個陣列印到檔案裏面
		for j := 0; j < len(program[i]); j++ {
			file.WriteString(program[i][j] + "\n")

		}
	}
}

func Getprogram(file []string) [][]string {
	var allprogram [][]string
	i, j := 0, 0 //i表示紀錄上一次被輸入的行數 j表示當前行數
	for j < len(file) {
		if strings.IndexAny(file[j], "O") == 0 { 
			allprogram = append(allprogram, file[i:j]) //將從上一次輸入到目前為止的程式碼丟進去
			i = j //紀錄當前行數為上次輸入
		}
		j++
	}
	allprogram = append(allprogram, file[i:j]) //將最後一段也丟進去
	return allprogram
}

func Getfile() string {
	fmt.Println("Please enter filename: ")
	var URL string
	fmt.Scanln(&URL)
	return URL
}

func Readfile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
