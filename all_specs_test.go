package yed_test

import (
	"testing"

	"github.com/orfjackal/gospec/src/gospec"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(YedSpec)
	gospec.MainGoTest(r, t)
}
