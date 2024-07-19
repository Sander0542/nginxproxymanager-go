# CreateRedirectionHostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | **[]string** | Domain Names separated by a comma | 
**ForwardHttpCode** | **int64** | Redirect HTTP Status Code | 
**ForwardScheme** | **string** |  | 
**ForwardDomainName** | **string** | Domain Name | 
**PreservePath** | Pointer to **bool** | Should the path be preserved | [optional] 
**CertificateId** | Pointer to [**GetProxyHosts200ResponseInnerCertificateId**](GetProxyHosts200ResponseInnerCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **bool** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **bool** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **bool** | Is HSTS applicable to all subdomains | [optional] 
**Http2Support** | Pointer to **bool** | HTTP2 Protocol Support | [optional] 
**BlockExploits** | Pointer to **bool** | Should we block common exploits | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateRedirectionHostRequest

`func NewCreateRedirectionHostRequest(domainNames []string, forwardHttpCode int64, forwardScheme string, forwardDomainName string, ) *CreateRedirectionHostRequest`

NewCreateRedirectionHostRequest instantiates a new CreateRedirectionHostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRedirectionHostRequestWithDefaults

`func NewCreateRedirectionHostRequestWithDefaults() *CreateRedirectionHostRequest`

NewCreateRedirectionHostRequestWithDefaults instantiates a new CreateRedirectionHostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainNames

`func (o *CreateRedirectionHostRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *CreateRedirectionHostRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *CreateRedirectionHostRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetForwardHttpCode

`func (o *CreateRedirectionHostRequest) GetForwardHttpCode() int64`

GetForwardHttpCode returns the ForwardHttpCode field if non-nil, zero value otherwise.

### GetForwardHttpCodeOk

`func (o *CreateRedirectionHostRequest) GetForwardHttpCodeOk() (*int64, bool)`

GetForwardHttpCodeOk returns a tuple with the ForwardHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHttpCode

`func (o *CreateRedirectionHostRequest) SetForwardHttpCode(v int64)`

SetForwardHttpCode sets ForwardHttpCode field to given value.


### GetForwardScheme

`func (o *CreateRedirectionHostRequest) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *CreateRedirectionHostRequest) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *CreateRedirectionHostRequest) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.


### GetForwardDomainName

`func (o *CreateRedirectionHostRequest) GetForwardDomainName() string`

GetForwardDomainName returns the ForwardDomainName field if non-nil, zero value otherwise.

### GetForwardDomainNameOk

`func (o *CreateRedirectionHostRequest) GetForwardDomainNameOk() (*string, bool)`

GetForwardDomainNameOk returns a tuple with the ForwardDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDomainName

`func (o *CreateRedirectionHostRequest) SetForwardDomainName(v string)`

SetForwardDomainName sets ForwardDomainName field to given value.


### GetPreservePath

`func (o *CreateRedirectionHostRequest) GetPreservePath() bool`

GetPreservePath returns the PreservePath field if non-nil, zero value otherwise.

### GetPreservePathOk

`func (o *CreateRedirectionHostRequest) GetPreservePathOk() (*bool, bool)`

GetPreservePathOk returns a tuple with the PreservePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservePath

`func (o *CreateRedirectionHostRequest) SetPreservePath(v bool)`

SetPreservePath sets PreservePath field to given value.

### HasPreservePath

`func (o *CreateRedirectionHostRequest) HasPreservePath() bool`

HasPreservePath returns a boolean if a field has been set.

### GetCertificateId

`func (o *CreateRedirectionHostRequest) GetCertificateId() GetProxyHosts200ResponseInnerCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CreateRedirectionHostRequest) GetCertificateIdOk() (*GetProxyHosts200ResponseInnerCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CreateRedirectionHostRequest) SetCertificateId(v GetProxyHosts200ResponseInnerCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CreateRedirectionHostRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *CreateRedirectionHostRequest) GetSslForced() bool`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *CreateRedirectionHostRequest) GetSslForcedOk() (*bool, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *CreateRedirectionHostRequest) SetSslForced(v bool)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *CreateRedirectionHostRequest) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *CreateRedirectionHostRequest) GetHstsEnabled() bool`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *CreateRedirectionHostRequest) GetHstsEnabledOk() (*bool, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *CreateRedirectionHostRequest) SetHstsEnabled(v bool)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *CreateRedirectionHostRequest) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *CreateRedirectionHostRequest) GetHstsSubdomains() bool`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *CreateRedirectionHostRequest) GetHstsSubdomainsOk() (*bool, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *CreateRedirectionHostRequest) SetHstsSubdomains(v bool)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *CreateRedirectionHostRequest) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *CreateRedirectionHostRequest) GetHttp2Support() bool`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *CreateRedirectionHostRequest) GetHttp2SupportOk() (*bool, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *CreateRedirectionHostRequest) SetHttp2Support(v bool)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *CreateRedirectionHostRequest) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetBlockExploits

`func (o *CreateRedirectionHostRequest) GetBlockExploits() bool`

GetBlockExploits returns the BlockExploits field if non-nil, zero value otherwise.

### GetBlockExploitsOk

`func (o *CreateRedirectionHostRequest) GetBlockExploitsOk() (*bool, bool)`

GetBlockExploitsOk returns a tuple with the BlockExploits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExploits

`func (o *CreateRedirectionHostRequest) SetBlockExploits(v bool)`

SetBlockExploits sets BlockExploits field to given value.

### HasBlockExploits

`func (o *CreateRedirectionHostRequest) HasBlockExploits() bool`

HasBlockExploits returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *CreateRedirectionHostRequest) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *CreateRedirectionHostRequest) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *CreateRedirectionHostRequest) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *CreateRedirectionHostRequest) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetMeta

`func (o *CreateRedirectionHostRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateRedirectionHostRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateRedirectionHostRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateRedirectionHostRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


