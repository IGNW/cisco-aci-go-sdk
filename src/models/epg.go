package models

const EPG_RESOURCE_NAME_PREFIX = "epg"
const EPG_OBJECT_CLASS = "fvAEPg"

// Represents an ACI Endpoint Group (EPG).
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-fvAEPg.html
type EPG struct {
	ResourceAttributes
	IsAttributeBased       bool
	IsMultiSite            bool
	PreferredPolicyControl string `oneof=enforced unenforced`
	LabelMatchCriteria     string `oneof=All AtleastOne AtmostOne None`
	IsPreferredGroupMember string `oneof=include exclude`
}

func (e *EPG) GetObjectClass() string {
	return EPG_OBJECT_CLASS
}

func (e *EPG) GetResourcePrefix() string {
	return EPG_OBJECT_CLASS
}

func (e *EPG) HasParent() bool {
	return true
}

func (e *EPG) ToMap() map[string]string {
	var model = e.ResourceAttributes.ToMap()

	model["isAttrBasedEPg"] = e.FormatBool(e.IsAttributeBased)

	model["isSharedSrvMsiteEPg"] = e.FormatBool(e.IsMultiSite)

	// The preferred policy control.
	model["pcEnfPref"] = e.PreferredPolicyControl

	// The provider label match criteria.
	model["matchT"] = e.LabelMatchCriteria

	// Represents parameter used to determine if EPg is part of a group that does not a contract for communication
	model["prefGrMemb"] = e.IsPreferredGroupMember

	return model
}

// NewEPG will construct a Contract from a string map.
func NewEPG(model map[string]string) *EPG {

	e := EPG{NewResourceAttributes(model),
		false,
		false,
		"",
		"",
		"",
	}

	e.IsAttributeBased = e.ParseBool(model["isAttrBasedEPg"])

	e.IsMultiSite = e.ParseBool(model["isSharedSrvMsiteEPg"])

	// The preferred policy control.
	e.PreferredPolicyControl = model["pcEnfPref"]

	// The provider label match criteria.
	e.LabelMatchCriteria = model["matchT"]

	// Represents parameter used to determine if EPg is part of a group that does not a contract for communication
	e.IsPreferredGroupMember = model["prefGrMemb"]

	return &e
}

// NewEPGMap will construct a string map from reading EPG values that can be converted to the type.
func NewEPGMap() map[string]string {

	m := NewResourceAttributesMap()

	m["isAttrBasedEPg"] = "no"
	m["isSharedSrvMsiteEPg"] = "no"
	m["pcEnfPref"] = "unenforced"
	m["matchT"] = "AtleastOne"
	m["prefGrMemb"] = "exclude"

	return m
}

//EPG has no objects it owns or relates to according to our model

/*
{
	"fvAEPg": {
		"attributes": {
			"childAction": "",
			"configIssues": "",
			"configSt": "applied",
			"descr": "",
			"extMngdBy": "",
			"fwdCtrl": "",
			"isAttrBasedEPg": "no",
			"isSharedSrvMsiteEPg": "no",
			"lcOwn": "local",
			"matchT": "AtleastOne",
			"modTs": "2018-02-23T04:36:27.126+00:00",
			"monPolDn": "uni/tn-common/monepg-default",
			"name": "server",
			"nameAlias": "",
			"pcEnfPref": "unenforced",
			"pcTag": "32773",
			"prefGrMemb": "exclude",
			"prio": "unspecified",
			"rn": "epg-server",
			"scope": "2850818",
			"status": "",
			"triggerSt": "triggerable",
			"txId": "9223372036854797235",
			"uid": "15374"
		},
*/
