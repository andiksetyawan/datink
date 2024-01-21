package swipe

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/swipe/entity"
)

func (r *repository) Create(ctx context.Context, db bun.IDB, swipe *entity.Swipe) error {
	_, err := db.NewInsert().Model(swipe).Exec(ctx)
	return err
}
