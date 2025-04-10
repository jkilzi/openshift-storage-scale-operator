// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// DNSTransportConfigApplyConfiguration represents a declarative configuration of the DNSTransportConfig type for use
// with apply.
type DNSTransportConfigApplyConfiguration struct {
	Transport *operatorv1.DNSTransport            `json:"transport,omitempty"`
	TLS       *DNSOverTLSConfigApplyConfiguration `json:"tls,omitempty"`
}

// DNSTransportConfigApplyConfiguration constructs a declarative configuration of the DNSTransportConfig type for use with
// apply.
func DNSTransportConfig() *DNSTransportConfigApplyConfiguration {
	return &DNSTransportConfigApplyConfiguration{}
}

// WithTransport sets the Transport field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Transport field is set to the value of the last call.
func (b *DNSTransportConfigApplyConfiguration) WithTransport(value operatorv1.DNSTransport) *DNSTransportConfigApplyConfiguration {
	b.Transport = &value
	return b
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *DNSTransportConfigApplyConfiguration) WithTLS(value *DNSOverTLSConfigApplyConfiguration) *DNSTransportConfigApplyConfiguration {
	b.TLS = value
	return b
}
