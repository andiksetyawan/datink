package swipe

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/swipe/entity"
	"datink/internal/resource"
)

type Repository interface {
	Create(ctx context.Context, db bun.IDB, swipe *entity.Swipe) error
	GetByUserIdOrFriendID(ctx context.Context, db bun.IDB, userID, friendID int64) (*entity.Swipe, error)
	Update(ctx context.Context, db bun.IDB, swipe *entity.Swipe) error
	IncSwipeCount(ctx context.Context, db bun.IDB, swipeCount *entity.SwipeCount) error
	GetSwipeCountNow(ctx context.Context, db bun.IDB, userID int64) (int64, error)
}

type repository struct {
	resource resource.Resource
}

func NewRepository(resource resource.Resource) (Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
