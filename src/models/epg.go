package models

const EPG_RESOURCE_NAME_PREFIX = "epg"
const EPG_OBJECT_CLASS = "fvAEPg"

// Represents an ACI Endpoint Group (EPG).
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-fvAEPg.html
type EPG struct {
	ResourceAttributes
	IsAttributeBased       bool
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
		"",
		"",
		"",
	}

	e.IsAttributeBased = e.ParseBool(model["isAttrBasedEPg"])

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
	m["pcEnfPref"] = "unenforced"
	m["matchT"] = "AtleastOne"
	m["prefGrMemb"] = "exclude"

	return m
}
