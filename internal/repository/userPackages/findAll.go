package userPackages

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/userPackage/entity"
)

func (r repository) FindAll(ctx context.Context, db bun.IDB) ([]entity.UserPackage, error) {
	var userPackages []entity.UserPackage
	if err := db.NewSelect().
		Model(&userPackages).
		Scan(ctx, &userPackages); err != nil {
		return nil, err
	}

	return userPackages, nil
}
