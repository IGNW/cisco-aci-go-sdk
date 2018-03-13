package cage

import "fmt"

type Contract struct {
	ResourceAttributes
	Subjects []*Subject
	EPGs     []*EPG
}

/* New creates a new Contract with the appropriate default values */
func (Contract) New(name string, alias string, descr string) Contract {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	t := Contract{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "@TODO",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	}}
	//Do any additional construction logic here.
	return t
}

// AddSubject adds a Subject to the Contract Subject list and sets the Parent prop of the Subject to the Contract it was called from
func (c *Contract) AddSubject(s *Subject) {
	s.Parent = &t
	c.Subjects = append(t.Subjects, s)
}

// AddEPG adds a EPG to the Contract EPG list
func (c *Contract) AddEPG(e *EPG) {
	c.EPGs = append(c.EPGs, e)
}
