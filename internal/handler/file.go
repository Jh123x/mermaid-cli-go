package handler

import (
	"io"
	"os"
)

func getInputData(inputFile string) (string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
