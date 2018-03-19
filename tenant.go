package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type Tenant struct {
	ResourceAttributes
	L3NetIdentifier string
	VRFs            []*VRF
	BridgeDomains   []*BridgeDomain
	AppProfiles     []*AppProfile
	Contracts       []*Contract
	Filters         []*Filter
}

/* New creates a new Tenant with the appropriate default values */
func NewTenant(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("tn-%s", name)

	t := Tenant{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvTenant",
		ResourceName: resourceName,
	},
		"",
		nil,
		nil,
		nil,
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &t
}

func TenantFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}

// AddVRF adds a VRF to the Tenants VRF list and sets the Parent prop of the VRF to the Tenant it was called from
func (t *Tenant) AddVRF(v *VRF) *Tenant {
	v.SetParent(t)
	t.VRFs = append(t.VRFs, v)
	return t
}

// AddBridgeDomain adds a Domain to the Tenants BridgeDomain list and sets the Parent prop of the BridgeDomain to the Tenant it was called from
func (t *Tenant) AddBridgeDomain(bd *BridgeDomain) *Tenant {
	bd.SetParent(t)
	t.BridgeDomains = append(t.BridgeDomains, bd)

	return t
}

// AddAppProfile adds a Domain to the Tenants AppProfile list and sets the Parent prop of the AppProfile to the Tenant it was called from
func (t *Tenant) AddAppProfile(ap *AppProfile) *Tenant {
	ap.SetParent(t)
	t.AppProfiles = append(t.AppProfiles, ap)

	return t
}

// AddContract adds a Domain to the Tenants Contract list and sets the Parent prop of the Contract to the Tenant it was called from
func (t *Tenant) AddContract(c *Contract) *Tenant {
	c.SetParent(t)
	t.Contracts = append(t.Contracts, c)

	return t
}

// AddFilter adds a Domain to the Tenants Filter list and sets the Parent prop of the Filter to the Tenant it was called from
func (t *Tenant) AddFilter(f *Filter) *Tenant {
	f.SetParent(t)
	t.Filters = append(t.Filters, f)

	return t
}
