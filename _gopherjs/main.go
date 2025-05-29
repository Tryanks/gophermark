package main

import (
	"bytes"
	"github.com/Tryanks/gophermark"
	"github.com/gopherjs/gopherjs/js"
)

var markdownParser gophermark.Markdown

func main() {
	markdownParser = gophermark.New()
	js.Global.Set("convertMarkdownToHTML", ConvertMarkdownToHTML)
}

// ConvertMarkdownToHTML converts markdown string to HTML string
func ConvertMarkdownToHTML(markdown string) string {
	var buf bytes.Buffer
	if err := markdownParser.Convert([]byte(markdown), &buf); err != nil {
		return "Error converting markdown: " + err.Error()
	}
	return buf.String()
}
