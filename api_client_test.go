package main

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

	if c.AuthToken.IsValid() == false {
		fmt.Printf("TOKEN EXPIRES: %v\n", c.AuthToken.Expiry)
	} else {
		fmt.Printf("TOKEN is Good: %s", c.AuthToken.Token)
	}
}
