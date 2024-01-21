package packages

import (
	"context"

	"datink/internal/models/package/entity"
)

func (s service) FindAll(ctx context.Context) ([]entity.Package, error) {
	return s.repository.Packages.FindAll(ctx, s.resource.Database.GetMaster())
}
