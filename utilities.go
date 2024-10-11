package utils

import (
	"encoding/json"
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

	log.Printf(
		"Splitting list of %d items into %d batches of %d items (max) each\n",
		len(*list),
		len(*list)/batchSize,
		batchSize,
	)

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

func SaveGECsToFile[T any](name string, data *[]T) error {
	jsondata, err := json.Marshal(data)

	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.json", name)
	log.Printf("Trying to create file '%s'", fileName)

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	log.Printf("Trying to write data to file '%s'", fileName)
	_, err = file.Write(jsondata)

	if err != nil {
		return err
	}

	log.Println("File successfully written. Closing...")
	return nil
}
