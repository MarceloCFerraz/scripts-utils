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
		return nil, fmt.Errorf("error reading token file: %v\n", err)
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
	if batchSize <= 0 {
		batchSize = 100
	}

	log.Printf(
		"Splitting list of %d items into batches of %d items (max) each\n",
		len(*list),
		batchSize,
	)

	var batches [][]T

	for i := 0; i < len(*list); i += batchSize {
		end := i + batchSize
		if end > len(*list) {
			end = len(*list)
		}
		batches = append(batches, (*list)[i:end])
	}

	return &batches, nil
}


func SaveDataToFile(name, extension string, data []byte) error {
	fileName := fmt.Sprintf("%s%s", name, extension)
	log.Printf("Trying to create file '%s'\n", fileName)

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	log.Printf("Trying to write data to file '%s'\n", fileName)
	_, err = file.Write(data)

	if err != nil {
		return err
	}

	log.Println("File successfully written. Closing...")
	return nil
}
