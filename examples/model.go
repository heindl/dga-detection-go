package examples

import "strings"

type Class int

const (
	Unknown Class = 0
	DGA     Class = 1
	Legit   Class = 2
)

func (Ω Class) String() string {
	switch Ω {
	case Unknown:
		return "unknown"
	case DGA:
		return "dga"
	case Legit:
		return "legit"
	}
	return ""
}

type Source string

func (Ω Source) Escape() string {
	return strings.Replace(string(Ω), "/", "", -1)
}
