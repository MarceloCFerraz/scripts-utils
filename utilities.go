package utils

import (
	"io"
	"log"
	"os"
)

func GetContentsFromFile(fileName string) ([]byte, error) {
	log.Printf("Opening and reading contents of file '%s'\n", fileName)

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
