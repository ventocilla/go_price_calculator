package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to open file")
	}
	file.Close()
	return lines, nil

}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	// Encode: data -> json ..
	// Encoder is used to serialize Go data structures into JSON format.
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON.")
	}

	file.Close()
	return nil

}
