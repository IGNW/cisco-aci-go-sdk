package models

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

//Filter represents the Filter resource type in ACI
type Filter struct {
	ResourceAttributes
	Subjects []*Subject
	Entries  []*Entry
}

type Entry struct {
	Protocol           string
	Source, Desination ToFrom
}

type ToFrom struct {
	To, From string
}

// NewFilter creates a new Filter with the appropriate default values
func NewFilter(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	f := Filter{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "@TODO",
		ResourceName: resourceName,
	},
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &f
}
func FilterFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
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
