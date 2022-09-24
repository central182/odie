//go:build integration

// The odapi_resty package can't be unit-tested: it depends on something that is not in our control.
// The purpose of this file is merely to drive development: we'll automate the setup steps
// and do manual inspections on the results.
package odapi_resty_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/central182/odie/internal/adapter/outbound/common/odapi"
	odapi_resty "github.com/central182/odie/internal/adapter/outbound/common/odapi/resty"
)

func TestClient(t *testing.T) {
	c := odapi_resty.New(os.Getenv("APP_ID"), os.Getenv("APP_KEY"))
	resp, err := c.GetEntries(odapi.GetEntriesRequest{
		SourceLang: "en-gb",
		WordId:     "documentation",
	})
	fmt.Printf("err = %+v\n", err)
	fmt.Printf("resp = %+v\n", resp)
}
