package lessons

import "fmt"

func NonBlockingChannels() {
	defaultingChannel := make(chan string)

	select {
	case <-defaultingChannel:
		fmt.Println("Successfully unloaded defaultingChannel")
	default:
		fmt.Println("Failure at unloading defaultingChannel")
	}

	select {
	case defaultingChannel <- "defaultingChannel attempting to insert message":
		fmt.Println("Success: ", <-defaultingChannel)
	default:
		fmt.Println("Failure to load defaultingChannel")
	}

	successChannel1 := make(chan string)
	successChannel2 := make(chan string)

	go func() {
		successChannel1 <- "msg1"
		successChannel2 <- "msg2"
	}()

	select {
	case msg1 := <-successChannel1:
		fmt.Println(msg1)
	case msg2 := <-successChannel2:
		fmt.Println(msg2)
	default:
		fmt.Println("Default")
	}
}
