package main

import (
	"fmt"
)

func main() {
	input := "anagram"
	output := "nagaram"

	inputByte := []rune(input)
	outputByte := []rune(output)

	if anagram(inputByte, outputByte) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func anagram(input, output []rune) bool {
	if len(input) != len(output) {
		return false
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(output); j++ {
			if input[i] == output[j] {
				output[j] = 0
				break
			}
		}
	}
	return isAllZero(output)
}

func isAllZero(output []rune) bool {
	for _, r := range output {
		if r != 0 {
			return false
		}
	}
	return true
}
