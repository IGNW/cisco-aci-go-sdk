// +build examples

package main

import (
	"fmt"
	"os"

	"github.com/ignw/cisco-aci-go-sdk"
)

func main() {
	var client *service.Client
	// Make tenant an array of Tenants, not just one.
	var tenant []*models.Tenant
	var appProfile *models.AppProfile
	var tenantName string = "Example-Tenant"
	var appProName string = "Example-AppProfile"
	var err error

	// Setup a Client to maintain a connection to the APIC
	client = service.GetClient()
	// Get a array of Tenants to put our new Applicationtion Profile in
	tenant, err = client.Tenants.GetByName(tenantName)

	// Make sure we actually got a tenant we can
	if len(tenant) < 1 {
		fmt.Printf("Found no Tenants looking for: %s\n", tenantName)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error finding Tenant: %s\n", err.Error())
	} else {
		fmt.Printf("Found and assigned Tenant: %s\n", tenantName)
	}

	//  Create a new Application Profile
	//  Define the new Application Profile
	appProfile = client.AppProfiles.New(appProName, "This is an example Application Profile")
	// Link the Application Profile to the Tenant we found above
	tenant[0].AddAppProfile(appProfile)
	// Save the new Application Profile to the APIC
	_, err = client.AppProfiles.Save(appProfile)

	if err != nil {
		fmt.Printf("Error creating AppProfile: %s\n", err.Error())
	} else {
		fmt.Printf("Successfully created a new AppProfile: %s\n", appProName)
	}
}
