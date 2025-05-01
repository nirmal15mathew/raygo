package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

func establishConnectionUDP(address string) (*net.UDPConn, bool) {
	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		fmt.Println("Error occured while resolving address: ", err)
	}
	conn, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println("Couldn't connect to address")
		return conn, false
	}
	return conn, true
}

func receiveDataAndDecodeUDP(conn *net.UDPConn) (map[string]Position, bool) {
	buff := make([]byte, 1024)
	var cData map[string]Position

	n, _, err := conn.ReadFromUDP(buff)
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

func encodeAndSendUDP(conn *net.UDPConn, x float32, y float32) {
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
