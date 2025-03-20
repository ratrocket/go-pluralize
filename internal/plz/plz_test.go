package plz_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"md0.org/pluralize/internal/plz"
	"md0.org/reporoot"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"pluralize": plz.Main,
	}))
}

func TestScript(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: filepath.Join(reporoot.MustRoot(), "testdata/script"),
	})
}
