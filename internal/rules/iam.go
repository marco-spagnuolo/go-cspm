package rules

import "github.com/aws/aws-sdk-go/service/iam"

// Verifica che gli utenti IAM abbiano l'autenticazione a due fattori abilitata
func CheckIAMUserMFA(user *iam.User, svc *iam.IAM) bool {
	input := &iam.ListMFADevicesInput{
		UserName: user.UserName,
	}
	result, err := svc.ListMFADevices(input)
	if err != nil {
		return false
	}

	return len(result.MFADevices) > 0
}

// Verifica che le politiche IAM non siano troppo permissive
func CheckIAMPolicy(policy *iam.Policy) bool {
	// Implementare la logica per verificare le politiche IAM
	// Esempio: Controllare che le politiche non includano "Action": "*" e "Resource": "*"
	return true
}
