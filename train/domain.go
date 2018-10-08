package train

import (
	"golang.org/x/net/publicsuffix"
	"strings"
)


type Domain string

func (Ω Domain) IsICANN() bool {
	_, icann := publicsuffix.PublicSuffix(string(Ω))
	return icann
}

func (Ω Domain) TLD() string {
	s, _ := publicsuffix.PublicSuffix(string(Ω))
	return s
}

func (Ω Domain) Domain() string {
 	s, _ := publicsuffix.EffectiveTLDPlusOne(string(Ω))
 	return strings.Split(s, ".")[0]
}
