package models

import (
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
	Subject	  string	`json:"subject" db:"subject"`
	Content   string    `json:"content" db:"content"`
	AuthorID  uuid.UUID `json:"author_id" db:"author_id"`
	Author    User      `belongs_to:"user"`
	CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
	Category  Category  `belongs_to:"category"`
	TopImageID uuid.UUID `json:"attachment_id" db:"attachment_id"`
	TopImage string `json:"top_image_url" db:"top_image_url" form:"attachment_url"`
}

type Posts []Post

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (p *Post) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Title, Name: "Title"},
		&validators.StringIsPresent{Field: p.Content, Name: "Content"},
	), nil
}
