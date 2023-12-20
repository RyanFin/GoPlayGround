package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {

	var iArr []int

	for _, e := range arr {
		iArr = append(iArr, int(e))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(iArr)))
	// Write your code here
	minSlice := iArr[0:4]
	maxSlice := iArr[1:5]

	a := addSlice(minSlice)
	b := addSlice(maxSlice)

	fmt.Printf("%v %v", b, a)
}

func addSlice(arr []int) int {
	var ans int
	for _, e := range arr {
		ans += e
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
