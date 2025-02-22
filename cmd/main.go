package main

import (
	"XReader/internal/fileops"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "", "Path to the file you want to read")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Error: Please provide a file path using -file")
		flag.Usage()
		os.Exit(1)
	}
	content, err := fileops.ReadFile(*filePath)

	if err != nil {
		fmt.Println("Error: %v\\n", err)
		os.Exit(1)
	}

	fmt.Println("File Content:")
	fmt.Println(content)

}
