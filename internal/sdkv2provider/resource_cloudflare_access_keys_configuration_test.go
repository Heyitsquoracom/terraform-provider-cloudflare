package sdkv2provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessKeysConfiguration_WithKeyRotationIntervalDaysSet(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_keys_configuration.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessKeysConfigurationWithKeyRotationIntervalDays(rnd, accountID, 60),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "key_rotation_interval_days", "60"),
				),
			},
		},
	})
}

func testAccessKeysConfigurationWithKeyRotationIntervalDays(rnd, accountID string, days int) string {
	return fmt.Sprintf(`
resource "cloudflare_access_keys_configuration" "%[1]s" {
  account_id = "%[2]s"
  key_rotation_interval_days = "%[3]d"
}`, rnd, accountID, days)
}

func TestAccCloudflareAccessKeysConfiguration_WithoutKeyRotationIntervalDaysSet(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_access_keys_configuration.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccessKeysConfigurationWithoutKeyRotationIntervalDays(rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "key_rotation_interval_days"),
				),
			},
		},
	})
}

func testAccessKeysConfigurationWithoutKeyRotationIntervalDays(rnd, accountID string) string {
	return fmt.Sprintf(`
resource "cloudflare_access_keys_configuration" "%[1]s" {
  account_id = "%[2]s"
}`, rnd, accountID)
}
