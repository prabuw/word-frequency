package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
	"unicode/utf8"
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

func splitFunc (data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) || isSpecial(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isSpace(r) || isSpecial(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}


func isSpecial(r rune) bool {
	return unicode.IsSymbol(r) || unicode.IsPunct(r) || unicode.IsNumber(r)
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

func printMap(wordCountMap map[string]int, maxWordLength int){
	for key, value := range wordCountMap {
		fmt.Printf("%*v : %v\n", -maxWordLength, key, value)
	}

	fmt.Println(len(wordCountMap))
}
