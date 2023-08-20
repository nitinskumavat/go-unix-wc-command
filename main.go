package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lineCount, bytesCount, wordCount int64
	var fileName string
	fileName = "test.txt"
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file stats", err)
		return
	}

	if fileName == "" {
		fileName = fileInfo.Name()
	}
	bytesCount = fileInfo.Size()

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordCount += (int64)(len(strings.Fields(scanner.Text())))
		lineCount++
	}
	fmt.Println(lineCount, wordCount, bytesCount)
}
