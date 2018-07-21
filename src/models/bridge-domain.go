package models

const BD_RESOURCE_PREFIX = "BD"
const BD_OBJECT_CLASS = "fvBD"

// Represents an ACI Bridge Domain.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-fvBD.html
type BridgeDomain struct {
	ResourceAttributes
	Type                     string `oneof=regular fc`
	ArpFlood                 bool
	OptimizeWan              bool
	MoveDetectMode           string `oneof=garp`
	AllowIntersiteBumTraffic bool
	IntersiteL2Stretch       bool
	IpLearning               bool
	LimitIpToSubnets         bool
	LLIpAddress              string
	MAC                      string
	MultiDestForwarding      string `oneof=bd-flood encap-flood drop`
	Multicast                bool
	UnicastRoute             bool
	UnknownUnicastMAC        string `oneof=flood proxy`
	UnknownMulticastMAC      string `oneof=flood opt-flood`
	VirtualMAC               string
	Subnets                  []*Subnet
	EPGs                     []*EPG
}

func (bd *BridgeDomain) GetObjectClass() string {
	return BD_OBJECT_CLASS
}

func (bd *BridgeDomain) GetResourcePrefix() string {
	return BD_RESOURCE_PREFIX
}

func (bd *BridgeDomain) HasParent() bool {
	return true
}

func (bd *BridgeDomain) ToMap() map[string]string {
	var model = bd.ResourceAttributes.ToMap()

	model["type"] = bd.Type
	model["OptimizeWanBandwidth"] = bd.FormatBool(bd.OptimizeWan)
	model["arpFlood"] = bd.FormatBool(bd.ArpFlood)
	model["epMoveDetectMode"] = bd.MoveDetectMode
	model["intersiteBumTrafficAllow"] = bd.FormatBool(bd.IntersiteL2Stretch)
	model["intersiteL2Stretch"] = bd.FormatBool(bd.IntersiteL2Stretch)
	model["ipLearning"] = bd.FormatBool(bd.IpLearning)
	model["limitIpLearnToSubnets"] = bd.FormatBool(bd.LimitIpToSubnets)
	model["llAddr"] = bd.LLIpAddress
	model["mac"] = bd.MAC
	model["multiDstPktAct"] = bd.MultiDestForwarding
	model["mcastAllow"] = bd.FormatBool(bd.Multicast)
	model["unicastRoute"] = bd.FormatBool(bd.UnicastRoute)
	model["unkMacUcastAct"] = bd.UnknownUnicastMAC
	model["unkMcastAct"] = bd.UnknownMulticastMAC
	model["vmac"] = bd.VirtualMAC

	return model
}

// NewBridgeDomain will construct a BridgeDomain from a string map.
func NewBridgeDomain(model map[string]string) *BridgeDomain {

	bd := BridgeDomain{NewResourceAttributes(model),
		"",
		false,
		false,
		"",
		false,
		false,
		false,
		false,
		"",
		"",
		"",
		false,
		false,
		"",
		"",
		"",
		nil,
		nil,
	}

	bd.Type = model["type"]
	bd.OptimizeWan = bd.ParseBool(model["OptimizeWanBandwidth"])
	bd.ArpFlood = bd.ParseBool(model["arpFlood"])
	bd.MoveDetectMode = model["epMoveDetectMode"]
	bd.AllowIntersiteBumTraffic = bd.ParseBool(model["intersiteBumTrafficAllow"])
	bd.IntersiteL2Stretch = bd.ParseBool(model["intersiteL2Stretch"])
	bd.IpLearning = bd.ParseBool(model["ipLearning"])
	bd.LimitIpToSubnets = bd.ParseBool(model["limitIpLearnToSubnets"])
	bd.LLIpAddress = model["llAddr"]
	bd.MAC = model["mac"]
	bd.MultiDestForwarding = model["multiDstPktAct"]
	bd.Multicast = bd.ParseBool(model["mcastAllow"])
	bd.UnicastRoute = bd.ParseBool(model["unicastRoute"])
	bd.UnknownUnicastMAC = model["unkMacUcastAct"]
	bd.UnknownMulticastMAC = model["unkMcastAct"]
	bd.VirtualMAC = model["vmac"]

	return &bd
}

// NewBridgeDomainMap will construct a string map from reading BridgeDomain values that can be converted to the type.
func NewBridgeDomainMap() map[string]string {

	m := NewResourceAttributesMap()

	m["type"] = "regular"
	m["OptimizeWanBandwidth"] = "no"
	m["arpFlood"] = "no"
	m["epMoveDetectMode"] = ""
	m["intersiteBumTrafficAllow"] = "no"
	m["intersiteL2Stretch"] = "no"
	m["ipLearning"] = "yes"
	m["limitIpLearnToSubnets"] = "yes"
	m["llAddr"] = ""
	m["mac"] = "00:22:BD:F8:19:FF"
	m["multiDstPktAct"] = "bd-flood"
	m["mcastAllow"] = "no"
	m["unicastRoute"] = "yes"
	m["unkMacUcastAct"] = "proxy"
	m["unkMcastAct"] = "flood"
	m["vmac"] = ""

	return m
}

// AddSubnet adds a Subnet to the BridgeDomain Subnet list and sets the Parent prop of the Subnet to the BridgeDomain it was called from
func (bd *BridgeDomain) AddSubnet(s *Subnet) *BridgeDomain {
	s.SetParent(bd)
	bd.Subnets = append(bd.Subnets, s)

	return bd
}

