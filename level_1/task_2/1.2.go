package main

import (
	"fmt"
	"sync"
)

func squaresWaitGroup(nums []int) {
	var wg sync.WaitGroup // примитив синхронизации, как будто бы семафор, но инкрементирует и декрементирует наоборот
	for _, n := range nums {
		wg.Add(1) // ну или в самом начале wg.Add(len(nums))
		go func(num int) { // каждое число обрабатываем в отдельной горутине
			defer wg.Done() // декрементирует счётчик
			fmt.Println(num * num)
		}(n)
	}
	wg.Wait() // ждет, счётчик = 0
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squaresWaitGroup(nums)
}
