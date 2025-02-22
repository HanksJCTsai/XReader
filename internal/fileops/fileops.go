package fileops

import (
	"io"
	"os"
)

func ReadFile(inputFilePath string) (string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(content), err
}
