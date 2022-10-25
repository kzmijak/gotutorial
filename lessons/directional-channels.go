package lessons

import "fmt"

func writer(channel chan<- string, content string) {
	channel <- content
}

func reader(channel <-chan string, content *string) {
	cnt := <-channel
	// fmt.Println(cnt)
	content = &cnt
}

func redirector(srcChannel <-chan string, destChannel chan<- string) {
	// destChannel<- <-srcChannel
	content := <-srcChannel
	destChannel <- content
}

func DirectionalChannels() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	var outcome string = "null"
	go writer(channel1, "message")
	go redirector(channel1, channel2)
	go reader(channel2, &outcome)
	go writer(channel1, outcome)

	fmt.Println(<-channel1)
}
