package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Position struct {
	X float32
	Y float32
}

var userName string = "random"

type ClientData struct {
	Position Position
	UserId   string
}

func establishConnection() (net.Conn, bool) {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return conn, false
	}
	return conn, true
}

func encodeAndSend(conn net.Conn, x float32, y float32) {
	cData := ClientData{
		Position: Position{X: x, Y: y},
		UserId:   userName,
	}

	var b bytes.Buffer

	enc := gob.NewEncoder(&b)

	if err := enc.Encode(cData); err != nil {
		fmt.Println("Error in encoding: ", err)
		return
	}

	serialD := b.Bytes()
	conn.Write(serialD)
}

func receiveDataAndDecode(conn net.Conn) (map[string]Position, bool) {
	buff := make([]byte, 1024)
	var cData map[string]Position

	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return cData, false

	}
	b := bytes.NewBuffer(buff[:n])

	dec := gob.NewDecoder(b)

	if err := dec.Decode(&cData); err != nil {
		fmt.Println("error decoding: ", err)
		return cData, false
	}

	return cData, true
}

func setUserName(uname string) {
	userName = uname
}

func sendData(conn net.Conn) {
	_, err := conn.Write([]byte("Data"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
