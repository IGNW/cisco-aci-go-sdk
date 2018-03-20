package main

import (
	"fmt"
	"testing"
)

func Test_TenantServiceNew(t *testing.T) {
	c := GetClient()

	if c == nil {
		fmt.Printf("Could not get Client, therefore tests coudl not start")
		t.Fail()
	}

	ten := c.Tenants.New("IGNW", "IGNW Test Tenant", "A Testing tenant made by IGNW")
	if _, ok := ten.(ResourceInterface); !ok {

		fmt.Printf("\nTenant was not created status was : %v", ok)
		fmt.Printf("\nTenant was not created as expected got: %v", ten)
		t.Fail()
	}

}

func Test_TenantServiceSave(t *testing.T) {

	client := GetClient()

	if client == nil {
		fmt.Printf("Could not get Client, therefore tests could not start")
		t.Fail()
	}

	ten := client.Tenants.New("IGNW", "IGNW Test Tenant", "A Testing tenant made by IGNW")

	data, response, err := client.Tenants.Save(ten)

	fmt.Printf("DATA: %v\n", data)
	fmt.Printf("RESP %v\n", response)

	if err != nil {
		t.Fail()
	}

}
