package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan bool, 100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		ch <- true
		wg.Add(1)
		go func() {
			defer wg.Done()
			printAndWait()
			<-ch
		}()
	}
	close(ch)
	wg.Wait()
}

func printAndWait() {
	fmt.Println("Doing ...")
	time.Sleep(2 * time.Second)
}
