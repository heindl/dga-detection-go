package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleParser(t *testing.T) {
 for _, e := range []struct{
	 raw         string
	 class       Class
	 source      Source
	 domain      URI
	 domainValue string
	 tldValue    string
 }{
 	{"bmskafjfoenhf.rucryptolockerdga", DGA, "cryptolocker", "bmskafjfoenhf.ru", "ru", "bmskafjfoenhf"},
 	{"hollywoodreporter.comn/alegit", Legit, "n/a","hollywoodreporter.com", "com", "hollywoodreporter"},
 	{"mail.orlandoscholarguidesmarkets.sczoreodga", DGA, "zoreo", "mail.orlandoscholarguidesmarkets.sc", "sc", "orlandoscholarguidesmarkets"},
 	{"unblockyoutube.co.ukn/alegit", Legit, "n/a", "unblockyoutube.co.uk", "co.uk", "unblockyoutube"},
 } {

 	ex := NewExample(e.raw)

 	assert.Equal(t, ex.Class(), e.class)
	 assert.Equal(t, ex.Source(), e.source)
	 assert.Equal(t, ex.Domain(), e.domain)
 	assert.Equal(t, ex.Domain().TLD(), e.domainValue)
	 assert.Equal(t, ex.Domain().Domain(), e.tldValue)

 }
}

//func TestReadExamples(t *testing.T) {
//  examples, err := ReadExamples("./dga-dataset.txt")
//  if err != nil {
//	  t.Fatal(err)
//  }
//	if len(examples) != 157927 {
//		t.Fatalf("example length: %d", len(examples))
//	}
//
//  classes := examples.Classes()
//  if len(classes[Unknown]) != 1 {
//	  t.Fatalf("expected one unknown item, found %d", len(examples))
//  }
//
//  sources := examples.Sources()
//  for k, _ := range sources {
//  	fmt.Println("key", k)
//  }
//  fmt.Println("sources", len(sources))
//  fmt.Println("average", float64(len(examples)) / float64(len(classes)))
//
//  _ = examples.TLDs()
//  //	fmt.Println("suffix", i, v.PercentDGA())
//  //}
//
//}