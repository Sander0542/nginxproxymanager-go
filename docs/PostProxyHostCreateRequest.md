# PostProxyHostCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | **[]string** | Domain Names separated by a comma | 
**ForwardScheme** | **string** |  | 
**ForwardHost** | **string** |  | 
**ForwardPort** | **int64** |  | 
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

### NewPostProxyHostCreateRequest

`func NewPostProxyHostCreateRequest(domainNames []string, forwardScheme string, forwardHost string, forwardPort int64, ) *PostProxyHostCreateRequest`

NewPostProxyHostCreateRequest instantiates a new PostProxyHostCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProxyHostCreateRequestWithDefaults

`func NewPostProxyHostCreateRequestWithDefaults() *PostProxyHostCreateRequest`

NewPostProxyHostCreateRequestWithDefaults instantiates a new PostProxyHostCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainNames

`func (o *PostProxyHostCreateRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *PostProxyHostCreateRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *PostProxyHostCreateRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetForwardScheme

`func (o *PostProxyHostCreateRequest) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *PostProxyHostCreateRequest) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *PostProxyHostCreateRequest) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.


### GetForwardHost

`func (o *PostProxyHostCreateRequest) GetForwardHost() string`

GetForwardHost returns the ForwardHost field if non-nil, zero value otherwise.

### GetForwardHostOk

`func (o *PostProxyHostCreateRequest) GetForwardHostOk() (*string, bool)`

GetForwardHostOk returns a tuple with the ForwardHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHost

`func (o *PostProxyHostCreateRequest) SetForwardHost(v string)`

SetForwardHost sets ForwardHost field to given value.


### GetForwardPort

`func (o *PostProxyHostCreateRequest) GetForwardPort() int64`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *PostProxyHostCreateRequest) GetForwardPortOk() (*int64, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *PostProxyHostCreateRequest) SetForwardPort(v int64)`

SetForwardPort sets ForwardPort field to given value.


### GetCertificateId

`func (o *PostProxyHostCreateRequest) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *PostProxyHostCreateRequest) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *PostProxyHostCreateRequest) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *PostProxyHostCreateRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *PostProxyHostCreateRequest) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *PostProxyHostCreateRequest) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *PostProxyHostCreateRequest) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *PostProxyHostCreateRequest) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *PostProxyHostCreateRequest) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *PostProxyHostCreateRequest) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *PostProxyHostCreateRequest) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *PostProxyHostCreateRequest) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *PostProxyHostCreateRequest) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *PostProxyHostCreateRequest) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *PostProxyHostCreateRequest) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *PostProxyHostCreateRequest) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *PostProxyHostCreateRequest) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *PostProxyHostCreateRequest) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *PostProxyHostCreateRequest) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *PostProxyHostCreateRequest) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *PostProxyHostCreateRequest) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *PostProxyHostCreateRequest) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *PostProxyHostCreateRequest) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *PostProxyHostCreateRequest) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetCachingEnabled

`func (o *PostProxyHostCreateRequest) GetCachingEnabled() boolAsInt`

GetCachingEnabled returns the CachingEnabled field if non-nil, zero value otherwise.

### GetCachingEnabledOk

`func (o *PostProxyHostCreateRequest) GetCachingEnabledOk() (*boolAsInt, bool)`

GetCachingEnabledOk returns a tuple with the CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingEnabled

`func (o *PostProxyHostCreateRequest) SetCachingEnabled(v boolAsInt)`

SetCachingEnabled sets CachingEnabled field to given value.

### HasCachingEnabled

`func (o *PostProxyHostCreateRequest) HasCachingEnabled() bool`

HasCachingEnabled returns a boolean if a field has been set.

### GetAllowWebsocketUpgrade

`func (o *PostProxyHostCreateRequest) GetAllowWebsocketUpgrade() boolAsInt`

GetAllowWebsocketUpgrade returns the AllowWebsocketUpgrade field if non-nil, zero value otherwise.

### GetAllowWebsocketUpgradeOk

`func (o *PostProxyHostCreateRequest) GetAllowWebsocketUpgradeOk() (*boolAsInt, bool)`

GetAllowWebsocketUpgradeOk returns a tuple with the AllowWebsocketUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebsocketUpgrade

`func (o *PostProxyHostCreateRequest) SetAllowWebsocketUpgrade(v boolAsInt)`

SetAllowWebsocketUpgrade sets AllowWebsocketUpgrade field to given value.

### HasAllowWebsocketUpgrade

`func (o *PostProxyHostCreateRequest) HasAllowWebsocketUpgrade() bool`

HasAllowWebsocketUpgrade returns a boolean if a field has been set.

### GetAccessListId

`func (o *PostProxyHostCreateRequest) GetAccessListId() int64`

GetAccessListId returns the AccessListId field if non-nil, zero value otherwise.

### GetAccessListIdOk

`func (o *PostProxyHostCreateRequest) GetAccessListIdOk() (*int64, bool)`

GetAccessListIdOk returns a tuple with the AccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListId

`func (o *PostProxyHostCreateRequest) SetAccessListId(v int64)`

SetAccessListId sets AccessListId field to given value.

### HasAccessListId

`func (o *PostProxyHostCreateRequest) HasAccessListId() bool`

HasAccessListId returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *PostProxyHostCreateRequest) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *PostProxyHostCreateRequest) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *PostProxyHostCreateRequest) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *PostProxyHostCreateRequest) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *PostProxyHostCreateRequest) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PostProxyHostCreateRequest) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PostProxyHostCreateRequest) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PostProxyHostCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *PostProxyHostCreateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PostProxyHostCreateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PostProxyHostCreateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PostProxyHostCreateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLocations

`func (o *PostProxyHostCreateRequest) GetLocations() []PutProxyHostUpdateRequestLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *PostProxyHostCreateRequest) GetLocationsOk() (*[]PutProxyHostUpdateRequestLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *PostProxyHostCreateRequest) SetLocations(v []PutProxyHostUpdateRequestLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *PostProxyHostCreateRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


