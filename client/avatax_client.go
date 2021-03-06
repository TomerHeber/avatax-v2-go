// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/client/accounts"
	"github.com/TomerHeber/avatax-v2-go/client/addresses"
	"github.com/TomerHeber/avatax-v2-go/client/advanced_rules"
	"github.com/TomerHeber/avatax-v2-go/client/ava_file_forms"
	"github.com/TomerHeber/avatax-v2-go/client/batches"
	"github.com/TomerHeber/avatax-v2-go/client/cert_express_invites"
	"github.com/TomerHeber/avatax-v2-go/client/certificates"
	"github.com/TomerHeber/avatax-v2-go/client/companies"
	"github.com/TomerHeber/avatax-v2-go/client/compliance"
	"github.com/TomerHeber/avatax-v2-go/client/contacts"
	"github.com/TomerHeber/avatax-v2-go/client/customers"
	"github.com/TomerHeber/avatax-v2-go/client/data_sources"
	"github.com/TomerHeber/avatax-v2-go/client/definitions"
	"github.com/TomerHeber/avatax-v2-go/client/distance_thresholds"
	"github.com/TomerHeber/avatax-v2-go/client/e_commerce_token"
	"github.com/TomerHeber/avatax-v2-go/client/filing_calendars"
	"github.com/TomerHeber/avatax-v2-go/client/filings"
	"github.com/TomerHeber/avatax-v2-go/client/firm_client_linkages"
	"github.com/TomerHeber/avatax-v2-go/client/free"
	"github.com/TomerHeber/avatax-v2-go/client/funding_requests"
	"github.com/TomerHeber/avatax-v2-go/client/items"
	"github.com/TomerHeber/avatax-v2-go/client/jurisdiction_overrides"
	"github.com/TomerHeber/avatax-v2-go/client/locations"
	"github.com/TomerHeber/avatax-v2-go/client/multi_document"
	"github.com/TomerHeber/avatax-v2-go/client/nexus"
	"github.com/TomerHeber/avatax-v2-go/client/notices"
	"github.com/TomerHeber/avatax-v2-go/client/notifications"
	"github.com/TomerHeber/avatax-v2-go/client/provisioning"
	"github.com/TomerHeber/avatax-v2-go/client/registrar"
	"github.com/TomerHeber/avatax-v2-go/client/reports"
	"github.com/TomerHeber/avatax-v2-go/client/settings"
	"github.com/TomerHeber/avatax-v2-go/client/subscriptions"
	"github.com/TomerHeber/avatax-v2-go/client/tax_codes"
	"github.com/TomerHeber/avatax-v2-go/client/tax_content"
	"github.com/TomerHeber/avatax-v2-go/client/tax_rules"
	"github.com/TomerHeber/avatax-v2-go/client/transactions"
	"github.com/TomerHeber/avatax-v2-go/client/upcs"
	"github.com/TomerHeber/avatax-v2-go/client/user_defined_fields"
	"github.com/TomerHeber/avatax-v2-go/client/users"
	"github.com/TomerHeber/avatax-v2-go/client/utilities"
)

// Default avatax HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new avatax HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Avatax {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new avatax HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Avatax {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new avatax client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Avatax {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Avatax)
	cli.Transport = transport
	cli.Accounts = accounts.New(transport, formats)
	cli.Addresses = addresses.New(transport, formats)
	cli.AdvancedRules = advanced_rules.New(transport, formats)
	cli.AvaFileForms = ava_file_forms.New(transport, formats)
	cli.Batches = batches.New(transport, formats)
	cli.CertExpressInvites = cert_express_invites.New(transport, formats)
	cli.Certificates = certificates.New(transport, formats)
	cli.Companies = companies.New(transport, formats)
	cli.Compliance = compliance.New(transport, formats)
	cli.Contacts = contacts.New(transport, formats)
	cli.Customers = customers.New(transport, formats)
	cli.DataSources = data_sources.New(transport, formats)
	cli.Definitions = definitions.New(transport, formats)
	cli.DistanceThresholds = distance_thresholds.New(transport, formats)
	cli.ECommerceToken = e_commerce_token.New(transport, formats)
	cli.FilingCalendars = filing_calendars.New(transport, formats)
	cli.Filings = filings.New(transport, formats)
	cli.FirmClientLinkages = firm_client_linkages.New(transport, formats)
	cli.Free = free.New(transport, formats)
	cli.FundingRequests = funding_requests.New(transport, formats)
	cli.Items = items.New(transport, formats)
	cli.JurisdictionOverrides = jurisdiction_overrides.New(transport, formats)
	cli.Locations = locations.New(transport, formats)
	cli.MultiDocument = multi_document.New(transport, formats)
	cli.Nexus = nexus.New(transport, formats)
	cli.Notices = notices.New(transport, formats)
	cli.Notifications = notifications.New(transport, formats)
	cli.Provisioning = provisioning.New(transport, formats)
	cli.Registrar = registrar.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.Settings = settings.New(transport, formats)
	cli.Subscriptions = subscriptions.New(transport, formats)
	cli.TaxCodes = tax_codes.New(transport, formats)
	cli.TaxContent = tax_content.New(transport, formats)
	cli.TaxRules = tax_rules.New(transport, formats)
	cli.Transactions = transactions.New(transport, formats)
	cli.Upcs = upcs.New(transport, formats)
	cli.UserDefinedFields = user_defined_fields.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Utilities = utilities.New(transport, formats)
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

