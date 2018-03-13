package cage

import (
	"fmt"
	"time"
)

type AuthToken struct {
	Token     string
	CreatedAt time.Time
	Expiry    time.Time
}

/** IsNotExpired evalutes the tokens expiry against the current time
and returns a bool true f
*/
func (t AuthToken) IsValid() bool {
	fmt.Printf("IS NOW: %v\n", time.Now().Unix())
	fmt.Printf("EXPIRES:  %v\n", t.Expiry.Unix())

	if t.Expiry.Unix() > time.Now().Unix() {
		return true
	}
	return false
}
