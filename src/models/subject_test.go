// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SubjectTestSuite struct {
	suite.Suite
}

func (suite *SubjectTestSuite) TestSubjectToMap() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":          "",
		"status":      "",
		"descr":       "",
		"name":        "TestSubjectMap",
		"consMatchT":  "All",
		"provMatchT":  "All",
		"prio":        "level1",
		"revFltPorts": "yes",
		"targetDscp":  "CS0",
	}

	Subject := Subject{
		ResourceAttributes{Name: "TestSubjectMap"},
		"All",
		"All",
		"level1",
		"CS0",
		true,
	}

	assert.Equal(sMap, Subject.ToMap())

}

func (suite *SubjectTestSuite) TestNewSubjectFromDefaults() {
	assert := assert.New(suite.T())

	subject := NewSubject(NewSubjectMap())

	expected := Subject{
		ResourceAttributes{Name: ""},
		"AtleastOne",
		"AtleastOne",
		"unspecified",
		"unspecified",
		true,
	}

	assert.Equal(&expected, subject)
}

func (suite *SubjectTestSuite) TestNewSubjectFromMap() {
	assert := assert.New(suite.T())

	sMap := map[string]string{
		"dn":          "",
		"status":      "",
		"descr":       "",
		"name":        "TestSubjectMap",
		"consMatchT":  "All",
		"provMatchT":  "All",
		"prio":        "level1",
		"revFltPorts": "yes",
		"targetDscp":  "CS0",
	}

	assert.Equal(&Subject{
		ResourceAttributes{Name: "TestSubjectMap"},
		"All",
		"All",
		"level1",
		"CS0",
		true,
	}, NewSubject(sMap))

}

func TestSubjectTestSuite(t *testing.T) {
	suite.Run(t, new(SubjectTestSuite))
}
