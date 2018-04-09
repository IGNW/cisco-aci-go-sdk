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

type FilterServiceTestSuite struct {
	suite.Suite
}

func (suite *FilterServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten := client.Tenants.New("IGNW-FT", "A Filter testing tenant made by IGNW")

	assert.NotNil(ten)

	err := client.Tenants.Save(ten)

	assert.Nil(err)

	v := client.Filters.New("IGNW-F1", "A testing Filter made by IGNW")

	ten.AddFilter(v)

	err = client.Filters.Save(v)

	assert.Nil(err)
}

func (suite *FilterServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	err := client.Filters.Delete("uni/tn-IGNW-FT/F-IGNW-F1")

	assert.Nil(err)

	err = client.Tenants.Delete("uni/tn-IGNW-FT")

	assert.Nil(err)
}

func (suite *FilterServiceTestSuite) TestFilterServiceGet() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	f, err := client.Filters.Get("uni/tn-IGNW-FT/F-IGNW-F1")

	assert.Nil(err)

	if assert.NotNil(f) {

		assert.Equal("IGNW-F1", f.Name)
		assert.Equal("F-IGNW-F1", f.ResourceName)
		assert.Equal("uni/tn-IGNW-FT/F-IGNW-F1", f.DomainName)
		assert.Equal("A testing Filter made by IGNW", f.Description)
		assert.Empty(f.Status)

	}

}

func (suite *FilterServiceTestSuite) TestFilterServiceGetByName() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	Filters, err := client.Filters.GetByName("IGNW-F1")

	assert.Nil(err)

	if assert.NotEmpty(Filters) {

		assert.Len(Filters, 1)

		assert.Contains(Filters, &models.Filter{
			models.ResourceAttributes{
				Name:         "IGNW-F1",
				ResourceName: "F-IGNW-F1",
				DomainName:   "uni/tn-IGNW-VT/F-IGNW-F1",
				Description:  "A testing Filter made by IGNW",
				ObjectClass:  "fvFilter",
				Status:       "",
			},
			nil,
			nil,
		})
	}
}

func (suite *FilterServiceTestSuite) TestFilterServiceGetAll() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	data, err := client.Filters.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Filter{
			models.ResourceAttributes{
				Name:         "IGNW-F1",
				ResourceName: "F-IGNW-F1",
				DomainName:   "uni/tn-IGNW-VT/F-IGNW-F1",
				Description:  "A testing Filter made by IGNW",
				ObjectClass:  "fvFilter",
				Status:       "",
			},
			nil,
			nil,
		})

		suite.T().Log(fmt.Printf("Got These Filters: %#v", data))

		for key, f := range data {
			suite.T().Log(fmt.Printf("\nFilter #%s has Name %s\n", strconv.Itoa(key), f.Name))
		}
	}
}

func TestFilterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FilterServiceTestSuite))
}
