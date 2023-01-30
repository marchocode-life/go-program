package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counters := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// get one line from stdand input
	// return true if get some charest
	for input.Scan() {
		// get line content 
		counters[input.Text()]++
	}

	for line, value := range counters {

		if value > 1 {
			fmt.Printf("line=%s,count number=%d\n", line, value)
		}

	}

}
