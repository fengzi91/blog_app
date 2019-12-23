package actions

import (
	"github.com/fengzi91/blog_app/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

func CategoriesCreateIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	categories := &models.Categories{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all Posts from the DB
	if err := q.All(categories); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("categories", categories)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("categories/index"))
}

func CategoriesShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	category := &models.Category{}

	if err := tx.Eager("Posts").Where("slug = (?)", c.Param("slug")).Last(category); err != nil {
		return c.Error(404, err)
	}
	c.Set("category", category)

	c.Set("posts", category.Posts)
	return c.Render(200, r.HTML("categories/show"))
}

//Inserted
func CategoriesCreateGet(c buffalo.Context) error {
	c.Set("category", &models.Category{})
	return c.Render(200, r.HTML("categories/create"))
}


func CategoriesCreatePost(c buffalo.Context) error {
	// Allocate an empty Post
	category := &models.Category{}
	// Bind post to the html form elements
	if err := c.Bind(category); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(category)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("category", category)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("posts/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "分类创建成功.")
	// and redirect to the index page
	return c.Redirect(302, "/categories/index")
}
func CategoriesEditGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	category := &models.Category{}
	if err := tx.Find(category, c.Param("cid")); err != nil {
		return c.Error(404, err)
	}
	c.Set("category", category)
	return c.Render(200, r.HTML("categories/edit.html"))
}

func CategoriesEditPost(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	category := &models.Category{}
	if err := tx.Find(category, c.Param("cid")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(category); err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(category)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("category", category)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("categories/edit.html"))
	}
	c.Flash().Add("success", "Category was updated successfully.")
	return c.Redirect(302, "/categories/%s", category.Slug)
}
// Delete Category
func CategoriesDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	category := &models.Category{}
	if err := tx.Find(category, c.Param("cid")); err != nil {
		return c.Error(404, err)
	}
	if err := tx.Destroy(category); err != nil {
		return errors.WithStack(err)
	}
	c.Flash().Add("success", "Category was successfully deleted.")
	return c.Redirect(302, "/categories/index")
}
/*
// PostsDetail displays a single post.
func PostsDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	post := &models.Post{}
	if err := tx.Find(post, c.Param("pid")); err != nil {
		return c.Error(404, err)
	}
	author := &models.User{}
	if err := tx.Find(author, post.AuthorID); err != nil {
		return c.Error(404, err)
	}
	c.Set("post", post)
	c.Set("author", author)
	comment := &models.Comment{}
	c.Set("comment", comment)
	comments := models.Comments{}
	if err := tx.BelongsTo(post).All(&comments); err != nil {
		return errors.WithStack(err)
	}
	for i := 0; i < len(comments); i++ {
		u := models.User{}
		if err := tx.Find(&u, comments[i].AuthorID); err != nil {
			return c.Error(404, err)
		}
		comments[i].Author = u
	}
	c.Set("comments", comments)
	return c.Render(200, r.HTML("posts/detail"))
}
 */
// 设置分类
func SetCategories(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		categories := &models.Categories{}
		tx := c.Value("tx").(*pop.Connection)
		err := tx.Order("`orders` asc").Order("updated_at desc").All(categories)
		if err != nil {
			return errors.WithStack(err)
		}
		c.Set("_categories", categories)
		return next(c)
	}
}
