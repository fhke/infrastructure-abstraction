package concurrency

// OptimisticLock is an optimistic locking implementation using version numbers.
// This is intended to be embedded in other structs.
type OptimisticLock struct {
	ResourceVersion int `json:"resourceVersion"`
}

// PreUpdate must be called prior to updating the resource.
func (o *OptimisticLock) PreUpdate(newResourceVersion int) error {
	if o.ResourceVersion != newResourceVersion {
		return ErrOutdated
	}

	o.ResourceVersion++
	return nil
}
