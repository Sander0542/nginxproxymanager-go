/*
Nginx Proxy Manager API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.12.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nginxproxymanager

import (
	"encoding/json"
	"fmt"
)

// checks if the GetProxyHosts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProxyHosts200ResponseInner{}

// GetProxyHosts200ResponseInner Proxy Host object
type GetProxyHosts200ResponseInner struct {
	// Unique identifier
	Id int64 `json:"id"`
	// Date and time of creation
	CreatedOn string `json:"created_on"`
	// Date and time of last update
	ModifiedOn string `json:"modified_on"`
	// User ID
	OwnerUserId int64 `json:"owner_user_id"`
	// Domain Names separated by a comma
	DomainNames []string `json:"domain_names"`
	ForwardHost string `json:"forward_host"`
	ForwardPort int64 `json:"forward_port"`
	// Access List ID
	AccessListId int64 `json:"access_list_id"`
	CertificateId GetProxyHosts200ResponseInnerCertificateId `json:"certificate_id"`
	// Is SSL Forced
	SslForced bool `json:"ssl_forced"`
	// Should we cache assets
	CachingEnabled bool `json:"caching_enabled"`
	// Should we block common exploits
	BlockExploits bool `json:"block_exploits"`
	AdvancedConfig string `json:"advanced_config"`
	Meta map[string]interface{} `json:"meta"`
	// Allow Websocket Upgrade for all paths
	AllowWebsocketUpgrade bool `json:"allow_websocket_upgrade"`
	// HTTP2 Protocol Support
	Http2Support bool `json:"http2_support"`
	ForwardScheme string `json:"forward_scheme"`
	// Is Enabled
	Enabled bool `json:"enabled"`
	Locations []GetProxyHosts200ResponseInnerLocationsInner `json:"locations"`
	// Is HSTS Enabled
	HstsEnabled bool `json:"hsts_enabled"`
	// Is HSTS applicable to all subdomains
	HstsSubdomains bool `json:"hsts_subdomains"`
	Certificate *GetProxyHosts200ResponseInnerCertificate `json:"certificate,omitempty"`
	Owner *GetAccessLists200ResponseInnerOwner `json:"owner,omitempty"`
	AccessList *GetProxyHosts200ResponseInnerAccessList `json:"access_list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetProxyHosts200ResponseInner GetProxyHosts200ResponseInner

// NewGetProxyHosts200ResponseInner instantiates a new GetProxyHosts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProxyHosts200ResponseInner(id int64, createdOn string, modifiedOn string, ownerUserId int64, domainNames []string, forwardHost string, forwardPort int64, accessListId int64, certificateId GetProxyHosts200ResponseInnerCertificateId, sslForced bool, cachingEnabled bool, blockExploits bool, advancedConfig string, meta map[string]interface{}, allowWebsocketUpgrade bool, http2Support bool, forwardScheme string, enabled bool, locations []GetProxyHosts200ResponseInnerLocationsInner, hstsEnabled bool, hstsSubdomains bool) *GetProxyHosts200ResponseInner {
	this := GetProxyHosts200ResponseInner{}
	this.Id = id
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	this.OwnerUserId = ownerUserId
	this.DomainNames = domainNames
	this.ForwardHost = forwardHost
	this.ForwardPort = forwardPort
	this.AccessListId = accessListId
	this.CertificateId = certificateId
	this.SslForced = sslForced
	this.CachingEnabled = cachingEnabled
	this.BlockExploits = blockExploits
	this.AdvancedConfig = advancedConfig
	this.Meta = meta
	this.AllowWebsocketUpgrade = allowWebsocketUpgrade
	this.Http2Support = http2Support
	this.ForwardScheme = forwardScheme
	this.Enabled = enabled
	this.Locations = locations
	this.HstsEnabled = hstsEnabled
	this.HstsSubdomains = hstsSubdomains
	return &this
}

// NewGetProxyHosts200ResponseInnerWithDefaults instantiates a new GetProxyHosts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProxyHosts200ResponseInnerWithDefaults() *GetProxyHosts200ResponseInner {
	this := GetProxyHosts200ResponseInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetProxyHosts200ResponseInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetProxyHosts200ResponseInner) SetId(v int64) {
	o.Id = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *GetProxyHosts200ResponseInner) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *GetProxyHosts200ResponseInner) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *GetProxyHosts200ResponseInner) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *GetProxyHosts200ResponseInner) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetOwnerUserId returns the OwnerUserId field value
func (o *GetProxyHosts200ResponseInner) GetOwnerUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetOwnerUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUserId, true
}

// SetOwnerUserId sets field value
func (o *GetProxyHosts200ResponseInner) SetOwnerUserId(v int64) {
	o.OwnerUserId = v
}

// GetDomainNames returns the DomainNames field value
func (o *GetProxyHosts200ResponseInner) GetDomainNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetDomainNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainNames, true
}

// SetDomainNames sets field value
func (o *GetProxyHosts200ResponseInner) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetForwardHost returns the ForwardHost field value
func (o *GetProxyHosts200ResponseInner) GetForwardHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForwardHost
}

// GetForwardHostOk returns a tuple with the ForwardHost field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetForwardHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwardHost, true
}

// SetForwardHost sets field value
func (o *GetProxyHosts200ResponseInner) SetForwardHost(v string) {
	o.ForwardHost = v
}

// GetForwardPort returns the ForwardPort field value
func (o *GetProxyHosts200ResponseInner) GetForwardPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ForwardPort
}

// GetForwardPortOk returns a tuple with the ForwardPort field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetForwardPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwardPort, true
}

// SetForwardPort sets field value
func (o *GetProxyHosts200ResponseInner) SetForwardPort(v int64) {
	o.ForwardPort = v
}

// GetAccessListId returns the AccessListId field value
func (o *GetProxyHosts200ResponseInner) GetAccessListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccessListId
}

// GetAccessListIdOk returns a tuple with the AccessListId field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetAccessListIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessListId, true
}

// SetAccessListId sets field value
func (o *GetProxyHosts200ResponseInner) SetAccessListId(v int64) {
	o.AccessListId = v
}

// GetCertificateId returns the CertificateId field value
func (o *GetProxyHosts200ResponseInner) GetCertificateId() GetProxyHosts200ResponseInnerCertificateId {
	if o == nil {
		var ret GetProxyHosts200ResponseInnerCertificateId
		return ret
	}

	return o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetCertificateIdOk() (*GetProxyHosts200ResponseInnerCertificateId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateId, true
}

// SetCertificateId sets field value
func (o *GetProxyHosts200ResponseInner) SetCertificateId(v GetProxyHosts200ResponseInnerCertificateId) {
	o.CertificateId = v
}

// GetSslForced returns the SslForced field value
func (o *GetProxyHosts200ResponseInner) GetSslForced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SslForced
}

// GetSslForcedOk returns a tuple with the SslForced field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetSslForcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SslForced, true
}

// SetSslForced sets field value
func (o *GetProxyHosts200ResponseInner) SetSslForced(v bool) {
	o.SslForced = v
}

// GetCachingEnabled returns the CachingEnabled field value
func (o *GetProxyHosts200ResponseInner) GetCachingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CachingEnabled
}

// GetCachingEnabledOk returns a tuple with the CachingEnabled field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetCachingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CachingEnabled, true
}

// SetCachingEnabled sets field value
func (o *GetProxyHosts200ResponseInner) SetCachingEnabled(v bool) {
	o.CachingEnabled = v
}

// GetBlockExploits returns the BlockExploits field value
func (o *GetProxyHosts200ResponseInner) GetBlockExploits() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BlockExploits
}

// GetBlockExploitsOk returns a tuple with the BlockExploits field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetBlockExploitsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockExploits, true
}

// SetBlockExploits sets field value
func (o *GetProxyHosts200ResponseInner) SetBlockExploits(v bool) {
	o.BlockExploits = v
}

// GetAdvancedConfig returns the AdvancedConfig field value
func (o *GetProxyHosts200ResponseInner) GetAdvancedConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvancedConfig
}

// GetAdvancedConfigOk returns a tuple with the AdvancedConfig field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetAdvancedConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvancedConfig, true
}

// SetAdvancedConfig sets field value
func (o *GetProxyHosts200ResponseInner) SetAdvancedConfig(v string) {
	o.AdvancedConfig = v
}

// GetMeta returns the Meta field value
func (o *GetProxyHosts200ResponseInner) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetProxyHosts200ResponseInner) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetAllowWebsocketUpgrade returns the AllowWebsocketUpgrade field value
func (o *GetProxyHosts200ResponseInner) GetAllowWebsocketUpgrade() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowWebsocketUpgrade
}

// GetAllowWebsocketUpgradeOk returns a tuple with the AllowWebsocketUpgrade field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetAllowWebsocketUpgradeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowWebsocketUpgrade, true
}

// SetAllowWebsocketUpgrade sets field value
func (o *GetProxyHosts200ResponseInner) SetAllowWebsocketUpgrade(v bool) {
	o.AllowWebsocketUpgrade = v
}

// GetHttp2Support returns the Http2Support field value
func (o *GetProxyHosts200ResponseInner) GetHttp2Support() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Http2Support
}

// GetHttp2SupportOk returns a tuple with the Http2Support field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetHttp2SupportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Http2Support, true
}

// SetHttp2Support sets field value
func (o *GetProxyHosts200ResponseInner) SetHttp2Support(v bool) {
	o.Http2Support = v
}

// GetForwardScheme returns the ForwardScheme field value
func (o *GetProxyHosts200ResponseInner) GetForwardScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForwardScheme
}

// GetForwardSchemeOk returns a tuple with the ForwardScheme field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetForwardSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwardScheme, true
}

// SetForwardScheme sets field value
func (o *GetProxyHosts200ResponseInner) SetForwardScheme(v string) {
	o.ForwardScheme = v
}

// GetEnabled returns the Enabled field value
func (o *GetProxyHosts200ResponseInner) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetProxyHosts200ResponseInner) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLocations returns the Locations field value
func (o *GetProxyHosts200ResponseInner) GetLocations() []GetProxyHosts200ResponseInnerLocationsInner {
	if o == nil {
		var ret []GetProxyHosts200ResponseInnerLocationsInner
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetLocationsOk() ([]GetProxyHosts200ResponseInnerLocationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *GetProxyHosts200ResponseInner) SetLocations(v []GetProxyHosts200ResponseInnerLocationsInner) {
	o.Locations = v
}

// GetHstsEnabled returns the HstsEnabled field value
func (o *GetProxyHosts200ResponseInner) GetHstsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HstsEnabled
}

// GetHstsEnabledOk returns a tuple with the HstsEnabled field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetHstsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HstsEnabled, true
}

// SetHstsEnabled sets field value
func (o *GetProxyHosts200ResponseInner) SetHstsEnabled(v bool) {
	o.HstsEnabled = v
}

// GetHstsSubdomains returns the HstsSubdomains field value
func (o *GetProxyHosts200ResponseInner) GetHstsSubdomains() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HstsSubdomains
}

// GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field value
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetHstsSubdomainsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HstsSubdomains, true
}

// SetHstsSubdomains sets field value
func (o *GetProxyHosts200ResponseInner) SetHstsSubdomains(v bool) {
	o.HstsSubdomains = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *GetProxyHosts200ResponseInner) GetCertificate() GetProxyHosts200ResponseInnerCertificate {
	if o == nil || IsNil(o.Certificate) {
		var ret GetProxyHosts200ResponseInnerCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetCertificateOk() (*GetProxyHosts200ResponseInnerCertificate, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *GetProxyHosts200ResponseInner) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given GetProxyHosts200ResponseInnerCertificate and assigns it to the Certificate field.
func (o *GetProxyHosts200ResponseInner) SetCertificate(v GetProxyHosts200ResponseInnerCertificate) {
	o.Certificate = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GetProxyHosts200ResponseInner) GetOwner() GetAccessLists200ResponseInnerOwner {
	if o == nil || IsNil(o.Owner) {
		var ret GetAccessLists200ResponseInnerOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetOwnerOk() (*GetAccessLists200ResponseInnerOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GetProxyHosts200ResponseInner) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given GetAccessLists200ResponseInnerOwner and assigns it to the Owner field.
func (o *GetProxyHosts200ResponseInner) SetOwner(v GetAccessLists200ResponseInnerOwner) {
	o.Owner = &v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *GetProxyHosts200ResponseInner) GetAccessList() GetProxyHosts200ResponseInnerAccessList {
	if o == nil || IsNil(o.AccessList) {
		var ret GetProxyHosts200ResponseInnerAccessList
		return ret
	}
	return *o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProxyHosts200ResponseInner) GetAccessListOk() (*GetProxyHosts200ResponseInnerAccessList, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *GetProxyHosts200ResponseInner) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given GetProxyHosts200ResponseInnerAccessList and assigns it to the AccessList field.
func (o *GetProxyHosts200ResponseInner) SetAccessList(v GetProxyHosts200ResponseInnerAccessList) {
	o.AccessList = &v
}

func (o GetProxyHosts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProxyHosts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["owner_user_id"] = o.OwnerUserId
	toSerialize["domain_names"] = o.DomainNames
	toSerialize["forward_host"] = o.ForwardHost
	toSerialize["forward_port"] = o.ForwardPort
	toSerialize["access_list_id"] = o.AccessListId
	toSerialize["certificate_id"] = o.CertificateId
	toSerialize["ssl_forced"] = o.SslForced
	toSerialize["caching_enabled"] = o.CachingEnabled
	toSerialize["block_exploits"] = o.BlockExploits
	toSerialize["advanced_config"] = o.AdvancedConfig
	toSerialize["meta"] = o.Meta
	toSerialize["allow_websocket_upgrade"] = o.AllowWebsocketUpgrade
	toSerialize["http2_support"] = o.Http2Support
	toSerialize["forward_scheme"] = o.ForwardScheme
	toSerialize["enabled"] = o.Enabled
	toSerialize["locations"] = o.Locations
	toSerialize["hsts_enabled"] = o.HstsEnabled
	toSerialize["hsts_subdomains"] = o.HstsSubdomains
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.AccessList) {
		toSerialize["access_list"] = o.AccessList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetProxyHosts200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_on",
		"modified_on",
		"owner_user_id",
		"domain_names",
		"forward_host",
		"forward_port",
		"access_list_id",
		"certificate_id",
		"ssl_forced",
		"caching_enabled",
		"block_exploits",
		"advanced_config",
		"meta",
		"allow_websocket_upgrade",
		"http2_support",
		"forward_scheme",
		"enabled",
		"locations",
		"hsts_enabled",
		"hsts_subdomains",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetProxyHosts200ResponseInner := _GetProxyHosts200ResponseInner{}

	err = json.Unmarshal(data, &varGetProxyHosts200ResponseInner)

	if err != nil {
		return err
	}

	*o = GetProxyHosts200ResponseInner(varGetProxyHosts200ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "modified_on")
		delete(additionalProperties, "owner_user_id")
		delete(additionalProperties, "domain_names")
		delete(additionalProperties, "forward_host")
		delete(additionalProperties, "forward_port")
		delete(additionalProperties, "access_list_id")
		delete(additionalProperties, "certificate_id")
		delete(additionalProperties, "ssl_forced")
		delete(additionalProperties, "caching_enabled")
		delete(additionalProperties, "block_exploits")
		delete(additionalProperties, "advanced_config")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "allow_websocket_upgrade")
		delete(additionalProperties, "http2_support")
		delete(additionalProperties, "forward_scheme")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "hsts_enabled")
		delete(additionalProperties, "hsts_subdomains")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "access_list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetProxyHosts200ResponseInner struct {
	value *GetProxyHosts200ResponseInner
	isSet bool
}

func (v NullableGetProxyHosts200ResponseInner) Get() *GetProxyHosts200ResponseInner {
	return v.value
}

func (v *NullableGetProxyHosts200ResponseInner) Set(val *GetProxyHosts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProxyHosts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProxyHosts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProxyHosts200ResponseInner(val *GetProxyHosts200ResponseInner) *NullableGetProxyHosts200ResponseInner {
	return &NullableGetProxyHosts200ResponseInner{value: val, isSet: true}
}

func (v NullableGetProxyHosts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProxyHosts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


