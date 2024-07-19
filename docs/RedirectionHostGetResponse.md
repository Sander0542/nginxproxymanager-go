# RedirectionHostGetResponse

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

### NewRedirectionHostGetResponse

`func NewRedirectionHostGetResponse() *RedirectionHostGetResponse`

NewRedirectionHostGetResponse instantiates a new RedirectionHostGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectionHostGetResponseWithDefaults

`func NewRedirectionHostGetResponseWithDefaults() *RedirectionHostGetResponse`

NewRedirectionHostGetResponseWithDefaults instantiates a new RedirectionHostGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RedirectionHostGetResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedirectionHostGetResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedirectionHostGetResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RedirectionHostGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *RedirectionHostGetResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *RedirectionHostGetResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *RedirectionHostGetResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *RedirectionHostGetResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *RedirectionHostGetResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *RedirectionHostGetResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *RedirectionHostGetResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *RedirectionHostGetResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetDomainNames

`func (o *RedirectionHostGetResponse) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *RedirectionHostGetResponse) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *RedirectionHostGetResponse) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *RedirectionHostGetResponse) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetForwardHttpCode

`func (o *RedirectionHostGetResponse) GetForwardHttpCode() int64`

GetForwardHttpCode returns the ForwardHttpCode field if non-nil, zero value otherwise.

### GetForwardHttpCodeOk

`func (o *RedirectionHostGetResponse) GetForwardHttpCodeOk() (*int64, bool)`

GetForwardHttpCodeOk returns a tuple with the ForwardHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHttpCode

`func (o *RedirectionHostGetResponse) SetForwardHttpCode(v int64)`

SetForwardHttpCode sets ForwardHttpCode field to given value.

### HasForwardHttpCode

`func (o *RedirectionHostGetResponse) HasForwardHttpCode() bool`

HasForwardHttpCode returns a boolean if a field has been set.

### GetForwardScheme

`func (o *RedirectionHostGetResponse) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *RedirectionHostGetResponse) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *RedirectionHostGetResponse) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.

### HasForwardScheme

`func (o *RedirectionHostGetResponse) HasForwardScheme() bool`

HasForwardScheme returns a boolean if a field has been set.

### GetForwardDomainName

`func (o *RedirectionHostGetResponse) GetForwardDomainName() string`

GetForwardDomainName returns the ForwardDomainName field if non-nil, zero value otherwise.

### GetForwardDomainNameOk

`func (o *RedirectionHostGetResponse) GetForwardDomainNameOk() (*string, bool)`

GetForwardDomainNameOk returns a tuple with the ForwardDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDomainName

`func (o *RedirectionHostGetResponse) SetForwardDomainName(v string)`

SetForwardDomainName sets ForwardDomainName field to given value.

### HasForwardDomainName

`func (o *RedirectionHostGetResponse) HasForwardDomainName() bool`

HasForwardDomainName returns a boolean if a field has been set.

### GetPreservePath

`func (o *RedirectionHostGetResponse) GetPreservePath() boolAsInt`

GetPreservePath returns the PreservePath field if non-nil, zero value otherwise.

### GetPreservePathOk

`func (o *RedirectionHostGetResponse) GetPreservePathOk() (*boolAsInt, bool)`

GetPreservePathOk returns a tuple with the PreservePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservePath

`func (o *RedirectionHostGetResponse) SetPreservePath(v boolAsInt)`

SetPreservePath sets PreservePath field to given value.

### HasPreservePath

`func (o *RedirectionHostGetResponse) HasPreservePath() bool`

HasPreservePath returns a boolean if a field has been set.

### GetCertificateId

`func (o *RedirectionHostGetResponse) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *RedirectionHostGetResponse) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *RedirectionHostGetResponse) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *RedirectionHostGetResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *RedirectionHostGetResponse) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *RedirectionHostGetResponse) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *RedirectionHostGetResponse) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *RedirectionHostGetResponse) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *RedirectionHostGetResponse) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *RedirectionHostGetResponse) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *RedirectionHostGetResponse) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *RedirectionHostGetResponse) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *RedirectionHostGetResponse) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *RedirectionHostGetResponse) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *RedirectionHostGetResponse) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *RedirectionHostGetResponse) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *RedirectionHostGetResponse) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *RedirectionHostGetResponse) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *RedirectionHostGetResponse) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *RedirectionHostGetResponse) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *RedirectionHostGetResponse) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *RedirectionHostGetResponse) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *RedirectionHostGetResponse) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *RedirectionHostGetResponse) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *RedirectionHostGetResponse) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *RedirectionHostGetResponse) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *RedirectionHostGetResponse) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *RedirectionHostGetResponse) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *RedirectionHostGetResponse) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RedirectionHostGetResponse) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RedirectionHostGetResponse) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RedirectionHostGetResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *RedirectionHostGetResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RedirectionHostGetResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RedirectionHostGetResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RedirectionHostGetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


