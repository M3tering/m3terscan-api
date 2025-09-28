package blockchain

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func GetClient() (*ethclient.Client, error) {
	// Load env once (optional, better at main.go)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		// still continue, maybe env already set
	}

	infuraAPIKey := os.Getenv("INFURA_API_KEY")
	if infuraAPIKey == "" {
		return nil, fmt.Errorf("INFURA_API_KEY not set")
	}

	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", infuraAPIKey)

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create eth client: %w", err)
	}

	return client, nil
}
