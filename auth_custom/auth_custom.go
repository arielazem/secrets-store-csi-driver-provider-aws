package auth_custom

import (
	"fmt"
	"strings"
)

func ValidatePermissions(objectName string, namespace string) error {
	if strings.HasPrefix(objectName, "123") {
		return fmt.Errorf("Object name is: %s", objectName)
	}
	return nil
}

func GetRootAWSAttributes() (string, string, string) {
	serviceAccount := "aws-secrets-manager-demo"
	nameSpace := "aws-secrets-manager-demo"
	Arn := "arn:aws:iam::414158246671:role/vault/vault-secrets-manager-us-west-2"

	return serviceAccount, nameSpace, Arn
}
