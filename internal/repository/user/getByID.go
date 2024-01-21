package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/user/entity"
)

func (r *repository) GetByID(ctx context.Context, db bun.IDB, userID int64) (*entity.User, error) {
	var user entity.User
	if err := db.NewSelect().
		Model(&user).
		Where("id = ?", userID).
		Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
