# - service
mkdir -p mocks/internal

# - packages
mkdir -p mocks/internal/service/packages
mockgen --source=internal/service/packages/base.go -package=packages --destination=mocks/internal/service/packages/mock_handler.go
