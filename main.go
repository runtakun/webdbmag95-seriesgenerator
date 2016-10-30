package main

import (
	"fmt"
	"sync"
)

func GenerateSeries(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < max; i++ {
			c <- i
		}
	}()
	return c
}

func ReadSeries(wg *sync.WaitGroup, name string, c <-chan int) {
	defer wg.Done()
	for i := range c {
		fmt.Printf("%s Read %d\n", name, i)
	}
}

func ExampleGenerator() {
	c := GenerateSeries(100)

	var wg sync.WaitGroup

	wg.Add(2)
	go ReadSeries(&wg, "a", c)
	go ReadSeries(&wg, "b", c)
	wg.Wait()
}

func main() {
	ExampleGenerator()
}
