package cage

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Jeffail/gabs"
)

// Represents a way to connect to the Cisco ACI API
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
	AuthToken  AuthToken
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

	err = json.NewDecoder(resp.Body).Decode(&v)

	return resp, err
}

// Authenticate makes a login request with the provided name and password
func (c *Client) Authenticate(name string, pwd string) error {
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

	method := "POST"
	path := "/api/mo/aaaLogin.json"
	req, err := c.newRequest(method, path, json)
	if err != nil {
		return err
	}

	var response interface{}

	c.do(req, &response)

	data, err := gabs.Consume(response)

	var valuePath, token string
	var createdAt, expiresIn int64
	var expiry time.Time

	if valuePath = "imdata.aaaLogin.attributes.token"; data.ExistsP(valuePath) {
		token = data.Path(valuePath).Data().([]interface{})[0].(string)
		fmt.Printf("TOKEN: %v\n", token)
	} else {
		return fmt.Errorf("Token was not found in response,\n was expected at %v", valuePath)
	}

	if valuePath = "imdata.aaaLogin.attributes.creationTime"; data.ExistsP(valuePath) {
		createdAtStr := data.Path(valuePath).Data().([]interface{})[0].(string)
		createdAt, err = strconv.ParseInt(createdAtStr, 10, 64)
		if err != nil {
			return fmt.Errorf("Could not parse creationTime got error: %v", err)
		}
	} else {
		return fmt.Errorf("creationTime was not found in response,\n was expected at %v", valuePath)
	}

	if valuePath = "imdata.aaaLogin.attributes.refreshTimeoutSeconds"; data.ExistsP(valuePath) {
		expiresInStr := data.Path(valuePath).Data().([]interface{})[0].(string)
		fmt.Printf("%v", expiresIn)
		expiresIn, err = strconv.ParseInt(expiresInStr, 10, 64)
		if err != nil {
			return fmt.Errorf("Could not parse refreshTimeOutSeconds got error: %v", err)
		}
	} else {
		return fmt.Errorf("refreshTimeOutSeconds was not found in response,\n was expected at %v", valuePath)
	}
	fmt.Printf("CREATED AT: %v\n", createdAt)
	fmt.Printf("LIVE TIL: %#v\n", expiresIn)
	fmt.Printf("WILL EXP:  %v\n", (createdAt + expiresIn))

	expiry = time.Unix((createdAt + expiresIn), 0)

	c.AuthToken = AuthToken{
		Token:  token,
		Expiry: expiry,
	}
	return err
}

/** MakeInsecureHTTPClient returns a http.Client for use by the API Client
but with insecure HTTPS params, namely bypassing TLS Verification and
downgrading ciphers. ACI does not support TLS 1 and there seems to be an
issue upgrading to TLS 1.2
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
