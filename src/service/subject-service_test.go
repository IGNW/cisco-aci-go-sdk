// +build integration

package service

import (
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SubjectServiceTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *SubjectServiceTestSuite) SetupTest() {

	assert := assert.New(suite.T())

	suite.client = GetClient()

	assert.NotNil(suite.client, "\nCould not get Client, therefore tests could not start")

	ten := suite.client.Tenants.New("IGNW-SUT", "A Subject testing tenant made by IGNW")

	assert.NotNil(ten)

	_, err := suite.client.Tenants.Save(ten)

	assert.Nil(err)

	c := suite.client.Contracts.New("IGNW-C2", "A testing contract made by IGNW")

	ten.AddContract(c)

	_, err = suite.client.Contracts.Save(c)

	assert.Nil(err)

	f := suite.client.Filters.New("IGNW-F2", "A testing filter made by IGNW")

	ten.AddFilter(f)

	_, err = suite.client.Filters.Save(f)

	assert.Nil(err)

	s := suite.client.Subjects.New("IGNW-SU1", "A testing Subject made by IGNW")

	c.AddSubject(s)
	f.AddSubject(s)

	_, err = suite.client.Subjects.Save(s)

	assert.Nil(err)
}

func (suite *SubjectServiceTestSuite) TearDownTest() {
	assert := assert.New(suite.T())

	err := suite.client.Subjects.Delete("uni/tn-IGNW-SUT/brc-IGNW-C2/subj-IGNW-SU1")

	assert.Nil(err)

	err = suite.client.Contracts.Delete("uni/tn-IGNW-SUT/brc-IGNW-C2")

	assert.Nil(err)

	err = suite.client.Filters.Delete("uni/tn-IGNW-SUT/flt-IGNW-F2")

	assert.Nil(err)

	err = suite.client.Tenants.Delete("uni/tn-IGNW-SUT")

	assert.Nil(err)
}

func (suite *SubjectServiceTestSuite) TestSubjectServiceGet() {
	assert := assert.New(suite.T())

	s, err := suite.client.Subjects.Get("uni/tn-IGNW-SUT/brc-IGNW-C2/subj-IGNW-SU1")

	assert.Nil(err)

	if assert.NotNil(s) {

		assert.Equal("IGNW-SU1", s.Name)
		assert.Equal("subj-IGNW-SU1", s.ResourceName)
		assert.Equal("uni/tn-IGNW-SUT/brc-IGNW-C2/subj-IGNW-SU1", s.DomainName)
		assert.Equal("A testing Subject made by IGNW", s.Description)
		assert.Empty(s.Status)

	}

}

func (suite *SubjectServiceTestSuite) TestSubjectServiceGetByName() {
	assert := assert.New(suite.T())

	subjects, err := suite.client.Subjects.GetByName("IGNW-SU1")

	assert.Nil(err)

	if assert.NotEmpty(subjects) {

		assert.Len(subjects, 1)

		assert.Contains(subjects, &models.Subject{
			models.ResourceAttributes{
				Name:         "IGNW-SU1",
				ResourceName: "subj-IGNW-SU1",
				DomainName:   "uni/tn-IGNW-SUT/brc-IGNW-C2/subj-IGNW-SU1",
				Description:  "A testing Subject made by IGNW",
				ObjectClass:  "vzSubj",
				Status:       "",
			},
			"AtleastOne",
			"AtleastOne",
			"unspecified",
			"unspecified",
			true,
		})
	}
}

func (suite *SubjectServiceTestSuite) TestSubjectServiceGetAll() {
	assert := assert.New(suite.T())

	data, err := suite.client.Subjects.GetAll()

	assert.Nil(err)

	if assert.NotEmpty(data) {

		assert.Contains(data, &models.Subject{
			models.ResourceAttributes{
				Name:         "IGNW-SU1",
				ResourceName: "subj-IGNW-SU1",
				DomainName:   "uni/tn-IGNW-SUT/brc-IGNW-C2/subj-IGNW-SU1",
				Description:  "A testing Subject made by IGNW",
				ObjectClass:  "vzSubj",
				Status:       "",
			},
			"AtleastOne",
			"AtleastOne",
			"unspecified",
			"unspecified",
			true,
		})
	}
}

func TestSubjectServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SubjectServiceTestSuite))
}
