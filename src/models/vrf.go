package models

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

// Represents an ACI VRF.
type VRF struct {
	ResourceAttributes
	BridgeDomains []*BridgeDomain
}

// New creates a new VRF with the appropriate default values.
func NewVRF(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("ctx-%s", name)

	v := VRF{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvCtx",
		ResourceName: resourceName,
	},
		nil,
	}
	//Do any additional construction logic here.
	return &v
}

func VRFFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}

// AddBridgeDomain adds a BridgeDomain to the VRF BridgeDomain list
func (v *VRF) AddBridgeDomain(bd *BridgeDomain) *VRF {
	v.BridgeDomains = append(v.BridgeDomains, bd)

	return v
}
