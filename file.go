package main

import "os"

// ReadFile reads file content and returns as string type.
func ReadFile(dir string) (string, error) {
	b, err := os.ReadFile(dir)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
