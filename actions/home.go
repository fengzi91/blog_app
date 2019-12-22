package actions

import (
	"github.com/fengzi91/blog_app/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
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
	if err := q.Eager().Order("created_at asc").All(posts); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("posts", posts)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	c.Set("title", "这里是首页")
	return c.Render(200, r.HTML("index.html"))
}
