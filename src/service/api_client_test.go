// +build integration

package service

import (
	"github.com/Jeffail/gabs"
	log "github.com/golang/glog"
	"testing"
)

func GetDummyData() *gabs.Container {
	var fooBarGabs, _ = gabs.Consume("{foo: bar}")
	return fooBarGabs
}
func TestAuthenticate(t *testing.T) {
	host, name, pass, insecure, err := LookupClientEnvars()
	if err != nil {
		t.Logf("Error with Envvars: %s\n", err)
		t.FailNow()
	}

	c := InitializeClient(host, name, pass, insecure)

	if c == nil {
		t.Logf("AUTHENTICATION FAILED")
		t.Fail()
	} else if c.AuthToken.IsValid() == false {
		t.Logf("\nTOKEN EXPIRES: %v\n", c.AuthToken.Expiry)
		t.Fail()
	} else {
		log.Infof("\nTOKEN is Good: %s", c.AuthToken.Token)
	}
}
