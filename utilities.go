package utils

import (
	"fmt"
	"io"
	"os"
)

func GetContentsFromFile(fileName string) ([]byte, error) {
	fmt.Printf("Opening and reading contents of file '%s'", fileName)

	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}
