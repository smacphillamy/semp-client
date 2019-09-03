// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks msg vpn distributed cache cluster global caching home cluster topic prefix links
// swagger:model MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks
type MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks struct {

	// The URI of this Topic Prefix object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn distributed cache cluster global caching home cluster topic prefix links
func (m *MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
