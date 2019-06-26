package main

import (
	"fmt"
	"net/http"
	"sync"
)


func main(){
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.baidu.com",
		"https://www.toutiao.com",
		"https://www.jianshu.com",
	}

	// Increment the WaitGroup counter.
	wg.Add(5)
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		go func (wg *sync.WaitGroup,url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			res,_:=http.Get(url)
			fmt.Println(res)
		}(&wg,url)

		go func (wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(123213213)
		}(&wg)

		go func (wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(325235)
		}(&wg)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

