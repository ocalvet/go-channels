package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan bool, 5)
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
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
	fmt.Println("Doing some processing")
	t := rand.Intn(8)
	time.Sleep(time.Duration(t) * time.Second)
}
