package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	input := GetDataFromUI()
	fmt.Println(hex.EncodeToString(MerkleTreeRoot(input)))

	var end string
	fmt.Scan(&end)
	fmt.Println(end)

}

func GetDataFromUI() []string {
	var dataRaw string
	var dataString []string
	fmt.Println("Enter data to be packed in blockchain (seperated by ',')")
	fmt.Scan(&dataRaw)
	dataString = strings.Split(dataRaw, ",")
	return dataString
}
