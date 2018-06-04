// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TODO: fix GetByName & GetAll tests, seems to be problem with the parent ref or child collections
type ContractServiceTestSuite struct {
	suite.Suite
	client *Client
}

// TODO: Expand to include full contract with subject and filters
func (suite *ContractServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-CT", "A contract testing tenant made by IGNW")

	assert.NotNil(ten)

	_, err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	c := suite.client.Contracts.New("IGNW-C1", "A testing contract made by IGNW")

	c.DSCP = "unspecified"
	c.Scope = "context"

	ten.AddContract(c)

	_, err = suite.client.Contracts.Save(c)

	assert.Nil(err)
}

func (suite *ContractServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.Contracts.Delete("uni/tn-IGNW-CT/brc-IGNW-C1")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-CT")

	assert.Nil(err)
}

func (suite *ContractServiceTestSuite) TestContractServiceGet() {
	assert := assert.New(suite.T())

	c, err := suite.client.Contracts.Get("uni/tn-IGNW-CT/brc-IGNW-C1")

	assert.Nil(err)

	if assert.NotNil(c) {

		assert.Equal("IGNW-C1", c.Name)
		assert.Equal("brc-IGNW-C1", c.ResourceName)
		assert.Equal("uni/tn-IGNW-CT/brc-IGNW-C1", c.DomainName)
		assert.Equal("A testing contract made by IGNW", c.Description)
		assert.Empty(c.Status)

	}

}

func (suite *ContractServiceTestSuite) TestContractServiceGetByName() {
	assert := assert.New(suite.T())

	contracts, err := suite.client.Contracts.GetByName("IGNW-C1")

	assert.Nil(err)

	if assert.NotEmpty(contracts) {

		assert.Len(contracts, 1)

		assert.Contains(contracts, &models.Contract{
			models.ResourceAttributes{
				Name:         "IGNW-C1",
				ResourceName: "brc-IGNW-C1",
				DomainName:   "uni/tn-IGNW-CT/brc-IGNW-C1",
				Description:  "A testing contract made by IGNW",
				ObjectClass:  "vzBrCP",
				Status:       "",
			},
			"context",
			"unspecified",
			nil,
			nil,
		})
	}
}

func (suite *ContractServiceTestSuite) TestContractServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.Contracts.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Contract{
			models.ResourceAttributes{
				Name:         "IGNW-C1",
				ResourceName: "brc-IGNW-C1",
				DomainName:   "uni/tn-IGNW-CT/brc-IGNW-C1",
				Description:  "A testing contract made by IGNW",
				ObjectClass:  "vzBrCP",
				Status:       "",
			},
			"context",
			"unspecified",
			nil,
			nil,
		})

	}
}

func TestContractServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ContractServiceTestSuite))
}
