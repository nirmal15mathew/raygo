package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Position struct {
	X float32
	Y float32
}

type ClientData struct {
	Position Position
	UserId   string
}

var clients = make(map[string](Position))

var mu sync.Mutex

// func handleConnection(conn net.Conn) {
// 	var userId string
// 	defer func() {
// 		delete(clients, userId)
// 	}()
// 	defer conn.Close()

// 	// conn.SetDeadline(time.Now().Add(15 * time.Second))
// 	fmt.Println("Got connection, ", conn.LocalAddr().String())
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// fmt.Printf("Received: %s", string(buf[:n]))
// 		// data := strings.ToUpper(string(buf[:n]))
// 		b := bytes.NewBuffer(buf[:n])

// 		var cData ClientData
// 		dec := gob.NewDecoder(b)

// 		if err := dec.Decode(&cData); err != nil {
// 			fmt.Println("error decoding: ", err)
// 			return
// 		}
// 		userId = cData.UserId

// 		mu.Lock()
// 		clients[cData.UserId] = cData.Position
// 		mu.Unlock()

// 		var sendBytes bytes.Buffer

// 		enc := gob.NewEncoder(&sendBytes)

// 		if err := enc.Encode(clients); err != nil {
// 			fmt.Println("Error in encoding: ", err)
// 			return
// 		}

// 		_, err = conn.Write(sendBytes.Bytes())
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func main() {
// 	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ln, err := net.ListenTCP("tcp", addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ln.Close()
// 	fmt.Println("Listening on port 8000")
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		go handleConnection(conn)
// 	}
// }

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buf := make([]byte, 1024)

	for {
		var userId string
		n, addr, err := connection.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			return
		}
		// fmt.Printf("Received: %s", string(buf[:n]))
		// data := strings.ToUpper(string(buf[:n]))
		b := bytes.NewBuffer(buf[:n])

		var cData ClientData
		dec := gob.NewDecoder(b)

		if err := dec.Decode(&cData); err != nil {
			fmt.Println("error decoding: ", err)
			return
		}
		userId = cData.UserId

		mu.Lock()
		clients[cData.UserId] = cData.Position
		mu.Unlock()

		// if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
		// 	fmt.Println("Exiting UDP server!")
		// 	return
		// }

		// data := []byte(strconv.Itoa(random(1, 1001)))
		// fmt.Printf("data: %s\n", string(data))
		// _, err = connection.WriteToUDP(data, addr)
		// if err != nil {
		// fmt.Println(err)
		// return
		// }
		var sendBytes bytes.Buffer

		enc := gob.NewEncoder(&sendBytes)

		if err := enc.Encode(clients); err != nil {
			fmt.Println("Error in encoding: ", err)
			return
		}

		_, err = connection.WriteToUDP(sendBytes.Bytes(), addr)
		if err != nil {
			log.Println(err)
			return
		}

		if userId == "STOP" {
			return
		}
	}
}
