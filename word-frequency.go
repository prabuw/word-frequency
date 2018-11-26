package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("big.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCountMap := make(map[string]int)

	for scanner.Scan(){
		word := scanner.Text()
		wordCountMap[word] += 1
	}

	printMap(wordCountMap)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printMap(wordCountMap map[string]int){
	for key, value := range wordCountMap {
		fmt.Println(key, ": ", value)
	}
}
