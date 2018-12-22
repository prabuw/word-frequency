package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("big.txt")
	check(err)

	defer file.Close()

	now := time.Now()
	defer printTimeTaken(now)

	scanner := bufio.NewScanner(file)

	//scanner.Split(bufio.ScanWords)
	scanner.Split(splitFunc)

	maxWordLength := 0
	wordCountMap := make(map[string]int)

	for scanner.Scan(){
		word := scanner.Text()
		wordCountMap[word] += 1

		maxWordLength = updatedMaxWordLength(maxWordLength, word)
	}

	printMap(wordCountMap, maxWordLength)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printTimeTaken(now time.Time) {
	fmt.Println()
	fmt.Printf("Time taken: %v\n", time.Since(now))
}

func updatedMaxWordLength(currentMaxWordLength int, word string) int{
	wordLength := len(word)

	if wordLength > currentMaxWordLength {
		return wordLength
	}

	return currentMaxWordLength
}
