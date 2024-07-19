# PutAccessListUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Access List | [optional] 
**SatisfyAny** | Pointer to **boolAsInt** |  | [optional] 
**PassAuth** | Pointer to **boolAsInt** |  | [optional] 
**Items** | Pointer to [**[]PutAccessListUpdateRequestItemsInner**](PutAccessListUpdateRequestItemsInner.md) |  | [optional] 
**Clients** | Pointer to [**[]PutAccessListUpdateRequestClientsInner**](PutAccessListUpdateRequestClientsInner.md) |  | [optional] 

## Methods

### NewPutAccessListUpdateRequest

`func NewPutAccessListUpdateRequest() *PutAccessListUpdateRequest`

NewPutAccessListUpdateRequest instantiates a new PutAccessListUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutAccessListUpdateRequestWithDefaults

`func NewPutAccessListUpdateRequestWithDefaults() *PutAccessListUpdateRequest`

NewPutAccessListUpdateRequestWithDefaults instantiates a new PutAccessListUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutAccessListUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutAccessListUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutAccessListUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutAccessListUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSatisfyAny

`func (o *PutAccessListUpdateRequest) GetSatisfyAny() boolAsInt`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *PutAccessListUpdateRequest) GetSatisfyAnyOk() (*boolAsInt, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *PutAccessListUpdateRequest) SetSatisfyAny(v boolAsInt)`

SetSatisfyAny sets SatisfyAny field to given value.

### HasSatisfyAny

`func (o *PutAccessListUpdateRequest) HasSatisfyAny() bool`

HasSatisfyAny returns a boolean if a field has been set.

### GetPassAuth

`func (o *PutAccessListUpdateRequest) GetPassAuth() boolAsInt`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *PutAccessListUpdateRequest) GetPassAuthOk() (*boolAsInt, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *PutAccessListUpdateRequest) SetPassAuth(v boolAsInt)`

SetPassAuth sets PassAuth field to given value.

### HasPassAuth

`func (o *PutAccessListUpdateRequest) HasPassAuth() bool`

HasPassAuth returns a boolean if a field has been set.

### GetItems

`func (o *PutAccessListUpdateRequest) GetItems() []PutAccessListUpdateRequestItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PutAccessListUpdateRequest) GetItemsOk() (*[]PutAccessListUpdateRequestItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PutAccessListUpdateRequest) SetItems(v []PutAccessListUpdateRequestItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *PutAccessListUpdateRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetClients

`func (o *PutAccessListUpdateRequest) GetClients() []PutAccessListUpdateRequestClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *PutAccessListUpdateRequest) GetClientsOk() (*[]PutAccessListUpdateRequestClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *PutAccessListUpdateRequest) SetClients(v []PutAccessListUpdateRequestClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *PutAccessListUpdateRequest) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


