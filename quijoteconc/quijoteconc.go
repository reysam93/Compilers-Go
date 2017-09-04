package main

import (
	"io"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)


const BufSize = 100


type Line struct {
	text   string
	number int
}

type Dictionary struct {
	m    map[string][]Line
	name string
}


func splitWords(r rune) bool {
	return !unicode.IsLetter(r)
}

func addWordsToMap(lineChan chan Line, dictName string, dictChan  chan Dictionary) {
	dict := Dictionary{m: make(map[string][]Line), name: dictName}

	for line := range lineChan {
		words := strings.FieldsFunc(strings.ToLower(line.text), splitWords)
		if words == nil {
			continue
		}

		// deleting repeated words in a line
		wordsMap := make(map[string]struct{})
		for _, word := range words {
			wordsMap[word] = struct{}{}
		}
		for word := range wordsMap {
			dict.m[word] = append(dict.m[word], line)
		}
	}
	dictChan <- dict
}

func printMap(c chan Dictionary) {
	var keys []string
	dict := <-c

	for k := range dict.m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		lines := dict.m[k]
		fmt.Printf("%s\n", k)
		for _, l := range lines {
			fmt.Printf("\t%s:%d '%s'\n", dict.name, l.number, l.text)
		}
	}
}

func scanLines(reader io.Reader, c chan Line) {
	lineScanner := bufio.NewScanner(reader)
	for i := 1; lineScanner.Scan(); i++ {
		line := Line{}
		line.text = lineScanner.Text()
		line.number = i
		c<-line
	}
	close(c)
	if lineScanner.Err() != nil {
		log.Fatal(lineScanner.Err())
	}
}

func main() {
	lineChan := make(chan Line, BufSize)
	dictChan := make(chan Dictionary)

	if len(os.Args) > 2 {
		log.Fatal("usage: quijote [file_name]\nNote: if no file name is provided, the default name will be quijote.txt")		
	}
	name := "quijote.txt"
	if len(os.Args) == 2{
		name = os.Args[1]
		os.Args = os.Args[:1]
	}

	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	go scanLines(file, lineChan)
	go addWordsToMap(lineChan, name, dictChan)
	printMap(dictChan)
}
