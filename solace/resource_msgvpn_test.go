package solace

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ExalDraen/semp-client/client/msg_vpn"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const msgVpnResourceSimple = `
resource "solace_msgvpn" "test" {
    name = "go-test"
}
`

func TestMsgVpn_basic(t *testing.T) {
	resourceName := "solace_msgvpn.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckMsgVpnDestroy,
		Steps: []resource.TestStep{
			{
				Config: msgVpnResourceSimple,
				Check: resource.ComposeTestCheckFunc(
					testCheckMsgVpnExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckMsgVpnExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*Config)

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		name := rs.Primary.Attributes["name"]
		params := msg_vpn.NewGetMsgVpnParams()
		params.MsgVpnName = name

		resp, err := c.Client.MsgVpn.GetMsgVpn(params, c.Auth)

		if err != nil {
			return fmt.Errorf("Bad: Get on msg vpn: %+v", err)
		}

		if *resp.Payload.Meta.ResponseCode == http.StatusNotFound {
			return fmt.Errorf("Bad: MSG VPN %q does not exist", name)
		}

		return nil
	}
}

func testCheckMsgVpnDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "solace_msgvpn" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		params := msg_vpn.NewGetMsgVpnParams()
		params.MsgVpnName = name

		resp, err := c.Client.MsgVpn.GetMsgVpn(params, c.Auth)

		// don't fail test on account of our inability to read the resource
		if err != nil {
			return nil
		}

		if *resp.Payload.Meta.ResponseCode != http.StatusNotFound {
			return fmt.Errorf("MSG VPN %q still exists:\n%#v", name, resp.Payload.Data)
		}
	}

	return nil
}
