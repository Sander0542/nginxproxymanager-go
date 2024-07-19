# GetAccessLists200ResponseInner

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

### NewGetAccessLists200ResponseInner

`func NewGetAccessLists200ResponseInner(id int64, createdOn string, modifiedOn string, ownerUserId int64, name string, satisfyAny bool, passAuth bool, meta map[string]interface{}, ) *GetAccessLists200ResponseInner`

NewGetAccessLists200ResponseInner instantiates a new GetAccessLists200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessLists200ResponseInnerWithDefaults

`func NewGetAccessLists200ResponseInnerWithDefaults() *GetAccessLists200ResponseInner`

NewGetAccessLists200ResponseInnerWithDefaults instantiates a new GetAccessLists200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessLists200ResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessLists200ResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessLists200ResponseInner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAccessLists200ResponseInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAccessLists200ResponseInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAccessLists200ResponseInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAccessLists200ResponseInner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAccessLists200ResponseInner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAccessLists200ResponseInner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *GetAccessLists200ResponseInner) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *GetAccessLists200ResponseInner) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *GetAccessLists200ResponseInner) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetName

`func (o *GetAccessLists200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAccessLists200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAccessLists200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetSatisfyAny

`func (o *GetAccessLists200ResponseInner) GetSatisfyAny() bool`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *GetAccessLists200ResponseInner) GetSatisfyAnyOk() (*bool, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *GetAccessLists200ResponseInner) SetSatisfyAny(v bool)`

SetSatisfyAny sets SatisfyAny field to given value.


### GetPassAuth

`func (o *GetAccessLists200ResponseInner) GetPassAuth() bool`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *GetAccessLists200ResponseInner) GetPassAuthOk() (*bool, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *GetAccessLists200ResponseInner) SetPassAuth(v bool)`

SetPassAuth sets PassAuth field to given value.


### GetMeta

`func (o *GetAccessLists200ResponseInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAccessLists200ResponseInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAccessLists200ResponseInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetOwner

`func (o *GetAccessLists200ResponseInner) GetOwner() GetAccessLists200ResponseInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetAccessLists200ResponseInner) GetOwnerOk() (*GetAccessLists200ResponseInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetAccessLists200ResponseInner) SetOwner(v GetAccessLists200ResponseInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetAccessLists200ResponseInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetItems

`func (o *GetAccessLists200ResponseInner) GetItems() []GetAccessLists200ResponseInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetAccessLists200ResponseInner) GetItemsOk() (*[]GetAccessLists200ResponseInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetAccessLists200ResponseInner) SetItems(v []GetAccessLists200ResponseInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetAccessLists200ResponseInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetClients

`func (o *GetAccessLists200ResponseInner) GetClients() []GetAccessLists200ResponseInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *GetAccessLists200ResponseInner) GetClientsOk() (*[]GetAccessLists200ResponseInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *GetAccessLists200ResponseInner) SetClients(v []GetAccessLists200ResponseInnerClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *GetAccessLists200ResponseInner) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


