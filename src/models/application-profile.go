package models

const AP_RESOURCE_PREFIX = "ap"
const AP_OBJECT_CLASS = "fvAp"

// Represents an ACI Application Profile.
type AppProfile struct {
	ResourceAttributes
	EPGs []*EPG
}

func (ap *AppProfile) GetObjectClass() string {
	return AP_OBJECT_CLASS
}

func (ap *AppProfile) GetResourcePrefix() string {
	return AP_RESOURCE_PREFIX
}

func (ap *AppProfile) HasParent() bool {
	return true
}

func (ap *AppProfile) ToMap() map[string]string {
	var model = ap.ResourceAttributes.ToMap()
	return model
}

// NewAppProfile will construct a AppProfile from a string map.
func NewAppProfile(model map[string]string) *AppProfile {

	a := AppProfile{NewResourceAttributes(model),
		nil,
	}

	return &a
}

// NewAppProfileMap will construct a string map from reading AppProfile values that can be converted to the type.
func NewAppProfileMap() map[string]string {
	m := NewResourceAttributesMap()
	return m
}

// AddEPG adds a EPG to the AppProfile EPG list and sets the Parent prop of the EPG to the AppProfile it was called from
func (ap *AppProfile) AddEPG(e *EPG) *AppProfile {
	e.SetParent(ap)
	ap.EPGs = append(ap.EPGs, e)

	return ap
}
