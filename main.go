package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	s := flag.Args()
	fmt.Println("\n============================")
	fmt.Println("Word count:", Counter(s[0]))
	fmt.Println("============================")
}
