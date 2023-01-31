package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counters := make(map[string]int)
	args := os.Args[1:]

	if len(args) == 0 {
		countLines(os.Stdin, counters)
	} else {

		file, err := os.Open(args[0])

		if err != nil {
			fmt.Println("error")
		} else {
			countLines(file, counters)
			file.Close()
		}

	}
	printResult(counters)
}

func countLines(f *os.File, counters map[string]int) {

	input := bufio.NewScanner(f)

	for input.Scan() {
		counters[input.Text()]++
	}
}

func printResult(counters map[string]int) {

	for key, value := range counters {
		fmt.Printf("key=%s,value=%d\n", key, value)
	}

}
