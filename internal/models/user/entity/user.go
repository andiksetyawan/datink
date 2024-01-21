package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel

	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified"`
	Birthdate time.Time `json:"birthdate"`
	Pictures  []string  `json:"pictures" bun:",array"`
	IsPremium bool      `json:"isPremium"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bun:",soft_delete,nullzero"`
}