// Avatax is a client for avatax
type Avatax struct {
	Accounts accounts.ClientService

	Addresses addresses.ClientService

	AdvancedRules advanced_rules.ClientService

	AvaFileForms ava_file_forms.ClientService

	Batches batches.ClientService

	CertExpressInvites cert_express_invites.ClientService

	Certificates certificates.ClientService

	Companies companies.ClientService

	Compliance compliance.ClientService

	Contacts contacts.ClientService

	Customers customers.ClientService

	DataSources data_sources.ClientService

	Definitions definitions.ClientService

	DistanceThresholds distance_thresholds.ClientService

	ECommerceToken e_commerce_token.ClientService

	FilingCalendars filing_calendars.ClientService

	Filings filings.ClientService

	FirmClientLinkages firm_client_linkages.ClientService

	Free free.ClientService

	FundingRequests funding_requests.ClientService

	Items items.ClientService

	JurisdictionOverrides jurisdiction_overrides.ClientService

	Locations locations.ClientService

	MultiDocument multi_document.ClientService

	Nexus nexus.ClientService

	Notices notices.ClientService

	Notifications notifications.ClientService

	Provisioning provisioning.ClientService

	Registrar registrar.ClientService

	Reports reports.ClientService

	Settings settings.ClientService

	Subscriptions subscriptions.ClientService

	TaxCodes tax_codes.ClientService

	TaxContent tax_content.ClientService

	TaxRules tax_rules.ClientService

	Transactions transactions.ClientService

	Upcs upcs.ClientService

	UserDefinedFields user_defined_fields.ClientService

	Users users.ClientService

	Utilities utilities.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Avatax) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Accounts.SetTransport(transport)
	c.Addresses.SetTransport(transport)
	c.AdvancedRules.SetTransport(transport)
	c.AvaFileForms.SetTransport(transport)
	c.Batches.SetTransport(transport)
	c.CertExpressInvites.SetTransport(transport)
	c.Certificates.SetTransport(transport)
	c.Companies.SetTransport(transport)
	c.Compliance.SetTransport(transport)
	c.Contacts.SetTransport(transport)
	c.Customers.SetTransport(transport)
	c.DataSources.SetTransport(transport)
	c.Definitions.SetTransport(transport)
	c.DistanceThresholds.SetTransport(transport)
	c.ECommerceToken.SetTransport(transport)
	c.FilingCalendars.SetTransport(transport)
	c.Filings.SetTransport(transport)
	c.FirmClientLinkages.SetTransport(transport)
	c.Free.SetTransport(transport)
	c.FundingRequests.SetTransport(transport)
	c.Items.SetTransport(transport)
	c.JurisdictionOverrides.SetTransport(transport)
	c.Locations.SetTransport(transport)
	c.MultiDocument.SetTransport(transport)
	c.Nexus.SetTransport(transport)
	c.Notices.SetTransport(transport)
	c.Notifications.SetTransport(transport)
	c.Provisioning.SetTransport(transport)
	c.Registrar.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.Settings.SetTransport(transport)
	c.Subscriptions.SetTransport(transport)
	c.TaxCodes.SetTransport(transport)
	c.TaxContent.SetTransport(transport)
	c.TaxRules.SetTransport(transport)
	c.Transactions.SetTransport(transport)
	c.Upcs.SetTransport(transport)
	c.UserDefinedFields.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Utilities.SetTransport(transport)
}
