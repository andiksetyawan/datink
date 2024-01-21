package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/user/entity"
)

func (r *repository) Create(ctx context.Context, db bun.IDB, user *entity.User) error {
	_, err := db.NewInsert().Model(user).Exec(ctx)
	return err
}
