package actions

import (
	"github.com/fengzi91/blog_app/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"strings"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	posts := &models.Posts{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all Posts from the DB
	if err := q.Eager().Scope(OrderByCreatedAt()).All(posts); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("posts", posts)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	c.Set("title", "这里是首页")
	return c.Render(200, r.HTML("index.html"))
}

func OrderByCreatedAt() pop.ScopeFunc {
	return func(q *pop.Query) *pop.Query {
		return q.Order("created_at desc")
	}
}

func SetCurrentRouter(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		path := c.Value("current_route").(buffalo.RouteInfo).PathName
		var newPath string
		newPath = snakeString((strings.Replace(path,"Path" ,"-page", 1)))
		c.Set("current_path_class", newPath)
		return next(c)
	}
}
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '-')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
