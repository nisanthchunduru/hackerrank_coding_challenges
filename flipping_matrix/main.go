package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func maxPossibleValueForPosition(matrix [][]int, rowIndex int, columnIndex int) int {
	maxPossibleValue := matrix[rowIndex][columnIndex]
	rowCount := len(matrix)
	columnCount := len(matrix[0])

	if matrix[rowIndex][columnCount - columnIndex - 1] > maxPossibleValue {
		maxPossibleValue = matrix[rowIndex][columnCount - columnIndex - 1]
	}

	if matrix[rowCount - rowIndex - 1][columnIndex] > maxPossibleValue {
		maxPossibleValue = matrix[rowCount - rowIndex - 1][columnIndex]
	}

	if matrix[rowCount - rowIndex - 1][columnCount - columnIndex - 1] > maxPossibleValue {
		maxPossibleValue = matrix[rowCount - rowIndex - 1][columnCount - columnIndex - 1]
	}

	return maxPossibleValue;
}

func flippingMatrix(matrix [][]int) int {
    // Write your code here

		rowCount := len(matrix)
		columnCount := len(matrix[0])
		firstQuardantMaxSum := 0;

		for i := 0; i < (rowCount / 2); i++ {
			for j := 0; j < (columnCount / 2); j++ {
				firstQuardantMaxSum = firstQuardantMaxSum + maxPossibleValueForPosition(matrix, i, j)
			}
		}

		return firstQuardantMaxSum;
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int(nTemp)

        var matrix [][]int
        for i := 0; i < 2 * int(n); i++ {
            matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

            var matrixRow []int
            for _, matrixRowItem := range matrixRowTemp {
                matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
                checkError(err)
                matrixItem := int(matrixItemTemp)
                matrixRow = append(matrixRow, matrixItem)
            }

            if len(matrixRow) != 2 * int(n) {
                panic("Bad input")
            }

            matrix = append(matrix, matrixRow)
        }

        result := flippingMatrix(matrix)

        fmt.Fprintf(writer, "%d\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
