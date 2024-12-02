package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

type Block struct {
    Index        int
    Timestamp    string
    Transactions []Transaction
    PreviousHash string
    Hash         string
    Nonce        int
}

type Transaction struct {
    From   string
    To     string
    Amount float64
}

type Blockchain struct {
    Blocks []Block
}

func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprint(block.Transactions) + block.PreviousHash + strconv.Itoa(block.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, transactions []Transaction) Block {
    var newBlock Block

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = time.Now().String()
    newBlock.Transactions = transactions
    newBlock.PreviousHash = oldBlock.Hash

    for i := 0; ; i++ {
        newBlock.Nonce = i
        newBlock.Hash = calculateHash(newBlock)
        if isValidHash(newBlock.Hash) {
            break
        }
    }

    return newBlock
}

func isValidHash(hash string) bool {
    return hash[:4] == "0000"
}

func createGenesisBlock() Block {
    genesisBlock := Block{
        Index:        0,
        Timestamp:    time.Now().String(),
        Transactions: []Transaction{},
        PreviousHash: "",
        Hash:         "",
    }
    genesisBlock.Hash = calculateHash(genesisBlock)
    return genesisBlock
}

func main() {
    genesisBlock := createGenesisBlock()
    blockchain := Blockchain{[]Block{genesisBlock}}

    transactions := []Transaction{
        {From: "Alice", To: "Bob", Amount: 10.5},
        {From: "Bob", To: "Charlie", Amount: 3.0},
    }

    newBlock := generateBlock(genesisBlock, transactions)
    blockchain.Blocks = append(blockchain.Blocks, newBlock)

    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Transactions: %v\n", block.Transactions)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Println()
    }
}
