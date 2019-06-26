package main

//取最快的结果
func main() {
	var chan_a,chan_b,chan_def chan int
	select {
	case <-chan_a: // 希望cond_a为真时才在chan_a上等待
		// do something
	case <-chan_b: // 希望cond_b为真时才在chan_b上等待
		// do something
	case <-chan_def:
		// do something
	}
}