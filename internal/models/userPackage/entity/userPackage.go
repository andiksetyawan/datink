package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type UserPackage struct {
	bun.BaseModel

	ID          int64     `json:"id" bun:"id,pk,autoincrement"`
	UserID      int64     `json:"userID"`
	PackageID   int64     `json:"packageID"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	TotalAmount float64   `json:"totalAmount"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bun:",soft_delete,nullzero"`
}
