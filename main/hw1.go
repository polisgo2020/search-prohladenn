package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pathToFolderWithFiles := os.Args[1]
	data := make(map[string][]int)
	files, err := ioutil.ReadDir(pathToFolderWithFiles)
	check(err)
	for fileIndex, file := range files {
		fileBytes, err := ioutil.ReadFile(pathToFolderWithFiles + "/" + file.Name())
		check(err)
		fileText := string(fileBytes)
		words := strings.Split(fileText, " ")
		for i := 0; i < len(words); i++ {
			word := words[i]
			if data[word] == nil {
				sliceWithOneField := []int{fileIndex}
				data[word] = sliceWithOneField
			} else if data[word][len(data[word])-1] != fileIndex {
				data[word] = append(data[word], fileIndex)
			}
		}
	}
	outFile, err := os.Create(pathToFolderWithFiles + "/output.txt")
	check(err)
	defer outFile.Close()
	for key, value := range data {
		stringSlice := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(value)), ","), "[]")
		outLine := fmt.Sprintf("\"%s\":{%s}\n", key, stringSlice)
		_, err := outFile.WriteString(outLine)
		check(err)
	}
}
