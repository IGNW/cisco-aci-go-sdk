package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Jeffail/gabs"
)

type AuthToken struct {
	Token         string
	Expiry        time.Time
	apicCreatedAt time.Time
	realCreatedAt time.Time
	offset        int64
}

func NewAuthToken(data *gabs.Container) (*AuthToken, error) {
	var valuePath, token string
	var createdAt, expiresIn int64
	var err error

	if valuePath = "imdata.aaaLogin.attributes.token"; data.ExistsP(valuePath) {
		token = data.Path(valuePath).Data().([]interface{})[0].(string)
		//fmt.Printf("TOKEN: %v\n", token)
	} else {
		return nil, fmt.Errorf("Token was not found in response,\n was expected at %v", valuePath)
	}

	if valuePath = "imdata.aaaLogin.attributes.creationTime"; data.ExistsP(valuePath) {
		createdAtStr := data.Path(valuePath).Data().([]interface{})[0].(string)
		createdAt, err = strconv.ParseInt(createdAtStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not parse creationTime got error: %v", err)
		}
	} else {
		return nil, fmt.Errorf("creationTime was not found in response,\n was expected at %v", valuePath)
	}

	if valuePath = "imdata.aaaLogin.attributes.refreshTimeoutSeconds"; data.ExistsP(valuePath) {
		expiresInStr := data.Path(valuePath).Data().([]interface{})[0].(string)
		fmt.Printf("%v", expiresIn)
		expiresIn, err = strconv.ParseInt(expiresInStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not parse refreshTimeOutSeconds got error: %v", err)
		}
	} else {
		return nil, fmt.Errorf("refreshTimeOutSeconds was not found in response,\n was expected at %v", valuePath)
	}
	//fmt.Printf("CREATED AT: %v\n", createdAt)
	//fmt.Printf("LIVE TIL: %#v\n", expiresIn)
	//fmt.Printf("WILL EXP:  %v\n", (createdAt + expiresIn))

	at := AuthToken{
		Token:         token,
		apicCreatedAt: time.Unix(createdAt, 0),
		realCreatedAt: time.Now(),
	}

	at.calculateExpiry(expiresIn)
	at.caclulateOffset()

	return &at, nil

}

/** IsNotExpired evalutes the tokens expiry against the current time after checking if the value is set
and returns a bool true if the token has not reach it's expiry
*/
func (t *AuthToken) IsValid() bool {
	//fmt.Printf("IS NOW: %v\n", time.Now().Unix())
	//fmt.Printf("EXPIRES:  %v\n", t.Expiry.Unix())

	if t.IsSet() == true && t.Expiry.Unix() > t.estimateAPICTime() {
		return true
	}
	return false
}

func (t *AuthToken) IsSet() bool {
	if t.Token != "" {
		return true
	}
	return false
}
func (t *AuthToken) caclulateOffset() {
	t.offset = t.apicCreatedAt.Unix() - t.realCreatedAt.Unix()
}

func (t *AuthToken) estimateAPICTime() int64 {
	return time.Now().Unix() + t.offset
}
func (t *AuthToken) calculateExpiry(willExpire int64) {
	t.Expiry = time.Unix((t.apicCreatedAt.Unix() + willExpire), 0)
}
