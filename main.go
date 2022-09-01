package main

import (
	"bytes"
	"Crypto/sha256"
)

type BlockChain struct {
	blocks []*Block

}

type Block struct {
	Hash[]byte
	Data[]byte
	PrevHash[]byte

}

func (b *Block) DeriveHash() {
		info := bytes.Join([][]byte{b/Data, b.PrevHash}, []byte{})
		hash := sha256.Sum256(info)
		b.Hash = hash[:]

func CreateBlock(data string, prevHash[]byte) *Block {
		block := &Block{[]byte{}, []byte(data), prevhash}
		block.DeriveHash()
		return block 

}

func (chain *BlockChain) AddBlock(data string) {
		prevBlock := chain.blocks[len(chain.blocks)-1]
		new := CreateBlock(data, prevBlock.Hash)
		chain.blocks = append(chain.blocks, new)

}

func Genesis() * Block {
		return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *BlockChain {
	return &BlockChain{[] *Block{Genesis()}}

}

func main() {
		chian := InitBlockchain()

		chain.AddBlock("0x0000001")
		chain.AddBlock("0x0000002")
		chain.AddBlock("0x0000003")

		for _, block := range chain.blocks {

			fmt.Printf("Previous Hash: %x\n", block.PrevBlock)
			fmt.Printf("Data in Block: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
		}
}

	}