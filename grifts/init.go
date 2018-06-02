package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/website2/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
