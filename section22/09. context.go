package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()

	ctx0, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Println("sleep Start = ", time.Now())
		time.Sleep(time.Second * 10)
		fmt.Println("sleep End = ", time.Now())
		cancel()
	}()

	var i int

	for {
		select {
		case <-ctx0.Done():
			return
		default:
			fmt.Println("======= ", i)
			i += 1
			time.Sleep(time.Second * 2)
		}
	}

}
