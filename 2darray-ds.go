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

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
    var ret int32

    // Get all the top left values in the multi-dimensional array and run sumHourGlassTl() func
    fmt.Println(sumHourGlassTl(3,3, arr))

    var a []int32

    // co-ordinates (0,0) - (3,3)
    for i := 0; i < 4; i++{
        for j:= 0; j < 4; j++{
            fmt.Println(i, j)
            a = append(a, sumHourGlassTl(int32(i), int32(j), arr))
        }
    }

    fmt.Println(a)
    fmt.Println(len(a))

    var s []int

    // Convert each element in array from int32 to int
    for _ , e := range a{
      // append individually to array
        s = append(s, int(e))
    }

    sort.Sort(sort.Reverse(sort.IntSlice(s)))

    ret = int32(s[0])


    return ret
}

// Create a function that will sum up an hour glass given the top-left co-ordinate
// return the value as an int32
func sumHourGlassTl(row int32, col int32, arr [][]int32) int32{
    // get top row
    var sum int32

    // map out each corodinate value for an hourglass depending on the top-left value provided to the sumHourGlassTl() function
    sum = sum + arr[row][col]
    sum = sum + arr[row][col+1]
    sum = sum + arr[row][col+2]

    //get middle row
    sum = sum + arr[row+1][col+1]

    //get bottom row
    sum = sum + arr[row+2][col]
    sum = sum + arr[row+2][col+1]
    sum = sum + arr[row+2][col+2]

    fmt.Println(sum)

    return sum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

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
