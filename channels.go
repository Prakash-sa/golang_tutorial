package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

/*
func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 31
		wg.Done()
	}()
	wg.Wait()
}

*/

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
