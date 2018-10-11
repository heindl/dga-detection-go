package main

import "github.com/heindl/dga-detection-go/examples"

func Classify(addr examples.Address) examples.Class {

	wStats := examples.GenFeatureSet(addr)

	if wStats.DomainWordCount == 0 {
		return examples.DGA
	}

	if wStats.DomainWordMinCharacterCount >= 5 && wStats.DomainWordCount == 3 || wStats.DomainWordCount == 4 && wStats.DomainRatioMatched > 0.95 {
		return examples.DGA
	}

	return examples.Legit

}
