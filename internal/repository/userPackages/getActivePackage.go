package userPackages

import (
	"context"
	"time"

	"github.com/uptrace/bun"

	"datink/internal/models/userPackage/entity"
)

func (r repository) GetActivePackage(ctx context.Context, db bun.IDB, userID int64) (*entity.UserPackage, error) {
	var userPackage entity.UserPackage
	now := time.Now()
	if err := db.NewSelect().
		Model(&userPackage).
		Where("user_id = ?", userID).
		Where("start_date <= ?", now).
		Where("end_date >= ?", now).
		Scan(ctx, &userPackage); err != nil {
		return nil, err
	}

	return &userPackage, nil
}
