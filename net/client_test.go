package cage

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/Jeffail/gabs"
)

func TestAuthenticate(t *testing.T) {

	c := &Client{
		BaseURL: &url.URL{
			Scheme: "https",
			//Host:   "73.254.132.17:8480",
			Host: "73.254.132.17:8443",
		},
		UserAgent: "",
	}
	c.httpClient = c.MakeInsecureHTTPClient()

	response, err := c.Authenticate("admin", "password")

	fmt.Printf("%v\n", err)
	fmt.Printf("%v", response)

	if err != nil {
		t.Fail()
	}
}

func TestGabsBehavior(t *testing.T) {
	data, _ := gabs.ParseJSON([]byte(`{
		fvTenant: {
			"foo" : "bar",
			
		}	
	}`))

	data2, _ := gabs.ParseJSON([]byte(`{
		"one" : {
			"two" : "2",
			"three": "3"
	}`))

	data.Consume(data2)

	fmt.Printf(data)
}
