package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	// タイマー用channel
	duration := 1 * time.Second
	tc := time.After(duration)
	<-tc
	fmt.Println("timer is finished.")
}
