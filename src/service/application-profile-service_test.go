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

type ApplicationProfileServiceTestSuite struct {
	suite.Suite
}

func (suite *ApplicationProfileServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten := client.Tenants.New("IGNW-APT", "A AP testing tenant made by IGNW")

	assert.NotNil(ten)

	err := client.Tenants.Save(ten)

	assert.Nil(err)

	ap := client.AppProfiles.New("IGNW-AP1", "A testing application profile made by IGNW")

	ten.AddAppProfile(ap)

	err = client.AppProfiles.Save(ap)

	assert.Nil(err)
}

func (suite *ApplicationProfileServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	err := client.AppProfiles.Delete("uni/tn-IGNW-APT/AP-IGNW-AP1")

	assert.Nil(err)

	err = client.Tenants.Delete("uni/tn-IGNW-APT")

	assert.Nil(err)
}

func (suite *ApplicationProfileServiceTestSuite) TestApplicationProfileServiceGet() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ap, err := client.AppProfiles.Get("uni/tn-IGNW-APT/AP-IGNW-AP1")

	assert.Nil(err)

	if assert.NotNil(ap) {

		assert.Equal("IGNW-AP1", ap.Name)
		assert.Equal("AP-IGNW-AP1", ap.ResourceName)
		assert.Equal("uni/tn-IGNW-APT/AP-IGNW-AP1", ap.DomainName)
		assert.Equal("A testing application profile made by IGNW", ap.Description)
		assert.Empty(ap.Status)

	}

}

func (suite *ApplicationProfileServiceTestSuite) TestApplicationProfileServiceGetByName() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	profiles, err := client.Tenants.GetByName("IGNW-C1")

	assert.Nil(err)

	if assert.NotEmpty(profiles) {

		assert.Len(profiles, 1)

		assert.Contains(profiles, &models.AppProfile{
			models.ResourceAttributes{
				Name:         "IGNW-AP1",
				ResourceName: "AP-IGNW-AP1",
				DomainName:   "uni/tn-IGNW/AP-IGNW-AP1",
				Description:  "A testing application profile made by IGNW",
				ObjectClass:  "fvAP",
				Status:       "",
			},
			nil,
			nil,
		})

	}
}

func (suite *ApplicationProfileServiceTestSuite) TestApplicationProfileServiceGetAll() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	data, err := client.Tenants.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.AppProfile{
			models.ResourceAttributes{
				Name:         "IGNW-AP1",
				ResourceName: "AP-IGNW-AP1",
				DomainName:   "uni/tn-IGNW/AP-IGNW-AP1",
				Description:  "A testing application profile made by IGNW",
				ObjectClass:  "fvAP",
				Status:       "",
			},
			nil,
			nil,
		})

		suite.T().Log(fmt.Printf("Got These Application Profiles: %#v", data))

		for key, bd := range data {
			suite.T().Log(fmt.Printf("\nApplication Profile #%s has Name %s\n", strconv.Itoa(key), bd.Name))
		}
	}
}

func TestApplicationProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationProfileServiceTestSuite))
}
