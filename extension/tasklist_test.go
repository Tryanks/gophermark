package extension

import (
	"testing"

	"github.com/Tryanks/gophermark"
	"github.com/Tryanks/gophermark/renderer/html"
	"github.com/Tryanks/gophermark/testutil"
)

func TestTaskList(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/tasklist.txt", t, testutil.ParseCliCaseArg()...)
}
