package main

import (
	"bufio"
	"os"
)

var fileName = "./flag.txt"
var file *os.File
var err error

func main() {
	file = openFile(fileName)
	writeFile(file, "keep coding!!")
}

func openFile(fileName string) *os.File {
	if checkFileIsExist(fileName) {
		//如果文件存在
		file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		//创建文件
		file, err = os.Create(fileName)
	}
	check(err)
	return file
}

func writeFile(file *os.File, content string) {
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.WriteString("\r\n")
	writer.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}