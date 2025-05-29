package extension

import (
	"github.com/Tryanks/gophermark"
)

type gfm struct {
}

// GFM is an extension that provides Github Flavored markdown functionalities.
var GFM = &gfm{}

func (e *gfm) Extend(m gophermark.Markdown) {
	Linkify.Extend(m)
	Table.Extend(m)
	Strikethrough.Extend(m)
	TaskList.Extend(m)
}
