package cage

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type BridgeDomain struct {
	ResourceAttributes
	Subnets []*Subnet
	EPGs    []*EPG
}

/* New creates a new BridgeDomain with the appropriate default values */
func (BridgeDomain) New(name string, alias string, descr string) BridgeDomain {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	bd := BridgeDomain{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "@TODO",
		DomainName:   fmt.Sprintf("uni/%s", resourceName),
		ResourceName: resourceName,
	}}
	//Do any additional construction logic here.
	return bd
}

func (bd *BridgeDomain) CreateAPIPayload() *gabs.Container {
	return bd.CreateDefaultPayload()
}

// AddSubnet adds a Subnet to the BridgeDomain Subnet list and sets the Parent prop of the Subnet to the BridgeDomain it was called from
func (bd *BridgeDomain) AddSubnet(s *Subnet) {
	s.Parent = &t
	t.Subnets = append(t.Subnets, s)
}

// AddEPG adds a EPG to the BridgeDomain EPG list 
func (bd *BridgeDomain) AddEPG(e *EPG) {
	bd.EPGs = append(bd.EPGs, e)
}
