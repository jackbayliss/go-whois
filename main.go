package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-whois <domain>")
		os.Exit(1)
	}
	domain := os.Args[1]

	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("Error resolving domain: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("IP addresses for %s:\n", domain)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
