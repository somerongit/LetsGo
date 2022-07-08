package main

import "time"

func checkChannel() {

	c := make(chan int)
	// <name> chan <datatype>

	go func() {
		c <- 11
	}()

	println("Data receicved from channel(", c, ") : ", <-c)

	time.Sleep(time.Second * 3)

	go func() {
		c <- 12
	}()

	println("Data receicved from channel(", c, ") : ", <-c)
}

func checkBufferChannel() {

	// As we r using buffer channel so we need to close the channel
	// Buffer size 3, so after 3 publish it need to wait of consumtion then send rest
	c := make(chan int, 3)

	go func() {
		c <- 7
		c <- 8
		c <- 9
		c <- 10
		c <- 11
		c <- 12
		close(c)
	}()

	for i := range c {
		println("Data receicved from channel(", c, ") : ", i)
	}
}

type Human struct{ Name string }

func checkBufferChannelWithCustomDataType() {

	c := make(chan *Human, 3)

	go func() {
		c <- &Human{"Human  7 "}
		c <- &Human{"Human  8 "}
		c <- &Human{"Human  9 "}
		c <- &Human{"Human 10 "}
		c <- &Human{"Human 11 "}
		c <- &Human{"Human 12 "}
		close(c)
	}()

	for i := range c {
		println("Data receicved from channel(", c, ") : ", i.Name)
	}
}
