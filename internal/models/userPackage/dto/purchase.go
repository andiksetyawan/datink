package dto

type UserPackagePurchaseRequest struct {
	PackageId int64 `json:"packageID" validate:"required"`
}

type UserPackagePurchaseResponse struct {
	ID int64 `json:"id"`
}
