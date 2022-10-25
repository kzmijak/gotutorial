package lessons

import (
	"fmt"
	"time"
)

func Ticker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	doneTicking := make(chan bool)

	go func() {
		for {
			select {
			case <-doneTicking:
				return
			case tickerOutput := <-ticker.C:
				fmt.Println(tickerOutput)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	doneTicking <- true

	fmt.Println("Stopped")
}
