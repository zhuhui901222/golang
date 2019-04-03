package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)



type T int

func main() {

	dataCh := make(chan T, 100)
	toStop := make(chan string)
	stopCh := make(chan T)

	//简约版调度器
	go func() {
		if t, ok := <-toStop; ok {
			log.Println(t)
			close(stopCh)
		}
	}()
	//生产者
	for i := 0; i < 30; i++ {
		go func(i int) {
			for {
				id := strconv.Itoa(i)
				value := T(rand.Intn(10000))
				if value == 9999 {
					select {
					case toStop <- "sender sender sender sender sender sender sender sender # id:" + id + "to close":
					default:

					}
				}

				select {
				case <-stopCh:
					return
				default:

				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:

				}
			}

		}(i)
	}
	//消费者
	for i := 0; i < 20; i++ {
		go func(i int) {
			id := strconv.Itoa(i)
			for {
				select {
				case <-stopCh:
					return
				default:

				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == 9998 {
						select {
						case toStop <- "receiver# id:" + id + "to close":
							fmt.Println(id,"9998")
						default:

						}
					}
					log.Println("from:",id,"	receiver value :", value)
				}
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}

