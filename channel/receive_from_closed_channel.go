package channel

import "fmt"

func RecieveFromACloedChan() {
	toClose := make(chan interface{})
	done := make(chan struct{})

	go func() {
		reciever := <-toClose
		fmt.Println(reciever)
		done <- struct{}{}
	}()

	fmt.Println("Channel to be closed")
	close(toClose)
	<-done
}
