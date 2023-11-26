package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

func parseLogLine(logLine string) (string, string, string, string) {
	parts := strings.Split(logLine, " | ")

	if len(parts) != 4 {
		panic("A log line is invalid")
	}

	// Extract log level, message, and IP address
	timestamp := parts[0]
	logLevel := parts[1]
	message := parts[2]
	ip := strings.TrimPrefix(parts[3], "IP: ")

	return timestamp, logLevel, message, ip
}

/*
 * Complete the 'logParser' function below.
 *
 * The function accepts following parameters:
 *  1. STRING inputFileName
 *  2. STRING normalFileName
 *  3. STRING errorFileName
 */
func logParser(inputFileName string, normalFileName string, errorFileName string) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic("Failed to open input file")
	}
	defer inputFile.Close()

	normalFile, err := os.OpenFile(normalFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Failed to open normal file")
	}
	defer normalFile.Close()

	errorFile, err := os.OpenFile(errorFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Failed to open error file")
	}
	defer errorFile.Close()

	normalLogLinesChannel := make(chan string)
	errorLogLinesChannel := make(chan string)
	var logWriterWaitGroup sync.WaitGroup
	logWriterWaitGroup.Add(2)
	go func() {
		defer logWriterWaitGroup.Done()
		logWriter(normalFile, normalLogLinesChannel)
	}()
	go func() {
		defer logWriterWaitGroup.Done()
		logWriter(errorFile, errorLogLinesChannel)
	}()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		logLine := scanner.Text()

		_, logLevel, _, _ := parseLogLine(logLine)
		if logLevel == "ERROR" {
			errorLogLinesChannel <- logLine
		} else {
			normalLogLinesChannel <- logLine
		}
	}
	if err := scanner.Err(); err != nil {
		panic("Failed to read input file")
	}

	close(normalLogLinesChannel)
	close(errorLogLinesChannel)

	logWriterWaitGroup.Wait()
}

func logWriter(writer io.Writer, logLinesChannel chan string) {
	for logLine := range logLinesChannel {
		_, err := writer.Write([]byte("\n" + logLine))
		if err != nil {
			panic("Failed to write to log file")
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	file, err := os.Create(filename)
	checkError(err)
	_, err = os.Create(errorFilename)
	checkError(err)
	_, err = os.Create(normalFilename)
	checkError(err)
	for i := 0; i < n; i++ {
		inputString := readLine(reader)
		_, err = file.WriteString(inputString)
		checkError(err)
	}
	defer file.Close()
	logParser(filename, normalFilename, errorFilename)
	fmt.Printf("ERROR:")
	errorContent, errErr := ioutil.ReadFile(errorFilename)
	checkError(errErr)
	fmt.Println(string(errorContent))
	fmt.Printf("\nNORMAL:")
	normalContent, errNormal := ioutil.ReadFile(normalFilename)
	checkError(errNormal)
	fmt.Println(string(normalContent))

}

const filename = "output"
const normalFilename = "normal"
const errorFilename = "error"

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return string(str) + "\n"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
