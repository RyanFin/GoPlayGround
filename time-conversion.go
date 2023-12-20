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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	var res string
	// get last two characters
	dayTimeStr := s[len(s)-2:]
	timeStr := s[0:8]
	fmt.Println(timeStr)
	switch dayTimeStr {
	case "AM":
		// 12:01:00 AM
		t := s[0:2] // 12
		if t == "12" {
			t = "00"
		}

		res = fmt.Sprintf("%s%s", t, timeStr[2:])

	case "PM":
		// 07:05:45
		t := s[0:2] //07
		// convert to int
		iT, _ := strconv.Atoi(t)
		fmt.Println(iT)
		if iT != 12 {
			iT += 12
		}

		iTS := strconv.Itoa(iT)
		// 19
		res = fmt.Sprintf("%s%s", iTS, timeStr[2:])
		fmt.Println(res)

	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
