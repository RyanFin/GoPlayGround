package main

import (
  "fmt"
)

func main(){
  var n int

  fmt.Scan(&n)

  for i := 0; i < n; i++{
    // n-i will reverse
    for j := 0; j < n-i-1; j++{
      // fmt.Print("out: ",n-i-1)
      fmt.Print(" ")
    }

    for j := 0; j <= i; j++{
      // fmt.Print("i: " , i)
      fmt.Print("#")
    }
    fmt.Println()
  }

}
