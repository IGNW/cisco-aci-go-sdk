// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EPGTestSuite struct {
	suite.Suite
}

func (suite *EPGTestSuite) TestEPGToMap() {
	assert := assert.New(suite.T())

	epgMap := map[string]string{
		"dn":             "",
		"status":         "",
		"descr":          "",
		"name":           "TestEPGMap",
		"isAttrBasedEPg": "yes",
		"pcEnfPref":      "enforced",
		"prefGrMemb":     "include",
		"matchT":         "All",
	}

	epg := EPG{
		ResourceAttributes{Name: "TestEPGMap"},
		true,
		"enforced",
		"All",
		"include",
	}

	assert.Equal(epgMap, epg.ToMap())

}

func (suite *EPGTestSuite) TestNewEPGFromDefaults() {
	assert := assert.New(suite.T())

	epg := NewEPG(NewEPGMap())

	expected := EPG{
		ResourceAttributes{Name: ""},
		false,
		"unenforced",
		"AtleastOne",
		"exclude",
	}

	assert.Equal(&expected, epg)
}

func (suite *EPGTestSuite) TestNewEPGFromMap() {
	assert := assert.New(suite.T())

	epgMap := map[string]string{
		"dn":             "",
		"status":         "",
		"descr":          "",
		"name":           "TestEPGMap",
		"isAttrBasedEPg": "yes",
		"pcEnfPref":      "enforced",
		"prefGrMemb":     "include",
		"matchT":         "All",
	}

	assert.Equal(&EPG{
		ResourceAttributes{Name: "TestEPGMap"},
		true,
		"enforced",
		"All",
		"include",
	}, NewEPG(epgMap))

}

func TestEPGTestSuite(t *testing.T) {
	suite.Run(t, new(EPGTestSuite))
}
