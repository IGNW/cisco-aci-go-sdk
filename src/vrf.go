package cage

import "fmt"

// import "fmt"

type VRF struct {
	ResourceAttributes
	BridgeDomains []*BridgeDomain
}

/* New creates a new Tenant with the appropriate default values */
func NewVRF(name string, alias string, descr string) Tenant {
	resourceName := fmt.Sprintf("ctz-%s", name)

	t := VRF{ResourceAttributes{
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
	return t
}

// AddBridgeDomain adds a BridgeDomain to the VRF BridgeDomain list
func (v *VRF) AddBridgeDomain(bd *BridgeDomain) {
	v.BridgeDomains = append(v.BridgeDomains, bd)
}
