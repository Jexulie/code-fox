package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// GetFileContents checks file, and it's size then gets the contents in byte
func GetFileContents(file string) ([]byte, error) {
	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, file)

	// Check file size if it's too big reject
	fileSize, err := checkAndGetSize(path)
	if err != nil {
		return nil, err
	}

	// If file is larger than 10 mb read by chunks
	if fileSize >= 1e7 {
		return readByChunks(path)
	}

	return readAtOnce(path)
}

// checkAndGetSize checks if file exists and returns the file size in bytes
func checkAndGetSize(filePath string) (int64, error) {
	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		// Can not find file
		if os.IsNotExist(err) {
			return 0, errors.New("file does not exist")
		}

		// Other issues like permission
		return 0, errors.New("error checking file")
	}

	return fileInfo.Size(), nil
}

// readByChunks reads big files by scanning and returns contents in bytes
func readByChunks(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open("example.txt")
	if err != nil {
		return nil, errors.New("error opening file")
	}

	defer file.Close()

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each line
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("error reading file")
	}

	return scanner.Bytes(), nil
}

// readAtOnce reads small files and returns contents in bytes
func readAtOnce(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}
