package solace

import (
	apiclient "github.com/ExalDraen/semp-client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Config is the configuration structure used to instantiate the Solace client as well as
// holding per-provider global data.
type Config struct {
	User     string
	Password string
	Host     string
	BasePath string
	MsgVPN   string

	Auth   runtime.ClientAuthInfoWriter
	Client *apiclient.SEMPSolaceElementManagementProtocol
}

// loadAndValidate validates the config provided and fully prepares the Config for use by
// the solace provider. This includes instantiating any API clients.
func (c *Config) loadAndValidate() error {

	// create the API client, enforcing http. go-swagger will always use https if available so
	// need to force http here.
	schemes := []string{"http"}
	auth := httptransport.BasicAuth(c.User, c.Password)
	client := apiclient.New(httptransport.New(c.Host, c.BasePath, schemes), strfmt.Default)
	c.Client = client
	c.Auth = auth

	return nil
}
