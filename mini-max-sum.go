package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {

    minSum, maxSum := 0, 0

    var s []int
    for _, e := range arr{
        // append each converted elem as an int
        s = append(s, int(e))
    }

    sort.Ints(s)
    // fmt.Println(s)

    for i := 0; i < len(s)-1; i++{
        minSum = minSum + s[i]
    }

    for i := len(s)-1; i > 0; i--{
        maxSum = maxSum + s[i]
    }

    fmt.Print(minSum , " " , maxSum)
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

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
