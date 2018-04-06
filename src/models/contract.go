package models

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type Contract struct {
	ResourceAttributes
	Subjects []*Subject
	EPGs     []*EPG
}

/* New creates a new Contract with the appropriate default values */
func NewContract(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	c := Contract{ResourceAttributes{
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
	return &c
}

func ContractFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
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
