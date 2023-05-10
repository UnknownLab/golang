package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for {
		c <- "ping"
		c <- "ping2"
		c <- "ping3"
		time.Sleep(time.Second * 10)
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
		time.Sleep(time.Second * 10)
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string, 2)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
