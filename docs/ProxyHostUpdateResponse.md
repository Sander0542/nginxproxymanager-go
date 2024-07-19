# ProxyHostUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** | Date and time of creation | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** | Date and time of last update | [optional] [readonly] 
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**ForwardScheme** | Pointer to **string** |  | [optional] 
**ForwardHost** | Pointer to **string** |  | [optional] 
**ForwardPort** | Pointer to **int64** |  | [optional] 
**CertificateId** | Pointer to [**PutProxyHostUpdateRequestCertificateId**](PutProxyHostUpdateRequestCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **boolAsInt** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **boolAsInt** | Is HSTS applicable to all subdomains | [optional] 
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

### NewProxyHostUpdateResponse

`func NewProxyHostUpdateResponse() *ProxyHostUpdateResponse`

NewProxyHostUpdateResponse instantiates a new ProxyHostUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyHostUpdateResponseWithDefaults

`func NewProxyHostUpdateResponseWithDefaults() *ProxyHostUpdateResponse`

NewProxyHostUpdateResponseWithDefaults instantiates a new ProxyHostUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProxyHostUpdateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProxyHostUpdateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProxyHostUpdateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProxyHostUpdateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ProxyHostUpdateResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ProxyHostUpdateResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ProxyHostUpdateResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ProxyHostUpdateResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *ProxyHostUpdateResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ProxyHostUpdateResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ProxyHostUpdateResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *ProxyHostUpdateResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetDomainNames

`func (o *ProxyHostUpdateResponse) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *ProxyHostUpdateResponse) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *ProxyHostUpdateResponse) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *ProxyHostUpdateResponse) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetForwardScheme

`func (o *ProxyHostUpdateResponse) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *ProxyHostUpdateResponse) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *ProxyHostUpdateResponse) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.

### HasForwardScheme

`func (o *ProxyHostUpdateResponse) HasForwardScheme() bool`

HasForwardScheme returns a boolean if a field has been set.

### GetForwardHost

`func (o *ProxyHostUpdateResponse) GetForwardHost() string`

GetForwardHost returns the ForwardHost field if non-nil, zero value otherwise.

### GetForwardHostOk

`func (o *ProxyHostUpdateResponse) GetForwardHostOk() (*string, bool)`

GetForwardHostOk returns a tuple with the ForwardHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHost

`func (o *ProxyHostUpdateResponse) SetForwardHost(v string)`

SetForwardHost sets ForwardHost field to given value.

### HasForwardHost

`func (o *ProxyHostUpdateResponse) HasForwardHost() bool`

HasForwardHost returns a boolean if a field has been set.

### GetForwardPort

`func (o *ProxyHostUpdateResponse) GetForwardPort() int64`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *ProxyHostUpdateResponse) GetForwardPortOk() (*int64, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *ProxyHostUpdateResponse) SetForwardPort(v int64)`

SetForwardPort sets ForwardPort field to given value.

### HasForwardPort

`func (o *ProxyHostUpdateResponse) HasForwardPort() bool`

HasForwardPort returns a boolean if a field has been set.

### GetCertificateId

`func (o *ProxyHostUpdateResponse) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *ProxyHostUpdateResponse) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *ProxyHostUpdateResponse) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *ProxyHostUpdateResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *ProxyHostUpdateResponse) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *ProxyHostUpdateResponse) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *ProxyHostUpdateResponse) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *ProxyHostUpdateResponse) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *ProxyHostUpdateResponse) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *ProxyHostUpdateResponse) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *ProxyHostUpdateResponse) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *ProxyHostUpdateResponse) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *ProxyHostUpdateResponse) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *ProxyHostUpdateResponse) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *ProxyHostUpdateResponse) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *ProxyHostUpdateResponse) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *ProxyHostUpdateResponse) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *ProxyHostUpdateResponse) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *ProxyHostUpdateResponse) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *ProxyHostUpdateResponse) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *ProxyHostUpdateResponse) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *ProxyHostUpdateResponse) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *ProxyHostUpdateResponse) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *ProxyHostUpdateResponse) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetCachingEnabled

`func (o *ProxyHostUpdateResponse) GetCachingEnabled() boolAsInt`

GetCachingEnabled returns the CachingEnabled field if non-nil, zero value otherwise.

### GetCachingEnabledOk

`func (o *ProxyHostUpdateResponse) GetCachingEnabledOk() (*boolAsInt, bool)`

GetCachingEnabledOk returns a tuple with the CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingEnabled

`func (o *ProxyHostUpdateResponse) SetCachingEnabled(v boolAsInt)`

SetCachingEnabled sets CachingEnabled field to given value.

### HasCachingEnabled

`func (o *ProxyHostUpdateResponse) HasCachingEnabled() bool`

HasCachingEnabled returns a boolean if a field has been set.

### GetAllowWebsocketUpgrade

`func (o *ProxyHostUpdateResponse) GetAllowWebsocketUpgrade() boolAsInt`

GetAllowWebsocketUpgrade returns the AllowWebsocketUpgrade field if non-nil, zero value otherwise.

### GetAllowWebsocketUpgradeOk

`func (o *ProxyHostUpdateResponse) GetAllowWebsocketUpgradeOk() (*boolAsInt, bool)`

GetAllowWebsocketUpgradeOk returns a tuple with the AllowWebsocketUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebsocketUpgrade

`func (o *ProxyHostUpdateResponse) SetAllowWebsocketUpgrade(v boolAsInt)`

SetAllowWebsocketUpgrade sets AllowWebsocketUpgrade field to given value.

### HasAllowWebsocketUpgrade

`func (o *ProxyHostUpdateResponse) HasAllowWebsocketUpgrade() bool`

HasAllowWebsocketUpgrade returns a boolean if a field has been set.

### GetAccessListId

`func (o *ProxyHostUpdateResponse) GetAccessListId() int64`

GetAccessListId returns the AccessListId field if non-nil, zero value otherwise.

### GetAccessListIdOk

`func (o *ProxyHostUpdateResponse) GetAccessListIdOk() (*int64, bool)`

GetAccessListIdOk returns a tuple with the AccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListId

`func (o *ProxyHostUpdateResponse) SetAccessListId(v int64)`

SetAccessListId sets AccessListId field to given value.

### HasAccessListId

`func (o *ProxyHostUpdateResponse) HasAccessListId() bool`

HasAccessListId returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *ProxyHostUpdateResponse) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *ProxyHostUpdateResponse) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *ProxyHostUpdateResponse) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *ProxyHostUpdateResponse) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *ProxyHostUpdateResponse) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProxyHostUpdateResponse) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProxyHostUpdateResponse) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProxyHostUpdateResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *ProxyHostUpdateResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProxyHostUpdateResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProxyHostUpdateResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProxyHostUpdateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLocations

`func (o *ProxyHostUpdateResponse) GetLocations() []PutProxyHostUpdateRequestLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ProxyHostUpdateResponse) GetLocationsOk() (*[]PutProxyHostUpdateRequestLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ProxyHostUpdateResponse) SetLocations(v []PutProxyHostUpdateRequestLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ProxyHostUpdateResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


