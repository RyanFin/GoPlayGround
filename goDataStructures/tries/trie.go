package main

import "fmt"

// AlphabetSize represents the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	// set up initial trie with root node configured
	result := &Trie{root: &Node{}}

	return result
}

// Insert a word into the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	// start at root node
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		// fmt.Println(charIndex)
		// fmt.Println(currentNode.children[charIndex])
		// If the character node is null for specific character
		if currentNode.children[charIndex] == nil {
			// Insert a new node here
			currentNode.children[charIndex] = &Node{}
		}

		// Update the current node
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true // mark that this is the last node in the word
}

// Search

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	// start at root node
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		// If the character node is null for specific character
		if currentNode.children[charIndex] == nil {
			// return false as word is not in the trie
			return false
		}

		// Update the current node
		currentNode = currentNode.children[charIndex]
	}

	// Search will check for last node
	if currentNode.isEnd == true {
		return true
	}

	return false

}

func main() {
	// Generate basic trie
	InitTrie()

	myTrie := InitTrie()

	myWordList := []string{"hello", "dragon", "cat", "dog", "apple"}

	for _, e := range myWordList {
		// add words to trie
		myTrie.Insert(e)
	}

	res := myTrie.Search("dog")
	fmt.Println(res)

}
