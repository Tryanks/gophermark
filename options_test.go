package gophermark_test

import (
	"testing"

	. "github.com/Tryanks/gophermark"
	"github.com/Tryanks/gophermark/parser"
	"github.com/Tryanks/gophermark/testutil"
)

func TestAttributeAndAutoHeadingID(t *testing.T) {
	markdown := New(
		WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/options.txt", t, testutil.ParseCliCaseArg()...)
}
