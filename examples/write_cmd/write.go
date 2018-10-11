package main

import (
	"github.com/gocarina/gocsv"
	"github.com/heindl/dga-detection-go/examples"
	"github.com/heindl/dga-detection-go/modpath"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	exs, err := examples.ReadExamples()
	if err != nil {
		panic(err)
	}

	modPath, err := modpath.Abs()
	if err != nil {
		panic(err)
	}

	eg := errgroup.Group{}
	for _src, _sourceExamples := range exs.Sources() {
		src := _src
		sourceExamples := _sourceExamples
		eg.Go(func() error {
			fName := filepath.Join(modPath, "examples/traindata", src.Escape()+".csv")
			fss := []*examples.FeatureSet{}
			for _, e := range sourceExamples {
				fss = append(fss, examples.GenFeatureSet(e.Address()))
			}
			b, err := gocsv.MarshalBytes(fss)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(fName, b, os.ModePerm); err != nil {
				return errors.Wrapf(err, "could not write %s", fName)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		panic(err)
	}

}
