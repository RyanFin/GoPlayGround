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
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

//
func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	var arrZero []int32
	var arrOne []int32
	var lastAnswer int32
	var ret []int32
	lastAnswer = 0
	// fmt.Println(arrOne)
	// fmt.Println(queries)
	// time complexity runtimes: O(q * e)
	for _, e := range queries {

		// run query 1: if 1 x y
		if e[0] == 1 {
			x := e[1]
			if XorCalculation(x, lastAnswer, n) == 0 {
				arrZero = append(arrZero, e[2])
			}

			if XorCalculation(x, lastAnswer, n) == 1 {
				arrOne = append(arrOne, e[2])
			}

		}

		// fmt.Println("arr0: " , arrZero)
		// fmt.Println("arr1:" ,arrOne)

		// run query 2: if 2 x y
		if e[0] == 2 {
			x := e[1]
			// index 0 specified by y
			if e[2] == 0 {
				fmt.Println("x ", x)
				fmt.Println("lastans: ", lastAnswer)
				fmt.Println("xorTest: ", XorCalculation(x, lastAnswer, n))
				if XorCalculation(x, lastAnswer, n) == 0 {
					// y value is the index
					lastAnswer = arrZero[0]
					fmt.Println(lastAnswer)
					ret = append(ret, lastAnswer)
				}

				if XorCalculation(x, lastAnswer, n) == 1 {
					lastAnswer = arrOne[0]
					fmt.Println(lastAnswer)
					ret = append(ret, lastAnswer)
				}

			}

			// index 1 specified by y
			if e[2] == 1 {
				fmt.Println("x2: ", x)
				fmt.Println("lastans2: ", lastAnswer)
				fmt.Println("xorTest 2: ", XorCalculation(x, lastAnswer, n))
				if XorCalculation(x, lastAnswer, n) == 0 {
					// y value is the index
					lastAnswer = arrZero[1]
					fmt.Println(lastAnswer)
					ret = append(ret, lastAnswer)
				}
				if XorCalculation(x, lastAnswer, n) == 1 {
					lastAnswer = arrOne[1]
					fmt.Println(lastAnswer)
					ret = append(ret, lastAnswer)
				}

			}

		}

	}

	fmt.Println(arrZero)
	fmt.Println(arrOne)

	return arrZero
}

func XorCalculation(x int32, lastAnswer int32, n int32) int32 {
	var ret int32
	var a bool
	var b bool

	// x
	if x == 0 {
		a = false
	}
	if x == 1 {
		a = true
	}

	// last answer
	if lastAnswer == 1 {
		b = true
	}

	if lastAnswer == 0 {
		b = false
	}

	// XOR operation
	if (a || b) && !(a && b) {
		ret = 1
	} else {
		ret = 0
	}

	return ret % n
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := dynamicArray(n, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
		panic(err)
	}
}
