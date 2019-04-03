package main


type T int

func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}

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

