# RedirectionHostCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** | Date and time of creation | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** | Date and time of last update | [optional] [readonly] 
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**ForwardHttpCode** | Pointer to **int64** | Redirect HTTP Status Code | [optional] 
**ForwardScheme** | Pointer to **string** | RFC Protocol | [optional] 
**ForwardDomainName** | Pointer to **string** | Domain Name | [optional] 
**PreservePath** | Pointer to **boolAsInt** | Should the path be preserved | [optional] 
**CertificateId** | Pointer to [**PutProxyHostUpdateRequestCertificateId**](PutProxyHostUpdateRequestCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **boolAsInt** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **boolAsInt** | Is HSTS applicable to all subdomains | [optional] 
**Http2Support** | Pointer to **boolAsInt** | HTTP2 Protocol Support | [optional] 
**BlockExploits** | Pointer to **boolAsInt** | Should we block common exploits | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **boolAsInt** | Is Enabled | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRedirectionHostCreateResponse

`func NewRedirectionHostCreateResponse() *RedirectionHostCreateResponse`

NewRedirectionHostCreateResponse instantiates a new RedirectionHostCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectionHostCreateResponseWithDefaults

`func NewRedirectionHostCreateResponseWithDefaults() *RedirectionHostCreateResponse`

NewRedirectionHostCreateResponseWithDefaults instantiates a new RedirectionHostCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RedirectionHostCreateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedirectionHostCreateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedirectionHostCreateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RedirectionHostCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *RedirectionHostCreateResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *RedirectionHostCreateResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *RedirectionHostCreateResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *RedirectionHostCreateResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *RedirectionHostCreateResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *RedirectionHostCreateResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *RedirectionHostCreateResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *RedirectionHostCreateResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetDomainNames

`func (o *RedirectionHostCreateResponse) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *RedirectionHostCreateResponse) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *RedirectionHostCreateResponse) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *RedirectionHostCreateResponse) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetForwardHttpCode

`func (o *RedirectionHostCreateResponse) GetForwardHttpCode() int64`

GetForwardHttpCode returns the ForwardHttpCode field if non-nil, zero value otherwise.

### GetForwardHttpCodeOk

`func (o *RedirectionHostCreateResponse) GetForwardHttpCodeOk() (*int64, bool)`

GetForwardHttpCodeOk returns a tuple with the ForwardHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHttpCode

`func (o *RedirectionHostCreateResponse) SetForwardHttpCode(v int64)`

SetForwardHttpCode sets ForwardHttpCode field to given value.

### HasForwardHttpCode

`func (o *RedirectionHostCreateResponse) HasForwardHttpCode() bool`

HasForwardHttpCode returns a boolean if a field has been set.

### GetForwardScheme

`func (o *RedirectionHostCreateResponse) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *RedirectionHostCreateResponse) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *RedirectionHostCreateResponse) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.

### HasForwardScheme

`func (o *RedirectionHostCreateResponse) HasForwardScheme() bool`

HasForwardScheme returns a boolean if a field has been set.

### GetForwardDomainName

`func (o *RedirectionHostCreateResponse) GetForwardDomainName() string`

GetForwardDomainName returns the ForwardDomainName field if non-nil, zero value otherwise.

### GetForwardDomainNameOk

`func (o *RedirectionHostCreateResponse) GetForwardDomainNameOk() (*string, bool)`

GetForwardDomainNameOk returns a tuple with the ForwardDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDomainName

`func (o *RedirectionHostCreateResponse) SetForwardDomainName(v string)`

SetForwardDomainName sets ForwardDomainName field to given value.

### HasForwardDomainName

`func (o *RedirectionHostCreateResponse) HasForwardDomainName() bool`

HasForwardDomainName returns a boolean if a field has been set.

### GetPreservePath

`func (o *RedirectionHostCreateResponse) GetPreservePath() boolAsInt`

GetPreservePath returns the PreservePath field if non-nil, zero value otherwise.

### GetPreservePathOk

`func (o *RedirectionHostCreateResponse) GetPreservePathOk() (*boolAsInt, bool)`

GetPreservePathOk returns a tuple with the PreservePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservePath

`func (o *RedirectionHostCreateResponse) SetPreservePath(v boolAsInt)`

SetPreservePath sets PreservePath field to given value.

### HasPreservePath

`func (o *RedirectionHostCreateResponse) HasPreservePath() bool`

HasPreservePath returns a boolean if a field has been set.

### GetCertificateId

`func (o *RedirectionHostCreateResponse) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *RedirectionHostCreateResponse) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *RedirectionHostCreateResponse) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *RedirectionHostCreateResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *RedirectionHostCreateResponse) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *RedirectionHostCreateResponse) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *RedirectionHostCreateResponse) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *RedirectionHostCreateResponse) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *RedirectionHostCreateResponse) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *RedirectionHostCreateResponse) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *RedirectionHostCreateResponse) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *RedirectionHostCreateResponse) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *RedirectionHostCreateResponse) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *RedirectionHostCreateResponse) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *RedirectionHostCreateResponse) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *RedirectionHostCreateResponse) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *RedirectionHostCreateResponse) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *RedirectionHostCreateResponse) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *RedirectionHostCreateResponse) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *RedirectionHostCreateResponse) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *RedirectionHostCreateResponse) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *RedirectionHostCreateResponse) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *RedirectionHostCreateResponse) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *RedirectionHostCreateResponse) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *RedirectionHostCreateResponse) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *RedirectionHostCreateResponse) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *RedirectionHostCreateResponse) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *RedirectionHostCreateResponse) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *RedirectionHostCreateResponse) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RedirectionHostCreateResponse) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RedirectionHostCreateResponse) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RedirectionHostCreateResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *RedirectionHostCreateResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RedirectionHostCreateResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RedirectionHostCreateResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RedirectionHostCreateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


