package main

import (
  "fmt"
  "math"
)

func main(){
  var n int
  fmt.Scan(&n)

  fmt.Println("myoutput: " , n)

  // create a multi-dimensional array, with allocation for user-specified input size 'n'
  a := make([][]int, n)
  fmt.Println("a : " , a)

  // sum digits for final answer
  lsum, rsum := 0, 0

  for i := 0; i < n; i++{
    // allocate a second array with the user-specified value in array []int
    a[i] = make([]int, n)
    for j := 0; j < n; j++{
      // take the input for each character in array
      fmt.Scan(&a[i][j])
      // rule 1 (top-left to bottom-right diagonal)
      if(i == j){
        lsum += a[i][j]
      }

      //rule 2 (top-right to bottom-left diagonal)
      if (i + j == n - 1){
        rsum += a[i][j]
      }

    }
  }

// fmt.Println(a)

diff := math.Abs(float64(lsum - rsum))

fmt.Println(diff)


}
