package lessons

import (
	"fmt"
	"time"
)

func unpackTheDelivery(pack <-chan string, storage chan<- string, workerId int) {
	for content := range pack {
		fmt.Println("Worker ", workerId, " started  unpacking the delivery ", content)
		time.Sleep(2 * time.Second)
		fmt.Println("Worker ", workerId, " has unpacked ", content)
		storage <- content
	}
}

func Workers() {
	numJobs := 5
	packages := make(chan string, numJobs)
	storage := make(chan string, numJobs)

	// installing the workers
	for i := 1; i <= 3; i++ {
		go unpackTheDelivery(packages, storage, i)
	}

	for i := 1; i <= numJobs; i++ {
		fmt.Println("Postal delivered package", i)
		packages <- fmt.Sprintf("My little pony %d", i)
	}

	close(packages)

	for i := 1; i <= numJobs; i++ {
		content := <-storage
		fmt.Println("Package ", i, "has landed in the storage, content: ", content)
	}

}
