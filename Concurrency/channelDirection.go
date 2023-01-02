package main

import (
	"fmt"
	"time"
)

func pingGeneratorD(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
}

func pongGeneratorD(c chan<- string) {
	//Information can only be sent to the channel - a generator
	for i := 0; i < 5; i++ {
		c <- "pong"
	}
}

func outputD(c <-chan string) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out")
			quit <- true
		}
	}
}

func useChannelDirection() {
	quit = make(chan bool)
	c := make(chan string)
	go pingGeneratorD(c)
	go pongGeneratorD(c)
	go outputD(c)

	<-quit
}
