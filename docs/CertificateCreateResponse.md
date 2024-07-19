# CertificateCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** | Date and time of creation | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** | Date and time of last update | [optional] [readonly] 
**Provider** | Pointer to **string** |  | [optional] 
**NiceName** | Pointer to **string** | Nice Name for the custom certificate | [optional] 
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**ExpiresOn** | Pointer to **time.Time** | Date and time of expiration | [optional] [readonly] 
**Meta** | Pointer to [**PostCertificateCreateRequestMeta**](PostCertificateCreateRequestMeta.md) |  | [optional] 

## Methods

### NewCertificateCreateResponse

`func NewCertificateCreateResponse() *CertificateCreateResponse`

NewCertificateCreateResponse instantiates a new CertificateCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCreateResponseWithDefaults

`func NewCertificateCreateResponseWithDefaults() *CertificateCreateResponse`

NewCertificateCreateResponseWithDefaults instantiates a new CertificateCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateCreateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateCreateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateCreateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *CertificateCreateResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *CertificateCreateResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *CertificateCreateResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *CertificateCreateResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *CertificateCreateResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *CertificateCreateResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *CertificateCreateResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *CertificateCreateResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetProvider

`func (o *CertificateCreateResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CertificateCreateResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CertificateCreateResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CertificateCreateResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetNiceName

`func (o *CertificateCreateResponse) GetNiceName() string`

GetNiceName returns the NiceName field if non-nil, zero value otherwise.

### GetNiceNameOk

`func (o *CertificateCreateResponse) GetNiceNameOk() (*string, bool)`

GetNiceNameOk returns a tuple with the NiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiceName

`func (o *CertificateCreateResponse) SetNiceName(v string)`

SetNiceName sets NiceName field to given value.

### HasNiceName

`func (o *CertificateCreateResponse) HasNiceName() bool`

HasNiceName returns a boolean if a field has been set.

### GetDomainNames

`func (o *CertificateCreateResponse) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *CertificateCreateResponse) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *CertificateCreateResponse) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *CertificateCreateResponse) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetExpiresOn

`func (o *CertificateCreateResponse) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *CertificateCreateResponse) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *CertificateCreateResponse) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *CertificateCreateResponse) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetMeta

`func (o *CertificateCreateResponse) GetMeta() PostCertificateCreateRequestMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateCreateResponse) GetMetaOk() (*PostCertificateCreateRequestMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateCreateResponse) SetMeta(v PostCertificateCreateRequestMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateCreateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


