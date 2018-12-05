package main

import (
	apiclient "github.com/ExalDraen/semp-client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Config is per-provider, specifies where to connect to solace
type Config struct {
	User     string
	Password string
	Host     string
	BasePath string
}

// ClientAndAuth bundles all the information needed to connect and act on Solace
type ClientAndAuth struct {
	Auth   runtime.ClientAuthInfoWriter
	Client *apiclient.SEMPSolaceElementManagementProtocol
}

// Client returns a go-swagger API client to interact with the given solace instance
func (c *Config) Client() (interface{}, error) {

	// create the API client
	schemes := []string{"http"}
	auth := httptransport.BasicAuth(c.User, c.Password)
	client := apiclient.New(httptransport.New(c.Host, c.BasePath, schemes), strfmt.Default)

	return ClientAndAuth{Auth: auth, Client: client}, nil
}
