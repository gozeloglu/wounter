package main

import (
	"flag"
	"fmt"
	"os"
)

var filePath = flag.String("path", "", "File path for counting file content")

func main() {
	// Update flag.Usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usageText, os.Args[0], os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), footerText)
	}

	flag.Parse()

	s := flag.Args()

	if *helpFlag {
		flag.Usage()
		return
	}

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
		flag.Usage()
		os.Exit(1)
	}

	content, err := ReadFile(*filePath)
	if err != nil {
		fmt.Println("Error while reading file.")
		fmt.Println("\nSee 'wounder --help'")
		os.Exit(1)
	}

	fmt.Println("\n============================")
	fmt.Println("Word count:", Counter(content))
	fmt.Println("============================")

}
