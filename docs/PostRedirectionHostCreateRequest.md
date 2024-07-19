# PostRedirectionHostCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | **[]string** | Domain Names separated by a comma | 
**ForwardHttpCode** | **int64** | Redirect HTTP Status Code | 
**ForwardScheme** | **string** | RFC Protocol | 
**ForwardDomainName** | **string** | Domain Name | 
**PreservePath** | Pointer to **boolAsInt** | Should the path be preserved | [optional] 
**CertificateId** | Pointer to [**PutProxyHostUpdateRequestCertificateId**](PutProxyHostUpdateRequestCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **boolAsInt** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**Http2Support** | Pointer to **boolAsInt** | HTTP2 Protocol Support | [optional] 
**BlockExploits** | Pointer to **boolAsInt** | Should we block common exploits | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPostRedirectionHostCreateRequest

`func NewPostRedirectionHostCreateRequest(domainNames []string, forwardHttpCode int64, forwardScheme string, forwardDomainName string, ) *PostRedirectionHostCreateRequest`

NewPostRedirectionHostCreateRequest instantiates a new PostRedirectionHostCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRedirectionHostCreateRequestWithDefaults

`func NewPostRedirectionHostCreateRequestWithDefaults() *PostRedirectionHostCreateRequest`

NewPostRedirectionHostCreateRequestWithDefaults instantiates a new PostRedirectionHostCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainNames

`func (o *PostRedirectionHostCreateRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *PostRedirectionHostCreateRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *PostRedirectionHostCreateRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetForwardHttpCode

`func (o *PostRedirectionHostCreateRequest) GetForwardHttpCode() int64`

GetForwardHttpCode returns the ForwardHttpCode field if non-nil, zero value otherwise.

### GetForwardHttpCodeOk

`func (o *PostRedirectionHostCreateRequest) GetForwardHttpCodeOk() (*int64, bool)`

GetForwardHttpCodeOk returns a tuple with the ForwardHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHttpCode

`func (o *PostRedirectionHostCreateRequest) SetForwardHttpCode(v int64)`

SetForwardHttpCode sets ForwardHttpCode field to given value.


### GetForwardScheme

`func (o *PostRedirectionHostCreateRequest) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *PostRedirectionHostCreateRequest) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *PostRedirectionHostCreateRequest) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.


### GetForwardDomainName

`func (o *PostRedirectionHostCreateRequest) GetForwardDomainName() string`

GetForwardDomainName returns the ForwardDomainName field if non-nil, zero value otherwise.

### GetForwardDomainNameOk

`func (o *PostRedirectionHostCreateRequest) GetForwardDomainNameOk() (*string, bool)`

GetForwardDomainNameOk returns a tuple with the ForwardDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDomainName

`func (o *PostRedirectionHostCreateRequest) SetForwardDomainName(v string)`

SetForwardDomainName sets ForwardDomainName field to given value.


### GetPreservePath

`func (o *PostRedirectionHostCreateRequest) GetPreservePath() boolAsInt`

GetPreservePath returns the PreservePath field if non-nil, zero value otherwise.

### GetPreservePathOk

`func (o *PostRedirectionHostCreateRequest) GetPreservePathOk() (*boolAsInt, bool)`

GetPreservePathOk returns a tuple with the PreservePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservePath

`func (o *PostRedirectionHostCreateRequest) SetPreservePath(v boolAsInt)`

SetPreservePath sets PreservePath field to given value.

### HasPreservePath

`func (o *PostRedirectionHostCreateRequest) HasPreservePath() bool`

HasPreservePath returns a boolean if a field has been set.

### GetCertificateId

`func (o *PostRedirectionHostCreateRequest) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *PostRedirectionHostCreateRequest) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *PostRedirectionHostCreateRequest) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *PostRedirectionHostCreateRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *PostRedirectionHostCreateRequest) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *PostRedirectionHostCreateRequest) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *PostRedirectionHostCreateRequest) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *PostRedirectionHostCreateRequest) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *PostRedirectionHostCreateRequest) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *PostRedirectionHostCreateRequest) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *PostRedirectionHostCreateRequest) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *PostRedirectionHostCreateRequest) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *PostRedirectionHostCreateRequest) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *PostRedirectionHostCreateRequest) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *PostRedirectionHostCreateRequest) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *PostRedirectionHostCreateRequest) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *PostRedirectionHostCreateRequest) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *PostRedirectionHostCreateRequest) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *PostRedirectionHostCreateRequest) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *PostRedirectionHostCreateRequest) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *PostRedirectionHostCreateRequest) GetBlockExploits() boolAsInt`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *PostRedirectionHostCreateRequest) GetBlockExploitsOk() (*boolAsInt, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *PostRedirectionHostCreateRequest) SetBlockExploits(v boolAsInt)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *PostRedirectionHostCreateRequest) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *PostRedirectionHostCreateRequest) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *PostRedirectionHostCreateRequest) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *PostRedirectionHostCreateRequest) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *PostRedirectionHostCreateRequest) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetMeta

`func (o *PostRedirectionHostCreateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PostRedirectionHostCreateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PostRedirectionHostCreateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PostRedirectionHostCreateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


