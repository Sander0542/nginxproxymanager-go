# CreateAccessList201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**OwnerUserId** | **int64** | User ID | 
**Name** | **string** |  | 
**SatisfyAny** | **bool** |  | 
**PassAuth** | **bool** |  | 
**Meta** | **map[string]interface{}** |  | 
**Owner** | Pointer to [**GetAccessLists200ResponseInnerOwner**](GetAccessLists200ResponseInnerOwner.md) |  | [optional] 
**Items** | Pointer to [**[]GetAccessLists200ResponseInnerItemsInner**](GetAccessLists200ResponseInnerItemsInner.md) |  | [optional] 
**Clients** | Pointer to [**[]GetAccessLists200ResponseInnerClientsInner**](GetAccessLists200ResponseInnerClientsInner.md) |  | [optional] 

## Methods

### NewCreateAccessList201Response

`func NewCreateAccessList201Response(id int64, createdOn string, modifiedOn string, ownerUserId int64, name string, satisfyAny bool, passAuth bool, meta map[string]interface{}, ) *CreateAccessList201Response`

NewCreateAccessList201Response instantiates a new CreateAccessList201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessList201ResponseWithDefaults

`func NewCreateAccessList201ResponseWithDefaults() *CreateAccessList201Response`

NewCreateAccessList201ResponseWithDefaults instantiates a new CreateAccessList201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAccessList201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccessList201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccessList201Response) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *CreateAccessList201Response) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *CreateAccessList201Response) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *CreateAccessList201Response) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *CreateAccessList201Response) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *CreateAccessList201Response) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *CreateAccessList201Response) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *CreateAccessList201Response) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *CreateAccessList201Response) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *CreateAccessList201Response) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetName

`func (o *CreateAccessList201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccessList201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccessList201Response) SetName(v string)`

SetName sets Name field to given value.


### GetSatisfyAny

`func (o *CreateAccessList201Response) GetSatisfyAny() bool`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *CreateAccessList201Response) GetSatisfyAnyOk() (*bool, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *CreateAccessList201Response) SetSatisfyAny(v bool)`

SetSatisfyAny sets SatisfyAny field to given value.


### GetPassAuth

`func (o *CreateAccessList201Response) GetPassAuth() bool`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *CreateAccessList201Response) GetPassAuthOk() (*bool, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *CreateAccessList201Response) SetPassAuth(v bool)`

SetPassAuth sets PassAuth field to given value.


### GetMeta

`func (o *CreateAccessList201Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateAccessList201Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateAccessList201Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetOwner

`func (o *CreateAccessList201Response) GetOwner() GetAccessLists200ResponseInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateAccessList201Response) GetOwnerOk() (*GetAccessLists200ResponseInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateAccessList201Response) SetOwner(v GetAccessLists200ResponseInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateAccessList201Response) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetItems

`func (o *CreateAccessList201Response) GetItems() []GetAccessLists200ResponseInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateAccessList201Response) GetItemsOk() (*[]GetAccessLists200ResponseInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateAccessList201Response) SetItems(v []GetAccessLists200ResponseInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateAccessList201Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetClients

`func (o *CreateAccessList201Response) GetClients() []GetAccessLists200ResponseInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *CreateAccessList201Response) GetClientsOk() (*[]GetAccessLists200ResponseInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *CreateAccessList201Response) SetClients(v []GetAccessLists200ResponseInnerClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *CreateAccessList201Response) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


