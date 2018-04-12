// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VRFServiceTestSuite struct {
	suite.Suite
}

func (suite *VRFServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	ten := client.Tenants.New("IGNW-VT", "A VRF testing tenant made by IGNW")

	assert.NotNil(ten)

	err := client.Tenants.Save(ten)

	assert.Nil(err)

	v := client.VRFs.New("IGNW-V1", "A testing VRF made by IGNW")

	ten.AddVRF(v)

	err = client.VRFs.Save(v)

	assert.Nil(err)
}

func (suite *VRFServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	err := client.VRFs.Delete("uni/tn-IGNW-VT/ctx-IGNW-V1")

	assert.Nil(err)

	err = client.Tenants.Delete("uni/tn-IGNW-VT")

	assert.Nil(err)
}

func (suite *VRFServiceTestSuite) TestVRFServiceGet() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	v, err := client.VRFs.Get("uni/tn-IGNW-VT/ctx-IGNW-V1")

	assert.Nil(err)

	if assert.NotNil(v) {

		assert.Equal("IGNW-V1", v.Name)
		assert.Equal("ctx-IGNW-V1", v.ResourceName)
		assert.Equal("uni/tn-IGNW-VT/ctx-IGNW-V1", v.DomainName)
		assert.Equal("A testing VRF made by IGNW", v.Description)
		assert.Empty(v.Status)

	}

}

func (suite *VRFServiceTestSuite) TestVRFServiceGetByName() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	VRFs, err := client.VRFs.GetByName("IGNW-V1")

	assert.Nil(err)

	if assert.NotEmpty(VRFs) {

		assert.Len(VRFs, 1)

		assert.Contains(VRFs, &models.VRF{
			models.ResourceAttributes{
				Name:         "IGNW-V1",
				ResourceName: "ctx-IGNW-V1",
				DomainName:   "uni/tn-IGNW-VT/C-IGNW-V1",
				Description:  "A testing VRF made by IGNW",
				ObjectClass:  "fvCtx",
				Status:       "",
			},
			nil,
		})
	}
}

func (suite *VRFServiceTestSuite) TestVRFServiceGetAll() {
	assert := assert.New(suite.T())

	client := GetClient()

	assert.NotNil(client, "\nCould not get Client, therefore tests could not start")

	data, err := client.VRFs.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.AppProfile{
			models.ResourceAttributes{
				Name:         "IGNW-V1",
				ResourceName: "ctx-IGNW-V1",
				DomainName:   "uni/tn-IGNW-VT/ctx-IGNW-V1",
				Description:  "A testing VRF made by IGNW",
				ObjectClass:  "fvCtx",
				Status:       "",
			},
			nil,
			nil,
		})

	}
}

func TestVRFServiceTestSuite(t *testing.T) {
	suite.Run(t, new(VRFServiceTestSuite))
}
