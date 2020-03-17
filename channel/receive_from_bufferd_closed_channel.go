package channel

import "fmt"

func RecieveFromAnAlreadyClosedButBufferedChan(bufferSize int) {
	fmt.Println("Buffer size:", bufferSize)
	toClose := make(chan int, bufferSize)
	startSignal := make(chan struct{})
	completeSignal := make(chan struct{})

	go func() {
		fmt.Println("Produce data to buffered channel")
		for i := 0; i < bufferSize; i++ {
			toClose <- i
		}
		fmt.Println("Closed buffered channel")
		close(toClose)
		fmt.Println("Notify consume to start consume")
		close(startSignal)
	}()

	go func() {
		defer close(completeSignal)
		<-startSignal
		fmt.Println("Begin to consume")
		for {
			select {
			case n, ok := <-toClose:
				if !ok {
					fmt.Println("No more to consume")
					return
				}
				fmt.Println("Consumed:", n)
			}
		}
	}()

	<-completeSignal
	fmt.Println("Test end")
}
