package main

import (
	"bytes"
	"crypto/sha256"
)

//MerkleTreeRoot : Calculating Merkle Tree Root
func MerkleTreeRoot(content []string) (headRoot []byte) {
	var data [][]byte
	for i := 0; i < len(content); i++ {
		data = append(data, []byte(content[i]))
	}
	var root Node
	return root.GenerateRoot(data).NodeHash
}

// Node : Node of Merkle Tree
type Node struct {
	NodeData []byte
	NodeHash []byte
}

// CalSHA256Hash : Calculate a sha256 hash
func (n Node) CalSHA256Hash(input []byte) []byte {
	h := sha256.New()
	h.Write(input)
	return h.Sum(nil)
}

// GenerateRoot : Create Merkle Tree
func (n Node) GenerateRoot(data [][]byte) (RootNode *Node) {

	// Prepare Leaf Node
	var nodes []*Node

	if len(data)%2 == 1 {
		nodes = append(nodes, &Node{
			NodeData: data[len(data)-1],
			NodeHash: n.CalSHA256Hash(data[len(data)-1]),
		})
	}

	//Building Tree from Bottom Layer
	for {
		if len(nodes) != 1 {
			var tempNodes []*Node
			for i := 0; i < len(nodes); i = i + 2 {
				tempNodes = append(tempNodes, &Node{
					NodeData: bytes.Join([][]byte{nodes[i].NodeHash, nodes[i+1].NodeHash}, []byte{}),
					NodeHash: n.CalSHA256Hash(bytes.Join([][]byte{nodes[i].NodeHash, nodes[i+1].NodeHash}, []byte{})),
				})
			}
			nodes = tempNodes
			if len(nodes)%2 == 1 {
				nodes = append(nodes, &Node{
					NodeData: data[len(data)-1],
					NodeHash: n.CalSHA256Hash(data[len(data)-1]),
				})
			}
		} else {
			break
		}
	}
	RootNode = nodes[0]
	return RootNode
}
