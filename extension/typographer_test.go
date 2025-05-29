package extension

import (
	"testing"

	"github.com/Tryanks/gophermark"
	"github.com/Tryanks/gophermark/renderer/html"
	"github.com/Tryanks/gophermark/testutil"
)

func TestTypographer(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			Typographer,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/typographer.txt", t, testutil.ParseCliCaseArg()...)
}
