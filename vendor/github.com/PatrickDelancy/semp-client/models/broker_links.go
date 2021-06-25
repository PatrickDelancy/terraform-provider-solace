// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BrokerLinks broker links
//
// swagger:model BrokerLinks
type BrokerLinks struct {

	// The URI of this Broker's About object.
	AboutURI string `json:"aboutUri,omitempty"`

	// The URI of this Broker's collection of Certificate Authority objects. Deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.
	CertAuthoritiesURI string `json:"certAuthoritiesUri,omitempty"`

	// The URI of this Broker's collection of Client Certificate Authority objects. Available since 2.19.
	ClientCertAuthoritiesURI string `json:"clientCertAuthoritiesUri,omitempty"`

	// The URI of this Broker's collection of Cluster objects. Available since 2.11.
	DmrClustersURI string `json:"dmrClustersUri,omitempty"`

	// The URI of this Broker's collection of Domain Certificate Authority objects. Available since 2.19.
	DomainCertAuthoritiesURI string `json:"domainCertAuthoritiesUri,omitempty"`

	// The URI of this Broker's collection of Message VPN objects. Available since 2.0.
	MsgVpnsURI string `json:"msgVpnsUri,omitempty"`

	// The URI of this Broker's System Information object. Deprecated since 2.2. /systemInformation was replaced by /about/api.
	SystemInformationURI string `json:"systemInformationUri,omitempty"`

	// The URI of this Broker object.
	URI string `json:"uri,omitempty"`

	// The URI of this Broker's collection of Virtual Hostname objects. Available since 2.17.
	VirtualHostnamesURI string `json:"virtualHostnamesUri,omitempty"`
}

// Validate validates this broker links
func (m *BrokerLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broker links based on context it is used
func (m *BrokerLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BrokerLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrokerLinks) UnmarshalBinary(b []byte) error {
	var res BrokerLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}