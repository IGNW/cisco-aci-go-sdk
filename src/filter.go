package cage

import "fmt"

//Filter represents the Filter resource type in ACI
type Filter struct {
	ResourceAttributes
	Subjects []*Subject
	Parent   *Tenant
}

// NewFilter creates a new Filter with the appropriate default values
func NewFilter(name string, alias string, descr string) Filter {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	f := Filter{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "@TODO",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	},
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return f
}

// AddSubject adds a Subject to the Filters Subject list and sets the Parent prop of the Subject to the Filter it was called from
func (f *Filter) AddSubject(s *Subject) *Filter {
	f.Subjects = append(f.Subjects, s)
	return f
}
