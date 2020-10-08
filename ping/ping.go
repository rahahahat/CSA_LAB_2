package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"
	"fmt"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	fmt.Println("Foo is sending: ping")
	channel <- "ping"
	for {
		select {
		case a := <- channel:
			fmt.Println("Foo has received: ", a)
			fmt.Println("Foo is sending: ping")
			channel <- "ping"
		}
	}
}


func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
	for {
		select {
		case a := <- channel:
			fmt.Println("Bar has received:", a)
			fmt.Println("Bar is sending: pong")
			channel <- "pong"
		}
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	// Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	channel := make(chan string)
	go foo(channel) 
	go bar(channel)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
