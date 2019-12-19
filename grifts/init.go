package grifts

import (
	"github.com/fengzi91/blog_app/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
