package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

// minerSendMsg : Just send message to a node. Get reply.
func minerSendMsg(conn net.Conn, msg []byte) (reply []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Miner:	Error writing:")
		fmt.Println("Miner:	", err)
		return
	}

	fmt.Println("Miner:	...sending message to nearby node")

	buf := make([]byte, 8192)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Miner:	...Error Reading:")
		fmt.Println("Miner:	...", err)
		return
	}

	fmt.Println("Miner:	...received message from nearby node")
	return bytes.TrimRight(buf, "\x00")
}

// minerGetDataFromUI : Receive data, in format of string, from user.
func minerGetDataFromUI() []string {
	var dataRaw string
	var dataString []string
	fmt.Println("Miner:	Enter data to be packed in blockchain (seperated by ',')")
	fmt.Scan(&dataRaw)
	dataString = strings.Split(dataRaw, ",")
	return dataString
}

// minerPrintBlock : Print a block in command line interface.
func minerPrintBlock(block *Block) {
	fmt.Printf("Miner:	Block Information\n")
	fmt.Printf("	 > TimeStamp     : %d\n", block.Timestamp)
	fmt.Printf("	 > PrevBlockHash : %x\n", block.PrevBlockHash)
	fmt.Printf("	 > MerkleTreeRoot: %x\n", block.MerkleTreeRoot)
	fmt.Printf("	 > Nonce         : %05d\n", block.Nonce)
	fmt.Printf("	 > CurrBlockHash : %x\n", block.CurrentBlockHash)
	fmt.Printf("	 > Data          : %s\n", block.Data)
	return
}
