package main



// Sendthe sequence 2, 3, 4, ... to channel 'ch'.

func Generate(ch chan<- int) {

	for i := 2;; i++ {
		ch<- i // Send 'i' to channel 'ch'.
	}

}

// Copythe values from channel 'in' to channel 'out',

//removing those divisible by 'prime'.

func Filter(in <-chan int, out chan<- int, prime int) {

	for {
		i := <-in // Receive valuefrom 'in'.
		if i%prime != 0 {
			out <- i // Send'i' to 'out'.
		}
	}

}

// Theprime sieve: Daisy-chain Filter processes.

func main() {

	ch := make(chan int) // Create a newchannel.

	go Generate(ch) // Launch Generate goroutine.

	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

}
