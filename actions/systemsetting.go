package actions

import "github.com/gobuffalo/buffalo"

// SystemsettingCreate default implementation.
func SystemsettingCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("systemsetting/create.html"))
}

// SystemsettingEdit default implementation.
func SystemsettingEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("systemsetting/edit.html"))
}

