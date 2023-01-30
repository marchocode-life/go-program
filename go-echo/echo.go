package main

import (
	"fmt"
	"os"
)

func main() {
	// get args length
	var length = len(os.Args)
	fmt.Println(length)

	// [1,length)
	var str = os.Args[1:]

	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(str[i])
	// }

	for index, v := range str {
		fmt.Printf("index=%d,value=%s\n", index, v)
	}

}
