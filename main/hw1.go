package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func splitter(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect number of input arguments")
	}
	pathToFolderWithFiles := os.Args[1]
	data := make(map[string][]string)
	files, err := ioutil.ReadDir(pathToFolderWithFiles)
	check(err)
	for _, file := range files {
		pathToInputFile := filepath.Join(pathToFolderWithFiles, file.Name())
		fileBytes, err := ioutil.ReadFile(pathToInputFile)
		check(err)
		fileText := string(fileBytes)
		words := strings.FieldsFunc(fileText, splitter)
		for i := range words {
			word := strings.ToLower(words[i])
			if data[word] == nil {
				data[word] = []string{file.Name()}
			} else if data[word][len(data[word])-1] != file.Name() {
				data[word] = append(data[word], file.Name())
			}
		}
	}
	pathToOutputFile := filepath.Join(pathToFolderWithFiles, "/output.txt")
	outFile, err := os.Create(pathToOutputFile)
	check(err)
	defer outFile.Close()
	for key, value := range data {
		stringSlice := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(value)), ","), "[]")
		outLine := fmt.Sprintf("%s:{%s}\n", key, stringSlice)
		_, err := outFile.WriteString(outLine)
		check(err)
	}
}
