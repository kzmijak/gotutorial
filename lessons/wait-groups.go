package lessons

import (
	"fmt"
	"sync"
	"time"
)

func performHeavyLifting(operationId int, waitGroup *sync.WaitGroup) {
	fmt.Printf("Operation %d started\n", operationId)
	
	time.Sleep(time.Duration(operationId) * time.Second)
	
	fmt.Printf("Operation %d finished\n", operationId)
	
	waitGroup.Done()
}

func WaitGroups() {
	var waitGroup sync.WaitGroup
	
	for iterator := 5; iterator > 0; iterator-- {
		waitGroup.Add(1)
		iterator := iterator
		go performHeavyLifting(iterator, &waitGroup)
	}

	fmt.Println("Waiting for everyone to end the task")
	waitGroup.Wait()
	fmt.Println("All workers done")

}
