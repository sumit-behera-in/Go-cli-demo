package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)

	// notify incoming signals to the channel
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case res := <-sig:
			signal.Stop(sig)
			fmt.Println(res, "signal received")
			os.Exit(0)
		}
	}
}
