package google_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleKmsCryptoKey_basic(t *testing.T) {
	kms := BootstrapKMSKey(t)

	// Name in the KMS client is in the format projects/<project>/locations/<location>/keyRings/<keyRingName>/cryptoKeys/<keyId>
	keyParts := strings.Split(kms.CryptoKey.Name, "/")
	cryptoKeyId := keyParts[len(keyParts)-1]

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleKmsCryptoKey_basic(kms.KeyRing.Name, cryptoKeyId),
				Check:  resource.TestMatchResourceAttr("data.google_kms_crypto_key.kms_crypto_key", "id", regexp.MustCompile(kms.CryptoKey.Name)),
			},
		},
	})
}

func testAccDataSourceGoogleKmsCryptoKey_basic(keyRingName, cryptoKeyName string) string {
	return fmt.Sprintf(`
data "google_kms_crypto_key" "kms_crypto_key" {
  key_ring = "%s"
  name     = "%s"
}
`, keyRingName, cryptoKeyName)
}