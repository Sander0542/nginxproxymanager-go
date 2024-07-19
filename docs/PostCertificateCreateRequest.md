# PostCertificateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**NiceName** | Pointer to **string** | Nice Name for the custom certificate | [optional] 
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**Meta** | Pointer to [**PostCertificateCreateRequestMeta**](PostCertificateCreateRequestMeta.md) |  | [optional] 

## Methods

### NewPostCertificateCreateRequest

`func NewPostCertificateCreateRequest(provider string, ) *PostCertificateCreateRequest`

NewPostCertificateCreateRequest instantiates a new PostCertificateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCertificateCreateRequestWithDefaults

`func NewPostCertificateCreateRequestWithDefaults() *PostCertificateCreateRequest`

NewPostCertificateCreateRequestWithDefaults instantiates a new PostCertificateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *PostCertificateCreateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PostCertificateCreateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PostCertificateCreateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetNiceName

`func (o *PostCertificateCreateRequest) GetNiceName() string`

GetNiceName returns the NiceName field if non-nil, zero value otherwise.

### GetNiceNameOk

`func (o *PostCertificateCreateRequest) GetNiceNameOk() (*string, bool)`

GetNiceNameOk returns a tuple with the NiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiceName

`func (o *PostCertificateCreateRequest) SetNiceName(v string)`

SetNiceName sets NiceName field to given value.

### HasNiceName

`func (o *PostCertificateCreateRequest) HasNiceName() bool`

HasNiceName returns a boolean if a field has been set.

### GetDomainNames

`func (o *PostCertificateCreateRequest) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *PostCertificateCreateRequest) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *PostCertificateCreateRequest) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *PostCertificateCreateRequest) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetMeta

`func (o *PostCertificateCreateRequest) GetMeta() PostCertificateCreateRequestMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PostCertificateCreateRequest) GetMetaOk() (*PostCertificateCreateRequestMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PostCertificateCreateRequest) SetMeta(v PostCertificateCreateRequestMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PostCertificateCreateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


