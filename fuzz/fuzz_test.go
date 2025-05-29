package fuzz

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/Tryanks/gophermark"
	"github.com/Tryanks/gophermark/extension"
	"github.com/Tryanks/gophermark/parser"
	"github.com/Tryanks/gophermark/renderer/html"
	"github.com/Tryanks/gophermark/util"
)

func fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		markdown := gophermark.New(
			gophermark.WithParserOptions(
				parser.WithAutoHeadingID(),
				parser.WithAttribute(),
			),
			gophermark.WithRendererOptions(
				html.WithUnsafe(),
				html.WithXHTML(),
			),
			gophermark.WithExtensions(
				extension.DefinitionList,
				extension.Footnote,
				extension.GFM,
				extension.Typographer,
				extension.Linkify,
				extension.Table,
				extension.TaskList,
			),
		)
		var b bytes.Buffer
		if err := markdown.Convert(util.StringToReadOnlyBytes(orig), &b); err != nil {
			panic(err)
		}
	})
}

func FuzzDefault(f *testing.F) {
	bs, err := os.ReadFile("../_test/spec.json")
	if err != nil {
		panic(err)
	}
	var testCases []map[string]interface{}
	if err := json.Unmarshal(bs, &testCases); err != nil {
		panic(err)
	}
	for _, c := range testCases {
		f.Add(c["markdown"])
	}
	fuzz(f)
}
