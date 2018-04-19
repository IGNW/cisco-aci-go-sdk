// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BridgeDomainServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *BridgeDomainServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-BDT", "A BD testing tenant made by IGNW")

	assert.NotNil(ten)

	err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	bd := suite.client.BridgeDomains.New("IGNW-BD1", "A testing bridge domain made by IGNW")

	ten.AddBridgeDomain(bd)

	err = suite.client.BridgeDomains.Save(bd)

	assert.Nil(err)
}

func (suite *BridgeDomainServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.BridgeDomains.Delete("uni/tn-IGNW-BDT/BD-IGNW-BD1")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-BDT")

	assert.Nil(err)
}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGet() {
	assert := assert.New(suite.T())

	bd, err := suite.client.BridgeDomains.Get("uni/tn-IGNW-BDT/BD-IGNW-BD1")

	assert.Nil(err)

	if assert.NotNil(bd) {

		assert.Equal("IGNW-BD1", bd.Name)
		assert.Equal("BD-IGNW-BD1", bd.ResourceName)
		assert.Equal("uni/tn-IGNW-BDT/BD-IGNW-BD1", bd.DomainName)
		assert.Equal("A testing bridge domain made by IGNW", bd.Description)
		assert.Empty(bd.Status)

	}

}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGetByName() {
	assert := assert.New(suite.T())

	domains, err := suite.client.BridgeDomains.GetByName("IGNW-BD1")

	assert.Nil(err)

	if assert.NotEmpty(domains) {

		assert.Len(domains, 1)

		assert.Contains(domains, &models.BridgeDomain{
			models.ResourceAttributes{
				Name:         "IGNW-BD1",
				ResourceName: "BD-IGNW-BD1",
				DomainName:   "uni/tn-IGNW-BDT/BD-IGNW-BD1",
				Description:  "A testing bridge domain made by IGNW",
				ObjectClass:  "fvBD",
				Status:       "",
			},
			"",
			false,
			false,
			"",
			false,
			false,
			false,
			false,
			"",
			"",
			"",
			false,
			false,
			"",
			"",
			"",
			nil,
			nil,
		})

	}
}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.BridgeDomains.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.BridgeDomain{
			models.ResourceAttributes{
				Name:         "IGNW-BD1",
				ResourceName: "BD-IGNW-BD1",
				DomainName:   "uni/tn-IGNW-BDT/BD-IGNW-BD1",
				Description:  "A testing bridge domain made by IGNW",
				ObjectClass:  "fvBD",
				Status:       "",
			},
			"",
			false,
			false,
			"",
			false,
			false,
			false,
			false,
			"",
			"",
			"",
			false,
			false,
			"",
			"",
			"",
			nil,
			nil,
		})
	}
}

func TestBridgeDomainServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BridgeDomainServiceTestSuite))
}
