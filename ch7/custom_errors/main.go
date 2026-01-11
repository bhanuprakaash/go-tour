package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Path   string
	Reason string
	Size   int64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s (%s)", e.Path, e.Reason)
}

func ValidateFile(path string, size int64) error {
	if path == "" {
		return errors.New("path is invalid")
	}
	const max_size = 100 * 1024 * 1024
	if size > max_size {
		return &ValidationError{
			Path:   path,
			Reason: "file limit exceeded",
			Size:   size,
		}
	}

	return nil
}

func main() {
	files := []struct {
		path string
		size int64
	}{
		{"", 0},
		{"resume.pdf", 500},
		{"movie.4k", 900000000000},
	}

	for _, file := range files {
		err := ValidateFile(file.path, file.size)
		if err == nil {
			fmt.Println("[OK]")
			continue
		}

		switch e := err.(type) {
		case *ValidationError:
			fmt.Printf("[SKIPPED] Too big by %d MB\n", (e.Size/1024/1024)-100)
		default:
			fmt.Printf("[CRITICAL] %v\n", e)
		}
	}

}
