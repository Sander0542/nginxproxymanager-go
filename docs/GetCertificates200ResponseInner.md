# GetCertificates200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**OwnerUserId** | **int64** | User ID | 
**Provider** | **string** |  | 
**NiceName** | **string** | Nice Name for the custom certificate | 
**DomainNames** | **[]string** | Domain Names separated by a comma | 
**ExpiresOn** | **string** | Date and time of expiration | [readonly] 
**Owner** | Pointer to [**GetAccessLists200ResponseInnerOwner**](GetAccessLists200ResponseInnerOwner.md) |  | [optional] 
**Meta** | [**GetCertificates200ResponseInnerMeta**](GetCertificates200ResponseInnerMeta.md) |  | 

## Methods

### NewGetCertificates200ResponseInner

`func NewGetCertificates200ResponseInner(id int64, createdOn string, modifiedOn string, ownerUserId int64, provider string, niceName string, domainNames []string, expiresOn string, meta GetCertificates200ResponseInnerMeta, ) *GetCertificates200ResponseInner`

NewGetCertificates200ResponseInner instantiates a new GetCertificates200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertificates200ResponseInnerWithDefaults

`func NewGetCertificates200ResponseInnerWithDefaults() *GetCertificates200ResponseInner`

NewGetCertificates200ResponseInnerWithDefaults instantiates a new GetCertificates200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCertificates200ResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCertificates200ResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCertificates200ResponseInner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetCertificates200ResponseInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetCertificates200ResponseInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetCertificates200ResponseInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetCertificates200ResponseInner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetCertificates200ResponseInner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetCertificates200ResponseInner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *GetCertificates200ResponseInner) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *GetCertificates200ResponseInner) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *GetCertificates200ResponseInner) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetProvider

`func (o *GetCertificates200ResponseInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetCertificates200ResponseInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetCertificates200ResponseInner) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetNiceName

`func (o *GetCertificates200ResponseInner) GetNiceName() string`

GetNiceName returns the NiceName field if non-nil, zero value otherwise.

### GetNiceNameOk

`func (o *GetCertificates200ResponseInner) GetNiceNameOk() (*string, bool)`

GetNiceNameOk returns a tuple with the NiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiceName

`func (o *GetCertificates200ResponseInner) SetNiceName(v string)`

SetNiceName sets NiceName field to given value.


### GetDomainNames

`func (o *GetCertificates200ResponseInner) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *GetCertificates200ResponseInner) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *GetCertificates200ResponseInner) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetExpiresOn

`func (o *GetCertificates200ResponseInner) GetExpiresOn() string`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *GetCertificates200ResponseInner) GetExpiresOnOk() (*string, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *GetCertificates200ResponseInner) SetExpiresOn(v string)`

SetExpiresOn sets ExpiresOn field to given value.


### GetOwner

`func (o *GetCertificates200ResponseInner) GetOwner() GetAccessLists200ResponseInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetCertificates200ResponseInner) GetOwnerOk() (*GetAccessLists200ResponseInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetCertificates200ResponseInner) SetOwner(v GetAccessLists200ResponseInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetCertificates200ResponseInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetMeta

`func (o *GetCertificates200ResponseInner) GetMeta() GetCertificates200ResponseInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCertificates200ResponseInner) GetMetaOk() (*GetCertificates200ResponseInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCertificates200ResponseInner) SetMeta(v GetCertificates200ResponseInnerMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


