package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

// import "fmt"

type EPG struct {
	ResourceAttributes
	Parent *AppProfile
}

/* New creates a new EPG with the appropriate default values */
func NewEPG(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("epg-%s", name)

	e := EPG{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvAEPg",
		ResourceName: resourceName,
	},
		nil,
	}

	//Do any additional construction logic specific to the EPG here
	return &e
}
func EPGFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}

//EPG has no objects it owns or relates to according to our model
