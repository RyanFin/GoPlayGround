package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the matchingStrings function below.
func matchingStrings(strings []string, queries []string) []int32 {

	var ret []int32
	var counter int32

	for _, q := range queries {
		counter = 0
		for _, e := range strings {
			if q == e {
				counter++
			}
		}
		ret = append(ret, counter)
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	stringsCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var strings []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := readLine(reader)
		strings = append(strings, stringsItem)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(strings, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
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
		panic(err)
	}
}
