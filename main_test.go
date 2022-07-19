package main

// Basic imports
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context

type TurboTestSuite struct {
	suite.Suite
	// arr []int;
	// str_arr []string;
	// VariableThatShouldStartAtFive int
}

// func (suite *TurboTestSuite) SetupTest() {}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTurboTestSuite(t *testing.T) {
	suite.Run(t, new(TurboTestSuite))
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *TurboTestSuite) TestOne() {
	assert.Equal(suite.T(), 4, 4)
	suite.Equal(5, 5)
}

func (suite *TurboTestSuite) TestTwo() {
	suite.Equal(5, 5)
	suite.DirExists("/tmp")
}
