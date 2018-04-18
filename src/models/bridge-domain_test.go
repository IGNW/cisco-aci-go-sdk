// +build unit

package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BridgeDomainTestSuite struct {
	suite.Suite
}

func (suite *BridgeDomainTestSuite) TestBridgeDomainToMap() {
	assert := assert.New(suite.T())

	bdMap := map[string]string{
		"dn":                       "",
		"status":                   "",
		"descr":                    "",
		"name":                     "TestBridgeDomainMap",
		"type":                     "fc",
		"OptimizeWanBandwidth":     "yes",
		"arpFlood":                 "yes",
		"epMoveDetectMode":         "garp",
		"intersiteBumTrafficAllow": "yes",
		"intersiteL2Stretch":       "yes",
		"ipLearning":               "yes",
		"limitIpLearnToSubnets":    "yes",
		"llAddr":                   "10.1.1.51",
		"mac":                      "00:22:BD:F8:19:FF",
		"multiDstPktAct":           "drop",
		"mcastAllow":               "yes",
		"unicastRoute":             "yes",
		"unkMacUcastAct":           "proxy",
		"unkMcastAct":              "flood",
		"vmac":                     "00:22:BD:F8:19:AA",
	}

	bd := BridgeDomain{
		ResourceAttributes{Name: "TestBridgeDomainMap"},
		"fc",
		true,
		true,
		"garp",
		true,
		true,
		true,
		true,
		"10.1.1.51",
		"00:22:BD:F8:19:FF",
		"drop",
		true,
		true,
		"proxy",
		"flood",
		"00:22:BD:F8:19:AA",
		nil,
		nil,
	}

	assert.Equal(bdMap, bd.ToMap())

}

func (suite *BridgeDomainTestSuite) TestNewBridgeDomainFromDefaults() {
	assert := assert.New(suite.T())

	bd := NewBridgeDomain(NewBridgeDomainMap())

	expected := BridgeDomain{
		ResourceAttributes{Name: ""},
		"regular",
		false,
		false,
		"",
		false,
		false,
		true,
		true,
		"",
		"280487012409856",
		"bd-flood",
		false,
		true,
		"proxy",
		"flood",
		"",
		nil,
		nil,
	}

	assert.Equal(&expected, bd)
}

func (suite *BridgeDomainTestSuite) TestNewBridgeDomainFromMap() {
	assert := assert.New(suite.T())

	bdMap := map[string]string{
		"dn":                       "",
		"status":                   "",
		"descr":                    "",
		"name":                     "TestBridgeDomainMap",
		"type":                     "fc",
		"OptimizeWanBandwidth":     "yes",
		"arpFlood":                 "yes",
		"epMoveDetectMode":         "garp",
		"intersiteBumTrafficAllow": "yes",
		"intersiteL2Stretch":       "yes",
		"ipLearning":               "yes",
		"limitIpLearnToSubnets":    "yes",
		"llAddr":                   "10.1.1.51",
		"mac":                      "00:22:BD:F8:19:FF",
		"multiDstPktAct":           "drop",
		"mcastAllow":               "yes",
		"unicastRoute":             "yes",
		"unkMacUcastAct":           "proxy",
		"unkMcastAct":              "flood",
		"vmac":                     "00:22:BD:F8:19:AA",
	}

	assert.Equal(&BridgeDomain{
		ResourceAttributes{Name: "TestBridgeDomainMap"},
		"fc",
		true,
		true,
		"garp",
		true,
		true,
		true,
		true,
		"10.1.1.51",
		"00:22:BD:F8:19:FF",
		"drop",
		true,
		true,
		"proxy",
		"flood",
		"00:22:BD:F8:19:AA",
		nil,
		nil,
	}, NewBridgeDomain(bdMap))

}

func TestBridgeDomainTestSuite(t *testing.T) {
	suite.Run(t, new(BridgeDomainTestSuite))
}
