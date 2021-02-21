package main

import (
  "fmt"
  "strings"
)

func main(){
  staircase(6)
}


func staircase(n int32) {

  // convert int32 to int
  var k = int(n)
  if(k > 0){
    for i := 1; i < k+1; i++{

      for j := k+1; j > i; j--{
        fmt.Println(j)
      }
    // Add 'i' spaces to the repeat string
    var spacer = strings.Repeat(" ", i)
    var hashes = strings.Repeat("#", i)
    // fmt.Println(strings.Repeat("#", i))

    // join together
    s := []string{spacer, hashes}
    fmt.Println(strings.Join(s, ""))

    }
  }
}
