package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// conn.SetDeadline(time.Now().Add(15 * time.Second))
	fmt.Println("Got connection, ", conn.LocalAddr().String())
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
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

		mu.Lock()
		clients[cData.UserId] = cData.Position
		mu.Unlock()

		var sendBytes bytes.Buffer

		enc := gob.NewEncoder(&sendBytes)

		if err := enc.Encode(clients); err != nil {
			fmt.Println("Error in encoding: ", err)
			return
		}

		_, err = conn.Write(sendBytes.Bytes())
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Listening on port 8000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
