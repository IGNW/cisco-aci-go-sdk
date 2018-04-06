// +build integration

package service

import (
	"fmt"
	"testing"

	"github.com/Jeffail/gabs"
)

func GetDummyData() *gabs.Container {
	var fooBarGabs, _ = gabs.Consume("{foo: bar}")
	return fooBarGabs
}
func TestAuthenticate(t *testing.T) {
	host, name, pass, insecure, err := LookupClientEnvars()
	if err != nil {
		fmt.Printf("Error with Envvars: %s\n", err)
	}

	c := InitializeClient(host, name, pass, insecure)

	if c == nil {
		t.Logf("AUTHENTICATION FAILED")
		t.Fail()
	} else if c.AuthToken.IsValid() == false {
		t.Logf("\nTOKEN EXPIRES: %v\n", c.AuthToken.Expiry)
		t.Fail()
	} else {
		fmt.Printf("\nTOKEN is Good: %s", c.AuthToken.Token)
	}
}
