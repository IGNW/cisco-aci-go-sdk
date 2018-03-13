package cage

import "fmt"

type Contract struct {
	ResourceAttributes
	Subjects []*Subject
	EPGs     []*EPG
}

/* New creates a new Contract with the appropriate default values */
func NewContract(name string, alias string, descr string) Contract {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	c := Contract{ResourceAttributes{
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
	return c
}

// AddSubject adds a Subject to the Contract Subject list and sets the Parent prop of the Subject to the Contract it was called from
func (c *Contract) AddSubject(s *Subject) {
	s.Parent = c
	c.Subjects = append(c.Subjects, s)
}

// AddEPG adds a EPG to the Contract EPG list
func (c *Contract) AddEPG(e *EPG) {
	c.EPGs = append(c.EPGs, e)
}
