package repository

import (
	"go.uber.org/dig"

	"datink/internal/repository/packages"
	"datink/internal/repository/swipe"
	"datink/internal/repository/user"
	"datink/internal/repository/userPackages"
)

type Repository struct {
	dig.In

	User        user.Repository
	Swipe       swipe.Repository
	Packages    packages.Repository
	UserPackage userPackages.Repository
}
