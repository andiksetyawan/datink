package swipe

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/swipe/entity"
)

func (r *repository) GetByUserIdOrFriendID(ctx context.Context, db bun.IDB, userID, friendID int64) (*entity.Swipe, error) {
	var swipe entity.Swipe
	if err := db.NewSelect().
		Model(&swipe).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("user_id = ?", userID).Where("friend_id = ?", friendID)
		}).
		WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("friend_id = ?", userID).Where("user_id = ?", friendID)
		}).
		Scan(ctx, &swipe); err != nil {
		return nil, err
	}

	return &swipe, nil
}
