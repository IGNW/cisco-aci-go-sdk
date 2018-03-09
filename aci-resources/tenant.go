package cage

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type Tenant struct {
	ResourceAttributes
}

/* NewTenant creates a new Tenant with the appropriate default values */
func NewTenant(name string, alias string, descr string) *Tenant {
	resourceName := fmt.Sprintf("tn-%s", name)

	t := Tenant{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvTenant",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	}}
	//Do any additional construction logic here.
	return &t
}

func (t *Tenant) CreateAPIPayload() *gabs.Container {
	return t.CreateDefaultPayload()
}
