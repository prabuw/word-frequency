package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	defer printTimeTaken(now)

	file, err := os.Open("big.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	maxWordLength := 0
	wordCountMap := make(map[string]int)

	for scanner.Scan(){
		word := scanner.Text()
		wordCountMap[word] += 1

		maxWordLength = updatedMaxWordLength(maxWordLength, word)
	}

	printMap(wordCountMap, maxWordLength)
}

func printTimeTaken(now time.Time) {
	fmt.Println()
	fmt.Printf("Time taken: %v\n", time.Since(now))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func updatedMaxWordLength(currentMaxWordLength int, word string) int{
	wordLength := len(word)

	if wordLength > currentMaxWordLength {
		return wordLength
	}

	return currentMaxWordLength
}

func printMap(wordCountMap map[string]int, maxWordLength int){
	for key, value := range wordCountMap {
		fmt.Printf("%*v : %v\n", -maxWordLength, key, value)
	}
}
