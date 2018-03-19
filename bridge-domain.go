package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type BridgeDomain struct {
	ResourceAttributes
	Subnets []*Subnet
	EPGs    []*EPG
}

/* NewBridgeDomain creates a new BridgeDomain with the appropriate default values */
func NewBridgeDomain(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("@TODO-%s", name)

	bd := BridgeDomain{ResourceAttributes{
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
	return &bd
}

func BridgeDomainFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}

// AddSubnet adds a Subnet to the BridgeDomain Subnet list and sets the Parent prop of the Subnet to the BridgeDomain it was called from
func (bd *BridgeDomain) AddSubnet(s *Subnet) *BridgeDomain {
	s.SetParent(bd)
	bd.Subnets = append(bd.Subnets, s)

	return bd
}

// AddEPG adds a EPG to the BridgeDomain EPG list
func (bd *BridgeDomain) AddEPG(e *EPG) *BridgeDomain {
	bd.EPGs = append(bd.EPGs, e)

	return bd
}
