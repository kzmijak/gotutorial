package lessons

import (
	"fmt"
	"time"
)

func ClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job no.", j)
			} else {
				fmt.Println("Finished receiving jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 6; i++ {
		time.Sleep(time.Second)
		fmt.Println("Packing job no.", i)
		jobs <- i
	}

	fmt.Println("Sent all jobs.")
	close(jobs)

	<-done
}
