package packages

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/package/entity"
)

func (r repository) GetByID(ctx context.Context, db bun.IDB, id int64) (*entity.Package, error) {
	var packages entity.Package
	if err := db.NewSelect().
		Model(&packages).
		Where("id = ?", id).
		Scan(ctx, &packages); err != nil {
		return nil, err
	}

	return &packages, nil
}
