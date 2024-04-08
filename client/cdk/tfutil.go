package cdk

import "fmt"

// converts a Terraform reference to a list of arns to a list of IDs
func tfRefARNsToIDs(ref string) string {
	return fmt.Sprintf(`${[for arn in %s: reverse(split("/", arn))[0]]}`, ref)
}
