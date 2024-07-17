package main

import (
	"flag"
	"fmt"
	"log"

	"cspm-go/check"
	"cspm-go/config"
	"cspm-go/report"
	"cspm-go/rules"
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

	checkFlag := flag.Bool("check", false, "Run security checks")
	reportFlag := flag.Bool("report", false, "Generate security report")
	rulesFile := flag.String("rules", "", "Path to custom rules file")

	flag.Parse()

	if *checkFlag {
		if *rulesFile != "" {
			loadedRules, err := rules.LoadRules(*rulesFile)
			if err != nil {
				log.Fatalf("Error loading rules: %v", err)
			}
			rules.RunCustomChecks(loadedRules)
		} else {
			check.RunChecks()
		}
	}
	if *reportFlag {
		report.GenerateReport()
	}

	if !*checkFlag && !*reportFlag {
		fmt.Println("No command provided. Use -check or -report.")
	}
}
