package swipe

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/swipe/entity"
)

func (r *repository) Update(ctx context.Context, db bun.IDB, swipe *entity.Swipe) error {
	_, err := db.NewUpdate().Model(swipe).OmitZero().WherePK().Exec(ctx)
	return err
}
