package parse

import "strings"

type Source string

func (Ω Source) Escape() string {
	return strings.Replace(string(Ω), "/", "", -1)
}
