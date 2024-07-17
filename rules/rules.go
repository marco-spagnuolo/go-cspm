package rules

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Rule represents a security rule
type Rule struct {
	ID          string `yaml:"id" json:"id"`
	Description string `yaml:"description" json:"description"`
	Check       string `yaml:"check" json:"check"`
}

// LoadRules loads security rules from a YAML file
func LoadRules(filename string) ([]Rule, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read rules file: %w", err)
	}

	var rules []Rule
	ext := filepath.Ext(filename)
	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &rules)
	case ".json":
		err = json.Unmarshal(data, &rules)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal rules: %w", err)
	}

	return rules, nil
}

// RunCustomChecks executes the security rules
func RunCustomChecks(rules []Rule) {
	for _, rule := range rules {
		fmt.Printf("Running check: %s - %s\n", rule.ID, rule.Description)
		switch rule.Check {
		case "s3:CheckBucketEncryption":
			CheckS3Buckets()
		case "iam:CheckUnusedKeys":
			CheckUnusedIAMKeys()
		case "ec2:CheckInstanceTypes":
			CheckEC2InstanceTypes()
		case "rds:CheckBackupRetention":
			CheckRDSBackupRetention()
		default:
			fmt.Printf("Unknown check: %s\n", rule.Check)
		}
	}
}

// CheckS3Buckets checks the encryption of S3 buckets
func CheckS3Buckets() {
	fmt.Println("Checking S3 bucket encryption...")

	cmd := exec.Command("aws", "s3api", "list-buckets", "--query", "Buckets[].Name", "--output", "text")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to list S3 buckets: %v", err)
	}

	buckets := strings.Fields(string(output))
	fmt.Printf("Found %d S3 buckets\n", len(buckets))

	for _, bucket := range buckets {
		checkBucketEncryption(bucket)
	}
}

// checkBucketEncryption verifies the encryption of a specific bucket
func checkBucketEncryption(bucket string) {
	cmd := exec.Command("aws", "s3api", "get-bucket-encryption", "--bucket", bucket)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Bucket %s is not encrypted or encryption could not be determined\n", bucket)
	} else {
		fmt.Printf("Bucket %s is encrypted\n", bucket)
	}
}

// CheckUnusedIAMKeys checks for unused IAM keys
func CheckUnusedIAMKeys() {
	fmt.Println("Checking for unused IAM keys...")
	// Implement the logic to check for unused IAM keys
}

// CheckEC2InstanceTypes checks EC2 instance types
func CheckEC2InstanceTypes() {
	fmt.Println("Checking EC2 instance types...")
	// Implement the logic to check EC2 instance types
}

// CheckRDSBackupRetention checks RDS backup retention policy
func CheckRDSBackupRetention() {
	fmt.Println("Checking RDS backup retention policy...")
	// Implement the logic to check RDS backup retention policy
}
