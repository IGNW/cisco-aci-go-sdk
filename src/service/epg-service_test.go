// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EPGServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *EPGServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-ET", "A EPG testing tenant made by IGNW")

	assert.NotNil(ten)

	_, err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	ap := suite.client.AppProfiles.New("IGNW-AP2", "A testing application profile made by IGNW")

	ten.AddAppProfile(ap)

	_, err = suite.client.AppProfiles.Save(ap)

	assert.Nil(err)

	e := suite.client.EPGs.New("IGNW-E1", "A testing EPG made by IGNW")

	e.IsAttributeBased = true
	e.PreferredPolicyControl = "enforced"
	e.LabelMatchCriteria = "All"
	e.IsPreferredGroupMember = "include"

	ap.AddEPG(e)

	_, err = suite.client.EPGs.Save(e)

	assert.Nil(err)
}

func (suite *EPGServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.EPGs.Delete("uni/tn-IGNW-ET/ap-IGNW-AP2/epg-IGNW-E1")

	assert.Nil(err)

	err = suite.client.AppProfiles.Delete("uni/tn-IGNW-ET/ap-IGNW-AP2")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-ET")

	assert.Nil(err)
}

func (suite *EPGServiceTestSuite) TestEPGServiceGet() {
	assert := assert.New(suite.T())

	e, err := suite.client.EPGs.Get("uni/tn-IGNW-ET/ap-IGNW-AP2/epg-IGNW-E1")

	assert.Nil(err)

	if assert.NotNil(e) {

		assert.Equal("IGNW-E1", e.Name)
		assert.Equal("epg-IGNW-E1", e.ResourceName)
		assert.Equal("uni/tn-IGNW-ET/ap-IGNW-AP2/epg-IGNW-E1", e.DomainName)
		assert.Equal("A testing EPG made by IGNW", e.Description)
		assert.Empty(e.Status)

	}

}

func (suite *EPGServiceTestSuite) TestEPGServiceGetByName() {
	assert := assert.New(suite.T())

	EPGs, err := suite.client.EPGs.GetByName("IGNW-E1")

	assert.Nil(err)

	if assert.NotEmpty(EPGs) {

		assert.Len(EPGs, 1)

		assert.Contains(EPGs, &models.EPG{
			models.ResourceAttributes{
				Name:         "IGNW-E1",
				ResourceName: "epg-IGNW-E1",
				DomainName:   "uni/tn-IGNW-ET/ap-IGNW-AP2/epg-IGNW-E1",
				Description:  "A testing EPG made by IGNW",
				ObjectClass:  "fvAEPg",
				Status:       "",
			},
			true,
			"enforced",
			"All",
			"include",
		})
	}
}

func (suite *EPGServiceTestSuite) TestEPGServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.EPGs.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.EPG{
			models.ResourceAttributes{
				Name:         "IGNW-E1",
				ResourceName: "epg-IGNW-E1",
				DomainName:   "uni/tn-IGNW-ET/ap-IGNW-AP2/epg-IGNW-E1",
				Description:  "A testing EPG made by IGNW",
				ObjectClass:  "fvAEPg",
				Status:       "",
			},
			true,
			"enforced",
			"All",
			"include",
		})

	}
}

func TestEPGServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EPGServiceTestSuite))
}
