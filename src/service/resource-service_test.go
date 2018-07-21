// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ResourceServiceTestSuite struct {
	suite.Suite
}

func (suite *ResourceServiceTestSuite) Test_getResourceName() {
	assert := assert.New(suite.T())

	service := new(ResourceService)

	service.ResourceNamePrefix = "yy"

	assert.Equal("yy-thing", service.getResourceName("thing"))

}

func (suite *ResourceServiceTestSuite) Test_getResourcePath() {
	assert := assert.New(suite.T())

	service := new(ResourceService)

	service.ResourceNamePrefix = "yy"

	t := &models.Tenant{models.ResourceAttributes{
		Name:         "IGNW-T1",
		ResourceName: "t-IGNW-T1",
		Status:       "",
	},
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	ap := &models.AppProfile{models.ResourceAttributes{
		Name:         "IGNW-A1",
		ResourceName: "ap-IGNW-A1",
		Status:       "",
	},
		nil,
	}

	epg := &models.EPG{models.ResourceAttributes{
		Name:         "IGNW-EP1",
		ResourceName: "ep-IGNW-EP1",
		Status:       "",
	},
		false,
		"",
		"",
		"",
	}

	t.AddAppProfile(ap)
	ap.AddEPG(epg)

	assert.Equal("/api/node/mo/uni/t-IGNW-T1.json", service.getResourcePath(t, ""))
	assert.Equal("/api/node/mo/uni/t-IGNW-T1/ap-IGNW-A1.json", service.getResourcePath(ap, ""))
	assert.Equal("/api/node/mo/uni/t-IGNW-T1/ap-IGNW-A1/ep-IGNW-EP1.json", service.getResourcePath(epg, ""))
}

func TestResourceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceServiceTestSuite))
}
