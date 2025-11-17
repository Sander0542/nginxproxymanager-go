# CreateAccessListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SatisfyAny** | Pointer to **bool** |  | [optional] 
**PassAuth** | Pointer to **bool** |  | [optional] 
**Items** | Pointer to [**[]CreateAccessListRequestItemsInner**](CreateAccessListRequestItemsInner.md) |  | [optional] 
**Clients** | Pointer to [**[]CreateAccessListRequestClientsInner**](CreateAccessListRequestClientsInner.md) |  | [optional] 

## Methods

### NewCreateAccessListRequest

`func NewCreateAccessListRequest(name string, ) *CreateAccessListRequest`

NewCreateAccessListRequest instantiates a new CreateAccessListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessListRequestWithDefaults

`func NewCreateAccessListRequestWithDefaults() *CreateAccessListRequest`

NewCreateAccessListRequestWithDefaults instantiates a new CreateAccessListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAccessListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccessListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccessListRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSatisfyAny

`func (o *CreateAccessListRequest) GetSatisfyAny() bool`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *CreateAccessListRequest) GetSatisfyAnyOk() (*bool, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *CreateAccessListRequest) SetSatisfyAny(v bool)`

SetSatisfyAny sets SatisfyAny field to given value.

### HasSatisfyAny

`func (o *CreateAccessListRequest) HasSatisfyAny() bool`

HasSatisfyAny returns a boolean if a field has been set.

### GetPassAuth

`func (o *CreateAccessListRequest) GetPassAuth() bool`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *CreateAccessListRequest) GetPassAuthOk() (*bool, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *CreateAccessListRequest) SetPassAuth(v bool)`

SetPassAuth sets PassAuth field to given value.

### HasPassAuth

`func (o *CreateAccessListRequest) HasPassAuth() bool`

HasPassAuth returns a boolean if a field has been set.

### GetItems

`func (o *CreateAccessListRequest) GetItems() []CreateAccessListRequestItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateAccessListRequest) GetItemsOk() (*[]CreateAccessListRequestItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateAccessListRequest) SetItems(v []CreateAccessListRequestItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateAccessListRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetClients

`func (o *CreateAccessListRequest) GetClients() []CreateAccessListRequestClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *CreateAccessListRequest) GetClientsOk() (*[]CreateAccessListRequestClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *CreateAccessListRequest) SetClients(v []CreateAccessListRequestClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *CreateAccessListRequest) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


