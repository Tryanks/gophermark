package extension

import (
	"testing"

	"github.com/Tryanks/gophermark"
	"github.com/Tryanks/gophermark/ast"
	east "github.com/Tryanks/gophermark/extension/ast"
	"github.com/Tryanks/gophermark/parser"
	"github.com/Tryanks/gophermark/renderer/html"
	"github.com/Tryanks/gophermark/testutil"
	"github.com/Tryanks/gophermark/text"
	"github.com/Tryanks/gophermark/util"
)

func TestTable(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
			html.WithXHTML(),
		),
		gophermark.WithExtensions(
			Table,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/table.txt", t, testutil.ParseCliCaseArg()...)
}

func TestTableWithAlignDefault(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignDefault),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          1,
			Description: "Cell with TableCellAlignDefault and XHTML should be rendered as an align attribute",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th align="center">abc</th>
<th align="right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td align="center">bar</td>
<td align="right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)

	markdown = gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignDefault),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          2,
			Description: "Cell with TableCellAlignDefault and HTML5 should be rendered as a style attribute",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th style="text-align:center">abc</th>
<th style="text-align:right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align:center">bar</td>
<td style="text-align:right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)
}

func TestTableWithAlignAttribute(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignAttribute),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          1,
			Description: "Cell with TableCellAlignAttribute and XHTML should be rendered as an align attribute",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th align="center">abc</th>
<th align="right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td align="center">bar</td>
<td align="right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)

	markdown = gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignAttribute),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          2,
			Description: "Cell with TableCellAlignAttribute and HTML5 should be rendered as an align attribute",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th align="center">abc</th>
<th align="right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td align="center">bar</td>
<td align="right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)
}

type tableStyleTransformer struct {
}

func (a *tableStyleTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	cell := node.FirstChild().FirstChild().FirstChild().(*east.TableCell)
	cell.SetAttributeString("style", []byte("font-size:1em"))
}

func TestTableWithAlignStyle(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignStyle),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          1,
			Description: "Cell with TableCellAlignStyle and XHTML should be rendered as a style attribute",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th style="text-align:center">abc</th>
<th style="text-align:right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align:center">bar</td>
<td style="text-align:right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)

	markdown = gophermark.New(
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignStyle),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          2,
			Description: "Cell with TableCellAlignStyle and HTML5 should be rendered as a style attribute",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th style="text-align:center">abc</th>
<th style="text-align:right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align:center">bar</td>
<td style="text-align:right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)

	markdown = gophermark.New(
		gophermark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&tableStyleTransformer{}, 0),
			),
		),
		gophermark.WithRendererOptions(
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignStyle),
			),
		),
	)

	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          3,
			Description: "Styled cell should not be broken the style by the alignments",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th style="font-size:1em;text-align:center">abc</th>
<th style="text-align:right">defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align:center">bar</td>
<td style="text-align:right">baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)
}

func TestTableWithAlignNone(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(
				WithTableCellAlignMethod(TableCellAlignNone),
			),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          1,
			Description: "Cell with TableCellAlignStyle and XHTML should not be rendered",
			Markdown: `
| abc | defghi |
:-: | -----------:
bar | baz
`,
			Expected: `<table>
<thead>
<tr>
<th>abc</th>
<th>defghi</th>
</tr>
</thead>
<tbody>
<tr>
<td>bar</td>
<td>baz</td>
</tr>
</tbody>
</table>`,
		},
		t,
	)
}

func TestTableFuzzedPanics(t *testing.T) {
	markdown := gophermark.New(
		gophermark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
		gophermark.WithExtensions(
			NewTable(),
		),
	)
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          1,
			Description: "This should not panic",
			Markdown:    "* 0\n-|\n\t0",
			Expected: `<ul>
<li>
<table>
<thead>
<tr>
<th>0</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
</tr>
</tbody>
</table>
</li>
</ul>`,
		},
		t,
	)
}
