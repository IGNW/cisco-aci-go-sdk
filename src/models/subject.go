package models

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type Subject struct {
	ResourceAttributes
}

func NewSubject(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("sbj-%s", name)

	s := Subject{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvSubject",
		ResourceName: resourceName,
	}}

	//Do any additional construction logic here.
	return &s
}

func SubjectFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}
