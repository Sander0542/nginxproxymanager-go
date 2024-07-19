# Create404HostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | **[]string** | Domain Names separated by a comma | 
**CertificateId** | Pointer to [**GetProxyHosts200ResponseInnerCertificateId**](GetProxyHosts200ResponseInnerCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **bool** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **bool** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **bool** | Is HSTS applicable to all subdomains | [optional] 
**Http2Support** | Pointer to **bool** | HTTP2 Protocol Support | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreate404HostRequest

`func NewCreate404HostRequest(domainNames []string, ) *Create404HostRequest`

NewCreate404HostRequest instantiates a new Create404HostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreate404HostRequestWithDefaults

`func NewCreate404HostRequestWithDefaults() *Create404HostRequest`

NewCreate404HostRequestWithDefaults instantiates a new Create404HostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainNames

`func (o *Create404HostRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *Create404HostRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *Create404HostRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetCertificateId

`func (o *Create404HostRequest) GetCertificateId() GetProxyHosts200ResponseInnerCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *Create404HostRequest) GetCertificateIdOk() (*GetProxyHosts200ResponseInnerCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *Create404HostRequest) SetCertificateId(v GetProxyHosts200ResponseInnerCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *Create404HostRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *Create404HostRequest) GetSslForced() bool`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *Create404HostRequest) GetSslForcedOk() (*bool, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *Create404HostRequest) SetSslForced(v bool)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *Create404HostRequest) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *Create404HostRequest) GetHstsEnabled() bool`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *Create404HostRequest) GetHstsEnabledOk() (*bool, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *Create404HostRequest) SetHstsEnabled(v bool)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *Create404HostRequest) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *Create404HostRequest) GetHstsSubdomains() bool`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *Create404HostRequest) GetHstsSubdomainsOk() (*bool, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *Create404HostRequest) SetHstsSubdomains(v bool)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *Create404HostRequest) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *Create404HostRequest) GetHttp2Support() bool`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *Create404HostRequest) GetHttp2SupportOk() (*bool, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *Create404HostRequest) SetHttp2Support(v bool)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *Create404HostRequest) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *Create404HostRequest) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *Create404HostRequest) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *Create404HostRequest) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *Create404HostRequest) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetMeta

`func (o *Create404HostRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Create404HostRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Create404HostRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Create404HostRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


