package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	// Define Host and Port
	serverHost := "localhost"
	serverPort := "3000"
	nearbyHost := "localhost"
	nearbyPort := "3001"
	if len(os.Args) == 2 {
		serverPort = os.Args[1]
	}
	if len(os.Args) == 3 {
		serverPort = os.Args[1]
		nearbyPort = os.Args[2]
	}
	serverAddr, err := net.ResolveTCPAddr("tcp", serverHost+":"+serverPort)
	errorMsg(err)
	nearbyAddr, err := net.ResolveTCPAddr("tcp", nearbyHost+":"+nearbyPort)
	errorMsg(err)

	// Choose Function - Either be a miner, or a nodecontroller
	// **Becauses it would be Peer2Peer if becomes miner and nodecontroller at the same time.
	// **We need to made the same TCP socket "Dial" and "Listen".
	// **By default, it is not supported in golang.
	fmt.Println("Self Node port", serverPort, "; Nearby Node port ", nearbyPort)
	fmt.Println("Enter 1 to become a Node Controller;")
	fmt.Println("Enter 2 to become a Miner")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		// Connect to Full Node and update Self blockchain

		// Listening
		listener, err := net.ListenTCP("tcp", serverAddr)
		errorMsg(err)
		fmt.Println("Node:	Server Listening on port", serverPort)

		// Create socket if a connection is accepted
		// golang allows multiple connection by default (non-blocking)
		for {
			fullNodeDownloadBC()
			conn, err := listener.Accept()
			errorMsg(err)
			go handleMsg(conn)
		}

	case "2":
		// Connect to nearby Node
		conn, err := net.DialTCP("tcp", serverAddr, nearbyAddr)
		errorMsg(err)
		fmt.Printf("Miner:	%s <--> %s\n", serverAddr.String(), nearbyAddr.String())

		// Send message
		var message string
		fmt.Println("Miner:	Enter data to be packed in blockchain.")
		fmt.Scan(&message)
		sendMsg(conn, message)

		conn.Close()
	}
}

func handleMsg(conn net.Conn) {

	fmt.Printf("Node:	Connected to: %s\n", conn.RemoteAddr().String())

	for {
		// Receive Message, maximum length is 1024 bytes
		bufReceive := make([]byte, 1024)
		bufSend := make([]byte, 1024)
		_, err := conn.Read(bufReceive)
		if err == io.EOF {
			fmt.Println("Node:	Error reading:")
			fmt.Println("Node:	", err)
		}
		if err != nil {
			fmt.Println("Node:	Error reading:")
			fmt.Println("Node:	", err)
		}

		// Choose action depending on message header
		if string(bytes.Trim(bufReceive, "\x00")) == "exit" {
			break
		}
		switch string(bytes.Trim(bufReceive, "\x00")) {

		case "mining":
			// "mining":	1. Return prevBlock to miner by conn.Write()
			//				3. Receive newBlock from miner by conn.Read()
			//				4. Check if the newBlock is valid, i.e. prevHash match & nonce is valid
			break
		case "getBCh":
			// "getBCh":	1. Send blockchain in memory to nearby node.
			break

		case "chckBK":
			// "chckBK":	1. Check if the newBlock is valid, i.e. prevHash match & nonce match
			break

		default:
			// other:	echo server
			fmt.Println("Node:	[", string(bytes.Trim(bufReceive, "\x00")), "] received from", conn.RemoteAddr().String())

			bufSend = bufReceive
			if bufSend != nil {
				_, err = conn.Write(bufSend)
			}
			if err != nil {
				fmt.Println("Node:	Error writing:")
				fmt.Println("Node:	", err)
				break
			}
			fmt.Println("Node:	[", string(bytes.Trim(bufSend, "\x00")), "] sent to", conn.RemoteAddr().String())
			conn.Close()
			fmt.Println("Node:	Connection from ", conn.RemoteAddr().String(), "is successively closed.")
			fullNodeUploadBC()
			return
		}
		conn.Close()
		fmt.Println("Node:	Connection from ", conn.RemoteAddr().String(), "is UNEXPECTEDLY closed.")
		return
	}
}

func sendMsg(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Miner:	Error writing:")
		fmt.Println("Miner:	", err)
		return
	}

	fmt.Println("Miner:	Sent", msg)

	buf := make([]byte, 512)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Miner:	Error Reading:")
		fmt.Println("Miner:	", err)
		return
	}

	fmt.Println("Miner:	Got ", string(bytes.Trim(buf, "\x00")))
}

func fullNodeDownloadBC() {
	//To Update
	fmt.Println("Node:	Full Node is downloaded.")
	return
}

func fullNodeUploadBC() {
	//To Update
	fmt.Println("Node:	Full Node is Updated.")
	return
}

func errorMsg(err error) {
	if err != nil {
		fmt.Println("Connection Error:	", err)
		os.Exit(1)
	}
}
