package examples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleParser(t *testing.T) {
	for _, e := range []struct {
		raw         string
		class       Class
		source      Source
		domain      Address
		tldValue    string
		domainValue string
	}{
		{"bmskafjfoenhf.rucryptolockerdga", DGA, "cryptolocker", "bmskafjfoenhf.ru", "ru", "bmskafjfoenhf"},
		{"hollywoodreporter.comn/alegit", Legit, "n/a", "hollywoodreporter.com", "com", "hollywoodreporter"},
		{"mail.orlandoscholarguidesmarkets.sczoreodga", DGA, "zoreo", "mail.orlandoscholarguidesmarkets.sc", "sc", "orlandoscholarguidesmarkets"},
		{"unblockyoutube.co.ukn/alegit", Legit, "n/a", "unblockyoutube.co.uk", "co.uk", "unblockyoutube"},
		{"www1.functionalstarterlimitedbecoming.npzoreodga", DGA, "zoreo", "www1.functionalstarterlimitedbecoming.np", "np", "functionalstarterlimitedbecoming"},
		{"fibrecarpetpermit.mmzoreodga", DGA, "zoreo", "fibrecarpetpermit.mm", "mm", "fibrecarpetpermit"},
		{"economiesbelievedphones.ckzoreodga", DGA, "zoreo", "economiesbelievedphones.ck", "ck", "economiesbelievedphones"},
		{"fighteradaptationqualifications.erzoreodga", DGA, "zoreo", "fighteradaptationqualifications.er", "er", "fighteradaptationqualifications"},
		{"informationscreensaversstickersmanufactured.bnzoreodga", DGA, "zoreo", "informationscreensaversstickersmanufactured.bn", "bn", "informationscreensaversstickersmanufactured"},
		{"fundamentalultrastainless.jmzoreodga", DGA, "zoreo", "fundamentalultrastainless.jm", "jm", "fundamentalultrastainless"},
	} {

		ex := NewExample(e.raw)

		assert.Equal(t, ex.Class(), e.class)
		assert.Equal(t, ex.Source(), e.source)
		assert.Equal(t, ex.Address(), e.domain)
		assert.Equal(t, ex.Address().TLD(), e.tldValue)
		assert.Equal(t, ex.Address().Domain(), e.domainValue)

	}
}

func TestReadExamples(t *testing.T) {
	examples, err := ReadExamples()
	if err != nil {
		t.Fatal(err)
	}
	if len(examples) != 157927 {
		t.Fatalf("example length: %d", len(examples))
	}

}
