package train

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadExamples(filePath string) (Examples, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 600)
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
	return  (dgas * 100) / float64(len(Ω))
}

func (Ω Examples) CSV() []byte {
	y := []string{}
	for _, e := range Ω {
		domain := e.Domain()
		line := fmt.Sprintf("%s,%s", domain.Domain(), domain.TLD())
		y = append(y, line)
	}
	return []byte(strings.Join(y, "\n"))
}

func (Ω Examples) Classes() map[Class]Examples {
	y := map[Class]Examples{
		Unknown: Examples{},
		DGA: Examples{},
		Legit: Examples{},
	}
	for _, e := range Ω {
		c := e.Class()
		if _, ok := y[c]; !ok {
			fmt.Println("Missing class", c)
			continue
		}
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
		c := e.Domain().TLD()
		i := strings.Index(c, ".")
		if i == -1 {
			fmt.Println(e, c)
			continue
		}
		v := c[i:]
		if _, ok := y[v]; !ok {
			y[v] = Examples{}
		}
		y[v] = append(y[v], e)
	}
	return y
}


