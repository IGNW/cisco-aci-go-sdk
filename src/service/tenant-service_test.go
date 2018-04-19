// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TenantServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *TenantServiceTestSuite) SetupSuite() {
	suite.T().Log("SetupSuite")

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW", "A Testing tenant made by IGNW")

	assert.NotNil(ten)

	err := suite.client.Tenants.Save(ten)

	assert.Nil(err)
}

func (suite *TenantServiceTestSuite) TearDownSuite() {
	suite.T().Log("TearDownSuite")

	assert := assert.New(suite.T())

	err := suite.client.Tenants.Delete("uni/tn-IGNW")

	assert.Nil(err)
}

func (suite *TenantServiceTestSuite) TestTenantServiceGet() {
	suite.T().Log("TestTenantServiceGet")

	assert := assert.New(suite.T())

	ten, err := suite.client.Tenants.Get("uni/tn-IGNW")

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
	suite.T().Log("TestTenantServiceGetByName")

	assert := assert.New(suite.T())

	tenants, err := suite.client.Tenants.GetByName("IGNW")

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
			nil,
			nil,
			nil,
			nil,
			nil,
		})
	}
}

func (suite *TenantServiceTestSuite) TestTenantServiceGetAll() {
	suite.T().Log("TestTenantServiceGetAll")

	assert := assert.New(suite.T())

	data, err := suite.client.Tenants.GetAll()

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
			nil,
			nil,
			nil,
			nil,
			nil,
		})
	}
}

func TestTenantServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TenantServiceTestSuite))
}
