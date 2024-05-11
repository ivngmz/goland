package _2concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func start() {
	go process()
	println("Starting...")
}

func process() {
	println("Processing...")
}

func generateRandomNumber() {
	randomNumber := rand.Int()
	time.Sleep(1 * time.Second)
	println(randomNumber)
}

func main() {
	/* Concurrency:
	- Concurrency is the composition of independently executing processes.
	- Concurrency is a way to structure a program by breaking it into pieces that can be executed independently.
	- Concurrency is not parallelism.
	- Parallelism is the simultaneous execution of multiple things, possibly on multiple cores.
	*/

	println("### Concurrency basic example")
	initialTime := time.Now()
	for i := 0; i < 10; i++ {
		//generateRandomNumber()
		go generateRandomNumber()
	}
	time.Sleep(2 * time.Second)
	timeTaken := time.Since(initialTime)
	fmt.Printf("Time taken: %v\n", timeTaken)

	println("### go routine process example")
	go start()
	time.Sleep(1 * time.Second)

	/* Anonymous go routine
	- Anonymous functions are useful when you want to define a function inline without having to name it.
	*/
	println("### go anonymous routine process example")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Processing anonymous... %v\n", i)
			time.Sleep(1 * time.Second)
		}()
	}
	time.Sleep(2 * time.Second)

	/* Wait groups
	- A WaitGroup waits for a collection of goroutines to finish.
	- The main goroutine calls Add to set the number of goroutines to wait for.
	- Then each of the goroutines runs and calls Done when finished.
	- At the same time, Wait can be used to block until all goroutines have finished.
	*/

	println("### go wait group process example")
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Processing wait group... %v\n", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Done wait group... %v\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	println("All done!")
	time.Sleep(5 * time.Second)
}
