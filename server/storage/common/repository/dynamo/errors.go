package dynamo

import (
	"net/http"

	awsHttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
)

func ErrIsNotFound(err error) bool {
	opErr, ok := err.(*smithy.OperationError)
	if !ok {
		return false
	}

	httpErr, ok := opErr.Err.(*awsHttp.ResponseError)
	return ok && httpErr.HTTPStatusCode() == http.StatusNotFound
}
