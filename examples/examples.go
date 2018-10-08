package parse

import (
	"bufio"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"sort"
	"strings"
	"sync"
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
	locker := sync.Mutex{}
	eg := errgroup.Group{}
	for _, _e := range Ω {
		e := _e
		eg.Go(func() error {
			line := ExampleToCSV(e)
			locker.Lock()
			defer locker.Unlock()
			y = append(y, line)
			return nil
		})
	}
	_ = eg.Wait()
	sort.Strings(y)
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
		tld := e.Domain().TLD()
		if _, ok := y[tld]; !ok {
			y[tld] = Examples{}
		}
		y[tld] = append(y[tld], e)
	}
	return y
}


