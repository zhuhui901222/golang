package main

func main(){

	Try(func() {
		panic("foo")
	}, func(e interface{}) {
		print(e)
	})

}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

