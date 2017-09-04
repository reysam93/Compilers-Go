package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type Line struct {
	text   string
	number int
}

type Dictionary struct {
	m    map[string][]Line
	name string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func splitWords(r rune) bool {
	return !unicode.IsLetter(r)
}

func addWordsToMap(m map[string][]Line, l Line) {
	words := strings.FieldsFunc(strings.ToLower(l.text), splitWords)
	if words == nil {
		return
	}
	// deleting repeated words in a line
	wordsMap := make(map[string]struct{})
	for _, word := range words {
		wordsMap[word] = struct{}{}
	}

	for word := range wordsMap {
		m[word] = append(m[word], l)
	}
}

func printMap(dict Dictionary) {
	var keys []string

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

func scanText(file *os.File, m map[string][]Line) {
	lineScanner := bufio.NewScanner(file)
	for i := 1; lineScanner.Scan(); i++ {
		line := Line{}
		line.text = lineScanner.Text()
		line.number = i
		addWordsToMap(m, line)
	}
	checkErr(lineScanner.Err())
}

func main() {
	dict := Dictionary{m: make(map[string][]Line)}

	switch len(os.Args) {
	case 1:
		dict.name = "quijote.txt"
	case 2:
		dict.name = os.Args[1]
	default:
		log.Fatal("usage: quijote [file_name]\nNote: if no file name is provided, the default name will be quijote.txt")
	}

	file, err := os.Open(dict.name)
	defer file.Close()
	checkErr(err)
	scanText(file, dict.m)
	printMap(dict)
}
