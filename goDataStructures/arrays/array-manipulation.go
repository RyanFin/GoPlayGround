package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {

	var ret int64

	// create a slice with n amount of zeros
	mainSlice := []int32{}
	var i int32
	for i = 0; i <= n; i++ {
		// append 'n' number of 0's to slice to begin
		mainSlice = append(mainSlice, 0)
	}

	mainSlice = addKValsToArr(mainSlice, queries)

	temp := mainSlice[0]
	for i = 0; i <= n; i++ {
		if mainSlice[i] > temp {
			temp = mainSlice[i]
		}
	}

	ret = int64(temp)

	return ret
}

func addKValsToArr(arr []int32, queries [][]int32) []int32 {
	for _, e := range queries {
		fmt.Println(e)
		a := e[0] - 1
		b := e[1] - 1
		k := e[2]

		// iterate array and change values
		var i int32
		for i = a; i <= b; i++ {
			arr[i] += k
		}

	}

	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

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
