package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs"
	log "github.com/golang/glog"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
)

// Client should be used to interact with the ACI Endpoint.
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
	AuthToken  *models.AuthToken
	username   *string
	password   *string
	Services
}

var clientInstance *Client

// InitializeClient will instantiate a new Client that has been authenticated for working with ACI.
func InitializeClient(apicURL string, user string, pass string, insecure bool) *Client {
	clientURL, err := url.Parse(apicURL)
	if err != nil {
		log.Fatal(err)
	}

	cookieJar, _ := cookiejar.New(nil)
	httpClient := http.DefaultClient
	httpClient.Jar = cookieJar

	clientInstance = &Client{
		BaseURL:    clientURL,
		UserAgent:  "go-lang-main-sdk",
		username:   &user,
		password:   &pass,
		httpClient: httpClient,
		Services: Services{
			AppProfiles:   GetAppProfileService(clientInstance),
			BridgeDomains: GetBridgeDomainService(clientInstance),
			Contracts:     GetContractService(clientInstance),
			VRFs:          GetVRFService(clientInstance),
			EPGs:          GetEPGService(clientInstance),
			Filters:       GetFilterService(clientInstance),
			Entries:       GetEntryService(clientInstance),
			Subjects:      GetSubjectService(clientInstance),
			Subnets:       GetSubnetService(clientInstance),
			Tenants:       GetTenantService(clientInstance),
		},
	}

	if insecure {
		clientInstance.useInsecureHTTPClient()
	}

	err = clientInstance.Authenticate()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return clientInstance
}

// GetClient will get the singleton instance or instantiate a new Client and cache it to a singleton.
func GetClient() *Client {
	if clientInstance == nil {
		host, name, pass, insecure, err := LookupClientEnvars()

		if err != nil {
			log.Fatal("Error with Envvars: %s\n", err)
			return nil
		}

		clientInstance = InitializeClient(host, name, pass, insecure)
	}
	return clientInstance
}

// LookupClientEnvars will configure the Client based on environment variables.
func LookupClientEnvars() (host string, user string, pass string, insecure bool, err error) {

	var exists bool
	var ins string
	var errs *multierror.Error

	if host, exists = os.LookupEnv("APIC_HOST"); !exists || host == "" {
		errs = multierror.Append(errs, fmt.Errorf("Envar 'APIC_HOST' is not set"))
	}

	if user, exists = os.LookupEnv("APIC_USER"); !exists || user == "" {
		errs = multierror.Append(errs, fmt.Errorf("Envar 'APIC_USER' is not set. "))
	}

	if pass, exists = os.LookupEnv("APIC_PASS"); !exists || pass == "" {
		errs = multierror.Append(errs, fmt.Errorf("Envar 'APIC_PASS' is not set. "))
	}

	if ins, exists = os.LookupEnv("APIC_ALLOW_INSECURE"); exists {

		insecure, err = strconv.ParseBool(ins)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("Envar 'APIC_FORCE_HTTPS_INSECURE' is not a bool "))
		}

	} else {
		insecure = false
	}

	err = errs.ErrorOrNil()

	return host, user, pass, insecure, err
}

/** MakeInsecureHTTPClient returns a http.Client for use by the API Client
but with insecure HTTPS params, namely bypassing TLS Verification and
downgrading ciphers. ACI does not support TLS 1 and there seems to be an
issue upgrading to TLS 1.2
*/
func (c *Client) useInsecureHTTPClient() {
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
	cookieJar, _ := cookiejar.New(nil)

	c.UserAgent = "go-lang-cage-sdk-insecure"
	c.httpClient = &http.Client{
		Transport: tr,
		Jar:       cookieJar,
	}

}

// newRequest will create a new HTTP request will all the appropriate headers and will automatically convert your gabs.Container payload to a JSON byte stream.
func (c *Client) newRequest(method string, path string, body *gabs.Container) (*http.Request, error) {

	rel, err := url.Parse(path)

	if err != nil {
		return nil, err
	}

	log.Infof("\nHTTP Body: %s ", body.String())

	u := c.BaseURL.ResolveReference(rel)
	bodyBytes := []byte(body.String())

	buf := bytes.NewBuffer(bodyBytes)

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

// newRequest will create a new HTTP request with authentication headers and and will automatically convert your gabs.Container payload to a JSON byte stream.
func (c Client) newAuthdRequest(method string, path string, body *gabs.Container) (*http.Request, error) {
	if c.AuthToken != nil && !c.AuthToken.IsValid() {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}

	req, err := c.newRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("APIC-Challenge", c.AuthToken.Token)

	return req, nil
}

// do  will execute the supplied request using the Client and return the result as a gabs.Container, the HTTP response and optionally an error.
func (c *Client) do(req *http.Request) (*gabs.Container, *http.Response, error) {

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var result interface{}

	fmt.Printf("\nHTTP Request: %s %s", req.Method, req.URL.String())
	fmt.Printf("\nHTTP Response: %d / %s", resp.StatusCode, resp.Status)

	err = json.NewDecoder(resp.Body).Decode(&result)

	data, err := gabs.Consume(result)
	log.Infof("Client.do-> JSON %#v\n", result)
	log.Infof("Client.do-> DATA %#v\n", data)
	log.Infof("Client.do-> ERROR %#v\n", err)

	return data, resp, err
}

// Authenticate will create a new ACI auth session based on the login name and password provided. It will then parse the token and store it on the Client for future API calls.
func (c *Client) Authenticate() error {
	// @TODO break this into a few smaller private methods
	json, err := gabs.ParseJSON([]byte(`{
		"aaaUser" : {
			"attributes" : {
				"name" : "",
				"pwd" : ""
			}
		}
	}`))

	//fmt.Printf("USER: %#v\n", c.username)
	//fmt.Printf("PASS: %#v\n", c.password)

	json.SetP(c.username, "aaaUser.attributes.name")
	json.SetP(c.password, "aaaUser.attributes.pwd")

	method := "POST"
	path := "/api/mo/aaaLogin.json"
	req, err := c.newRequest(method, path, json)
	if err != nil {
		return err
	}

	data, response, err := c.do(req)

	log.Infof("Client.Authenticate-> RESPONSE %v", response)

	token, err := models.NewAuthToken(data)
	if err != nil {
		return err
	}
	log.Infof("Client.Authenticate-> %s", token.Token)
	c.AuthToken = token

	return nil
}

func (c Client) convertMapToQueryParams(params map[string]string) string {

	queryString := "?"
	paramCount := 0

	for key, value := range params {
		if key != "" && value != "" {
			if paramCount > 0 {
				queryString += "&"
			}

			key := url.QueryEscape(key)
			value := url.QueryEscape(value)

			queryString += fmt.Sprintf("%s=%s", key, value)

		}
	}

	return queryString

}
