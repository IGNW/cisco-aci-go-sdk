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

type ContractServiceTestSuite struct {
	suite.Suite
}

func (suite *ContractServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten := client.Tenants.New("IGNW-CT", "A contract testing tenant made by IGNW")

	assert.NotNil(ten)

	err := client.Tenants.Save(ten)

	assert.Nil(err)

	c := client.Contracts.New("IGNW-C1", "A testing contract made by IGNW")

	ten.AddContract(c)

	err = client.Contracts.Save(c)

	assert.Nil(err)
}

func (suite *ContractServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	err := client.Contracts.Delete("uni/tn-IGNW-CT/C-IGNW-C1")

	assert.Nil(err)

	err = client.Tenants.Delete("uni/tn-IGNW-CT")

	assert.Nil(err)
}

func (suite *ContractServiceTestSuite) TestContractServiceGet() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	c, err := client.Contracts.Get("uni/tn-IGNW-CT/AP-IGNW-C1")

	assert.Nil(err)

	if assert.NotNil(c) {

		assert.Equal("IGNW-C1", c.Name)
		assert.Equal("C-IGNW-C1", c.ResourceName)
		assert.Equal("uni/tn-IGNW-CT/AP-IGNW-C1", c.DomainName)
		assert.Equal("A testing contract made by IGNW", c.Description)
		assert.Empty(c.Status)

	}

}

func (suite *ContractServiceTestSuite) TestContractServiceGetByName() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	contracts, err := client.Contracts.GetByName("IGNW-C1")

	assert.Nil(err)

	if assert.NotEmpty(contracts) {

		assert.Len(contracts, 1)

		assert.Contains(contracts, &models.Contract{
			models.ResourceAttributes{
				Name:         "IGNW-C1",
				ResourceName: "C-IGNW-C1",
				DomainName:   "uni/tn-IGNW-CT/C-IGNW-C1",
				Description:  "A testing contract made by IGNW",
				ObjectClass:  "fvContract",
				Status:       "",
			},
			nil,
			nil,
		})
	}
}

func (suite *ContractServiceTestSuite) TestContractServiceGetAll() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	data, err := client.Contracts.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.AppProfile{
			models.ResourceAttributes{
				Name:         "IGNW-C1",
				ResourceName: "C-IGNW-C1",
				DomainName:   "uni/tn-IGNW-CT/C-IGNW-C1",
				Description:  "A testing contract made by IGNW",
				ObjectClass:  "fvContract",
				Status:       "",
			},
			nil,
			nil,
		})

		suite.T().Log(fmt.Printf("Got These Contracts: %#v", data))

		for key, c := range data {
			suite.T().Log(fmt.Printf("\nContract #%s has Name %s\n", strconv.Itoa(key), c.Name))
		}
	}
}

func TestContractServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ContractServiceTestSuite))
}
