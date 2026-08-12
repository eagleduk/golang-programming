package main

import "fmt"

func main() {
	even := make(chan int)
	oddn := make(chan int)
	quit := make(chan int)

	go send(even, oddn, quit)

	recive(even, oddn, quit)
}

func recive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("======== EVEN:", v)
		case v := <-o:
			fmt.Println("-------- ODDN:", v)
		case v := <-q:
			fmt.Println("++++++++ QUIT: ", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
