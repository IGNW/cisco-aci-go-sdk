package models

// Represents an ACI Contract Subject.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzSubj.html
type Subject struct {
	ResourceAttributes
	ConsumerMatch      string `oneof=All AtleastOne AtmostOne None`
	ProviderMatch      string `oneof=All AtleastOne AtmostOne None`
	Priority           string `oneof=unspecified level1 level2 level3`
	DSCP               string `oneof=unspecified CS0 CS1 AF11 AF12 AF13 CS2 AF21 AF22 AF23 CS3 AF31 AF32 AF33 CS4 AF41 AF42 AF43 CS5 VA EF CS6 CS7`
	ReverseFilterPorts bool
}

func (s *Subject) ToMap() map[string]string {
	var model = s.ResourceAttributes.ToMap()

	// define additional fields
	model["consMatchT"] = s.ConsumerMatch

	model["prio"] = s.Priority

	model["provMatchT"] = s.ProviderMatch

	model["revFltPorts"] = s.FormatBool(s.ReverseFilterPorts)

	model["targetDscp"] = s.DSCP

	return model
}

// NewSubject will construct a Subject from a string map.
func NewSubject(model map[string]string) *Subject {

	m := Subject{NewResourceAttributes(model),
		"",
		"",
		"",
		"",
		true,
	}

	// The subject match criteria across consumers / "consumer" match
	m.ConsumerMatch = model["consMatchT"]

	// The priority level of a sub application running behind an endpoint group, such as an Exchange server.
	m.Priority = model["prio"]

	// The subject match criteria across consumers / "provider" match
	m.ProviderMatch = model["provMatchT"]

	// Enables the filter to apply on both ingress and egress traffic.
	m.ReverseFilterPorts = m.ParseBool(model["revFltPorts"])

	// The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.
	m.DSCP = model["targetDscp"]

	return &m
}

// NewBridgeDomainMap will construct a string map from reading BridgeDomain values that can be converted to the type.
func NewSubjectMap() map[string]string {

	m := NewResourceAttributesMap()

	m["consMatchT"] = "AtleastOne"

	m["prio"] = "unspecified"

	m["provMatchT"] = "AtleastOne"

	m["revFltPorts"] = "yes"

	m["targetDscp"] = "unspecified"

	return m
}

/*
{
	"vzSubj": {
		"attributes": {
			"childAction": "",
			"configIssues": "",
			"consMatchT": "AtleastOne",
			"descr": "",
			"extMngdBy": "",
			"lcOwn": "local",
			"modTs": "2018-02-23T04:36:26.045+00:00",
			"monPolDn": "uni/tn-common/monepg-default",
			"name": "D2_Diff-IP-Sub_Subject",
			"nameAlias": "",
			"prio": "unspecified",
			"provMatchT": "AtleastOne",
			"revFltPorts": "yes",
			"rn": "subj-D2_Diff-IP-Sub_Subject",
			"status": "",
			"targetDscp": "unspecified",
			"uid": "15374"
		},
*/
