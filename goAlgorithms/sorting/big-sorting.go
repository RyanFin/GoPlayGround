// hackerrank - Big Sorting
/*
-- Test Input ---
6 (number of elems in array)
31415926535897932384626433832795
1
3
10
3
5

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// strconv.Atoi()

// Complete the bigSorting function below.

// Bubblesort algorithm
func bigSorting(unsorted []string) []string {
	var out int
	var in int
	for out = len(unsorted) - 1; out > 1; out-- {
		for in = 0; in < out; in++ {
			a, _ := strconv.Atoi(unsorted[in])
			b, _ := strconv.Atoi(unsorted[in+1])
			if a > b {
				temp := unsorted[in]
				unsorted[in] = unsorted[in+1]
				unsorted[in+1] = temp
			}
		}
	}

	fmt.Println("BubbleSort Result: ", unsorted)
	return unsorted

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var unsorted []string

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		unsorted = append(unsorted, unsortedItem)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
		fmt.Println(err)
	}
}
