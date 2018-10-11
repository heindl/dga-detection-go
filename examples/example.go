package examples

import (
	"golang.org/x/net/publicsuffix"
	"strings"
)

type Example interface {
	Address() Address
	Source() Source
	Class() Class
}

func NewExample(s string) Example {
	if strings.ContainsRune(s, 1) {
		return splitExample(strings.Split(s, string(rune(1))))
	}
	if strings.Contains(s, " ") {
		return splitExample(strings.Fields(s))
	}
	return unsplitExample(s)
}

type splitExample []string

func (Ω splitExample) Address() Address {
	if len(Ω) == 0 {
		return Address("")
	}
	return Address(Ω[0])
}

func (Ω splitExample) Source() Source {
	if len(Ω) < 2 {
		return Source("")
	}
	return Source(strings.ToLower(Ω[1]))
}

func (Ω splitExample) Class() Class {
	if len(Ω) < 3 {
		return Unknown
	}
	switch strings.ToLower(Ω[2]) {
	case "dga", "dg", "dgf", "dgb", "da", "dgs", "dha":
		return DGA
	case "legit", "lgit", "lenit", "legip", "lefit":
		return Legit
	default:
		return Unknown
	}
}

type unsplitExample string

func (Ω unsplitExample) Source() Source {
	ex := Ω.trimDomain()
	ex = ex.trimClass()
	return Source(strings.TrimSpace(string(ex)))
}

func (Ω unsplitExample) trimClass() unsplitExample {
	e := strings.TrimSpace(strings.ToLower(string(Ω)))
	if strings.HasSuffix(e, "dga") {
		return unsplitExample(strings.TrimSuffix(e, "dga"))
	}
	if strings.HasSuffix(e, "legit") {
		return unsplitExample(strings.TrimSuffix(e, "legit"))
	}
	return unsplitExample(e)
}

func (Ω unsplitExample) Class() Class {
	raw := strings.TrimSpace(strings.ToLower(string(Ω)))
	trimmed := strings.TrimSuffix(raw, "dga")
	if len(trimmed) < len(raw) {
		return DGA
	}
	trimmed = strings.TrimSuffix(raw, "legit")
	if len(trimmed) < len(raw) {
		return Legit
	}
	return Unknown
}

func (Ω unsplitExample) Address() Address {
	raw := string(Ω.trimClass())
	for {
		if len(raw) == 0 {
			break
		}
		_, icann := publicsuffix.PublicSuffix(string(raw))
		if !icann {
			raw = strings.TrimSpace(raw[:len(raw)-1])
			continue
		}
		return Address(raw)
	}
	return Address("")
}

func (Ω unsplitExample) trimDomain() unsplitExample {
	ex := string(Ω)
	do := string(Ω.Address())
	return unsplitExample(strings.TrimPrefix(strings.TrimSpace(ex), do))
}
