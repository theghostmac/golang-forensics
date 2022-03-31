package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printSyntax() {
	fmt.Println("Usage: " + os.Args[0] + " <filepath>")
	fmt.Println("Example: " + os.Args[0] + " filename.txt")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printSyntax()
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	filename := checkArgs()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
}
