package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 5)
	for i := 0; i < 10; i++ {
		ch <- true
		go func() {
			printAndWait()
			<-ch
		}()
	}
}

func printAndWait() {
	fmt.Println("Doing ...")
	time.Sleep(2 * time.Second)
}
