package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)


var fileName = "./muke.txt"
var file *os.File
var err error

func main() {

	file := openFile(fileName)

	url := "http://www.imooc.com/u/"

	// Instantiate default collector
	c := colly.NewCollector(

		)

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		4, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 1000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		fmt.Println(reflect.TypeOf(r.URL))
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("#main > div.bg-other.user-head-info > div > div.user-info-right > h3 > span", func(e *colly.HTMLElement) {

		fmt.Printf("Image found:  %s\n", e.Text )

		writeFile(file,  e.Text)
	})

	c.OnHTML("#main > div.bg-other.user-head-info > div > div.user-pic > div > img", func(e *colly.HTMLElement) {

		src:=e.Attr("src")
		// Print src

		fmt.Printf("Image found:  %s\n", src )

		writeFile(file, src+"\n")
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})


	for i := 4676745; i < 4676750; i++ {
		// Add URLs to the queue
		q.AddURL(fmt.Sprintf("%s/%d/courses", url, i))
	}
	// Consume URLs
	q.Run(c)

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