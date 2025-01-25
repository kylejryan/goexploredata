package dataframe

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gota/gota/dataframe"
)

type DataFrame struct {
	DF dataframe.DataFrame
}

// Load reads a CSV file and returns a DataFrame object.
func Load(filePath string) (*DataFrame, error) {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Get absolute path
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Debug prints
	fmt.Printf("Current Working Directory: %s\n", cwd)
	fmt.Printf("Provided File Path: %s\n", filePath)
	fmt.Printf("Absolute File Path: %s\n", absPath)

	// Check if file exists
	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", absPath)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)
	if df.Err != nil {
		return nil, fmt.Errorf("failed to read CSV: %v", df.Err)
	}

	return &DataFrame{DF: df}, nil
}
