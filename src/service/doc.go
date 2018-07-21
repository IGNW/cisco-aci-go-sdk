/* Package service provides and API client that makes working with the Cisco ACI API much easier.

Getting Started

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

SDK Client

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

Making your first tenant

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
*/
package service
