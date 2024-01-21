package packages

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	*require.Assertions

	ctrl            *gomock.Controller
	resource        shared.Resource
	planCityRepo    *mock_planCity.MockRepository
	planCityService planCity.Service
	dbManager       *mock_database.MockDatabaseManager
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.planCityRepo = mock_planCity.NewMockRepository(s.ctrl)
	s.resource = shared.Resource{}
	s.dbManager = mock_database.NewMockDatabaseManager(s.ctrl)
	s.planCityService, _ = New(s.resource, s.planCityRepo, s.dbManager)
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
	helperTime.ResetMock()
}

func (s *ServiceSuite) TestService_New() {
	s.Run("when instantiate", func() {
		result, err := New(s.resource, s.planCityRepo, s.dbManager)
		s.Nil(err)
		s.Equal(&service{
			resource:   s.resource,
			repository: s.planCityRepo,
			dbManager:  s.dbManager,
		}, result)
	})
}
