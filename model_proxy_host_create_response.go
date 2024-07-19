/*
Nginx Proxy Manager API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.x.x
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nginxproxymanager

import (
	"encoding/json"
	"time"
)

// checks if the ProxyHostCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyHostCreateResponse{}

// ProxyHostCreateResponse struct for ProxyHostCreateResponse
type ProxyHostCreateResponse struct {
	// Unique identifier
	Id *int64 `json:"id,omitempty"`
	// Date and time of creation
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// Date and time of last update
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	// Domain Names separated by a comma
	DomainNames []string `json:"domain_names,omitempty"`
	ForwardScheme *string `json:"forward_scheme,omitempty"`
	ForwardHost *string `json:"forward_host,omitempty"`
	ForwardPort *int64 `json:"forward_port,omitempty"`
	CertificateId *PutProxyHostUpdateRequestCertificateId `json:"certificate_id,omitempty"`
	// Is SSL Forced
	SslForced *boolAsInt `json:"ssl_forced,omitempty"`
	// Is HSTS Enabled
	HstsEnabled *boolAsInt `json:"hsts_enabled,omitempty"`
	// Is HSTS applicable to all subdomains
	HstsSubdomains *boolAsInt `json:"hsts_subdomains,omitempty"`
	// HTTP2 Protocol Support
	Http2Support *boolAsInt `json:"http2_support,omitempty"`
	// Should we block common exploits
	BlockExploits *boolAsInt `json:"block_exploits,omitempty"`
	// Should we cache assets
	CachingEnabled *boolAsInt `json:"caching_enabled,omitempty"`
	// Allow Websocket Upgrade for all paths
	AllowWebsocketUpgrade *boolAsInt `json:"allow_websocket_upgrade,omitempty"`
	// Access List ID
	AccessListId *int64 `json:"access_list_id,omitempty"`
	AdvancedConfig *string `json:"advanced_config,omitempty"`
	// Is Enabled
	Enabled *boolAsInt `json:"enabled,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Locations []PutProxyHostUpdateRequestLocationsInner `json:"locations,omitempty"`
}

// NewProxyHostCreateResponse instantiates a new ProxyHostCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyHostCreateResponse() *ProxyHostCreateResponse {
	this := ProxyHostCreateResponse{}
	return &this
}

// NewProxyHostCreateResponseWithDefaults instantiates a new ProxyHostCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyHostCreateResponseWithDefaults() *ProxyHostCreateResponse {
	this := ProxyHostCreateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ProxyHostCreateResponse) SetId(v int64) {
	o.Id = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *ProxyHostCreateResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetModifiedOn() time.Time {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasModifiedOn() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *ProxyHostCreateResponse) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *ProxyHostCreateResponse) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetForwardScheme returns the ForwardScheme field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetForwardScheme() string {
	if o == nil || IsNil(o.ForwardScheme) {
		var ret string
		return ret
	}
	return *o.ForwardScheme
}

// GetForwardSchemeOk returns a tuple with the ForwardScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetForwardSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardScheme) {
		return nil, false
	}
	return o.ForwardScheme, true
}

// HasForwardScheme returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasForwardScheme() bool {
	if o != nil && !IsNil(o.ForwardScheme) {
		return true
	}

	return false
}

// SetForwardScheme gets a reference to the given string and assigns it to the ForwardScheme field.
func (o *ProxyHostCreateResponse) SetForwardScheme(v string) {
	o.ForwardScheme = &v
}

// GetForwardHost returns the ForwardHost field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetForwardHost() string {
	if o == nil || IsNil(o.ForwardHost) {
		var ret string
		return ret
	}
	return *o.ForwardHost
}

// GetForwardHostOk returns a tuple with the ForwardHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetForwardHostOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardHost) {
		return nil, false
	}
	return o.ForwardHost, true
}

// HasForwardHost returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasForwardHost() bool {
	if o != nil && !IsNil(o.ForwardHost) {
		return true
	}

	return false
}

// SetForwardHost gets a reference to the given string and assigns it to the ForwardHost field.
func (o *ProxyHostCreateResponse) SetForwardHost(v string) {
	o.ForwardHost = &v
}

// GetForwardPort returns the ForwardPort field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetForwardPort() int64 {
	if o == nil || IsNil(o.ForwardPort) {
		var ret int64
		return ret
	}
	return *o.ForwardPort
}

// GetForwardPortOk returns a tuple with the ForwardPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetForwardPortOk() (*int64, bool) {
	if o == nil || IsNil(o.ForwardPort) {
		return nil, false
	}
	return o.ForwardPort, true
}

// HasForwardPort returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasForwardPort() bool {
	if o != nil && !IsNil(o.ForwardPort) {
		return true
	}

	return false
}

// SetForwardPort gets a reference to the given int64 and assigns it to the ForwardPort field.
func (o *ProxyHostCreateResponse) SetForwardPort(v int64) {
	o.ForwardPort = &v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetCertificateId() PutProxyHostUpdateRequestCertificateId {
	if o == nil || IsNil(o.CertificateId) {
		var ret PutProxyHostUpdateRequestCertificateId
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool) {
	if o == nil || IsNil(o.CertificateId) {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasCertificateId() bool {
	if o != nil && !IsNil(o.CertificateId) {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given PutProxyHostUpdateRequestCertificateId and assigns it to the CertificateId field.
func (o *ProxyHostCreateResponse) SetCertificateId(v PutProxyHostUpdateRequestCertificateId) {
	o.CertificateId = &v
}

// GetSslForced returns the SslForced field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetSslForced() boolAsInt {
	if o == nil || IsNil(o.SslForced) {
		var ret boolAsInt
		return ret
	}
	return *o.SslForced
}

// GetSslForcedOk returns a tuple with the SslForced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetSslForcedOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.SslForced) {
		return nil, false
	}
	return o.SslForced, true
}

// HasSslForced returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasSslForced() bool {
	if o != nil && !IsNil(o.SslForced) {
		return true
	}

	return false
}

// SetSslForced gets a reference to the given boolAsInt and assigns it to the SslForced field.
func (o *ProxyHostCreateResponse) SetSslForced(v boolAsInt) {
	o.SslForced = &v
}

// GetHstsEnabled returns the HstsEnabled field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetHstsEnabled() boolAsInt {
	if o == nil || IsNil(o.HstsEnabled) {
		var ret boolAsInt
		return ret
	}
	return *o.HstsEnabled
}

// GetHstsEnabledOk returns a tuple with the HstsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetHstsEnabledOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.HstsEnabled) {
		return nil, false
	}
	return o.HstsEnabled, true
}

// HasHstsEnabled returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasHstsEnabled() bool {
	if o != nil && !IsNil(o.HstsEnabled) {
		return true
	}

	return false
}

// SetHstsEnabled gets a reference to the given boolAsInt and assigns it to the HstsEnabled field.
func (o *ProxyHostCreateResponse) SetHstsEnabled(v boolAsInt) {
	o.HstsEnabled = &v
}

// GetHstsSubdomains returns the HstsSubdomains field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetHstsSubdomains() boolAsInt {
	if o == nil || IsNil(o.HstsSubdomains) {
		var ret boolAsInt
		return ret
	}
	return *o.HstsSubdomains
}

// GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetHstsSubdomainsOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.HstsSubdomains) {
		return nil, false
	}
	return o.HstsSubdomains, true
}

// HasHstsSubdomains returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasHstsSubdomains() bool {
	if o != nil && !IsNil(o.HstsSubdomains) {
		return true
	}

	return false
}

// SetHstsSubdomains gets a reference to the given boolAsInt and assigns it to the HstsSubdomains field.
func (o *ProxyHostCreateResponse) SetHstsSubdomains(v boolAsInt) {
	o.HstsSubdomains = &v
}

// GetHttp2Support returns the Http2Support field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetHttp2Support() boolAsInt {
	if o == nil || IsNil(o.Http2Support) {
		var ret boolAsInt
		return ret
	}
	return *o.Http2Support
}

// GetHttp2SupportOk returns a tuple with the Http2Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetHttp2SupportOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.Http2Support) {
		return nil, false
	}
	return o.Http2Support, true
}

// HasHttp2Support returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasHttp2Support() bool {
	if o != nil && !IsNil(o.Http2Support) {
		return true
	}

	return false
}

// SetHttp2Support gets a reference to the given boolAsInt and assigns it to the Http2Support field.
func (o *ProxyHostCreateResponse) SetHttp2Support(v boolAsInt) {
	o.Http2Support = &v
}

// GetBlockExploits returns the BlockExploits field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetBlockExploits() boolAsInt {
	if o == nil || IsNil(o.BlockExploits) {
		var ret boolAsInt
		return ret
	}
	return *o.BlockExploits
}

// GetBlockExploitsOk returns a tuple with the BlockExploits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetBlockExploitsOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.BlockExploits) {
		return nil, false
	}
	return o.BlockExploits, true
}

// HasBlockExploits returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasBlockExploits() bool {
	if o != nil && !IsNil(o.BlockExploits) {
		return true
	}

	return false
}

// SetBlockExploits gets a reference to the given boolAsInt and assigns it to the BlockExploits field.
func (o *ProxyHostCreateResponse) SetBlockExploits(v boolAsInt) {
	o.BlockExploits = &v
}

// GetCachingEnabled returns the CachingEnabled field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetCachingEnabled() boolAsInt {
	if o == nil || IsNil(o.CachingEnabled) {
		var ret boolAsInt
		return ret
	}
	return *o.CachingEnabled
}

// GetCachingEnabledOk returns a tuple with the CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetCachingEnabledOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.CachingEnabled) {
		return nil, false
	}
	return o.CachingEnabled, true
}

// HasCachingEnabled returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasCachingEnabled() bool {
	if o != nil && !IsNil(o.CachingEnabled) {
		return true
	}

	return false
}

// SetCachingEnabled gets a reference to the given boolAsInt and assigns it to the CachingEnabled field.
func (o *ProxyHostCreateResponse) SetCachingEnabled(v boolAsInt) {
	o.CachingEnabled = &v
}

// GetAllowWebsocketUpgrade returns the AllowWebsocketUpgrade field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetAllowWebsocketUpgrade() boolAsInt {
	if o == nil || IsNil(o.AllowWebsocketUpgrade) {
		var ret boolAsInt
		return ret
	}
	return *o.AllowWebsocketUpgrade
}

// GetAllowWebsocketUpgradeOk returns a tuple with the AllowWebsocketUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetAllowWebsocketUpgradeOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.AllowWebsocketUpgrade) {
		return nil, false
	}
	return o.AllowWebsocketUpgrade, true
}

// HasAllowWebsocketUpgrade returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasAllowWebsocketUpgrade() bool {
	if o != nil && !IsNil(o.AllowWebsocketUpgrade) {
		return true
	}

	return false
}

// SetAllowWebsocketUpgrade gets a reference to the given boolAsInt and assigns it to the AllowWebsocketUpgrade field.
func (o *ProxyHostCreateResponse) SetAllowWebsocketUpgrade(v boolAsInt) {
	o.AllowWebsocketUpgrade = &v
}

// GetAccessListId returns the AccessListId field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetAccessListId() int64 {
	if o == nil || IsNil(o.AccessListId) {
		var ret int64
		return ret
	}
	return *o.AccessListId
}

// GetAccessListIdOk returns a tuple with the AccessListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetAccessListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessListId) {
		return nil, false
	}
	return o.AccessListId, true
}

// HasAccessListId returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasAccessListId() bool {
	if o != nil && !IsNil(o.AccessListId) {
		return true
	}

	return false
}

// SetAccessListId gets a reference to the given int64 and assigns it to the AccessListId field.
func (o *ProxyHostCreateResponse) SetAccessListId(v int64) {
	o.AccessListId = &v
}

// GetAdvancedConfig returns the AdvancedConfig field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetAdvancedConfig() string {
	if o == nil || IsNil(o.AdvancedConfig) {
		var ret string
		return ret
	}
	return *o.AdvancedConfig
}

// GetAdvancedConfigOk returns a tuple with the AdvancedConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetAdvancedConfigOk() (*string, bool) {
	if o == nil || IsNil(o.AdvancedConfig) {
		return nil, false
	}
	return o.AdvancedConfig, true
}

// HasAdvancedConfig returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasAdvancedConfig() bool {
	if o != nil && !IsNil(o.AdvancedConfig) {
		return true
	}

	return false
}

// SetAdvancedConfig gets a reference to the given string and assigns it to the AdvancedConfig field.
func (o *ProxyHostCreateResponse) SetAdvancedConfig(v string) {
	o.AdvancedConfig = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetEnabled() boolAsInt {
	if o == nil || IsNil(o.Enabled) {
		var ret boolAsInt
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetEnabledOk() (*boolAsInt, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given boolAsInt and assigns it to the Enabled field.
func (o *ProxyHostCreateResponse) SetEnabled(v boolAsInt) {
	o.Enabled = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ProxyHostCreateResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ProxyHostCreateResponse) GetLocations() []PutProxyHostUpdateRequestLocationsInner {
	if o == nil || IsNil(o.Locations) {
		var ret []PutProxyHostUpdateRequestLocationsInner
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyHostCreateResponse) GetLocationsOk() ([]PutProxyHostUpdateRequestLocationsInner, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ProxyHostCreateResponse) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []PutProxyHostUpdateRequestLocationsInner and assigns it to the Locations field.
func (o *ProxyHostCreateResponse) SetLocations(v []PutProxyHostUpdateRequestLocationsInner) {
	o.Locations = v
}

func (o ProxyHostCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyHostCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["created_on"] = o.CreatedOn
	}
	if !IsNil(o.ModifiedOn) {
		toSerialize["modified_on"] = o.ModifiedOn
	}
	if !IsNil(o.DomainNames) {
		toSerialize["domain_names"] = o.DomainNames
	}
	if !IsNil(o.ForwardScheme) {
		toSerialize["forward_scheme"] = o.ForwardScheme
	}
	if !IsNil(o.ForwardHost) {
		toSerialize["forward_host"] = o.ForwardHost
	}
	if !IsNil(o.ForwardPort) {
		toSerialize["forward_port"] = o.ForwardPort
	}
	if !IsNil(o.CertificateId) {
		toSerialize["certificate_id"] = o.CertificateId
	}
	if !IsNil(o.SslForced) {
		toSerialize["ssl_forced"] = o.SslForced
	}
	if !IsNil(o.HstsEnabled) {
		toSerialize["hsts_enabled"] = o.HstsEnabled
	}
	if !IsNil(o.HstsSubdomains) {
		toSerialize["hsts_subdomains"] = o.HstsSubdomains
	}
	if !IsNil(o.Http2Support) {
		toSerialize["http2_support"] = o.Http2Support
	}
	if !IsNil(o.BlockExploits) {
		toSerialize["block_exploits"] = o.BlockExploits
	}
	if !IsNil(o.CachingEnabled) {
		toSerialize["caching_enabled"] = o.CachingEnabled
	}
	if !IsNil(o.AllowWebsocketUpgrade) {
		toSerialize["allow_websocket_upgrade"] = o.AllowWebsocketUpgrade
	}
	if !IsNil(o.AccessListId) {
		toSerialize["access_list_id"] = o.AccessListId
	}
	if !IsNil(o.AdvancedConfig) {
		toSerialize["advanced_config"] = o.AdvancedConfig
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	return toSerialize, nil
}

type NullableProxyHostCreateResponse struct {
	value *ProxyHostCreateResponse
	isSet bool
}

func (v NullableProxyHostCreateResponse) Get() *ProxyHostCreateResponse {
	return v.value
}

func (v *NullableProxyHostCreateResponse) Set(val *ProxyHostCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyHostCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyHostCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyHostCreateResponse(val *ProxyHostCreateResponse) *NullableProxyHostCreateResponse {
	return &NullableProxyHostCreateResponse{value: val, isSet: true}
}

func (v NullableProxyHostCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyHostCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


