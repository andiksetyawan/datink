package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type SwipeDirection string

const (
	SwipeDirectionLike    SwipeDirection = "LIKE"
	SwipeDirectionMatch   SwipeDirection = "MATCH"
	SwipeDirectionDisLike SwipeDirection = "DISLIKE"
)

type Swipe struct {
	bun.BaseModel

	ID             int64          `json:"id" bun:"id,pk,autoincrement"`
	UserID         int64          `json:"userID"`
	FriendID       int64          `json:"friendID"`
	SwipeDirection SwipeDirection `json:"swipeDirection"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bun:",soft_delete,nullzero"`
}

type SwipeCount struct {
	bun.BaseModel

	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	UserID    int64     `json:"userID"`
	SwipeDate time.Time `json:"swipeDate" bun:",nullzero,notnull"`
	Total     int64     `json:"total"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bun:",soft_delete,nullzero"`
}
