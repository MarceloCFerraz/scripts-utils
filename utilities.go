package utils

import (
	"fmt"
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

func ReadJWTTokenFromFile(filePath string) (*string, error) {
	// Read the file contents
	tokenBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading token file: %v", err)
	}

	// Convert the byte slice to a string (JWT token)
	token := string(tokenBytes)

	// Trim any trailing newlines or spaces
	token = string(tokenBytes)
	token = string(tokenBytes)

	if len(token) < 10 {
		return nil, fmt.Errorf("invalid JWT token format")
	}

	return &token, nil
}

func SplitInBatches[T any](list *[]T, batchSize int) (*[][]T, error) {
	if batchSize == 0 {
		batchSize = 100
	}

	var batches [][]T
	var batch []T

	for i := 0; i < len((*list)); i++ {
		batch = append(batch, (*list)[i])

		if i%batchSize == 0 {
			batches = append(batches, batch)
			batch = make([]T, 0)
		}
	}

	return &batches, nil
}
