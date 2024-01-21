package swipe

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"datink/internal/models/swipe/entity"
)

func (r *repository) GetSwipeCountNow(ctx context.Context, db bun.IDB, userID int64) (int64, error) {
	var swipe entity.SwipeCount
	if err := db.NewSelect().
		Model(&swipe).
		Where("user_id = ?", userID).
		Where("swipe_date = ?", time.Now()).
		Scan(ctx, &swipe); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	return swipe.Total, nil
}
