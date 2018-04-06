// +build integration

package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

	err := client.Tenants.Delete("uni/tn-IGNW-BD")

	assert.Nil(err)

	err = client.BridgeDomains.Delete("uni/tn-IGNW-BD/BD-IGNW-BD1")

	assert.Nil(err)
}

//func TestBridgeDomainServiceTestSuite(t *testing.T) {
func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGet() {
	assert := assert.New(suite.T())
	assert.Fail("Not Implemented")

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	bd, err := client.BridgeDomains.Get("IGNW-BD1")

	assert.Nil(err)

	if assert.NotNil(bd) {

		assert.Equal("IGNW-BD1", bd.Name)
		assert.Equal("BD-IGNW-BD1", bd.ResourceName)
		assert.Equal("uni/tn-IGNW/BD-IGNW-BD1", bd.DomainName)
		assert.Empty(bd.Status)

	}

}

func (suite *BridgeDomainServiceTestSuite) TestBridgeDomainServiceGetAll() {
	assert := assert.New(suite.T())
	assert.Fail("Not Implemented")
	/*
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
	*/
}

func TestBridgeDomainServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BridgeDomainServiceTestSuite))
}
