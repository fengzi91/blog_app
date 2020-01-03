package models

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)
// Attachment is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Attachment struct {
    ID uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
    User User `belongs_to:"user"`
    UserID uuid.UUID `json:"author_id" db:"author_id"`
    Url string `json:"url" db:"url"`
    Size int64 `json:"size" db:"size"`
    File binding.File `db:"-" form:"file"`
    Posts Posts `has_many:"posts"`
}

// String is not required by pop and may be deleted
func (a Attachment) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Attachments is not required by pop and may be deleted
type Attachments []Attachment

// String is not required by pop and may be deleted
func (a Attachments) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Attachment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Attachment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&ImageFile{Field: "Image", Name: "Url", Value: a.File},
	), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Attachment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
type ImageFile struct {
	Name    string
	Field   string
	Expr    string
	Value	binding.File
	Message string
}

func (v *ImageFile) IsValid(errors *validate.Errors) {
	ext := path.Ext(v.Value.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		errors.Add(strings.ToLower(v.Field), fmt.Sprintf("%s must a image!", v.Field))
	}

}

func (a *Attachment) BeforeCreate(tx *pop.Connection) error {
	if !a.File.Valid() {
		return nil
	}
	dir := filepath.Join("./public", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	u1, nil := uuid.NewV4()
	localFile := filepath.Join(dir, u1.String() + path.Ext(a.File.Filename))
	f, err := os.Create(localFile)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(localFile)
	a.Url = strings.Replace(localFile, "public/", "/", 1)
	fmt.Println("转换后的结果")
	fmt.Println(a.Url)
	a.Size = a.File.Size
	defer f.Close()
	_, err = io.Copy(f, a.File)
	return err
}
