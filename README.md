# Cisco ACI Go Enabler (main)

A limited functionality SDK derived from the APIC GUI with the primary goal of enabling a Terraform provider for Cisco ACI

## TODOS

- []    Create service providers and defintions for remaining resources.
        These should be able to be done programatically, but I was having some weird python behavior.
        See `scripts/new-resource.py`, maybe you'll have better luck getting it working. If not, copying from the Tenants service will work with some search and replace.

- []    Add required properties to resource objects
        All resource objects currently extend the ResourceAttributes struct which contains the common properties shared between resource types. 
        Unique properties should be added to the resource struct definition (IE tenant.go)

- []    Add properties from previous task to JSON encoder/decoder methods. All resource services have a `fromJSON` method to take response json 
        and turn it into and SDK object. Currently all the "convert to json" functionality is held within `ResourceServices.GetAPIPayload`.
        For semantic reasons this should be aligned with the json verbage IE `toJSON` 

- []    Refactor `fromJSON` method into a ResourceService method. Currently each service has a fully duplicated instance, and there's plenty of 
        opportunity to share.
        Individual services can overide this with their own `fromJSON` but still access it via ResourceService. 
        For example tenants service (ts) would be able to access via `ts.ResourceService.fromJSON` with it's own `ts.fromJSON` implementation

## Building

First you will need to ensure you have installed and configured Go Lang version 1.10+.  You can check with `go version` if you already have a Go Lang environment setup.

Next you will need to the clone the repo into you Go path, typically `~/go/src/github.com`.

```bash
git clone git@github.com:IGNW/cisco-aci-go-sdk.git
cd cisco-aci-go-sdk
```

We've included a `Makefile` to make building and running tests easier.

- make - build sdk
- make unit - run unit tests
- make integration - run integration tests
- make test - run all tests
- make fmtcheck - check source code formatting
- make fmt - fix source code formatting
- make vet - vet source code for problems
- make errcheck - look for un-handled errors

### Testing Environment

In order to run `integration tests` you will need to setup some environement vairables.  You can do this by updating the `GNUmakefile`, exporting
them in your shell environment or by supplying them on the command line.

*Shell Export*
```bash
export APIC_HOST=https://host:port
export APIC_USER=admin
export APIC_PASS=password
export APIC_ALLOW_INSECURE=true

make integration
```

*Command Line*
```bash
APIC_HOST=https://host:port APIC_USER=admin APIC_PASS=password APIC_ALLOW_INSECURE=true make integration
```

## Getting Started

The CAGE SDK provides a way to interact with Cisco ACI without knowing a lot about the underlying API and authentication schemes.  Currently,
the SDK has a limited set of capabilities that cover the following object from the ACI model.
- Tenant
- Application Profile
- Endpoint Group
- Bridge Domain
- Subnet
- VRF
- Contract
- Subject
- Filter

### SDK Client

The SDK Client is used to interact with the ACI endpoint.  As the code illustrates above a call to `service.GetClient()` will
give you a fully instantiated and authenticated `client`.  

The `client` contains helper properties for working with the ACI model collection(s) as follows: 
- Tenant -> `client.Tenants`
- Application Profile -> `client.AppProfiles`
- Endpoint Group -> `client.EPGs`
- Bridge Domain -> `client.BridgeDomains`
- Subnet -> `client.Subnets`
- VRF -> `client.VRFs`
- Contract -> `client.Contracts`
- Subject -> `client.Subjects`
- Filter -> `client.Filters`

Every `client` model `collection` has the following methods:
- New(name, description) - _Create a new model object_
- Save(model) - _Save a model to the ACI endpoint_
- Delete(domainName) - _Delete a model from the ACI endpoint_
- Get(domainName) - _Get a specific model by it's unique domain name_
- GetById(id) - _Get a specific model by it's unique identifier_
- GetByName(name) - _Get model(s) by name_
- GetAll() - _Get all models for the collection_

### Making your first tenant

The following is a quick example on using the SDK to create an ACI Tenant.

```go
package examples

import (
	"fmt"
	"github.com/ignw/cisco-aci-go-sdk/src/service"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

func main() {
	var client *service.Client
	var tenant *models.Tenant
	var err error
	
	client = service.GetClient()
	
	tenant = client.Tenants.New("Example-Tenant", "This is an example tenant")
	
	err = client.Tenants.Save(tenant)
	
	if err != nil {
		fmt.Printf("Error creating tenant: %s", err.Error())
	} else {
		fmt.Println("Successfully created new tenant!")
	}
}
```

             	
             	
             	
             	
             	
             	

 


 
