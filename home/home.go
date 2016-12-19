package home

import (
	"github.com/DATA-DOG/godog"
	"github.com/bitlogic/bitlogic.io.test/browser"
)

var b *browser.Client

func imInterestedInBitlogicio() error {
	return nil
}

func iNavigateToHttpbitlogicio() (err error) {
	b, err = browser.NewClient("/")
	return
}

func theLandingPageIsLoaded() error {
	return b.CheckPage()
}

//FeatureContext exports steps related to home
func FeatureContext(s *godog.Suite) {
	s.Step(`^i\'m interested in bitlogic\.io$`, imInterestedInBitlogicio)
	s.Step(`^i navigate to http:\/\/bitlogic\.io$`, iNavigateToHttpbitlogicio)
	s.Step(`^the landing page is loaded$`, theLandingPageIsLoaded)
}
