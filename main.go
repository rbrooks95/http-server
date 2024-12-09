package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)
func main()  {
listener , err := net.Listen("tcp", ":8080")

if err != nil {
	log.Fatalf("Error creating listner: %v", err)

}
defer listener.Close()
log.Println("server is listening on port 8080")

for{
	conn, err := listener.Accept()
	if err != nil {
		log.Printf("Error accepting connection: %v", err)
		continue
	}
	go handleConnection(conn)
}
}

func handleConnection(conn net.Conn)  {
	defer conn.Close()
	log.Printf("New client connected: %s", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("client disconnected: %s", conn.RemoteAddr().String())
			return
		}
		response := strings.ToUpper(message)

		_, writeErr := conn.Write([]byte(response))

		if writeErr != nil {
			log.Printf("error writing to client: %v", writeErr)
			return 
		}
		fmt.Printf("recived : %sSent: %s",message, response)
	}
}
