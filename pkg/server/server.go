package server

import (
	"be/pkg/client"
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func StartServer(port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}
	fmt.Println("tcp server starting with port : ", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
		}
		go handleConnection(conn, port)
	}
}
func handleConnection(conn net.Conn, port string) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("client left..")
		conn.Close()
		return
	}
	message := string(bufferBytes)
	clientAddr := conn.RemoteAddr().String()
	arr := strings.Split(message, " ")
	if arr[0] == "client" && port == "8000" {
		fmt.Println("send log for node 2 && node 3")
		client.StoreLog("127.0.0.1:8001", message)
		client.StoreLog("127.0.0.1:8002", message)
	}
	response := fmt.Sprintf(message + " from " + clientAddr + "\n")
	log.Println(response)
	conn.Write([]byte("you sent: " + response))
	handleConnection(conn, port)
}
