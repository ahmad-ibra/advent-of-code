package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNumbers(c chan<- int) {
	defer close(c)
	defer wg.Done()

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		c <- i
	}

	fmt.Println("Closing numChan")
}

func putWords(c chan<- string) {
	defer close(c)
	defer wg.Done()

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		c <- fmt.Sprintf("hello %v", i)
	}

	fmt.Println("Closing strChan")
}

func readNumbers(c <-chan int, cur int) {
	defer wg.Done()
	for val := range c {
		fmt.Printf("read %v from goroutine %v\n", val, cur)
	}
}

func readFromMultipleChannels(numChan <-chan int, strChan <-chan string) {
	defer wg.Done()

	finishedNum := false
	finishedStr := false

	for {
		if finishedNum && finishedStr {
			break
		}
		select {
		case val, ok := <-numChan:
			if !ok {
				fmt.Println("numChan closed")
				finishedNum = true
			} else {
				fmt.Printf("read %v\n", val)
			}

		case val, ok := <-strChan:
			if !ok {
				fmt.Println("strChan closed")
				finishedStr = true
			} else {
				fmt.Printf("read %v\n", val)
			}
		}
	}
}

func main() {
	// ------------------------------------------
	// Multiple goroutines reading from 1 channel
	// ------------------------------------------
	numChan := make(chan int, 1)
	wg.Add(1)
	go putNumbers(numChan)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readNumbers(numChan, i)
	}
	wg.Wait()
	fmt.Println("DONE!")

	// --------------------------------------------
	// One gorouting reading from multiple channels
	// --------------------------------------------
	numChan2 := make(chan int)
	strChan := make(chan string)

	wg.Add(3)
	go putNumbers(numChan2)
	go putWords(strChan)
	go readFromMultipleChannels(numChan2, strChan)
	wg.Wait()

	fmt.Println("DONE!")
}
