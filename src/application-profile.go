package cage

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type AppProfile struct {
	ResourceAttributes
	EPGs   []*EPG
	Parent *Tenant
}

/* New creates a new AppProfile with the appropriate default values */
func (AppProfile) New(name string, alias string, descr string) AppProfile {
	resourceName := fmt.Sprintf("ap-%s", name)

	bd := AppProfile{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvApp",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	}}
	//Do any additional construction logic here.
	return bd
}

func (bd *AppProfile) CreateAPIPayload() *gabs.Container {
	return bd.CreateDefaultPayload()
}

// AddEPG adds a EPG to the AppProfile EPG list and sets the Parent prop of the EPG to the AppProfile it was called from
func (bd *AppProfile) AddEPG(e *EPG) {
	e.Parent = &bd
	bd.EPGs = append(bd.EPGs, e)
}
