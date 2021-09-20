package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("press q to exit")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		cleaningText := strings.ReplaceAll(text, "\n", "")

		if strings.TrimSpace(cleaningText) == "q" {
			break
		} else {
			UniqueLexicograph(strings.TrimSpace(cleaningText))
		}
	}
}

func UniqueLexicograph(words string) {
	uniqueCharacter := make([]string, 0, len(words))
	for _, character := range words {
		charString := string(character)
		if !contains(uniqueCharacter, charString) {
			uniqueCharacter = append(uniqueCharacter, charString)
		}
	}
	outputBuilder(uniqueCharacter, "First Occurrence From Left to Right: ")
	orderedLexicograph(uniqueCharacter)

}

func orderedLexicograph(uniqueSlice []string) {
	sort.Slice(uniqueSlice, func(i, j int) bool {
		return uniqueSlice[i] < uniqueSlice[j]
	})
	outputBuilder(uniqueSlice, "First in Lexicograph order: ")
}

func contains(uniqueSlice []string, word string) bool {
	for _, val := range uniqueSlice {
		if val == word {
			return true
		}
	}
	return false
}

func outputBuilder(finishedSlice []string, message string) {
	output := ""
	for _, character := range finishedSlice {
		output = output + character
	}
	fmt.Println(message + output)
}
