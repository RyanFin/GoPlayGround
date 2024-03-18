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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var positives, negatives, zeros float64

	// fmt.Println("array: ", arr)
	for _, e := range arr {

		// fmt.Println("elem: ", e)
		if e > 0 {
			positives += 1
		}

		if e < 0 {
			negatives += 1
		}

		if e == 0 {
			zeros += 1
		}
	}

	fmt.Println("positives: ", positives)
	fmt.Println("negatives: ", negatives)
	fmt.Println("zeros: ", zeros)

	po := positives / float64(len(arr))

	ne := negatives / float64(len(arr))

	ze := zeros / float64(len(arr))

	decimalSixFigures(po)
	decimalSixFigures(ne)
	decimalSixFigures(ze)
}

func decimalSixFigures(a float64) {
	fmt.Printf("%.6f", a)
	fmt.Println("")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
