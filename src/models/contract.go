package models

// Represents an ACI Contract.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzBrCP.html
type Contract struct {
	ResourceAttributes
	Scope    string `oneof=application-profile tenant context global`
	DSCP     string `oneof=unspecified CS0 CS1 AF11 AF12 AF13 CS2 AF21 AF22 AF23 CS3 AF31 AF32 AF33 CS4 AF41 AF42 AF43 CS5 VA EF CS6 CS7`
	Subjects []*Subject
	EPGs     []*EPG
}

func (c *Contract) ToMap() map[string]string {
	var model = c.ResourceAttributes.ToMap()

	// Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile
	model["scope"] = c.Scope
	model["targetDscp"] = c.DSCP

	return model
}

// NewContract will construct a Contract from a string map.
func NewContract(model map[string]string) *Contract {

	m := Contract{NewResourceAttributes(model),
		"",
		"",
		nil,
		nil,
	}

	m.Scope = model["scope"]
	m.DSCP = model["targetDscp"]

	return &m
}

// NewContractMap will construct a string map from reading Contract values that can be converted to the type.
func NewContractMap() map[string]string {

	m := NewResourceAttributesMap()

	m["scope"] = "context"
	m["targetDscp"] = "unspecified"

	return m
}

// AddSubject adds a Subject to the Contract Subject list and sets the Parent prop of the Subject to the Contract it was called from
func (c *Contract) AddSubject(s *Subject) *Contract {
	s.SetParent(c)
	c.Subjects = append(c.Subjects, s)

	return c
}

// AddEPG adds a EPG to the Contract EPG list
func (c *Contract) AddEPG(e *EPG) *Contract {
	c.EPGs = append(c.EPGs, e)

	return c
}

/*
{
	"vzBrCP": {
		"attributes": {
			"childAction": "",
			"configIssues": "",
			"descr": "",
			"extMngdBy": "",
			"lcOwn": "local",
			"modTs": "2018-02-23T04:36:26.469+00:00",
			"monPolDn": "uni/tn-common/monepg-default",
			"name": "D2_Diff-IP-Sub",
			"nameAlias": "",
			"ownerKey": "",
			"ownerTag": "",
			"prio": "unspecified",
			"reevaluateAll": "no",
			"rn": "brc-D2_Diff-IP-Sub",
			"scope": "application-profile",
			"status": "",
			"targetDscp": "unspecified",
			"uid": "15374"
		}
*/
