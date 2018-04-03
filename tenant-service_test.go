package cage

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_TenantServiceNew(t *testing.T) {
	c := GetClient()

	if c == nil {
		fmt.Printf("Could not get Client, therefore tests coudl not start")
		t.Fail()
	}

	ten := c.Tenants.New("IGNW", "A Testing tenant made by IGNW")

	if ten == nil {

		fmt.Printf("\nTenant was not created as expected got: %v", ten)
		t.Fail()
	} else {
		fmt.Printf("\nTenant was created successfully status was : %v", ten)
	}

}

func Test_TenantServiceSave(t *testing.T) {

	client := GetClient()

	if client == nil {
		t.Logf("Could not get Client, therefore test could not start")
		t.Fail()
	}

	ten := client.Tenants.New("IGNW", "A Testing tenant made by IGNW")

	err := client.Tenants.Save(ten)

	if err != nil {
		t.Logf("ERROR: Error saving the Tenant: %s", err)
		t.Fail()
	}
}

func Test_TenantServiceDelete(t *testing.T) {

	client := GetClient()

	if client == nil {
		t.Logf("Could not get Client, therefore test could not start")
		t.Fail()
	}

	err := client.Tenants.Delete("uni/tn-IGNW")

	if err != nil {
		t.Logf("ERROR: Error saving the Tenant: %s", err)
		t.Fail()
	}
}

func Test_TenantServiceGet(t *testing.T) {

	client := GetClient()

	if client == nil {
		t.Logf("Could not get Client, therefore test could not start")
		t.Fail()
	}

	_, err := client.Tenants.Get("uni/tn-IGNW")

	if err != nil {
		t.Logf("ERROR: Error saving the Tenant: %s", err)
		t.Fail()
	}
}

func Test_TenantServiceGetAll(t *testing.T) {

	client := GetClient()

	if client == nil {
		t.Logf("Could not get Client, therefore test could not start")
		t.Fail()
	}

	data, err := client.Tenants.GetAll()

	t.Log(fmt.Printf("Got These Tenants: %#v", data))

	for key, tenant := range data {

		t.Log(fmt.Printf("Tenant  #%s has Name %s\n", strconv.Itoa(key), tenant.Name))

	}

	if err != nil {
		t.Logf("ERROR: Error saving the Tenant: %s", err)
		t.Fail()
	}
}
