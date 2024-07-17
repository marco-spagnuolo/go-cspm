package rules

import "github.com/aws/aws-sdk-go/service/rds"

// Verifica che le istanze RDS abbiano la crittografia abilitata
func CheckRDSInstanceEncryption(dbInstance *rds.DBInstance) bool {
	return *dbInstance.StorageEncrypted
}

// Verifica che le istanze RDS abbiano backup automatici configurati
func CheckRDSInstanceBackup(dbInstance *rds.DBInstance) bool {
	return *dbInstance.BackupRetentionPeriod > 0
}
