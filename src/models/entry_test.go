// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EntryTestSuite struct {
	suite.Suite
}

func (suite *EntryTestSuite) TestEntryToMap() {
	assert := assert.New(suite.T())

	entryMap := map[string]string{
		"dn":          "",
		"status":      "",
		"descr":       "",
		"name":        "TestEntryMap",
		"applyToFrag": "no",
		"arpOpc":      "reply",
		"dFromPort":   "80",
		"dToPort":     "80",
		"etherT":      "ipv4",
		"icmpv4T":     "echo-rep",
		"icmpv6T":     "dst-unreach",
		"matchDscp":   "CS0",
		"prot":        "icmp",
		"sFromPort":   "8080",
		"sToPort":     "8080",
		"stateful":    "no",
		"tcpRules":    "est",
	}

	entry := Entry{
		ResourceAttributes{Name: "TestEntryMap"},
		"icmp",
		"reply",
		false,
		"ipv4",
		"echo-rep",
		"dst-unreach",
		"CS0",
		false,
		"est",
		&ToFrom{To: "8080", From: "8080"},
		&ToFrom{To: "80", From: "80"},
	}

	assert.Equal(entryMap, entry.ToMap())

}

func (suite *EntryTestSuite) TestNewEntryFromDefaults() {
	assert := assert.New(suite.T())

	entry := NewEntry(NewEntryMap())

	expected := Entry{
		ResourceAttributes{Name: ""},
		"unspecified",
		"unspecified",
		false,
		"unspecified",
		"unspecified",
		"unspecified",
		"unspecified",
		false,
		"unspecified",
		&ToFrom{To: "", From: ""},
		&ToFrom{To: "", From: ""},
	}

	assert.Equal(&expected, entry)
}

func (suite *EntryTestSuite) TestNewEntryFromMap() {
	assert := assert.New(suite.T())

	entryMap := map[string]string{
		"dn":          "",
		"status":      "",
		"descr":       "",
		"name":        "TestEntryMap",
		"applyToFrag": "no",
		"arpOpc":      "reply",
		"dFromPort":   "80",
		"dToPort":     "80",
		"etherT":      "ipv4",
		"icmpv4T":     "echo-rep",
		"icmpv6T":     "dst-unreach",
		"matchDscp":   "CS0",
		"prot":        "icmp",
		"sFromPort":   "8080",
		"sToPort":     "8080",
		"stateful":    "no",
		"tcpRules":    "est",
	}

	assert.Equal(&Entry{
		ResourceAttributes{Name: "TestEntryMap"},
		"icmp",
		"reply",
		false,
		"ipv4",
		"echo-rep",
		"dst-unreach",
		"CS0",
		false,
		"est",
		&ToFrom{To: "8080", From: "8080"},
		&ToFrom{To: "80", From: "80"},
	}, NewEntry(entryMap))

}

func TestEntryTestSuite(t *testing.T) {
	suite.Run(t, new(EntryTestSuite))
}
