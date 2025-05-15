package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := "anagraM"
	output := "nagaram"
	input = strings.ToLower(input)
	output = strings.ToLower(output)

	inputByte := []rune(input)
	outputByte := []rune(output)

	if anagramWithSorting(inputByte, outputByte) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

// ---------------------------------------------------------------------------
// func anagram(input, output []rune) bool {
// 	if len(input) != len(output) {
// 		return false
// 	}

// 	for i := 0; i < len(input); i++ {
// 		for j := 0; j < len(output); j++ {
// 			if input[i] == output[j] {
// 				output[j] = 0
// 				break
// 			}
// 		}
// 	}
// 	return isAllZero(output)
// }

// func isAllZero(output []rune) bool {
// 	for _, r := range output {
// 		if r != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

//-----------------------------------------------------------------------------------------

//way2 --> using HASH TABLE

// func anagram(input, output []rune) bool {
// 	if len(input) != len(output) {
// 		return false
// 	}

// 	counts := make(map[rune]int)
// 	for _, r := range input {
// 		counts[r]++
// 	}
// 	for _, r := range output {
// 		counts[r]--
// 		if counts[r] < 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

//--------------------------------------------------------

//Way 3 :: sorting technique

func anagramWithSorting(input, output []rune) bool {
	if len(input) != len(output) {
		return false
	}

	sort.Slice(input, func(i int, j int) bool {
		return input[i] < input[j]
	})

	sort.Slice(output, func(i int, j int) bool {
		return output[i] < output[j]
	})

	return strings.EqualFold(string(input), string(output))

}
