package main

import (
	"crypto/sha256"
	"math"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
	Hash  []byte
}

func MerkleTreeRoot(dataString []string) (headRoot []byte) {
	var dataRaw [][]byte
	for i := 0; i < len(dataString); i++ {
		dataRaw = append(dataRaw, []byte(dataString[i]))
	}
	var m MerkleTree
	headRoot = m.NMerkleTree(dataRaw).RootNode.Data
	return headRoot
}

func (m MerkleTree) NMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}

	if left == nil && right == nil {

		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}

func (m MerkleTree) NMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	level := 0
	for {
		if len(data) > int(math.Pow(2, float64(level))) {
			level = level + 1
		} else {
			break
		}
	}

	for i := len(data); i < int(math.Pow(2, float64(level))); i++ {
		data = append(data, data[len(data)-1])
	}

	for _, dat := range data {
		node := m.NMerkleNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(nodes)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := m.NMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}

	return &mTree

}
