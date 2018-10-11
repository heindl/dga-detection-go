package main

import (
	"fmt"
	"github.com/heindl/dga-detection-go/examples"
)

func main() {
	exs, err := examples.ReadExamples()
	if err != nil {
		panic(err)
	}
	correct := 0
	falseDGA := 0
	falseLegit := 0
	for _, e := range exs {
		class := Classify(e.Address())
		if e.Class() != class {
			if e.Class() == examples.DGA {
				//fmt.Println(e.Address().Domain(), ",", e.Address().TLD(), ",")
				falseLegit += 1
			} else {
				falseDGA += 1
			}
		} else {
			correct += 1
		}
	}

	fmt.Println("false dga", falseDGA)
	fmt.Println("false legit", falseLegit)
	fmt.Printf("accuracy : %.6f \n", float64(falseDGA+falseLegit)/float64(correct+falseDGA+falseLegit))
}
