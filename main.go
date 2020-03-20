package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	value := produce(done, 3)

	for {
		v, ok := <-value
		if !ok {
			break
		}

		fmt.Printf("%d", v)
	}

	fmt.Println()
}

func produce(done <-chan struct{}, num int) <-chan int {
	valueStream := make(chan int)

	go func() {
		defer close(valueStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case valueStream <- 6:

			}
		}
	}()

	return valueStream
}
