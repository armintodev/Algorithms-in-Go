package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup2 sync.WaitGroup
var quit chan bool

func pingGenerator(c chan string) {
	//defer waitGroup2.Done()

	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func output(c chan string) {
	//defer waitGroup2.Done()

	//for {
	//	value := <-c
	//	fmt.Println(value)
	//	//fmt.Println(<-c)  //also we can use this code.
	//}

	//avoid throw deadlock exception.
	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out")
			//waitGroup2.Done()
			quit <- true
		}
	}
}

func useChannel() {
	c := make(chan string)

	waitGroup2.Add(2)
	go pingGenerator(c)
	go output(c)

	waitGroup2.Wait()
}

func useChannelWithoutWaitGroup() {
	quit = make(chan bool)
	c := make(chan string)

	go pingGenerator(c)
	go output(c)
	<-quit
}
