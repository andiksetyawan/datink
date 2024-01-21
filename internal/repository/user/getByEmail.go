package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/user/entity"
)

func (r *repository) GetByEmail(ctx context.Context, db bun.IDB, email string) (*entity.User, error) {
	var user entity.User
	if err := db.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
