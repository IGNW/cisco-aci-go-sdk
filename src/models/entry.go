package models

const ENTRY_RESOURCE_PREFIX = "e"
const ENTRY_OBJECT_CLASS = "vzEntry"

// Represents an ACI Contract Filter Entry.
// https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzEntry.html
type Entry struct {
	ResourceAttributes
	Protocol            string
	Source, Destination *ToFrom
}

type ToFrom struct {
	To, From string
}

func (e *Entry) GetObjectClass() string {
	return ENTRY_OBJECT_CLASS
}

func (e *Entry) GetResourcePrefix() string {
	return ENTRY_RESOURCE_PREFIX
}

func (e *Entry) HasParent() bool {
	return true
}

func (e *Entry) ToMap() map[string]string {
	var model = e.ResourceAttributes.ToMap()
	return model
}

// NewFilterMap will construct a Filter from a string map.
func NewEntry(model map[string]string) *Entry {

	m := Entry{NewResourceAttributes(model),
		"",
		nil,
		nil,
	}

	return &m
}

// NewEntryMap will construct a string map from reading Filter Entry values that can be converted to the type.
func NewEntryMap() map[string]string {

	m := NewResourceAttributesMap()

	return m
}
