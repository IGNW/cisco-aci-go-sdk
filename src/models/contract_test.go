// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ContractTestSuite struct {
	suite.Suite
}

func (suite *ContractTestSuite) TestContractToMap() {
	assert := assert.New(suite.T())

	contractMap := map[string]string{
		"dn":         "",
		"status":     "",
		"descr":      "",
		"name":       "TestContractMapDSCP",
		"scope":      "global",
		"targetDscp": "unspecified",
	}

	contract := Contract{
		ResourceAttributes{Name: "TestContractMapDSCP"},
		"global",
		"unspecified",
		nil,
		nil,
	}

	assert.Equal(contractMap, contract.ToMap())

}

func (suite *ContractTestSuite) TestNewContractFromDefaults() {
	assert := assert.New(suite.T())

	contract := NewContract(NewContractMap())

	expected := Contract{
		ResourceAttributes{Name: ""},
		"context",
		"unspecified",
		nil,
		nil,
	}

	assert.Equal(&expected, contract)
}

func (suite *ContractTestSuite) TestNewContractFromMap() {
	assert := assert.New(suite.T())

	contractMap := map[string]string{
		"dn":         "",
		"status":     "",
		"descr":      "",
		"name":       "TestContractMapDSCP",
		"scope":      "global",
		"targetDscp": "unspecified",
	}

	assert.Equal(&Contract{
		ResourceAttributes{Name: "TestContractMapDSCP"},
		"global",
		"unspecified",
		nil,
		nil,
	}, NewContract(contractMap))

}

func TestContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}
