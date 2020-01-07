package actions

import (
  "encoding/json"
  "fmt"
  "github.com/fengzi91/blog_app/models"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "io/ioutil"
  "reflect"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Attachment)
// DB Table: Plural (attachments)
// Resource: Plural (Attachments)
// Path: Plural (/attachments)
// View Template Folder: Plural (/templates/attachments/)

// AttachmentsResource is the resource for the Attachment model
type AttachmentsResource struct{
  buffalo.Resource
}

// List gets all Attachments. This function is mapped to the path
// GET /attachments
func (v AttachmentsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  attachments := &models.Attachments{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Attachments from the DB
  if err := q.All(attachments); err != nil {
    return err
  }

  // Add the paginator to the context so it can be used in the template.
  c.Set("pagination", q.Paginator)

  return c.Render(200, r.Auto(c, attachments))
}

// Show gets the data for one Attachment. This function is mapped to
// the path GET /attachments/{attachment_id}
func (v AttachmentsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Attachment
  attachment := &models.Attachment{}

  // To find the Attachment the parameter attachment_id is used.
  if err := tx.Eager().Find(attachment, c.Param("attachment_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.Auto(c, attachment))
}

// New renders the form for creating a new Attachment.
// This function is mapped to the path GET /attachments/new
func (v AttachmentsResource) New(c buffalo.Context) error {
  return c.Render(200, r.Auto(c, &models.Attachment{}))
}
// Create adds a Attachment to the DB. This function is mapped to the
// path POST /attachments
func (v AttachmentsResource) Create(c buffalo.Context) error {
  // Allocate an empty Attachment
  attachment := &models.Attachment{}
  user := c.Value("current_user").(*models.User)
  // Bind attachment to the html form elements
  if err := c.Bind(attachment); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }
  attachment.UserID = user.ID
  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(attachment)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, attachment))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "attachment.created.success"))
  // and redirect to the attachments index page
  return c.Render(201, r.Auto(c, attachment))
}

// Edit renders a edit form for a Attachment. This function is
// mapped to the path GET /attachments/{attachment_id}/edit
func (v AttachmentsResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Attachment
  attachment := &models.Attachment{}

  if err := tx.Find(attachment, c.Param("attachment_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.Auto(c, attachment))
}
// Update changes a Attachment in the DB. This function is mapped to
// the path PUT /attachments/{attachment_id}
func (v AttachmentsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Attachment
  attachment := &models.Attachment{}

  if err := tx.Find(attachment, c.Param("attachment_id")); err != nil {
    return c.Error(404, err)
  }

  // Bind Attachment to the html form elements
  if err := c.Bind(attachment); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(attachment)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the edit.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, attachment))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "attachment.updated.success"))
  // and redirect to the attachments index page
  return c.Render(200, r.Auto(c, attachment))
}

// Destroy deletes a Attachment from the DB. This function is mapped
// to the path DELETE /attachments/{attachment_id}
func (v AttachmentsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Attachment
  attachment := &models.Attachment{}

  // To find the Attachment the parameter attachment_id is used.
  if err := tx.Find(attachment, c.Param("attachment_id")); err != nil {
    return c.Error(404, err)
  }

  if err := tx.Destroy(attachment); err != nil {
    return err
  }

  // If there are no errors set a flash message
  c.Flash().Add("success", T.Translate(c, "attachment.destroyed.success"))
  // Redirect to the attachments index page
  return c.Render(200, r.Auto(c, attachment))
}

func AttachmentsAdd(c buffalo.Context) error {
  // attachment := &models.Attachment{}
  // tx, ok := c.Value("tx").(*pop.Connection)
  /*
  if !ok {
    return fmt.Errorf("no transaction found")
  }
  */

  fmt.Println(c.Request().Body)
  body, err := ioutil.ReadAll(c.Request().Body)
  if err != nil {
    fmt.Printf("read body err, %v\n", err)
    return nil
  }
  println("json:", string(body))
  var mapResult map[string]interface{}
  errs := json.Unmarshal([]byte(string(body)), &mapResult)
  if errs != nil {
    fmt.Println("JsonToMapDemo err: ", errs)
  }
  // aa := mapResult["upload"]
  // aa := mapResult["upload"]
  // json.Unmarshal([]byte(string(aa)), &mapResult)
  fmt.Println(reflect.TypeOf(mapResult["upload"]))
  return c.Render(200, r.JSON(c.Params()))
}

