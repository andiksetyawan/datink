package swipe

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/swipe/entity"
)

func (r *repository) IncSwipeCount(ctx context.Context, db bun.IDB, swipeCount *entity.SwipeCount) error {
	_, err := db.NewInsert().Model(swipeCount).
		On("CONFLICT (user_id, swipe_date) DO UPDATE").
		Set("total = swipe_count.total+1").
		Exec(ctx)
	return err
}
