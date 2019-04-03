package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main(){
	var start,end int
	fmt.Print("请输入起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入起始页(>=起始页):")
	fmt.Scan(&end)
	DoWork(start,end)
}

func DoWork(start,end int){
	fmt.Printf("正在爬取 %d 到 %d 的页面",start,end)
	page:=make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i,page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个网页爬取完成",<-page)
	}
}

func SpiderPage(i int,page chan int)  {
	url:="https://www.dto.jp/gal/"+strconv.Itoa(i)
	result,err:=HttpGet(url)
	if err!=nil{
		fmt.Println("HttpGet error=",err)
		page<-i
		return
	}

	fmt.Print("len=",len(result))

	if len(result)<3000{
		return
	}

	//写入文件
	fileName:=strconv.Itoa(i)+".html"
	f,err2:=os.Create(fileName)
	if err2 !=nil {
		fmt.Println("Http get Err=",err2)
		page<-i
		return
	}

	f.WriteString(result)
	f.Close()
	page<-i
}

func HttpGet(url string)(result string, err error)  {
	resp,err1:=http.Get(url)
	if  err1!=nil{
		err=err1
		return
	}
	defer resp.Body.Close()


	buf:=make([]byte,1024*4)
	for{
		n,err:=resp.Body.Read(buf)
		if n==0{
			fmt.Println("resp.Body.Read err=",err)
			break
		}
		result+=string(buf[:n])
	}
	return


}


