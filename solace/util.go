package solace

import (
	"fmt"

	"github.com/PatrickDelancy/semp-client/models"
)

// TerraformResourceData represents a subset of the function from schema.ResourceData
// Approach copied from terraform-google-provider (utils.go)
// cf. https://godoc.org/github.com/hashicorp/terraform/helper/schema#ResourceData
type TerraformResourceData interface {
	HasChange(string) bool
	GetOk(string) (interface{}, bool)
	Set(string, interface{}) error
	SetId(string)
	Id() string
}

// formatError returns a useful string representation of the SempError provided
func formatError(err *models.SempError) string {
	return fmt.Sprintf("[%v] - %v: %v", *err.Code, *err.Description, *err.Status)
}

// getMsgVPN reads the "msg_vpn" field from the given resource data and falls
// back to the provider's value if not given. If the provider's value is not
// given, an error is returned.
func getMsgVPN(d TerraformResourceData, config *Config) (string, error) {
	return getMsgVPNFromSchema("msg_vpn", d, config)
}

func getMsgVPNFromSchema(msgVpnSchemaField string, d TerraformResourceData, config *Config) (string, error) {
	res, ok := d.GetOk(msgVpnSchemaField)
	if ok && msgVpnSchemaField != "" {
		return res.(string), nil
	}
	if config.MsgVPN != "" {
		return config.MsgVPN, nil
	}
	return "", fmt.Errorf("%s: required field is not set", msgVpnSchemaField)
}
