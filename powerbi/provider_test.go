package powerbi

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"powerbi": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	requiredEnvs := []string{
		"POWERBI_TENANT_ID",
		"POWERBI_CLIENT_ID",
		"POWERBI_CLIENT_SECRET",
		"POWERBI_USERNAME",
		"POWERBI_PASSWORD",
	}
	for _, requiredEnv := range requiredEnvs {
		if v := os.Getenv(requiredEnv); v == "" {
			t.Fatal(fmt.Sprintf("%s must be set for acceptance tests", requiredEnv))
		}
	}
}