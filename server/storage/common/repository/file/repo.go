package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// FileRepo is a convenience struct for file-based repository impls.
// T must be serializable as JSON.
type FileRepo[T any] struct {
	f    *os.File
	Data T
}

// New instantiates a new FileRepo.
func New[T any](f *os.File) (*FileRepo[T], error) {
	dataJSON, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading data from file: %w", err)
	}

	d := new(T)
	if err := json.Unmarshal(dataJSON, d); err != nil && len(dataJSON) > 0 {
		return nil, fmt.Errorf("error decoding data from JSON: %w", err)
	}

	return &FileRepo[T]{
		f:    f,
		Data: *d,
	}, nil
}

// WriteData writes the current Data struct to the file as JSON.
func (j *FileRepo[T]) WriteData() error {
	if _, err := j.f.Seek(0, 0); err != nil {
		return fmt.Errorf("file seek error: %w", err)
	}

	if err := j.f.Truncate(0); err != nil {
		return fmt.Errorf("file truncate error: %w", err)
	}

	dataJSON, err := json.Marshal(j.Data)
	if err != nil {
		return fmt.Errorf("error encoding data to JSON: %w", err)
	}

	if _, err := j.f.Write(dataJSON); err != nil {
		return fmt.Errorf("error writing data to file: %w", err)
	}

	return nil
}
