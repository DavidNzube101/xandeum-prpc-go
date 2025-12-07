package main

import (
	"fmt"
	"log"

	"github.com/xandeum/prpc-go"
)

func main() {
	// Test with a public pNode IP from Discord
	ip := "192.190.136.28"
	client := prpc.NewClient(ip)

	fmt.Printf("Testing pRPC client with IP: %s\n", ip)

	// Test GetPods
	fmt.Println("Fetching pods...")
	pods, err := client.GetPods()
	if err != nil {
		log.Printf("Error fetching pods: %v", err)
	} else {
		fmt.Printf("Pods: %+v\n", pods)
	}

	// Test GetStats
	fmt.Println("Fetching stats...")
	stats, err := client.GetStats()
	if err != nil {
		log.Printf("Error fetching stats: %v", err)
	} else {
		fmt.Printf("Stats: %+v\n", stats)
	}
}
