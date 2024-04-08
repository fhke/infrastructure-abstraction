package concurrency

import "errors"

var ErrOutdated = errors.New("resource is outdated, update failed")
