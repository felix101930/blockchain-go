package main

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct{

	blocks []*Block

}
type Block struct{

	Hash []byte
	Data []byte
	PrevHash []byte

}

func (chain *BlockChain) AddBlock(data string){

	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.PrevHash)
	chain.blocks = append(chain.blocks, new)

}

// create first block 
func Genesis() *Block {
	return CreateBlock("Genesis",[]byte{})
}


func InitBlockChain() *BlockChain{

	return &BlockChain{[]*Block{Genesis()}}

}

func (b *Block) DeriveHash() {

	info :=bytes.Join([][]byte{b.Data,b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {

	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block


}

func main() {

	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	


}