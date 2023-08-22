package cmdline

import (
	// nolint:gosec

	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdOpts(t *testing.T) {
	l := log.Default()
	fn1 := func(cwo *configWktOpt) error {
		return fmt.Errorf("should not come here")
	}
	fn2 := func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") }
	fn3 := func(cpo *configTemplateOpt) error { return fmt.Errorf("should not come here") }
	t.Run("asking help", func(t *testing.T) {
		_, err := parseMainOpts2(l, []string{"-h"}, fn1, fn2, fn3)()
		assert.ErrorContains(t, err, "flag: help requested")
	})
	t.Run("missing main options", func(t *testing.T) {
		_, err := parseMainOpts2(l, []string{}, fn1, fn2, fn3)()
		assert.ErrorContains(t, err, "pls choose \"wkt\" or \"props\" option")
	})
	t.Run("choosing wkt", func(t *testing.T) {
		_, err := parseMainOpts2(l, []string{"wkt"},
			func(cwo *configWktOpt) error {
				return fmt.Errorf("AAA")
			}, fn2, fn3)()
		assert.ErrorContains(t, err, "AAA")
	})
	t.Run("choosing props", func(t *testing.T) {
		_, err := parseMainOpts2(l, []string{"props"},
			fn1,
			func(cpo *configPropsOpt) error {
				return fmt.Errorf("AAAA")
			},
			fn3,
		)()
		assert.ErrorContains(t, err, "AAA")
	})
	t.Run("choosing template", func(t *testing.T) {
		_, err := parseMainOpts2(l, []string{"template", "-template=aa"},
			fn1,
			fn2,
			func(cpo *configTemplateOpt) error {
				return fmt.Errorf("AAAA")
			},
		)()
		assert.ErrorContains(t, err, "AAA")
	})
}
