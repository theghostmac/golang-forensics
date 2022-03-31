package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	filename := "stegan-image.jpg"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	buffReader := bufio.NewReader(file)

	fileStats, _ := file.Stat()
	for i := int64(0); i < fileStats.Size(); i++ {
		aByte, err := buffReader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		if aByte == '\x50' {
			byteSlice := make([]byte, 3)
			byteSlice, err = buffReader.Peek(3)
			if err != nil {
				log.Fatal(err)
			}
			if bytes.Equal(byteSlice, []byte{'\x4b', '\x03', '\x04'}) {
				log.Printf("Found a zip signature at byte %d.", i)
			}
		}
	}
}
