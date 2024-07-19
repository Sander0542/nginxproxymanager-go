# PutProxyHostUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**ForwardScheme** | Pointer to **string** |  | [optional] 
**ForwardHost** | Pointer to **string** |  | [optional] 
**ForwardPort** | Pointer to **int64** |  | [optional] 
**CertificateId** | Pointer to [**PutProxyHostUpdateRequestCertificateId**](PutProxyHostUpdateRequestCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **boolAsInt** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**Http2Support** | Pointer to **boolAsInt** | HTTP2 Protocol Support | [optional] 
**BlockExploits** | Pointer to **boolAsInt** | Should we block common exploits | [optional] 
**CachingEnabled** | Pointer to **boolAsInt** | Should we cache assets | [optional] 
**AllowWebsocketUpgrade** | Pointer to **boolAsInt** | Allow Websocket Upgrade for all paths | [optional] 
**AccessListId** | Pointer to **int64** | Access List ID | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **boolAsInt** | Is Enabled | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Locations** | Pointer to [**[]PutProxyHostUpdateRequestLocationsInner**](PutProxyHostUpdateRequestLocationsInner.md) |  | [optional] 

## Methods

### NewPutProxyHostUpdateRequest

`func NewPutProxyHostUpdateRequest() *PutProxyHostUpdateRequest`

NewPutProxyHostUpdateRequest instantiates a new PutProxyHostUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutProxyHostUpdateRequestWithDefaults

`func NewPutProxyHostUpdateRequestWithDefaults() *PutProxyHostUpdateRequest`

NewPutProxyHostUpdateRequestWithDefaults instantiates a new PutProxyHostUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainNames

`func (o *PutProxyHostUpdateRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *PutProxyHostUpdateRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *PutProxyHostUpdateRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *PutProxyHostUpdateRequest) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetForwardScheme

`func (o *PutProxyHostUpdateRequest) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *PutProxyHostUpdateRequest) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *PutProxyHostUpdateRequest) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.

### HasForwardScheme

`func (o *PutProxyHostUpdateRequest) HasForwardScheme() bool`

HasForwardScheme returns a boolean if a field has been set.

### GetForwardHost

`func (o *PutProxyHostUpdateRequest) GetForwardHost() string`

GetForwardHost returns the ForwardHost field if non-nil, zero value otherwise.

### GetForwardHostOk

`func (o *PutProxyHostUpdateRequest) GetForwardHostOk() (*string, bool)`

GetForwardHostOk returns a tuple with the ForwardHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHost

`func (o *PutProxyHostUpdateRequest) SetForwardHost(v string)`

SetForwardHost sets ForwardHost field to given value.

### HasForwardHost

`func (o *PutProxyHostUpdateRequest) HasForwardHost() bool`

HasForwardHost returns a boolean if a field has been set.

### GetForwardPort

`func (o *PutProxyHostUpdateRequest) GetForwardPort() int64`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *PutProxyHostUpdateRequest) GetForwardPortOk() (*int64, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *PutProxyHostUpdateRequest) SetForwardPort(v int64)`

SetForwardPort sets ForwardPort field to given value.

### HasForwardPort

`func (o *PutProxyHostUpdateRequest) HasForwardPort() bool`

HasForwardPort returns a boolean if a field has been set.

### GetCertificateId

`func (o *PutProxyHostUpdateRequest) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *PutProxyHostUpdateRequest) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *PutProxyHostUpdateRequest) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *PutProxyHostUpdateRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *PutProxyHostUpdateRequest) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *PutProxyHostUpdateRequest) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *PutProxyHostUpdateRequest) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *PutProxyHostUpdateRequest) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *PutProxyHostUpdateRequest) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *PutProxyHostUpdateRequest) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *PutProxyHostUpdateRequest) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *PutProxyHostUpdateRequest) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *PutProxyHostUpdateRequest) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *PutProxyHostUpdateRequest) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *PutProxyHostUpdateRequest) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *PutProxyHostUpdateRequest) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *PutProxyHostUpdateRequest) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *PutProxyHostUpdateRequest) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *PutProxyHostUpdateRequest) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *PutProxyHostUpdateRequest) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *PutProxyHostUpdateRequest) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *PutProxyHostUpdateRequest) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *PutProxyHostUpdateRequest) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *PutProxyHostUpdateRequest) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetCachingEnabled

`func (o *PutProxyHostUpdateRequest) GetCachingEnabled() boolAsInt`

GetCachingEnabled returns the CachingEnabled field if non-nil, zero value otherwise.

### GetCachingEnabledOk

`func (o *PutProxyHostUpdateRequest) GetCachingEnabledOk() (*boolAsInt, bool)`

GetCachingEnabledOk returns a tuple with the CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingEnabled

`func (o *PutProxyHostUpdateRequest) SetCachingEnabled(v boolAsInt)`

SetCachingEnabled sets CachingEnabled field to given value.

### HasCachingEnabled

`func (o *PutProxyHostUpdateRequest) HasCachingEnabled() bool`

HasCachingEnabled returns a boolean if a field has been set.

### GetAllowWebsocketUpgrade

`func (o *PutProxyHostUpdateRequest) GetAllowWebsocketUpgrade() boolAsInt`

GetAllowWebsocketUpgrade returns the AllowWebsocketUpgrade field if non-nil, zero value otherwise.

### GetAllowWebsocketUpgradeOk

`func (o *PutProxyHostUpdateRequest) GetAllowWebsocketUpgradeOk() (*boolAsInt, bool)`

GetAllowWebsocketUpgradeOk returns a tuple with the AllowWebsocketUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebsocketUpgrade

`func (o *PutProxyHostUpdateRequest) SetAllowWebsocketUpgrade(v boolAsInt)`

SetAllowWebsocketUpgrade sets AllowWebsocketUpgrade field to given value.

### HasAllowWebsocketUpgrade

`func (o *PutProxyHostUpdateRequest) HasAllowWebsocketUpgrade() bool`

HasAllowWebsocketUpgrade returns a boolean if a field has been set.

### GetAccessListId

`func (o *PutProxyHostUpdateRequest) GetAccessListId() int64`

GetAccessListId returns the AccessListId field if non-nil, zero value otherwise.

### GetAccessListIdOk

`func (o *PutProxyHostUpdateRequest) GetAccessListIdOk() (*int64, bool)`

GetAccessListIdOk returns a tuple with the AccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListId

`func (o *PutProxyHostUpdateRequest) SetAccessListId(v int64)`

SetAccessListId sets AccessListId field to given value.

### HasAccessListId

`func (o *PutProxyHostUpdateRequest) HasAccessListId() bool`

HasAccessListId returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *PutProxyHostUpdateRequest) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *PutProxyHostUpdateRequest) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *PutProxyHostUpdateRequest) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *PutProxyHostUpdateRequest) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *PutProxyHostUpdateRequest) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutProxyHostUpdateRequest) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutProxyHostUpdateRequest) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PutProxyHostUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *PutProxyHostUpdateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PutProxyHostUpdateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PutProxyHostUpdateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PutProxyHostUpdateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLocations

`func (o *PutProxyHostUpdateRequest) GetLocations() []PutProxyHostUpdateRequestLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *PutProxyHostUpdateRequest) GetLocationsOk() (*[]PutProxyHostUpdateRequestLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *PutProxyHostUpdateRequest) SetLocations(v []PutProxyHostUpdateRequestLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *PutProxyHostUpdateRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


