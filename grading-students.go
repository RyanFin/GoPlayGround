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
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
    var finalGrades  []int32
    // Write your code here
    for _ , e := range grades{
        // if grade is below 40
        if(e < 35){
            // Append the failing score directly to the return slice without rounding up
            finalGrades = append(finalGrades, e)
        } else{
            if(e % 5 == 0){
                // Do not round up
                finalGrades = append(finalGrades, e)
            }
            if(e % 5 == 1){
                // Do not round up
                finalGrades = append(finalGrades, e)
            }
            if(e % 5 == 2){
                // Do not round up
                finalGrades = append(finalGrades, e)
            }
            if(e % 5 == 3){
                // round up grade as ((grade - next multiple of 5) < 3)
                finalGrades = append(finalGrades, roundUpToNearestMultOfFive(e))
            }
             if(e % 5 == 4){
                // round up grade as ((grade - next multiple of 5) < 3)
                finalGrades = append(finalGrades, roundUpToNearestMultOfFive(e))
            }
        }
    }

    return finalGrades
}

func roundUpToNearestMultOfFive(n int32) int32{
    return (n + 4) / 5 * 5
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var grades []int32

    for i := 0; i < int(gradesCount); i++ {
        gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        gradesItem := int32(gradesItemTemp)
        grades = append(grades, gradesItem)
    }

    result := gradingStudents(grades)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
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
