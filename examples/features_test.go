package examples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestDictionary(t *testing.T) {
//
//	for word, exists := range map[string]bool{
//		"canada": true,
//	}{
//		y := englishDictionary.IsWord(word)
//		assert.Equal(t, exists, y, "%s", word)
//	}
//
//}

func TestSegment(t *testing.T) {

	for domain, words := range map[string][]string{
		"aircanada":                      {"canada"},
		"orlandoscholarguidesmarkets":    {"orlando", "scholar", "markets", "guides"},
		"vgtpjgusskvsdiwoxkhxwxln":       nil,
		"estimatesnetworksproteinsburns": {"estimates", "networks", "proteins", "burns"},
	} {
		y := EnglishWords(domain)
		assert.Equal(t, words, y, "%s", domain)
	}

}