// AddEPG adds a EPG to the BridgeDomain EPG list
func (bd *BridgeDomain) AddEPG(e *EPG) *BridgeDomain {
	bd.EPGs = append(bd.EPGs, e)

	return bd
}

/*
{
	"fvBD": {
		"attributes": {
			"OptimizeWanBandwidth": "no",
			"arpFlood": "no",
			"bcastP": "225.0.175.32",
			"childAction": "",
			"configIssues": "",
			"descr": "",
			"epClear": "no",
			"epMoveDetectMode": "",
			"extMngdBy": "",
			"intersiteBumTrafficAllow": "no",
			"intersiteL2Stretch": "no",
			"ipLearning": "yes",
			"lcOwn": "local",
			"limitIpLearnToSubnets": "yes",
			"llAddr": "::",
			"mac": "00:22:BD:F8:19:FF",
			"mcastAllow": "no",
			"modTs": "2018-02-23T04:36:17.595+00:00",
			"monPolDn": "uni/tn-common/monepg-default",
			"mtu": "inherit",
			"multiDstPktAct": "bd-flood",
			"name": "IGNW_BD",
			"nameAlias": "",
			"ownerKey": "",
			"ownerTag": "",
			"pcTag": "32770",
			"rn": "BD-IGNW_BD",
			"scope": "2850818",
			"seg": "16613251",
			"status": "",
			"type": "regular",
			"uid": "15374",
			"unicastRoute": "yes",
			"unkMacUcastAct": "proxy",
			"unkMcastAct": "flood",
			"vmac": "not-applicable"
		},
		"children": [
			{
				"fvSubnet": {
					"attributes": {
						"childAction": "",
						"ctrl": "nd",
						"descr": "",
						"extMngdBy": "",
						"ip": "192.168.103.1/24",
						"lcOwn": "local",
						"modTs": "2018-02-23T04:36:17.500+00:00",
						"monPolDn": "uni/tn-common/monepg-default",
						"name": "VLAN",
						"nameAlias": "",
						"preferred": "no",
						"rn": "subnet-[192.168.103.1/24]",
						"scope": "private",
						"status": "",
						"uid": "15374",
						"virtual": "no"
					}
				}
			},
			{
				"fvRtBd": {
					"attributes": {
						"childAction": "",
						"lcOwn": "local",
						"modTs": "2018-02-23T04:36:28.438+00:00",
						"rn": "rtbd-[uni/tn-IGNW/ap-D3_Multi-IP-Sub/epg-server]",
						"status": "",
						"tCl": "fvAEPg",
						"tDn": "uni/tn-IGNW/ap-D3_Multi-IP-Sub/epg-server"
					}
				}
			},
			{
				"fvRsIgmpsn": {
					"attributes": {
						"childAction": "",
						"forceResolve": "yes",
						"lcOwn": "local",
						"modTs": "2018-02-23T04:36:17.595+00:00",
						"monPolDn": "uni/tn-common/monepg-default",
						"rType": "mo",
						"rn": "rsigmpsn",
						"state": "formed",
						"stateQual": "default-target",
						"status": "",
						"tCl": "igmpSnoopPol",
						"tContextDn": "",
						"tDn": "uni/tn-common/snPol-default",
						"tRn": "snPol-default",
						"tType": "name",
						"tnIgmpSnoopPolName": "",
						"uid": "0"
					}
				}
			},
			{
				"fvRsCtx": {
					"attributes": {
						"childAction": "",
						"extMngdBy": "",
						"forceResolve": "yes",
						"lcOwn": "local",
						"modTs": "2018-02-23T04:36:17.500+00:00",
						"monPolDn": "uni/tn-common/monepg-default",
						"rType": "mo",
						"rn": "rsctx",
						"state": "formed",
						"stateQual": "none",
						"status": "",
						"tCl": "fvCtx",
						"tContextDn": "",
						"tDn": "uni/tn-IGNW/ctx-IGNW_VRF",
						"tRn": "ctx-IGNW_VRF",
						"tType": "name",
						"tnFvCtxName": "IGNW_VRF",
						"uid": "0"
					}
				}
			},
			{
				"fvRsBdToEpRet": {
					"attributes": {
						"childAction": "",
						"forceResolve": "yes",
						"lcOwn": "local",
						"modTs": "2018-02-23T04:36:17.595+00:00",
						"monPolDn": "uni/tn-common/monepg-default",
						"rType": "mo",
						"resolveAct": "resolve",
						"rn": "rsbdToEpRet",
						"state": "formed",
						"stateQual": "default-target",
						"status": "",
						"tCl": "fvEpRetPol",
						"tContextDn": "",
						"tDn": "uni/tn-common/epRPol-default",
						"tRn": "epRPol-default",
						"tType": "name",
						"tnFvEpRetPolName": "",
						"uid": "0"
					}
				}
			},
			{
				"fvRsBDToNdP": {
					"attributes": {
						"childAction": "",
						"forceResolve": "yes",
						"lcOwn": "local",
						"modTs": "2018-02-23T04:36:17.595+00:00",
						"monPolDn": "uni/tn-common/monepg-default",
						"rType": "mo",
						"rn": "rsBDToNdP",
						"state": "formed",
						"stateQual": "default-target",
						"status": "",
						"tCl": "ndIfPol",
						"tContextDn": "",
						"tDn": "uni/tn-common/ndifpol-default",
						"tRn": "ndifpol-default",
						"tType": "name",
						"tnNdIfPolName": "",
						"uid": "0"
					}
				}
			}
		]
	}
}
*/
