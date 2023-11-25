package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func zigZagIntegers(integers []int) {
	sort.Ints(integers)
	middleIndex := len(integers) / 2
	lastIndex := len(integers) - 1
	reverseSlice(integers[middleIndex:(lastIndex + 1)])
}

func joinIntArrayWithSpace(integers []int) string {
	stringifiedIntegers := make([]string, len(integers))
	for i, num := range integers {
		stringifiedIntegers[i] = strconv.Itoa(num)
	}
	return strings.Join(stringifiedIntegers, " ")
}

func main() {
	var testCaseCount int
	_, err := fmt.Scan(&testCaseCount)
	if err != nil {
		panic("Failed to read test case count")
	}
	output := ""
	for i := 0; i < testCaseCount; i++ {
		var integerCount int
		_, err = fmt.Scan(&integerCount)
		if err != nil {
			panic("Failed to read test case integer count")
		}

		integers := make([]int, integerCount)
		for j := 0; j < integerCount; j++ {
			var integer int
			// _, err = fmt.Scanf("%d", &integer); if err != nil {
			//     panic("Failed to read integer")
			// }
			_, err = fmt.Scan(&integer)
			if err != nil {
				panic("Failed to read integer")
			}
			integers[j] = integer
		}

		zigZagIntegers(integers)

		spaceSeparatedIntegers := joinIntArrayWithSpace(integers)
		if i == (testCaseCount - 1) {
			output = fmt.Sprintf("%s%s", output, spaceSeparatedIntegers)
		} else {
			output = fmt.Sprintf("%s%s\n", output, spaceSeparatedIntegers)
		}
	}

	outputPath := os.Getenv("OUTPUT_PATH")
	if outputPath != "" {
		ioutil.WriteFile(outputPath, []byte(output), 0644)
	} else {
		fmt.Printf("%s", output)
	}
}
