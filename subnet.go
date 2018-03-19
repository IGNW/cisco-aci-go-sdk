package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

type Subnet struct {
	ResourceAttributes
	Parent *BridgeDomain
}

func NewSubnet(name string, alias string, descr string) ResourceInterface {
	resourceName := fmt.Sprintf("sn-%s", name)

	s := Subnet{ResourceAttributes{
		Name:         name,
		NameAlias:    alias,
		Description:  descr,
		Status:       "created",
		ObjectClass:  "fvSubnet",
		ResourceName: resourceName,
	},
		nil,
	}
	//Do any additional construction logic here.
	return &s
}

func SubnetFromJSON(data *gabs.Container) (ResourceInterface, error) {
	return nil, nil
}
