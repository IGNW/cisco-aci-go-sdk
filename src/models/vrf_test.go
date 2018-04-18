// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VRFTestSuite struct {
	suite.Suite
}

func (suite *VRFTestSuite) TestVRFToMap() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":        "",
		"status":    "",
		"descr":     "",
		"name":      "TestVRFMap",
		"pcEnfPref": "enforced",
		"pcEnfDir":  "egress",
	}

	vrf := VRF{
		ResourceAttributes{Name: "TestVRFMap"},
		"enforced",
		"egress",
		nil,
	}

	assert.Equal(sMap, vrf.ToMap())

}

func (suite *VRFTestSuite) TestNewVRFFromDefaults() {
	assert := assert.New(suite.T())

	vrf := NewVRF(NewVRFMap())

	expected := VRF{
		ResourceAttributes{Name: ""},
		"unenforced",
		"ingress",
		nil,
	}

	assert.Equal(&expected, vrf)
}

func (suite *VRFTestSuite) TestNewVRFFromMap() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":        "",
		"status":    "",
		"descr":     "",
		"name":      "TestVRFMap",
		"pcEnfPref": "enforced",
		"pcEnfDir":  "egress",
	}

	assert.Equal(&VRF{
		ResourceAttributes{Name: "TestVRFMap"},
		"enforced",
		"egress",
		nil,
	}, NewVRF(sMap))

}

func TestVRFTestSuite(t *testing.T) {
	suite.Run(t, new(VRFTestSuite))
}
