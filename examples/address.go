package examples

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"
	"strings"
)

type Address string

func (Ω Address) format() Address {
	s := strings.TrimPrefix(string(Ω), "www1.")
	s = strings.TrimPrefix(s, "www.")
	return Address(s)
}

func (Ω Address) IsICANN() bool {
	Ω = Ω.format()
	_, icann := publicsuffix.PublicSuffix(string(Ω))
	return icann
}

func (Ω Address) TLD() string {
	s := Ω.effectiveTLDPlusOne()
	return strings.Join(strings.Split(s, ".")[1:], ".")
}

func (Ω Address) effectiveTLDPlusOne() string {
	Ω = Ω.format()
	s, err := publicsuffix.EffectiveTLDPlusOne(string(Ω))
	if err != nil {
		logrus.Errorf("%s: manually splitting", err.Error())
		return string(Ω)
	}
	return s
}

func (Ω Address) Domain() string {
	s := Ω.effectiveTLDPlusOne()
	return strings.Split(s, ".")[0]
}
