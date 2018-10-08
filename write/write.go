package main

import (
	"fmt"
	"github.com/heindl/dga/train"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var rootCmd = &cobra.Command{
	Use:   "dga",
	Run: func(cmd *cobra.Command, args []string) {

		examples, err := train.ReadExamples("../data/original/dga-dataset.txt")
		if err != nil {
			panic(err)
		}

		for src, sourceExamples := range examples.Sources() {
			for class, classExamples := range sourceExamples.Classes() {
				if class == train.Unknown {
					continue
				}
				fName := fmt.Sprintf("../data/organized/%s-%s.csv", src.Escape(), class.String())
				if err := ioutil.WriteFile(fName, classExamples.CSV(), 700); err != nil {
					panic(errors.Wrapf(err, "could not write %s", fName))
				}
			}

		}
	},
}
func main() {

	rootCmd.Execute()
	//f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 600)
	//if err != nil {
	//	return errors.WithStack(err)
	//}


}

