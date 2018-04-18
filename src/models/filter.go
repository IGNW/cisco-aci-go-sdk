package models

// Represents an ACI Contract Filter.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzFilter.html
type Filter struct {
	ResourceAttributes
	Subjects []*Subject
	Entries  []*Entry
}

// Represents an ACI Contract Filter Entry.
// https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzEntry.html
type Entry struct {
	Protocol           string
	Source, Desination ToFrom
}

type ToFrom struct {
	To, From string
}

func (s *Filter) ToMap() map[string]string {
	var model = s.ResourceAttributes.ToMap()
	return model
}

// NewFilterMap will construct a Filter from a string map.
func NewFilter(model map[string]string) *Filter {

	m := Filter{NewResourceAttributes(model),
		nil,
		nil,
	}

	return &m
}

// NewFilterMap will construct a string map from reading Filter values that can be converted to the type.
func NewFilterMap() map[string]string {

	m := NewResourceAttributesMap()

	return m
}

// AddSubject adds a Subject to the Filters Subject list and sets the Parent prop of the Subject to the Filter it was called from
func (f *Filter) AddSubject(s *Subject) *Filter {
	f.Subjects = append(f.Subjects, s)
	return f
}

// AddSubject adds a Subject to the Filters Subject list and sets the Parent prop of the Subject to the Filter it was called from
func (f *Filter) AddEntry(s *Entry) *Filter {
	f.Entries = append(f.Entries, s)
	return f
}

/*
"vzFilter": {
	"attributes": {
		"childAction": "",
		"descr": "",
		"extMngdBy": "",
		"fwdId": "173",
		"id": "implicit",
		"lcOwn": "local",
		"modTs": "2018-02-23T04:02:42.018+00:00",
		"monPolDn": "uni/tn-common/monepg-default",
		"name": "DB",
		"nameAlias": "",
		"ownerKey": "",
		"ownerTag": "",
		"revId": "174",
		"rn": "flt-DB",
		"status": "",
		"txId": "9223372036854796330",
		"uid": "15374",
		"usesIds": "yes"
	},
	"children": [
		{
			"vzEntry": {
				"attributes": {
					"applyToFrag": "no",
					"arpOpc": "unspecified",
					"childAction": "",
					"dFromPort": "3306",
					"dToPort": "3306",
					"descr": "",
					"etherT": "ip",
					"extMngdBy": "",
					"icmpv4T": "unspecified",
					"icmpv6T": "unspecified",
					"lcOwn": "local",
					"matchDscp": "unspecified",
					"modTs": "2018-02-23T04:02:42.003+00:00",
					"monPolDn": "uni/tn-common/monepg-default",
					"name": "MySQL",
					"nameAlias": "",
					"prot": "tcp",
					"rn": "e-MySQL",
					"sFromPort": "unspecified",
					"sToPort": "unspecified",
					"stateful": "yes",
					"status": "",
					"tcpRules": "",
					"uid": "15374"
				}
			}
		}]
}
*/
