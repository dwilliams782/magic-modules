package google_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleOrganization_byFullName(t *testing.T) {
	orgId := acctest.GetTestOrgFromEnv(t)
	name := "organizations/" + orgId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleOrganization_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_organization.org", "id", name),
					resource.TestCheckResourceAttr("data.google_organization.org", "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleOrganization_byShortName(t *testing.T) {
	orgId := acctest.GetTestOrgFromEnv(t)
	name := "organizations/" + orgId

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleOrganization_byName(orgId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_organization.org", "id", name),
					resource.TestCheckResourceAttr("data.google_organization.org", "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceGoogleOrganization_byDomain(t *testing.T) {
	name := acctest.RandString(t, 16) + ".com"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGoogleOrganization_byDomain(name),
				ExpectError: regexp.MustCompile("Organization not found: " + name),
			},
		},
	})
}

func testAccCheckGoogleOrganization_byName(name string) string {
	return fmt.Sprintf(`
data "google_organization" "org" {
  organization = "%s"
}
`, name)
}

func testAccCheckGoogleOrganization_byDomain(name string) string {
	return fmt.Sprintf(`
data "google_organization" "org" {
  domain = "%s"
}
`, name)
}