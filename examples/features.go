package examples

import (
	"bufio"
	"github.com/heindl/dga/modpath"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FeatureSet struct {
	//Label examples.Class `csv:"label" validate:"min=0,max=1"`
	TLD                         string  `csv:"tld" validate:"required"`
	DomainCharacterCount        int     `csv:"" validate:"min=1"`
	DomainWordCount             int     `csv:""`
	DomainRatioMatched          float64 `csv:""`
	DomainWordMinCharacterCount int     `csv:""`
	DomainWordMaxCharacterCount int     `csv:""`
	DomainWordAvgCharacterCount float64 `csv:""`
}

func GenFeatureSet(addr Address) *FeatureSet {

	wStats := Words(addr.Domain())

	return &FeatureSet{
		TLD:                         addr.TLD(),
		DomainCharacterCount:        int(wStats.TotalCharacterCount),
		DomainWordCount:             len(wStats.WordLengths),
		DomainRatioMatched:          floats.Sum(wStats.WordLengths) / wStats.TotalCharacterCount,
		DomainWordAvgCharacterCount: stat.Mean(append([]float64{0}, wStats.WordLengths...), nil),
		DomainWordMinCharacterCount: int(floats.Min(append([]float64{0}, wStats.WordLengths...))),
		DomainWordMaxCharacterCount: int(floats.Max(append([]float64{0}, wStats.WordLengths...))),
	}
}

const minWordLength = 4

type DomainWordStats struct {
	TotalCharacterCount float64
	Words               []string
	WordLengths         []float64
}

func Words(domain string) *DomainWordStats {

	sts := &DomainWordStats{
		TotalCharacterCount: float64(len(domain)),
		Words:               EnglishWords(domain),
		WordLengths:         []float64{},
	}

	if len(sts.Words) == 0 {
		return sts
	}

	for _, w := range sts.Words {
		sts.WordLengths = append(sts.WordLengths, float64(len(w)))
	}

	return sts
}

func EnglishWords(line string, scanLength ...int) []string {
	scanLen := len(line)
	if len(scanLength) > 0 {
		scanLen = scanLength[0]
	}
	if len(line) < minWordLength || scanLen < minWordLength {
		return nil
	}
	for i := 0; i <= len(line)-scanLen; i++ {
		word := line[i : i+scanLen]
		if !englishDictionary.IsWord(word) {
			continue
		}
		left := line[:i]
		right := line[i+scanLen:]
		y := append([]string{word}, EnglishWords(left, len(left))...)
		return append(y, EnglishWords(right, len(right))...)
	}
	return EnglishWords(line, scanLen-1)
}

type dictionary []string

func (Ω dictionary) IsWord(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	y := sort.Search(len(Ω)-1, func(i int) bool {
		return Ω[i] >= s
	})
	return Ω[y] == s
}

var englishDictionary dictionary

func init() {
	modPath, err := modpath.Abs()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(
		filepath.Join(modPath, "aoo-mozilla-en-dict/en_US (Kevin Atkinson)/wordlist_kevin_en_US_20180416_123131w.txt"),
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
		englishDictionary = append(englishDictionary, word)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	sort.Strings(englishDictionary)
}
