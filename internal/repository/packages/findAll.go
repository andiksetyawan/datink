package packages

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/package/entity"
)

func (r repository) FindAll(ctx context.Context, db bun.IDB) ([]entity.Package, error) {
	var packages []entity.Package
	if err := db.NewSelect().
		Model(&packages).
		Scan(ctx, &packages); err != nil {
		return nil, err
	}

	return packages, nil
}
