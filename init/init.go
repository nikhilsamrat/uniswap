package init

import (
	"log"
	"os"
	"path/filepath"
	"uniswap/pkg/utils"
	"gopkg.in/yaml.v2"
)


// Deployment
type Deployment struct {
	FactoryAddress   string `yaml:"factory_address"`
    StartingBlockNum int    `yaml:"starting_block_number"`
}

// Exchange represents an exchange on a blockchain
type UniswapExchange struct {
    Name             string `yaml:"name"`
    V2   Deployment `yaml:"v2"`
	V3   Deployment `yaml:"v3"`
}

// Blockchain represents a blockchain with its exchanges
type Blockchain struct {
    Name     string    `yaml:"name"`
    Exchanges []UniswapExchange `yaml:"exchanges"`
}

// Config represents the entire configuration
type Config struct {
    Blockchains []Blockchain `yaml:"blockchains"`
}


const deployments = ("mainnet")


func main() {
	// Get the current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatal("Error getting current directory:", err)
        return
    }

	db, err :=utils.CreateRocksDB("/home/nikhilsamrat/go-buffer/uniswap/database")
	if err != nil {
		log.Fatal("Failed to create rocksdb:", err)
	}
	db.Close()
}