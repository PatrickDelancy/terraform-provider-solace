package solace

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var reqEnvVars = []string{
	"SOLACE_USER",
	"SOLACE_PASSWORD",
	"SOLACE_HOST",
}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"solace": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	for _, e := range reqEnvVars {
		if v := os.Getenv(e); v == "" {
			t.Fatalf("%v must be set for acceptance tests", v)
		}
	}
}
