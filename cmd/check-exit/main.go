package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cthach/tor"
)

func main() {
	os.Exit(mainWithExit())
}

func mainWithExit() int {
	if len(os.Args) == 1 {
		log.Fatal("No IP address given")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ip := os.Args[1]

	exit, err := tor.IsExit(ctx, ip)
	if err != nil {
		fmt.Printf("Error checking %s is a exit node: %v\n", ip, err)
		return 1
	}

	if !exit {
		fmt.Printf("FAIL! %s IS NOT a known exit node\n", ip)
		return 1
	}

	fmt.Printf("PASS! %s IS a known exit node\n", ip)
	return 0
}
