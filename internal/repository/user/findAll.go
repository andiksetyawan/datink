package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/user/entity"
)

func (r *repository) FindAll(ctx context.Context, db bun.IDB) ([]entity.User, error) {
	var users []entity.User
	err := db.NewSelect().Model(&users).Scan(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
