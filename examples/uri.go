package parse

import (
	"github.com/heindl/dga/features"
	"golang.org/x/net/publicsuffix"
	"strings"
)


type URI string

func (Ω URI) IsICANN() bool {
	_, icann := publicsuffix.PublicSuffix(string(Ω))
	return icann
}

func (Ω URI) TLD() string {
	s, _ := publicsuffix.PublicSuffix(string(Ω))
	return s
}

func (Ω URI) Domain() string {
 	s, _ := publicsuffix.EffectiveTLDPlusOne(string(Ω))
 	return strings.Split(s, ".")[0]
}

func (Ω URI) Words() (words []string, wordCharacters, unknownCharacters float64) {
	words = features.EnglishWords(Ω.Domain())
	for _, w := range words {
		wordCharacters += float64(len(w))
	}
	unknownCharacters = float64(len(Ω.Domain())) - wordCharacters
	return
}
