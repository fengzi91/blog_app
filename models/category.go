package models

import (
"time"

"github.com/gobuffalo/pop"
"github.com/gobuffalo/validate"
"github.com/gobuffalo/validate/validators"
uuid "github.com/gobuffalo/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name   string    `json:"name" db:"name"`
	Slug   string    `json:"slug" db:"slug"`
	Order  int 		 `json:"order" db:"order"`
	Posts  Posts	 `has_many:"posts"`
}

type Categories []Category

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (c *Category) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "分类名称"},
		&validators.StringIsPresent{Field: c.Slug, Name: "分类别名"},
	), nil
}
