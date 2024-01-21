package userPackage

import (
	"context"
	"time"

	"datink/internal/models/auth"
	"datink/internal/models/userPackage/dto"
	"datink/internal/models/userPackage/entity"
)

func (s service) Purchase(ctx context.Context, user *auth.UserJwtAuth, req dto.UserPackagePurchaseRequest) (dto.UserPackagePurchaseResponse, error) {
	packages, err := s.repository.Packages.GetByID(ctx, s.resource.Database.GetMaster(), req.PackageId)
	if err != nil {
		return dto.UserPackagePurchaseResponse{}, err
	}

	newPremiumPackage := entity.UserPackage{
		UserID:      user.UserID,
		PackageID:   packages.ID,
		StartDate:   time.Now(),
		EndDate:     time.Now().AddDate(0, 1, 0),
		TotalAmount: packages.Price,
	}
	err = s.repository.UserPackage.Create(ctx, s.resource.Database.GetMaster(), &newPremiumPackage)
	if err != nil {
		return dto.UserPackagePurchaseResponse{}, err
	}

	return dto.UserPackagePurchaseResponse{ID: newPremiumPackage.ID}, nil
}
