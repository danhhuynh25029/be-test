package main

import (
	"be/pkg/server"
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := arg[1]
	server.StartServer(PORT)
}
