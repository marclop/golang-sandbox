package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Unbuffered
	c := make(chan int, 0)
	// Buffered
	// c := make(chan int, 20)
	//
	//   This creates a buffered channel with a capacity of 20.
	//   Normally channels are synchronous;
	//   both sides of the channel will wait until the other side is ready.
	//   A buffered channel is asynchronous;
	//   sending or receiving a message will not wait unless the channel is already full.
	//

	go func() {
		for {
			valueAssignment(randomGenerator(100), c)
		}
	}()
	go runner(c)
	var input string
	fmt.Scanln(&input)
}

func valueAssignment(value int, c chan int) {
	c <- value
	time.Sleep(time.Second * 2)
}

func busPrinter(value int) {
	fmt.Println(value)
	time.Sleep(time.Second)
}

func randomGenerator(lenght int) (r int) {
	timestamp := time.Now().UnixNano()
	r = rand.New(rand.NewSource(timestamp)).Intn(lenght)
	return
}

func runner(c chan int) {
	for {
		select {
		case msg := <-c:
			busPrinter(msg)
		case <-time.After(time.Second):
			fmt.Println("Timeout")
		}
	}
}
