# ProxyHostGetResponse

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

### NewProxyHostGetResponse

`func NewProxyHostGetResponse() *ProxyHostGetResponse`

NewProxyHostGetResponse instantiates a new ProxyHostGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyHostGetResponseWithDefaults

`func NewProxyHostGetResponseWithDefaults() *ProxyHostGetResponse`

NewProxyHostGetResponseWithDefaults instantiates a new ProxyHostGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProxyHostGetResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProxyHostGetResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProxyHostGetResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProxyHostGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ProxyHostGetResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ProxyHostGetResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ProxyHostGetResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ProxyHostGetResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *ProxyHostGetResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ProxyHostGetResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ProxyHostGetResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *ProxyHostGetResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetDomainNames

`func (o *ProxyHostGetResponse) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *ProxyHostGetResponse) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *ProxyHostGetResponse) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *ProxyHostGetResponse) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetForwardScheme

`func (o *ProxyHostGetResponse) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *ProxyHostGetResponse) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *ProxyHostGetResponse) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.

### HasForwardScheme

`func (o *ProxyHostGetResponse) HasForwardScheme() bool`

HasForwardScheme returns a boolean if a field has been set.

### GetForwardHost

`func (o *ProxyHostGetResponse) GetForwardHost() string`

GetForwardHost returns the ForwardHost field if non-nil, zero value otherwise.

### GetForwardHostOk

`func (o *ProxyHostGetResponse) GetForwardHostOk() (*string, bool)`

GetForwardHostOk returns a tuple with the ForwardHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHost

`func (o *ProxyHostGetResponse) SetForwardHost(v string)`

SetForwardHost sets ForwardHost field to given value.

### HasForwardHost

`func (o *ProxyHostGetResponse) HasForwardHost() bool`

HasForwardHost returns a boolean if a field has been set.

### GetForwardPort

`func (o *ProxyHostGetResponse) GetForwardPort() int64`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *ProxyHostGetResponse) GetForwardPortOk() (*int64, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *ProxyHostGetResponse) SetForwardPort(v int64)`

SetForwardPort sets ForwardPort field to given value.

### HasForwardPort

`func (o *ProxyHostGetResponse) HasForwardPort() bool`

HasForwardPort returns a boolean if a field has been set.

### GetCertificateId

`func (o *ProxyHostGetResponse) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *ProxyHostGetResponse) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *ProxyHostGetResponse) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *ProxyHostGetResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *ProxyHostGetResponse) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *ProxyHostGetResponse) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *ProxyHostGetResponse) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *ProxyHostGetResponse) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *ProxyHostGetResponse) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *ProxyHostGetResponse) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *ProxyHostGetResponse) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *ProxyHostGetResponse) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *ProxyHostGetResponse) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *ProxyHostGetResponse) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *ProxyHostGetResponse) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *ProxyHostGetResponse) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *ProxyHostGetResponse) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *ProxyHostGetResponse) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *ProxyHostGetResponse) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *ProxyHostGetResponse) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *ProxyHostGetResponse) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *ProxyHostGetResponse) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *ProxyHostGetResponse) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *ProxyHostGetResponse) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetCachingEnabled

`func (o *ProxyHostGetResponse) GetCachingEnabled() boolAsInt`

GetCachingEnabled returns the CachingEnabled field if non-nil, zero value otherwise.

### GetCachingEnabledOk

`func (o *ProxyHostGetResponse) GetCachingEnabledOk() (*boolAsInt, bool)`

GetCachingEnabledOk returns a tuple with the CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingEnabled

`func (o *ProxyHostGetResponse) SetCachingEnabled(v boolAsInt)`

SetCachingEnabled sets CachingEnabled field to given value.

### HasCachingEnabled

`func (o *ProxyHostGetResponse) HasCachingEnabled() bool`

HasCachingEnabled returns a boolean if a field has been set.

### GetAllowWebsocketUpgrade

`func (o *ProxyHostGetResponse) GetAllowWebsocketUpgrade() boolAsInt`

GetAllowWebsocketUpgrade returns the AllowWebsocketUpgrade field if non-nil, zero value otherwise.

### GetAllowWebsocketUpgradeOk

`func (o *ProxyHostGetResponse) GetAllowWebsocketUpgradeOk() (*boolAsInt, bool)`

GetAllowWebsocketUpgradeOk returns a tuple with the AllowWebsocketUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebsocketUpgrade

`func (o *ProxyHostGetResponse) SetAllowWebsocketUpgrade(v boolAsInt)`

SetAllowWebsocketUpgrade sets AllowWebsocketUpgrade field to given value.

### HasAllowWebsocketUpgrade

`func (o *ProxyHostGetResponse) HasAllowWebsocketUpgrade() bool`

HasAllowWebsocketUpgrade returns a boolean if a field has been set.

### GetAccessListId

`func (o *ProxyHostGetResponse) GetAccessListId() int64`

GetAccessListId returns the AccessListId field if non-nil, zero value otherwise.

### GetAccessListIdOk

`func (o *ProxyHostGetResponse) GetAccessListIdOk() (*int64, bool)`

GetAccessListIdOk returns a tuple with the AccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListId

`func (o *ProxyHostGetResponse) SetAccessListId(v int64)`

SetAccessListId sets AccessListId field to given value.

### HasAccessListId

`func (o *ProxyHostGetResponse) HasAccessListId() bool`

HasAccessListId returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *ProxyHostGetResponse) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *ProxyHostGetResponse) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *ProxyHostGetResponse) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *ProxyHostGetResponse) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *ProxyHostGetResponse) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProxyHostGetResponse) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProxyHostGetResponse) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProxyHostGetResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *ProxyHostGetResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProxyHostGetResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProxyHostGetResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProxyHostGetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLocations

`func (o *ProxyHostGetResponse) GetLocations() []PutProxyHostUpdateRequestLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ProxyHostGetResponse) GetLocationsOk() (*[]PutProxyHostUpdateRequestLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ProxyHostGetResponse) SetLocations(v []PutProxyHostUpdateRequestLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ProxyHostGetResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


