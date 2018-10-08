package parse

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

const minWordLength = 3

func MatchEnglishWords(line string, scanLength ...int) []string {
	scanLen := len(line)
	if len(scanLength) > 0 {
		scanLen = scanLength[0]
	}
	if len(line) < minWordLength || scanLen < minWordLength {
		return nil
	}
	for i := 0; i <= len(line)-scanLen; i++ {
		word := line[i : i+scanLen]
		if !EnglishDictionary.IsWord(word) {
			continue
		}
		y := []string{word}
		left := line[:i]
		y = append(y, MatchEnglishWords(left, len(left))...)
		right := line[i:]
		return append(y, MatchEnglishWords(right, len(right))...)
	}
	return MatchEnglishWords(line, scanLen - 1)
}

type dictionary []string

func (Ω dictionary) IsWord(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	y := sort.Search(len(Ω) - 1, func(i int) bool {
		return Ω[i] >= s
	})
	return Ω[y] == s
}

var EnglishDictionary dictionary

func init() {
	file, err := os.OpenFile(
		"../aoo-mozilla-en-dict/en_US (Kevin Atkinson)/wordlist_kevin_en_US_20180416_123131w.txt",
		os.O_RDONLY,
		600,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if strings.Contains(word, "'") {
			continue
		}
		if len(word) < minWordLength {
			continue
		}
		word = strings.ToLower(word)
		EnglishDictionary = append(EnglishDictionary, word)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	sort.Strings(EnglishDictionary)
}