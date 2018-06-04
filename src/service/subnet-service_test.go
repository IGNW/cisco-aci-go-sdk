// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SubnetServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *SubnetServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-ST", "A Subnet testing tenant made by IGNW")

	assert.NotNil(ten)

	_, err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	bd := suite.client.BridgeDomains.New("IGNW-BD2", "A testing bridge domain made by IGNW")

	ten.AddBridgeDomain(bd)

	_, err = suite.client.BridgeDomains.Save(bd)

	assert.Nil(err)

	s := suite.client.Subnets.New("IGNW-S1", "A testing Subnet made by IGNW")

	s.IpAddress = "10.0.0.1/24"

	bd.AddSubnet(s)

	_, err = suite.client.Subnets.Save(s)

	assert.Nil(err)
}

func (suite *SubnetServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.Subnets.Delete("uni/tn-IGNW-ST/BD-IGNW-BD2/subnet-[10.0.0.1/24]")

	assert.Nil(err)

	err = suite.client.BridgeDomains.Delete("uni/tn-IGNW-ST/BD-IGNW-BD2")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-ST")

	assert.Nil(err)
}

func (suite *SubnetServiceTestSuite) TestSubnetServiceGet() {
	assert := assert.New(suite.T())

	s, err := suite.client.Subnets.Get("uni/tn-IGNW-ST/BD-IGNW-BD2/subnet-[10.0.0.1/24]")

	assert.Nil(err)

	if assert.NotNil(s) {

		assert.Equal("IGNW-S1", s.Name)
		assert.Equal("subnet-[10.0.0.1/24]", s.ResourceName)
		assert.Equal("uni/tn-IGNW-ST/BD-IGNW-BD2/subnet-[10.0.0.1/24]", s.DomainName)
		assert.Equal("A testing Subnet made by IGNW", s.Description)
		assert.Empty(s.Status)

	}

}

func (suite *SubnetServiceTestSuite) TestSubnetServiceGetByName() {
	assert := assert.New(suite.T())

	subnets, err := suite.client.Subnets.GetByName("IGNW-S1")

	assert.Nil(err)

	if assert.NotEmpty(subnets) {

		assert.Len(subnets, 1)

		assert.Contains(subnets, &models.Subnet{
			models.ResourceAttributes{
				Name:         "IGNW-S1",
				ResourceName: "subnet-[10.0.0.1/24]",
				DomainName:   "uni/tn-IGNW-ST/BD-IGNW-BD2/subnet-[10.0.0.1/24]",
				Description:  "A testing Subnet made by IGNW",
				ObjectClass:  "fvSubnet",
				Status:       "",
			},
			"nd",
			"10.0.0.1/24",
			false,
			[]string{"private"},
			false,
		})
	}
}

func (suite *SubnetServiceTestSuite) TestSubnetServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.Subnets.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Subnet{
			models.ResourceAttributes{
				Name:         "IGNW-S1",
				ResourceName: "subnet-[10.0.0.1/24]",
				DomainName:   "uni/tn-IGNW-ST/BD-IGNW-BD2/subnet-[10.0.0.1/24]",
				Description:  "A testing Subnet made by IGNW",
				ObjectClass:  "fvSubnet",
				Status:       "",
			},
			"nd",
			"10.0.0.1/24",
			false,
			[]string{"private"},
			false,
		})
	}
}

func TestSubnetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SubnetServiceTestSuite))
}
