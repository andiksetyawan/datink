package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Package struct {
	bun.BaseModel

	ID          int64   `json:"id" bun:"id,pk,autoincrement"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bun:",soft_delete,nullzero"`
}
