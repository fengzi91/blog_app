package actions

import (
	"html/template"
	"time"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",
		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,
		// Add template helpers here:
		Helpers: render.Helpers{
			"csrf": func() template.HTML {
				return template.HTML("<input name=\"authenticity_token\" value=\"<%= authenticity_token %>\" type=\"hidden\">")
			},
			"formatDate": func(createdAt time.Time) string {
				str := createdAt.Format("2006-01-02")
				return str
			},
			"html": func(s string) template.HTML {
				return template.HTML(s)
			},
		},
	})
}
