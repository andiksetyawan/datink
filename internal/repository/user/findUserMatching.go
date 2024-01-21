package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/common"
	entity2 "datink/internal/models/swipe/entity"
	"datink/internal/models/user/entity"
)

func (r *repository) FindUserMatching(ctx context.Context, db bun.IDB, userID int64, opt *common.Paging) ([]entity.User, error) {
	var users []entity.User

	if err := db.NewSelect().
		Model(&users).
		Join("LEFT OUTER JOIN swipes AS s on (\"user\".id=\"s\".user_id and \"s\".friend_id = ? and swipe_direction = ?)", userID, entity2.SwipeDirectionLike).
		Where("\"user\".id <> ?", userID).
		OrderExpr("coalesce(s.friend_id , random()*-1000000) desc").
		Limit(opt.GetLimit()).
		Offset(opt.GetOffset()).
		Scan(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
