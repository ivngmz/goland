package _3concurrencyAdvanced

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

package main

import (
"fmt"
"strings"
"sync"
"time"
)

func main() {
	/* Channels:
	Info:
		- Channels are the pipes that connect concurrent goroutines.
		- You can send values into channels from one goroutine and receive those values into another goroutine.
		- By default, sends and receives block until the other side is ready.
		- Channels allows different goroutines to communicate with each other.
		- As Rob Pike said: "Do not communicate by sharing memory; instead, share memory by communicating."
		- communicate by sharing memory: Threads and mutexes.
		- share memory by communicating: Goroutines and channels.
		- Communication over channels:
			- Is bidirectional: Send and receive operations are symmetric.
			- Is synchronous: Both sides of the channel must be ready to communicate.
			- Is blocking: Send and receive operations block until the other side is ready.
			- Is unbuffered by default: Channels have no capacity.
			- Is typed: The type of data that can be sent or received is defined by the channel.
			- Is concurrent: Goroutines can send and receive data at the same time.
	- Syntax:
		- ch := make(chan int)
	- operations:
		- Send: ch <- value
		- Receive: value := <-ch
	    - Close: close(ch)
	    - Query Buffer: cap(ch)
	    - Quering length: len(ch) (it is not recommended to use it)

	Unbuffered channels:
		- An unbuffered channel is a channel without a capacity.
		- You can only send a value to the channel if there is a goroutine ready to receive the value.
		- When the channel is empty, the receive operation blocks.
		- When the channel is full, the send operation blocks.
		- Syntax:
			- ch := make(chan int)
	Buffered channels:
		- A buffered channel is a channel with a capacity.
		- You can send as many values as the capacity of the channel without blocking.
		- When the channel is full, the send operation blocks.
		- When the channel is empty, the receive operation blocks.
		- Syntax:
			- ch := make(chan int, 3)
	*/

	// Unbuffered channel
	ch := make(chan string)
	var wg sync.WaitGroup

	println("### Start of the program")
	wg.Add(2)
	content1 := "SOS"
	go sendMsg(content1, ch, &wg)
	go receiveMsg(ch, &wg)
	wg.Wait()

	// Buffered channel
	chBuffered := make(chan string, 9)
	wg.Add(2)
	println("### Start of the program")
	content2 := "SOS, mayday mayday thers a big problem over here, " +
		"please pass the message to the other side of the channel!"
	go sendMsg(content2, chBuffered, &wg)
	go receiveMsg(chBuffered, &wg)
	wg.Wait()
	println("### End of the program")

	/* Select statement:
	- The select statement lets a goroutine wait on multiple communication operations.
	- A select blocks until one of its cases can run, then it executes that case.
	- It chooses one at random if multiple are ready.
	- Syntax:
		select {
		case <-ch1:
			fmt.Println("Received from ch1")
		case ch2 <- 1:
			fmt.Println("Sent to ch2")
		default:
			fmt.Println("Default case")
		}
	*/
	selectMsg(ch, chBuffered)

	select {
	case <-ch:
		fmt.Println(">>>>>>>>>>>>> Received from ch on second select")
	case <-time.After(1 * time.Second):
		fmt.Println(">>>>>>>>>>>>> Timeout")
	}
}

// Send data to the channel
func sendMsg(content string, ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Sending message to the channel")
	chLength := len(ch)
	wordList := strings.Split(content, " ")
	// for range loop on channel
	for _, word := range wordList {
		ch <- word
		fmt.Printf("Sent word to the channel: %v\n", word)
		chLength = len(ch)
		fmt.Printf("Channel length equals to: %v\n", chLength)
	}
	defer wg.Done()
	wg.Done()
}

// Receive data to the channel
func receiveMsg(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for mesagges ...")
	// for range loop on channel
	for message := range ch {
		fmt.Printf("Received message from the channel: %v\n", message)
	}
	defer wg.Done()
	wg.Done()
}

func selectMsg(ch1, ch2 chan string) {
	select {
	case <-ch1:
		fmt.Println(">>>>>>>>>>>>> Received from ch1")
		fmt.Printf("Received from ch1: %v\n", <-ch1)
	case ch2 <- "message":
		fmt.Println(">>>>>>>>>>>>> Sent to ch2")
		fmt.Printf("Sent to ch2: %v\n", <-ch2)
	default:
		fmt.Println("Default case")
	}
	time.Sleep(1 * time.Second)
}
