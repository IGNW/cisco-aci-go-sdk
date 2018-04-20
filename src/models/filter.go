package models

const FILTER_RESOURCE_PREFIX = "flt"
const FILTER_OBJECT_CLASS = "vzFilter"

// Represents an ACI Contract Filter.
// See: https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzFilter.html
type Filter struct {
	ResourceAttributes
	Subjects []*Subject
	Entries  []*Entry
}

func (f *Filter) GetObjectClass() string {
	return FILTER_OBJECT_CLASS
}

func (f *Filter) GetResourcePrefix() string {
	return FILTER_RESOURCE_PREFIX
}

func (f *Filter) HasParent() bool {
	return true
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

func (f *Filter) ToMap() map[string]string {
	var model = f.ResourceAttributes.ToMap()
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
