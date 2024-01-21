package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/common"
	entitySwipe "datink/internal/models/swipe/entity"
	"datink/internal/models/user/entity"
)

func (r *repository) FindUserMatch(ctx context.Context, db bun.IDB, userID int64, opt *common.Paging) ([]entity.User, error) {
	var users []entity.User

	if err := db.NewSelect().
		Model(&users).
		Join("JOIN swipes AS s on (\"user\".id=\"s\".user_id and \"s\".friend_id = ? and swipe_direction = ?)", userID, entitySwipe.SwipeDirectionMatch).
		Where("s.user_id = ?", userID).
		WhereOr("s.friend_id = ?", userID).
		Limit(opt.GetLimit()).
		Offset(opt.GetOffset()).
		Scan(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
