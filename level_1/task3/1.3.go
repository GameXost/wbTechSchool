package main

import (
	"fmt"
)

func Workers(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	var cnt int
	fmt.Scan(&cnt)

	ch := make(chan int)
	defer close(ch)

	for i := 0; i < cnt; i++ {
		go func() {
			Workers(ch)
		}()
	}

	i := 0
	for {
		ch <- i
		i++
	}
}
