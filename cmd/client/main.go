package main

import (
	"be/pkg/client"
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}
	connect := arg[1]
	if len(arg) > 2 {
		if arg[2] == "client" {
			fmt.Println("client")
		}
	}
	client.CreateConnection(connect)
}
