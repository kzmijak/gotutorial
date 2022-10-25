package lessons

import (
	"fmt"
	"time"
)

func Timeouts() {
	lateChannel := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		lateChannel <- "I am in no hurry, but here I am."
	}()

	select {
	case lateMessage := <-lateChannel:
		fmt.Println(lateMessage)
	case <-time.After(1 * time.Second):
		fmt.Println("I am very impatient. Cut it.")
	}

	hurriedChannel := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		hurriedChannel <- "I am in big hurry, but I'm here."
	}()

	select {
	case hurriedMessage := <-hurriedChannel:
		fmt.Println(hurriedMessage)
	case <-time.After(3 * time.Second):
		fmt.Println("I am in no hurry, but I can't wait no longer")
	}
}
