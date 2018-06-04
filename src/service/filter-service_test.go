// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FilterServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *FilterServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-FT", "A Filter testing tenant made by IGNW")

	assert.NotNil(ten)

	_, err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	v := suite.client.Filters.New("IGNW-F1", "A testing Filter made by IGNW")

	ten.AddFilter(v)

	_, err = suite.client.Filters.Save(v)

	assert.Nil(err)
}

func (suite *FilterServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.Filters.Delete("uni/tn-IGNW-FT/flt-IGNW-F1")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-FT")

	assert.Nil(err)
}

func (suite *FilterServiceTestSuite) TestFilterServiceGet() {
	assert := assert.New(suite.T())

	f, err := suite.client.Filters.Get("uni/tn-IGNW-FT/flt-IGNW-F1")

	assert.Nil(err)

	if assert.NotNil(f) {

		assert.Equal("IGNW-F1", f.Name)
		assert.Equal("flt-IGNW-F1", f.ResourceName)
		assert.Equal("uni/tn-IGNW-FT/flt-IGNW-F1", f.DomainName)
		assert.Equal("A testing Filter made by IGNW", f.Description)
		assert.Empty(f.Status)

	}

}

func (suite *FilterServiceTestSuite) TestFilterServiceGetByName() {
	assert := assert.New(suite.T())

	filters, err := suite.client.Filters.GetByName("IGNW-F1")

	assert.Nil(err)

	if assert.NotEmpty(filters) {

		assert.Len(filters, 1)

		assert.Contains(filters, &models.Filter{
			models.ResourceAttributes{
				Name:         "IGNW-F1",
				ResourceName: "flt-IGNW-F1",
				DomainName:   "uni/tn-IGNW-FT/flt-IGNW-F1",
				Description:  "A testing Filter made by IGNW",
				ObjectClass:  "vzFilter",
				Status:       "",
			},
			nil,
			nil,
		})
	}
}

func (suite *FilterServiceTestSuite) TestFilterServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.Filters.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Filter{
			models.ResourceAttributes{
				Name:         "IGNW-F1",
				ResourceName: "flt-IGNW-F1",
				DomainName:   "uni/tn-IGNW-FT/flt-IGNW-F1",
				Description:  "A testing Filter made by IGNW",
				ObjectClass:  "vzFilter",
				Status:       "",
			},
			nil,
			nil,
		})
	}
}

func TestFilterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FilterServiceTestSuite))
}
