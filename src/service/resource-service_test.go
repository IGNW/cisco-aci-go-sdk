// +build integration-exclude

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

	t := models.Tenant{ResourceAttributes{
		Name:   "IGNW-T1",
		Status: "",
	},
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	ap := models.AppProfile{ResourceAttributes{
		Name:   "IGNW-A1",
		Status: "",
	},
		nil,
	}

	epg := models.EPG{ResourceAttributes{
		Name:   "IGNW-EP1",
		Status: "",
	},
		false,
		false,
		"",
		"",
		false,
	}

	assert.Equal("", service.getResourcePath(t))
	assert.Equal("", service.getResourcePath(ap))
	assert.Equal("", service.getResourcePath(epg))
}

func TestResourceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceServiceTestSuite))
}
