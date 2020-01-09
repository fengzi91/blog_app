package actions

import (
  "encoding/json"
  "fmt"
  "github.com/fengzi91/blog_app/models"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/gofrs/uuid"
  "github.com/pkg/errors"
  "io/ioutil"
  "github.com/gomodule/redigo/redis"
  "strings"
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
  // fmt.Println(c.Request().Body)
  body, err := ioutil.ReadAll(c.Request().Body)
  if err != nil {
    fmt.Printf("read body err, %v\n", err)
    return nil
  }
  var a Res
  if err = json.Unmarshal(body, &a); err != nil {
    fmt.Printf("Unmarshal err, %v\n", err)
    return nil
  }
  fmt.Printf("打印 Header, %v\n", c.Request().Header)
  fmt.Printf("Http Body 数据\n, %v\n", c.Request().Body)
  token := a.Upload.Meta.Token
  uid := a.Upload.Meta.UserID

  hookName := strings.Join(c.Request().Header["Hook-Name"], "")
  // 将文件加入 attachments
  if (hookName != "post-receive") {
    bools := ValidateToken(uid, token)
    if !bools {
      // 暂时跳过权限检查
      return c.Render(200, r.String("没有权限上传文件"))
    }
    return c.Render(200, r.String("上传文件"))
  }
  tx, ok := c.Value("tx").(*pop.Connection)

  if !ok {
    return fmt.Errorf("no transaction found")
  }
  attachment := models.Attachment{UserID: uid, Url: "https://zys-blog.cdn.bcebos.com/" + a.Upload.Storage.Key, Size: a.Upload.Size}
  verrs, err := tx.ValidateAndCreate(attachment)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Error(500, errors.New("文件保存附件失败!"))
  }
  return c.Render(200, r.JSON(attachment))
}

type Res struct {
  Upload      UploadModel  `json:"Upload"`
  HTTPRequest HttpModel `json:"HTTPRequest"`
}
type HttpModel struct {
  Method  string `json:"Method"`
}
type StorageModel struct {
  Bucket string `json:"Bucket"`
  Key string `json:"Key"`
}
type MetaDataModel struct {
  Token uuid.UUID `json:"token"`
  UserID uuid.UUID `json:"uid"`
}
type UploadModel struct {
  ID string `json:"ID"`
  Storage StorageModel `json:"Storage"`
  Meta  MetaDataModel `json:"MetaData"`
  Size int64 `json:"Size"`
}
func GenerateToken(uid uuid.UUID) (token uuid.UUID, err error) {
  token, _ = uuid.NewV4()
  conn, err := redis.Dial("tcp", ":6379")

  if err != nil {
    return token, errors.WithStack(err)
  }

  if _, err := conn.Do("SET", uid, token); err != nil {
    return token, errors.WithStack(err)
  }

  defer conn.Close()
  return token, nil
}

func ValidateToken(uid uuid.UUID, token uuid.UUID) bool {
  conn, err := redis.Dial("tcp", ":6379")
  defer conn.Close()
  if err != nil {
    return false
  }
  a, err := redis.String(conn.Do("GET", uid))
  if err != nil {
    return false
  }
  if uuid.FromStringOrNil(a) == token {
    return true
  }
  return false
}

