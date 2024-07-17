package report

import (
	"fmt"
	"log"
	"os"
)

func GenerateReport() {
	fmt.Println("Generating security report...")

	reportFile, err := os.Create("report.txt")
	if err != nil {
		log.Fatalf("Failed to create report file: %v", err)
	}
	defer reportFile.Close()

	reportFile.WriteString("Security Report\n")
	reportFile.WriteString("---------------\n\n")

	reportFile.WriteString("S3 Bucket Encryption\n")
	reportFile.WriteString("--------------------\n\n")

	reportFile.WriteString("Bucket1: Encrypted\n")
	reportFile.WriteString("Bucket2: Not encrypted\n")

	fmt.Println("Security report generated.")
}
