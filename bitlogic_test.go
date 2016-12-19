package bitlogic

import (
	"github.com/DATA-DOG/godog"
	"github.com/bitlogic/bitlogic.io.test/home"
)

func FeatureContext(s *godog.Suite) {
	home.FeatureContext(s)
}
