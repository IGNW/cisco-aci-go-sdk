package cage

import "fmt"

// import "fmt"

type VRF struct {
	ResourceAttributes
	BridgeDomains []*BridgeDomain
}

/* New creates a new Tenant with the appropriate default values */
func NewVRF(name string, alias string, descr string) VRF {
	resourceName := fmt.Sprintf("ctx-%s", name)

	v := VRF{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvCtx",
		DomainName:   "",
		ResourceName: resourceName,
	},
		nil,
	}
	//Do any additional construction logic here.
	return v
}

// AddBridgeDomain adds a BridgeDomain to the VRF BridgeDomain list
func (v *VRF) AddBridgeDomain(bd *BridgeDomain) *VRF {
	v.BridgeDomains = append(v.BridgeDomains, bd)

	return v
}
