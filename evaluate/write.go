package main
//
//import (
//	"fmt"
//	"github.com/heindl/dga/examples"
//	"github.com/heindl/dga/parse"
//	"github.com/pkg/errors"
//	"github.com/spf13/cobra"
//	"io/ioutil"
//	"os"
//	"path"
//	"sort"
//)
//
//var rootCmd = &cobra.Command{
//	Use:   "dga",
//	Run: func(cmd *cobra.Command, args []string) {
//		return
//	},
//}
//
//func expand(examples examples.Examples) error {
//
//	for src, sourceExamples := range examples.Sources() {
//		for class, classExamples := range sourceExamples.Classes() {
//			if len(classExamples) == 0 {
//				continue
//			}
//			if class == examples.Unknown {
//				continue
//			}
//			if class == examples.Legit {
//				d := []string{}
//				for k, v := range classExamples.TLDs() {
//					d = append(d, fmt.Sprintf("%d:%s", len(v), k))
//				}
//				sort.Strings(d)
//				fmt.Println(d)
//			}
//			fName :=
//				path.Join(
//					os.Getenv("HOME"),
//					fmt.Sprintf(
//						"Desktop/dga-train/%s-%s.csv",
//						src.Escape(),
//						class.String(),
//					),
//				)
//			data := classExamples.CSV()
//			if err := ioutil.WriteFile(fName, data, os.ModePerm); err != nil {
//				return errors.Wrapf(err, "could not write %s", fName)
//			}
//		}
//	}
//	return nil
//}
//
//func main() {
//
//	rootCmd.Execute()
//	//f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 600)
//	//if err != nil {
//	//	return errors.WithStack(err)
//	//}
//
//
//}
//
