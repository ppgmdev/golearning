package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("could not open file")
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	defer file.Close() // executes once the method is finished

	scanner := bufio.NewScanner(file) // read line by line from an open file

	var lines []string // lines slice

	for scanner.Scan() { // scan until theres no more content to scan
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		//file.Close()
		return nil, errors.New("failed to read line in file")
	}

	//file.Close()
	return lines, err
}

func (fm FileManager) WriteResult(data interface{}) error { // any type
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close() // executes once the method is finished

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		//file.Close()
		return errors.New("failed to convert data to json")
	}

	//file.Close()

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
