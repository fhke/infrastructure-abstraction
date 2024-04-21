package dynamo

import "fmt"

var ErrNoPrimaryKey = fmt.Errorf("scan did not return %s attribute", primaryKey)
