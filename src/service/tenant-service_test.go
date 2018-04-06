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

type TenantServiceTestSuite struct {
	suite.Suite
}

func (suite *TenantServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten := client.Tenants.New("IGNW", "A Testing tenant made by IGNW")

	assert.NotNil(ten)

	err := client.Tenants.Save(ten)

	assert.Nil(err)
}

func (suite *TenantServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	err := client.Tenants.Delete("uni/tn-IGNW")

	assert.Nil(err)
}

func (suite *TenantServiceTestSuite) TestTenantServiceGet() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten, err := client.Tenants.Get("uni/tn-IGNW")

	assert.Nil(err)

	if assert.NotNil(ten) {

		assert.Equal("IGNW", ten.Name)
		assert.Equal("tn-IGNW", ten.ResourceName)
		assert.Equal("uni/tn-IGNW", ten.DomainName)
		assert.Equal("A Testing tenant made by IGNW", ten.Description)
		assert.Empty(ten.Status)

	}
}

func (suite *TenantServiceTestSuite) TestTenantServiceGetByName() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	tenants, err := client.Tenants.GetByName("IGNW")

	assert.Nil(err)

	if assert.NotEmpty(tenants) {

		assert.Contains(tenants, &models.Tenant{
			models.ResourceAttributes{
				Name:         "IGNW",
				ResourceName: "tn-IGNW",
				DomainName:   "uni/tn-IGNW",
				Description:  "A Testing tenant made by IGNW",
				ObjectClass:  "fvTenant",
				Status:       "",
			},
			"",
			nil,
			nil,
			nil,
			nil,
			nil,
		})
	}
}

func (suite *TenantServiceTestSuite) TestTenantServiceGetAll() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	data, err := client.Tenants.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Tenant{
			models.ResourceAttributes{
				Name:         "IGNW",
				ResourceName: "tn-IGNW",
				DomainName:   "uni/tn-IGNW",
				Description:  "A Testing tenant made by IGNW",
				ObjectClass:  "fvTenant",
				Status:       "",
			},
			"",
			nil,
			nil,
			nil,
			nil,
			nil,
		})

		suite.T().Log(fmt.Printf("Got These Tenants: %#v", data))

		for key, tenant := range data {
			suite.T().Log(fmt.Printf("\nTenant  #%s has Name %s\n", strconv.Itoa(key), tenant.Name))
		}
	}
}

func TestTenantServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TenantServiceTestSuite))
}
