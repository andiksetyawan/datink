package userPackages

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/userPackage/entity"
)

func (r *repository) Create(ctx context.Context, db bun.IDB, user *entity.UserPackage) error {
	_, err := db.NewInsert().Model(user).Exec(ctx)
	return err
}
