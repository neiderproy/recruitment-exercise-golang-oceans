package factory

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSambleErrorAmount() {
	err := make(chan error)
	s.adapter.StartAssemblingProcess(-1, nil, err,nil)
	errorExpect := errors.New("not valid amount")
	s.Assert().Equal(errorExpect, err)
}
