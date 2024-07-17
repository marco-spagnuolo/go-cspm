package main

import (
	"cspm-go/config"
	"fmt"
	"log"
)

func main() {
	config, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if config.AWSAccessKeyID == "" || config.AWSSecretAccessKey == "" || config.AWSRegion == "" {
		log.Fatalf("Missing required AWS configuration")
	}

	fmt.Printf("AWS Access Key ID: %s\n", config.AWSAccessKeyID)
	fmt.Printf("AWS Secret Access Key: %s\n", config.AWSSecretAccessKey)
	fmt.Printf("AWS Region: %s\n", config.AWSRegion)
}
