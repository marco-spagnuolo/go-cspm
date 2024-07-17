package check

import (
	"fmt"
	"log"
	"os/exec"
)

func RunChecks() {
	fmt.Println("Running security checks...")
	// Aggiungi qui i controlli di sicurezza, ad esempio:
	CheckS3Buckets()
	fmt.Println("Security checks completed.")
}

func CheckS3Buckets() {
	fmt.Println("Checking S3 bucket encryption...")

	cmd := exec.Command("aws", "s3api", "list-buckets", "--query", "Buckets[].Name", "--output", "text")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to list S3 buckets: %v", err)
	}

	buckets := string(output)
	fmt.Printf("Found S3 buckets: %s\n", buckets)

	for _, bucket := range buckets {
		checkBucketEncryption(string(bucket))
	}
}

func checkBucketEncryption(bucket string) {
	cmd := exec.Command("aws", "s3api", "get-bucket-encryption", "--bucket", bucket)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Bucket %s is not encrypted or encryption could not be determined\n", bucket)
	} else {
		fmt.Printf("Bucket %s is encrypted\n", bucket)
	}
}
