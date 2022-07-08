package main

import "time"

// Pinger prints a ping and wait for pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		println("Channel Pinger(", ponger, "):Ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// Ponger prints a pong and wait for ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		println("Channel Ponger(", pinger, "): Pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func paly() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)
	// The play go routine starts the ping/pong by sending into the ping channel
	ping <- 1

	select {}
}
