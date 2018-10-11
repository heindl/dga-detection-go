package examples

import (
	"bufio"
	"fmt"
	"github.com/heindl/dga-detection-go/modpath"
	"os"
	"path/filepath"
)

func ReadExamples() (Examples, error) {

	modPath, err := modpath.Abs()
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filepath.Join(modPath, "./dataset/dga-dataset.txt"), os.O_RDONLY, 600)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	y := Examples{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		y = append(y, NewExample(scanner.Text()))
	}
	return y, nil
}

type Examples []Example

func (Ω Examples) PercentDGA() float64 {
	dgas := float64(0)
	for _, e := range Ω {
		if e.Class() == DGA {
			dgas += 1
		}
	}
	return (dgas * 100) / float64(len(Ω))
}

func (Ω Examples) Classes() map[Class]Examples {
	y := map[Class]Examples{
		Unknown: Examples{},
		DGA:     Examples{},
		Legit:   Examples{},
	}
	for _, e := range Ω {
		c := e.Class()
		y[c] = append(y[c], e)
	}
	return y
}

func (Ω Examples) Sources() map[Source]Examples {
	y := map[Source]Examples{}
	for _, e := range Ω {
		c := e.Source()
		if c == "" {
			fmt.Println("sourceless", e)
			continue
		}
		if _, ok := y[c]; !ok {
			y[c] = Examples{}
		}
		y[c] = append(y[c], e)
	}
	return y
}

func (Ω Examples) TLDs() map[string]Examples {
	y := map[string]Examples{}
	for _, e := range Ω {
		tld := e.Address().TLD()
		if _, ok := y[tld]; !ok {
			y[tld] = Examples{}
		}
		y[tld] = append(y[tld], e)
	}
	return y
}
