# PostAccessListCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Access List | 
**SatisfyAny** | Pointer to **boolAsInt** |  | [optional] 
**PassAuth** | Pointer to **boolAsInt** |  | [optional] 
**Items** | Pointer to [**[]PostAccessListCreateRequestItemsInner**](PostAccessListCreateRequestItemsInner.md) |  | [optional] 
**Clients** | Pointer to [**[]PutAccessListUpdateRequestClientsInner**](PutAccessListUpdateRequestClientsInner.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPostAccessListCreateRequest

`func NewPostAccessListCreateRequest(name string, ) *PostAccessListCreateRequest`

NewPostAccessListCreateRequest instantiates a new PostAccessListCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAccessListCreateRequestWithDefaults

`func NewPostAccessListCreateRequestWithDefaults() *PostAccessListCreateRequest`

NewPostAccessListCreateRequestWithDefaults instantiates a new PostAccessListCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostAccessListCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostAccessListCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostAccessListCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSatisfyAny

`func (o *PostAccessListCreateRequest) GetSatisfyAny() boolAsInt`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *PostAccessListCreateRequest) GetSatisfyAnyOk() (*boolAsInt, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *PostAccessListCreateRequest) SetSatisfyAny(v boolAsInt)`

SetSatisfyAny sets SatisfyAny field to given value.

### HasSatisfyAny

`func (o *PostAccessListCreateRequest) HasSatisfyAny() bool`

HasSatisfyAny returns a boolean if a field has been set.

### GetPassAuth

`func (o *PostAccessListCreateRequest) GetPassAuth() boolAsInt`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *PostAccessListCreateRequest) GetPassAuthOk() (*boolAsInt, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *PostAccessListCreateRequest) SetPassAuth(v boolAsInt)`

SetPassAuth sets PassAuth field to given value.

### HasPassAuth

`func (o *PostAccessListCreateRequest) HasPassAuth() bool`

HasPassAuth returns a boolean if a field has been set.

### GetItems

`func (o *PostAccessListCreateRequest) GetItems() []PostAccessListCreateRequestItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PostAccessListCreateRequest) GetItemsOk() (*[]PostAccessListCreateRequestItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PostAccessListCreateRequest) SetItems(v []PostAccessListCreateRequestItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *PostAccessListCreateRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetClients

`func (o *PostAccessListCreateRequest) GetClients() []PutAccessListUpdateRequestClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *PostAccessListCreateRequest) GetClientsOk() (*[]PutAccessListUpdateRequestClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *PostAccessListCreateRequest) SetClients(v []PutAccessListUpdateRequestClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *PostAccessListCreateRequest) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetMeta

`func (o *PostAccessListCreateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PostAccessListCreateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PostAccessListCreateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PostAccessListCreateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


