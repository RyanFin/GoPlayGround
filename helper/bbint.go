package main

import "fmt"

// main method
func main(){

    // variable for user input
    var input string
    fmt.Scan(&input)

    answerBool := isBalanced(input)

    fmt.Println(answerBool)
}

// desc - input: string
// return: bool
func isBalanced(a string) bool {

    // open and closed brace counters
    oBrace, cBrace := 0 , 0

    // for each character in the string
    for _, e := range a{
        // if the character is an open brace increase the value of the open brace var

        if(string(e) == "{"){
            oBrace ++
        }

        // if the character is a closed brace increase the value of the closed brace var
        if(string(e) == "}"){
            cBrace ++
        }

    }

    // if the amount of open and closed braces are the same return true
    if (oBrace == cBrace){
        return true
    } else{
        // otherwise return false
        return false
    }

    return false

}
