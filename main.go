package main

import (
	"flag"
	"fmt"
	"os"
)

var filePath = flag.String("path", "", "File path for counting file content")

func main() {
	flag.Parse()

	s := flag.Args()

	if *versionFlag {
		fmt.Println(version)
		return
	}

	if len(s) != 0 {
		fmt.Println("\n============================")
		fmt.Println("Word count:", Counter(s[0]))
		fmt.Println("============================")
		return
	}

	if len(*filePath) == 0 {
		// TODO Show usage
		fmt.Println("Invalid file path.")
		os.Exit(1)
	}

	content, err := ReadFile(*filePath)
	if err != nil {
		fmt.Println("Error while reading file.")
		os.Exit(1)
	}

	fmt.Println("\n============================")
	fmt.Println("Word count:", Counter(content))
	fmt.Println("============================")

}
