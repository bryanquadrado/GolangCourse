package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)
	
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()

	}()
	
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()

	}()

	go func() {
		escrever("Olá mundo 2!")
		waitGroup.Done()

	}()
	
	go func() {
		escrever("Programando em Go! 2")
		waitGroup.Done()

	}()

	waitGroup.Wait()

	

}

func escrever(texto string) {
	for i := 0; i < 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}