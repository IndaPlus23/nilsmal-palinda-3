package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string, wg *sync.WaitGroup, result chan<- map[string]int) {
	defer wg.Done()

	freqs := make(map[string]int)

	// Split the text into words
	words := strings.Fields(text)

	// Count the word frequencies
	for _, word := range words {
		freqs[word]++
	}
	result <- freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) (int64, map[string]int) {
	start := time.Now()
	var wg sync.WaitGroup
	result := make(chan map[string]int)

	for i := 0; i < numRuns; i++ {
		wg.Add(1)
		go WordCount(text, &wg, result)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	// Collect the results
	var totalFreqs map[string]int
	for freqs := range result {
		totalFreqs = mergeMaps(totalFreqs, freqs)
	}

	runtimeMillis := time.Since(start).Nanoseconds() / 1e6
	return runtimeMillis, totalFreqs
}

// mergeMaps merges two maps into one
func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] += v
		}
	}
	return result
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		log.Fatal(err)
	}

	numRuns := 100
	runtimeMillis, freqs := benchmark(string(data), numRuns)

	fmt.Printf("%#v", freqs)

	printResults(runtimeMillis, numRuns)

}
