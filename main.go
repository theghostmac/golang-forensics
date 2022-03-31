package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	fileInfo, err = os.Stat("target.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name of file is: ", fileInfo.Name())
	fmt.Println("Size of file in bytes is: ", fileInfo.Size())
	fmt.Println("Permissions on this file is: ", fileInfo.Mode())
	fmt.Println("Last time of modification of the file is: ", fileInfo.ModTime())
	fmt.Println("The file is a directory: ", fileInfo.IsDir())
	fmt.Println("The system interface type of the file is: %T\n", fileInfo.Sys())
	fmt.Println("The system information is: : ", fileInfo.Sys())
}
