package models

import (
	"fmt"
	"github.com/pkg/errors"
	"time"

	"github.com/gobuffalo/pop"
	uuid "github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Post struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	AuthorID  uuid.UUID `json:"author_id" db:"author_id"`
	Author    User      `belongs_to:"user"`
	CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
	Category  Category  `belongs_to:"category"`
	TopImageID uuid.UUID `json:"attachment_id" db:"attachment_id"`
	TopImage Attachment `belongs_to:"attachment"`
}

type Posts []Post

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (p *Post) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Title, Name: "Title"},
		&validators.StringIsPresent{Field: p.Content, Name: "Content"},
	), nil
}

func (p *Post) AfterFind(tx *pop.Connection) error {
	fmt.Println("打印post")
	fmt.Println(p.TopImage)
	return errors.WithStack(nil)
}
