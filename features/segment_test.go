package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary(t *testing.T) {

	for word, exists := range map[string]bool{
		"canada": true,
	}{
		y := MatchEnglishWords(word)
		assert.Equal(t, exists, y, "%s", word)
	}

}

func TestSegment(t *testing.T) {

	for domain, length := range map[string]int{
		"aircanada": 6,
		"orlandoscholarguidesmarkets": 7,
		"vgtpjgusskvsdiwoxkhxwxln": 0,
	}{
		y :=  MatchEnglishWords(domain)
		assert.Equal(t, length, len(y), "%s", domain)
	}

}
