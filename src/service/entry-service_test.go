// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TODO: add more test cases to cover interactions between settings based on EthernetType and Protcol

type EntryServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *EntryServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-ETT", "A Filter testing tenant made by IGNW")

	assert.NotNil(ten)

	err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	f := suite.client.Filters.New("IGNW-F2", "A testing Filter made by IGNW")

	ten.AddFilter(f)

	err = suite.client.Filters.Save(f)

	assert.Nil(err)

	e := suite.client.Entries.New("IGNW-E1", "A testing Entry made by IGNW")

	e.Protocol = "tcp"
	e.ArpOpCodes = "unspecified"
	e.ApplyToFrag = false
	e.EthernetType = "ip"
	e.ICMPv4Settings = "echo-rep"
	e.ICMPv6Settings = "dst-unreach"
	e.DSCP = "CS0"
	e.Stateful = false
	e.TcpFlags = "est"
	e.Source = &models.ToFrom{To: "8080", From: "8080"}

	// TODO: add more test cases to cover automatic conversions for ports
	// HACK: ACI API automatically converts this to http??
	e.Destination = &models.ToFrom{To: "80", From: "80"}

	f.AddEntry(e)

	err = suite.client.Entries.Save(e)

	assert.Nil(err)
}

func (suite *EntryServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.Filters.Delete("uni/tn-IGNW-ETT/flt-IGNW-F2")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-ETT")

	assert.Nil(err)
}

func (suite *EntryServiceTestSuite) TestEntryServiceGet() {
	assert := assert.New(suite.T())

	e, err := suite.client.Entries.Get("uni/tn-IGNW-ETT/flt-IGNW-F2/e-IGNW-E1")

	assert.Nil(err)

	if assert.NotNil(e) {

		assert.Equal("IGNW-E1", e.Name)
		assert.Equal("e-IGNW-E1", e.ResourceName)
		assert.Equal("uni/tn-IGNW-ETT/flt-IGNW-F2/e-IGNW-E1", e.DomainName)
		assert.Equal("A testing Entry made by IGNW", e.Description)
		assert.Empty(e.Status)

	}

}

func (suite *EntryServiceTestSuite) TestEntryServiceGetByName() {
	assert := assert.New(suite.T())

	entries, err := suite.client.Entries.GetByName("IGNW-E1")

	assert.Nil(err)

	if assert.NotEmpty(entries) {

		assert.Len(entries, 1)

		assert.Contains(entries, &models.Entry{
			models.ResourceAttributes{
				Name:         "IGNW-E1",
				ResourceName: "e-IGNW-E1",
				DomainName:   "uni/tn-IGNW-ETT/flt-IGNW-F2/e-IGNW-E1",
				Description:  "A testing Entry made by IGNW",
				ObjectClass:  "vzEntry",
				Status:       "",
			},
			"tcp",
			"unspecified",
			false,
			"ip",
			"echo-rep",
			"dst-unreach",
			"CS0",
			false,
			"est",
			&models.ToFrom{To: "8080", From: "8080"},
			&models.ToFrom{To: "http", From: "http"},
		})
	}
}

func (suite *EntryServiceTestSuite) TestEntryServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.Entries.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Entry{
			models.ResourceAttributes{
				Name:         "IGNW-E1",
				ResourceName: "e-IGNW-E1",
				DomainName:   "uni/tn-IGNW-ETT/flt-IGNW-F2/e-IGNW-E1",
				Description:  "A testing Entry made by IGNW",
				ObjectClass:  "vzEntry",
				Status:       "",
			},
			"tcp",
			"unspecified",
			false,
			"ip",
			"echo-rep",
			"dst-unreach",
			"CS0",
			false,
			"est",
			&models.ToFrom{To: "8080", From: "8080"},
			&models.ToFrom{To: "http", From: "http"},
		})
	}
}

func TestEntryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EntryServiceTestSuite))
}
