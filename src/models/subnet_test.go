// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SubnetTestSuite struct {
	suite.Suite
}

func (suite *SubnetTestSuite) TestSubnetToMap() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":        "",
		"status":    "",
		"descr":     "",
		"name":      "TestSubnetMap",
		"ctrl":      "unspecified",
		"ip":        "10.1.1.100",
		"preferred": "yes",
		"scope":     "public",
		"virtual":   "yes",
	}

	subnet := Subnet{
		ResourceAttributes{Name: "TestSubnetMap"},
		"unspecified",
		"10.1.1.100",
		true,
		[]string{"public"},
		true,
	}

	assert.Equal(sMap, subnet.ToMap())

}

func (suite *SubnetTestSuite) TestSubnetToMapMuliScope() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":        "",
		"status":    "",
		"descr":     "",
		"name":      "TestSubnetMap",
		"ctrl":      "unspecified",
		"ip":        "10.1.1.100",
		"preferred": "yes",
		"scope":     "public,shared",
		"virtual":   "yes",
	}

	subnet := Subnet{
		ResourceAttributes{Name: "TestSubnetMap"},
		"unspecified",
		"10.1.1.100",
		true,
		[]string{"public", "shared"},
		true,
	}

	assert.Equal(sMap, subnet.ToMap())

}

func (suite *SubnetTestSuite) TestNewSubnetFromDefaults() {
	assert := assert.New(suite.T())

	subnet := NewSubnet(NewSubnetMap())

	expected := Subnet{
		ResourceAttributes{Name: "", ResourceName: "subnet-[]"},
		"nd",
		"",
		false,
		[]string{"private"},
		false,
	}

	assert.Equal(&expected, subnet)
}

func (suite *SubnetTestSuite) TestNewSubnetFromMap() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":        "",
		"status":    "",
		"descr":     "",
		"name":      "TestSubnetMap",
		"ctrl":      "unspecified",
		"ip":        "10.1.1.101",
		"preferred": "yes",
		"scope":     "public",
		"virtual":   "yes",
	}

	assert.Equal(&Subnet{
		ResourceAttributes{Name: "TestSubnetMap", ResourceName: "subnet-[10.1.1.101]"},
		"unspecified",
		"10.1.1.101",
		true,
		[]string{"public"},
		true,
	}, NewSubnet(sMap))

}

func (suite *SubnetTestSuite) TestNewSubnetFromMapMultiScope() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":        "",
		"status":    "",
		"descr":     "",
		"name":      "TestSubnetMap",
		"ctrl":      "unspecified",
		"ip":        "10.1.1.101",
		"preferred": "yes",
		"scope":     "public,shared",
		"virtual":   "yes",
	}

	assert.Equal(&Subnet{
		ResourceAttributes{Name: "TestSubnetMap", ResourceName: "subnet-[10.1.1.101]"},
		"unspecified",
		"10.1.1.101",
		true,
		[]string{"public", "shared"},
		true,
	}, NewSubnet(sMap))

}

func TestSubnetTestSuite(t *testing.T) {
	suite.Run(t, new(SubnetTestSuite))
}
