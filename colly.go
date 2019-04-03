package main



import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

var fileName = "./flag.txt"
var file *os.File
var err error

func main() {

	file := openFile(fileName)

	// Instantiate default collector

	c := colly.NewCollector()



	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org

	c.AllowedDomains = []string{"www.dto.jp"}


	c.OnHTML(".com_gal.style1 li .inner .item.image", func(e *colly.HTMLElement) {

		alt := e.Attr("alt")

		src:=e.Attr("style")
		// Print src

		fmt.Printf("Image found: %s -> %s\n", alt,src )

		writeFile(file, src+"\n")
	})

	// On every a element which has href attribute call callback

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		// Print link

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link found on page

		// Only those links are visited which are in AllowedDomains

		c.Visit(e.Request.AbsoluteURL(link))

	})






	// Before making a request print "Visiting ..."

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL.String())

	})



	// Start scraping on https://hackerspaces.org

	c.Visit("https://www.dto.jp/shop/12605/gals")

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