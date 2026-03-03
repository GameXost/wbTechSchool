package main

import (
	"fmt"
)

func Workers(work int) chan int {
	ch := make(chan int)
	for i := 0; i < work; i++ {
		go func() {
			for n := range ch {
				fmt.Println(n)
			}

		}()
	}
	return ch
}
func main() {
	var cnt int
	fmt.Scan(&cnt)
	ch := Workers(cnt)
	defer close(ch)
	i := 0
	for {
		ch <- i
		i++
	}
}
