package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type AppProfile struct {
	ResourceAttributes
	EPGs   []*EPG
	Parent *Tenant
}

/* NewAppProfile creates a new AppProfile with the appropriate default values */
func NewAppProfile(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("ap-%s", name)

	ap := AppProfile{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvApp",
		ResourceName: resourceName,
	},
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &ap
}

func AppProfileFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}

// AddEPG adds a EPG to the AppProfile EPG list and sets the Parent prop of the EPG to the AppProfile it was called from
func (ap *AppProfile) AddEPG(e *EPG) *AppProfile {
	e.SetParent(ap)
	ap.EPGs = append(ap.EPGs, e)

	return ap
}
