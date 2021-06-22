package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



var (
    PARENTHESIS = map[string]string{
        "(":")",
        "{": "}",
        "[": "]",
    }
)

type Element struct {
    value interface{}
    next  *Element
}

type Stack struct {
    top  *Element
    size int
}

func (stack *Stack) Len() int {
    return stack.size
}

func (stack *Stack) Push(value interface{}) {
    stack.top = &Element{value, stack.top}
    stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
    if stack.size > 0 {
        value, stack.top = stack.top.value, stack.top.next
        stack.size--
        return
    }
    return nil
}

// Complete the braces function below.
func braces(values []string) []string {
    
    var ret []string
    var stack Stack
    for _, e := range values{
        fmt.Println("values: " , e)
        confirmer := []bool{}
        for _ , p := range strings.Split(e, ""){
            fmt.Println(p)
            switch p{
                case "{": 
                stack.Push(p)
                break
                case "[":
                stack.Push(p) 
                break
                case "(":
                stack.Push(p)
                break
                case "}":
                // fmt.Println(check(stack, p))
                confirmer = append(confirmer, check(stack, p))
                break
                case "]":
                // fmt.Println(check(stack, p))
                confirmer = append(confirmer, check(stack, p))
                break
                case ")":
                // fmt.Println(check(stack, p))
                confirmer = append(confirmer, check(stack, p))
                break
            }
            fmt.Println(confirmer)
            
            v := contains(confirmer, false)
            if(v){
                ret = append(ret , "NO")
            } else{
                ret = append(ret, "YES")
            }
              
        }
    }
    
    // fmt.Println("Stack: " , stack)
    fmt.Println("ret: " , ret)
    
    return ret
}

func check(stack Stack, p string) bool{
    if stack.Len() > 0{
        chx := stack.Pop().(string)
        if p == "}" && chx != "{" || p == "]" && chx != "[" || p == ")" && chx != "("{
            return false
        } else{
            return true
        }
    }
    return false
}

func contains(slice []bool, item bool) bool {
    set := make(map[bool]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}




func main() {