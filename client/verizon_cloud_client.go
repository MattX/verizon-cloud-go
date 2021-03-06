// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/client/account_info"
	"github.com/mattx/verizon-cloud-go/client/contacts"
	"github.com/mattx/verizon-cloud-go/client/favorites"
	"github.com/mattx/verizon-cloud-go/client/files_and_folders"
	"github.com/mattx/verizon-cloud-go/client/flashbacks"
	"github.com/mattx/verizon-cloud-go/client/playlists"
	"github.com/mattx/verizon-cloud-go/client/shares"
	"github.com/mattx/verizon-cloud-go/client/tags"
)

// Default verizon cloud HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.cloudapi.verizon.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/cloud/1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new verizon cloud HTTP client.
func NewHTTPClient(formats strfmt.Registry) *VerizonCloud {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new verizon cloud HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *VerizonCloud {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new verizon cloud client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *VerizonCloud {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(VerizonCloud)
	cli.Transport = transport
	cli.AccountInfo = account_info.New(transport, formats)
	cli.Contacts = contacts.New(transport, formats)
	cli.Favorites = favorites.New(transport, formats)
	cli.FilesAndFolders = files_and_folders.New(transport, formats)
	cli.Flashbacks = flashbacks.New(transport, formats)
	cli.Playlists = playlists.New(transport, formats)
	cli.Shares = shares.New(transport, formats)
	cli.Tags = tags.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// VerizonCloud is a client for verizon cloud
type VerizonCloud struct {
	AccountInfo account_info.ClientService

	Contacts contacts.ClientService

	Favorites favorites.ClientService

	FilesAndFolders files_and_folders.ClientService

	Flashbacks flashbacks.ClientService

	Playlists playlists.ClientService

	Shares shares.ClientService

	Tags tags.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *VerizonCloud) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccountInfo.SetTransport(transport)
	c.Contacts.SetTransport(transport)
	c.Favorites.SetTransport(transport)
	c.FilesAndFolders.SetTransport(transport)
	c.Flashbacks.SetTransport(transport)
	c.Playlists.SetTransport(transport)
	c.Shares.SetTransport(transport)
	c.Tags.SetTransport(transport)
}
