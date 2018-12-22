package main

import "fmt"

func printMap(wordCountMap map[string]int, maxWordLength int){
	for key, value := range wordCountMap {
		fmt.Printf("%*v : %v\n", -maxWordLength, key, value)
	}

	fmt.Println(len(wordCountMap))
}
