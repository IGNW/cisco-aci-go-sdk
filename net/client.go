package cage

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Jeffail/gabs"
)

// Represents a way to connect to the Cisco ACI API
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
	AuthToken  *Token
}

func (c *Client) newRequest(method, path string, body *gabs.Container) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	buf := bytes.NewBuffer(body.Bytes())

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if buf != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//err = json.NewDecoder(resp.Body).Decode(v)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Printf("%#v", bodyString)
	//err = xml.Unmarshal(string(resp.Body), &v)
	return resp, err
}

// Authenticate makes a login request with the provided name and password
func (c *Client) Authenticate(name string, pwd string) (*http.Response, error) {
	//authRoute := "/aaaLogin"
	//Used to refresh the session cookie
	//refreshRoute := "/aaaRefresh"
	json, err := gabs.ParseJSON([]byte(`{
		"aaaUser" : {
			"attributes" : {
				"name" : "",
				"pwd" : ""
			}
		}
	}`))

	json.SetP(name, "aaaUser.attributes.name")
	json.SetP(pwd, "aaaUser.attributes.pwd")

	fmt.Print(json.String())

	method := "POST"
	path := "/api/mo/aaaLogin.json"
	req, err := c.newRequest(method, path, json)
	if err != nil {
		return nil, err
	}

	var response interface{}
	return c.do(req, response)
}

/** MakeInsecureHTTPClient returns a http.Client for use by the API Client
but with insecure HTTPS params, namely bypassing TLS Verification and
downgrading ciphers. ACI does not support TLS 1 and there seems to be an
issue upgrading to TLS 1.2 so inhere we force the use of TLS 1.1
*/
func (c Client) MakeInsecureHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			},
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       true,
			MinVersion:               tls.VersionTLS11,
			MaxVersion:               tls.VersionTLS11,
		},
	}
	return &http.Client{Transport: tr}

}
