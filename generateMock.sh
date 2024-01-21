# - service
mkdir -p mocks/internal
# - file
mkdir -p mocks/internal/service/packages
mockgen --source=internal/service/packages/base.go -package=mock_file --destination=mocks/internal/service/packages/mock_handler.go
