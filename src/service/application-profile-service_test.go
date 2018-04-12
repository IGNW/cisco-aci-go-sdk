// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ApplicationProfileServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *ApplicationProfileServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-APT", "A AP testing tenant made by IGNW")

	assert.NotNil(ten)

	err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	ap := suite.client.AppProfiles.New("IGNW-AP1", "A testing application profile made by IGNW")

	ten.AddAppProfile(ap)

	err = suite.client.AppProfiles.Save(ap)

	assert.Nil(err)
}

func (suite *ApplicationProfileServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.AppProfiles.Delete("uni/tn-IGNW-APT/ap-IGNW-AP1")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-APT")

	assert.Nil(err)
}

func (suite *ApplicationProfileServiceTestSuite) TestApplicationProfileServiceGet() {
	assert := assert.New(suite.T())

	ap, err := suite.client.AppProfiles.Get("uni/tn-IGNW-APT/ap-IGNW-AP1")

	assert.Nil(err)

	if assert.NotNil(ap) {

		assert.Equal("IGNW-AP1", ap.Name)
		assert.Equal("ap-IGNW-AP1", ap.ResourceName)
		assert.Equal("uni/tn-IGNW-APT/ap-IGNW-AP1", ap.DomainName)
		assert.Equal("A testing application profile made by IGNW", ap.Description)
		assert.Empty(ap.Status)

	}

}

func (suite *ApplicationProfileServiceTestSuite) TestApplicationProfileServiceGetByName() {
	assert := assert.New(suite.T())

	profiles, err := suite.client.AppProfiles.GetByName("IGNW-AP1")

	assert.Nil(err)

	if assert.NotEmpty(profiles) {

		assert.Len(profiles, 1)

		assert.Contains(profiles, &models.AppProfile{
			models.ResourceAttributes{
				Name:         "IGNW-AP1",
				ResourceName: "ap-IGNW-AP1",
				DomainName:   "uni/tn-IGNW-APT/ap-IGNW-AP1",
				Description:  "A testing application profile made by IGNW",
				ObjectClass:  "fvAp",
				Status:       "",
			},
			nil,
			nil,
		})

	}
}

func (suite *ApplicationProfileServiceTestSuite) TestApplicationProfileServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.AppProfiles.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.AppProfile{
			models.ResourceAttributes{
				Name:         "IGNW-AP1",
				ResourceName: "ap-IGNW-AP1",
				DomainName:   "uni/tn-IGNW-APT/ap-IGNW-AP1",
				Description:  "A testing application profile made by IGNW",
				ObjectClass:  "fvAp",
				Status:       "",
			},
			nil,
			nil,
		})
	}
}

func TestApplicationProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationProfileServiceTestSuite))
}
