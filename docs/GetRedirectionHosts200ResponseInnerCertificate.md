# GetRedirectionHosts200ResponseInnerCertificate

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
**Owner** | Pointer to [**GetAuditLogs200ResponseInnerUser**](GetAuditLogs200ResponseInnerUser.md) |  | [optional] 
**Meta** | [**GetCertificates200ResponseInnerMeta**](GetCertificates200ResponseInnerMeta.md) |  | 

## Methods

### NewGetRedirectionHosts200ResponseInnerCertificate

`func NewGetRedirectionHosts200ResponseInnerCertificate(id int64, createdOn string, modifiedOn string, ownerUserId int64, provider string, niceName string, domainNames []string, expiresOn string, meta GetCertificates200ResponseInnerMeta, ) *GetRedirectionHosts200ResponseInnerCertificate`

NewGetRedirectionHosts200ResponseInnerCertificate instantiates a new GetRedirectionHosts200ResponseInnerCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedirectionHosts200ResponseInnerCertificateWithDefaults

`func NewGetRedirectionHosts200ResponseInnerCertificateWithDefaults() *GetRedirectionHosts200ResponseInnerCertificate`

NewGetRedirectionHosts200ResponseInnerCertificateWithDefaults instantiates a new GetRedirectionHosts200ResponseInnerCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetProvider

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetNiceName

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetNiceName() string`

GetNiceName returns the NiceName field if non-nil, zero value otherwise.

### GetNiceNameOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetNiceNameOk() (*string, bool)`

GetNiceNameOk returns a tuple with the NiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiceName

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetNiceName(v string)`

SetNiceName sets NiceName field to given value.


### GetDomainNames

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetExpiresOn

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetExpiresOn() string`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetExpiresOnOk() (*string, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetExpiresOn(v string)`

SetExpiresOn sets ExpiresOn field to given value.


### GetOwner

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetOwner() GetAuditLogs200ResponseInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetOwnerOk() (*GetAuditLogs200ResponseInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetOwner(v GetAuditLogs200ResponseInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetRedirectionHosts200ResponseInnerCertificate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetMeta

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetMeta() GetCertificates200ResponseInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRedirectionHosts200ResponseInnerCertificate) GetMetaOk() (*GetCertificates200ResponseInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRedirectionHosts200ResponseInnerCertificate) SetMeta(v GetCertificates200ResponseInnerMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


