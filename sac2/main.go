package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)

	// notify incoming signals to the channel
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go prompter()

	for {
		select {
		case res := <-sig:
			signal.Stop(sig)
			fmt.Println(res, "signal received")
			os.Exit(0)
		}
	}
}

func prompter() {
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Printf("<- you entered \"%s\"\n", line)
		fmt.Print(">> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
