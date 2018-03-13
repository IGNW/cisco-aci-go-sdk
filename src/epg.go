package cage

import "fmt"

// import "fmt"

type EPG struct {
	ResourceAttributes
	Parent        *AppProfile
	BridgeDomains []*BridgeDomain
}

/* New creates a new EPG with the appropriate default values */
func NewEPG(name string, alias string, descr string) EPG {
	resourceName := fmt.Sprintf("epg-%s", name)

	e := EPG{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvAEPg",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	},
		nil,
		nil,
	}

	//Do any additional construction logic specific to the EPG here
	return e
}
