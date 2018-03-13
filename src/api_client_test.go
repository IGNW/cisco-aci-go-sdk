package cage

import (
	"fmt"
	"net/url"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	c := &Client{
		BaseURL: &url.URL{
			Scheme: "https",
			//Host:   "73.254.132.17:8480",
			Host: "73.254.132.17:8443",
		},
		UserAgent: "go-test",
	}

	c.httpClient = c.MakeInsecureHTTPClient()

	err := c.Authenticate("admin", "password")

	fmt.Printf("AUTH: %#v\n", c.AuthToken)

	if c.AuthToken.IsValid() == true {
		fmt.Printf("TOKEN EXPIRED: %v\n", c.AuthToken.Expiry)
		fmt.Printf("ERROR: %v\n", err)
		t.Fail()
	}

}
