package main

import (
	"fmt"
	"strings"
)

var dataA = []string{
	"Golang is efficient",
	"Concurrency in Go",
	"Error handling in Go",
	"Go programming language",
	"Goroutines and channels",
	"Efficient code in Go",
}

var dataB = []string{
	"Java programming",
	"Object-oriented programming",
	"Java concurrency",
	"Big data and privacy",
	"Java vs. Go",
	"Java is versatile",
}

func wordCount(data []string) map[string]int {
	wordCounts := make(map[string]int)

	for _, line := range data {
		words := strings.Fields(line)
		for _, word := range words {
			cleanedWord := strings.ToLower(strings.Trim(word, ".,"))
			if cleanedWord != "" {
				wordCounts[cleanedWord]++
			}
		}
	}

	return wordCounts
}

func main() {
	result1 := wordCount(dataA)
	result2 := wordCount(dataB)

	fmt.Println("Word Count for Dataset 1:")
	for word, count := range result1 {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Println("\nWord Count for Dataset 2:")
	for word, count := range result2 {
		fmt.Printf("%s: %d\n", word, count)
	}

	// Common Word Count
	commonWords := make(map[string]struct{})
	for word := range result1 {
		if _, exists := result2[word]; exists {
			commonWords[word] = struct{}{}
		}
	}

	fmt.Println("\nCommon Word Count:")
	for word := range commonWords {
		count1 := result1[word]
		count2 := result2[word]
		fmt.Printf("%s: Dataset 1 - %d, Dataset 2 - %d\n", word, count1, count2)
	}
}
