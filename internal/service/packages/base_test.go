package packages

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"datink/internal/resource"
)

type ServiceSuite struct {
	suite.Suite
	*require.Assertions

	ctrl     *gomock.Controller
	resource resource.Resource
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

//todo
