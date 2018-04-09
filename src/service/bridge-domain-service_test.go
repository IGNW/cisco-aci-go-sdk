// +build integration

package service

import (
	"fmt"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type BridgeDomainServiceTestSuite struct {
	suite.Suite
}

func (suite *BridgeDomainServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten := client.Tenants.New("IGNW-BDT", "A BD testing tenant made by IGNW")

	assert.NotNil(ten)

	err := client.Tenants.Save(ten)

	assert.Nil(err)

	bd := client.BridgeDomains.New("IGNW-BD1", "A testing bridge domain made by IGNW")

	ten.AddBridgeDomain(bd)

	err = client.BridgeDomains.Save(bd)

	assert.Nil(err)
}

func (suite *BridgeDomainServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	err := client.BridgeDomains.Delete("uni/tn-IGNW-BDT/BD-IGNW-BD1")

	assert.Nil(err)

	err = client.Tenants.Delete("uni/tn-IGNW-BDT")

	assert.Nil(err)
}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGet() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	bd, err := client.BridgeDomains.Get("uni/tn-IGNW-BDT/BD-IGNW-BD1")

	assert.Nil(err)

	if assert.NotNil(bd) {

		assert.Equal("IGNW-BD1", bd.Name)
		assert.Equal("BD-IGNW-BD1", bd.ResourceName)
		assert.Equal("uni/tn-IGNW-BDT/BD-IGNW-BD1", bd.DomainName)
		assert.Equal("A BD testing tenant made by IGNW", bd.Description)
		assert.Empty(bd.Status)

	}

}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGetByName() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	domains, err := client.BridgeDomains.GetByName("IGNW-BD1")

	assert.Nil(err)

	if assert.NotEmpty(domains) {

		assert.Len(domains, 1)

		assert.Contains(domains, &models.BridgeDomain{
			models.ResourceAttributes{
				Name:         "IGNW-BD1",
				ResourceName: "BD-IGNW-BD1",
				DomainName:   "uni/tn-IGNW/BD-IGNW-BD1",
				Description:  "A BD testing tenant made by IGNW",
				ObjectClass:  "fvBD",
				Status:       "",
			},
			nil,
			nil,
		})

	}
}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGetAll() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	data, err := client.BridgeDomains.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.BridgeDomain{
			models.ResourceAttributes{
				Name:         "IGNW-BD1",
				ResourceName: "BD-IGNW-BD1",
				DomainName:   "uni/tn-IGNW/BD-IGNW-BD1",
				Description:  "A BD testing tenant made by IGNW",
				ObjectClass:  "fvBD",
				Status:       "",
			},
			nil,
			nil,
		})

		suite.T().Log(fmt.Printf("Got These Bridge Domains: %#v", data))

		for key, bd := range data {
			suite.T().Log(fmt.Printf("\nBridge Domain  #%s has Name %s\n", strconv.Itoa(key), bd.Name))
		}
	}
}

func TestBridgeDomainServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BridgeDomainServiceTestSuite))
}
