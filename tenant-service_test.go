package main

import "testing"

func Test_TenantServiceNew(t *testing.T) {
	c := GetClient()

	ten := c.Tenants.New("IGNW", "IGNW Test Tenant", "A Testing tenant made by IGNW")
	if _, ok := ten.(Tenant); !ok {
		t.Fail()
	}

}

func Test_TenantServiceSave(t *testing.T) {

}
