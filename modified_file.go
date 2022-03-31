package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FileStructure struct {
	CompletePath string
	Information  os.FileInfo
}

func insertSorted(fileList *list.List, fileStruct FileStructure) {
	if fileList.Len() == 0 {
		fileList.PushFront(fileStruct)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileStruct.Information.ModTime().Before(element.Value.(FileStructure).Information.ModTime()) {
			fileList.InsertBefore(fileStruct, element)
			return
		}
		fileList.PushBack(fileStruct)
	}
}

func FetchFilesInDirRecursivelyBySize(fileList *list.List, path string) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: " + err.Error())
	}
	for _, dirFiles := range dirFiles {
		CompletePath := filepath.Join(path, dirFiles.Name())
		if dirFiles.IsDir() {
			FetchFilesInDirRecursivelyBySize(fileList, filepath.Join(path, dirFiles.Name()))
		} else if dirFiles.Mode().IsRegular() {
			insertSorted(fileList, FileStructure{CompletePath: CompletePath, Information: dirFiles})
		}
	}
}

func main() {
	fileList := list.New()
	FetchFilesInDirRecursivelyBySize(fileList, "/")
	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(FileStructure).Information.ModTime())
		fmt.Printf("%s\n", element.Value.(FileStructure).CompletePath)
	}
}
