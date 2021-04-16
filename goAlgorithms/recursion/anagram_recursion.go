package main

import "fmt"

var (
	size    int
	count   int
	arrChar = []string{}
)

func displayWord() {
	if count < 99 {
		fmt.Print(" ")
	}
	if count < 9 {
		fmt.Print(" ")
	}
	count += 1
	fmt.Print(count, " ")

	for i := 0; i < size; i++ {
		fmt.Print(arrChar[i])
	}

	fmt.Print("   ")

	if count%6 == 0 {
		fmt.Println("")
	}

}

func rotate(newSize int) {
	var j int
	position := size - newSize
	temp := arrChar[position]

	for j = position + 1; j < size; j++ {
		arrChar[j-1] = arrChar[j]
	}
	arrChar[j-1] = temp

}

func doAnagram(newSize int) {
	if newSize == 1 {
		return
	}

	for j := 0; j < size; j++ {
		doAnagram(newSize - 1)
		displayWord()
		rotate(newSize)
	}
}

func main() {
	var input string
	fmt.Print("Enter a word...")
	fmt.Scan(&input)
	size = len(input)
	count = 0

	for i := 0; i < size; i++ {
		// fmt.Println(string(input[i]))
		arrChar = append(arrChar, string(input[i]))
	}

	doAnagram(size)

}
