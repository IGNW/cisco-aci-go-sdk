package cage

import (
	"fmt"
	"testing"
)

func TestTenantCreation(t *testing.T) {
	tenant := NewTenant("test", "testing-alias", "short description")
	container := tenant.CreateDefaultPayload()
	fmt.Printf(container.String())
}
