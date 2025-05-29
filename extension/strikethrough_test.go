package extension

import (
	"testing"

	"github.com/Tryanks/gophermark"
	"github.com/Tryanks/gophermark/renderer/html"
	"github.com/Tryanks/gophermark/testutil"
)

func TestStrikethrough(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			Strikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/strikethrough.txt", t, testutil.ParseCliCaseArg()...)
}
