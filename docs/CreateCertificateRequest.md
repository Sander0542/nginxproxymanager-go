# CreateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**NiceName** | Pointer to **string** | Nice Name for the custom certificate | [optional] 
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**Meta** | Pointer to [**GetCertificates200ResponseInnerMeta**](GetCertificates200ResponseInnerMeta.md) |  | [optional] 

## Methods

### NewCreateCertificateRequest

`func NewCreateCertificateRequest(provider string, ) *CreateCertificateRequest`

NewCreateCertificateRequest instantiates a new CreateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCertificateRequestWithDefaults

`func NewCreateCertificateRequestWithDefaults() *CreateCertificateRequest`

NewCreateCertificateRequestWithDefaults instantiates a new CreateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateCertificateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateCertificateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateCertificateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetNiceName

`func (o *CreateCertificateRequest) GetNiceName() string`

GetNiceName returns the NiceName field if non-nil, zero value otherwise.

### GetNiceNameOk

`func (o *CreateCertificateRequest) GetNiceNameOk() (*string, bool)`

GetNiceNameOk returns a tuple with the NiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiceName

`func (o *CreateCertificateRequest) SetNiceName(v string)`

SetNiceName sets NiceName field to given value.

### HasNiceName

`func (o *CreateCertificateRequest) HasNiceName() bool`

HasNiceName returns a boolean if a field has been set.

### GetDomainNames

`func (o *CreateCertificateRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *CreateCertificateRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *CreateCertificateRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *CreateCertificateRequest) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetMeta

`func (o *CreateCertificateRequest) GetMeta() GetCertificates200ResponseInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateCertificateRequest) GetMetaOk() (*GetCertificates200ResponseInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateCertificateRequest) SetMeta(v GetCertificates200ResponseInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateCertificateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


