package main


type T int


//closed 1成功发送 0异常，发送失败
func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
		}
	}()
	ch <- value
	return true
}

//closed 1成功关闭 0异常 已关闭
func SafeClose(ch chan T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
		}
	}()
	close(ch)
	return true
}

func main()  {
	c1:=make(chan T ,1)
	SafeSend(c1,2434)
	SafeClose(c1)

}

