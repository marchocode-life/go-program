package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	store := make(map[string]int)
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("file not found")
		return
	}

	data, err := ioutil.ReadFile(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "file read faild", err)
	}

	for _, str := range strings.Split(string(data), "\n") {
		store[str]++
	}

	for key, value := range store {
		fmt.Printf("key=%s,value=%d\n", key, value)
	}

}
