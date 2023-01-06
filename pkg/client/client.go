package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func CreateConnection(connect string) {
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected successful :", connect)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

func StoreLog(connect, messgae string) {
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected successful :", connect)
	fmt.Fprintf(c, messgae+"\n")
	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("->: " + message)
}
