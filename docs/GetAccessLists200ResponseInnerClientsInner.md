# GetAccessLists200ResponseInnerClientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**AccessListId** | **int64** | Access List ID | 
**Address** | [**GetAccessLists200ResponseInnerClientsInnerAddress**](GetAccessLists200ResponseInnerClientsInnerAddress.md) |  | 
**Directive** | **string** |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewGetAccessLists200ResponseInnerClientsInner

`func NewGetAccessLists200ResponseInnerClientsInner(id int64, createdOn string, modifiedOn string, accessListId int64, address GetAccessLists200ResponseInnerClientsInnerAddress, directive string, meta map[string]interface{}, ) *GetAccessLists200ResponseInnerClientsInner`

NewGetAccessLists200ResponseInnerClientsInner instantiates a new GetAccessLists200ResponseInnerClientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessLists200ResponseInnerClientsInnerWithDefaults

`func NewGetAccessLists200ResponseInnerClientsInnerWithDefaults() *GetAccessLists200ResponseInnerClientsInner`

NewGetAccessLists200ResponseInnerClientsInnerWithDefaults instantiates a new GetAccessLists200ResponseInnerClientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessLists200ResponseInnerClientsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessLists200ResponseInnerClientsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAccessLists200ResponseInnerClientsInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAccessLists200ResponseInnerClientsInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAccessLists200ResponseInnerClientsInner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAccessLists200ResponseInnerClientsInner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetAccessListId

`func (o *GetAccessLists200ResponseInnerClientsInner) GetAccessListId() int64`

GetAccessListId returns the AccessListId field if non-nil, zero value otherwise.

### GetAccessListIdOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetAccessListIdOk() (*int64, bool)`

GetAccessListIdOk returns a tuple with the AccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListId

`func (o *GetAccessLists200ResponseInnerClientsInner) SetAccessListId(v int64)`

SetAccessListId sets AccessListId field to given value.


### GetAddress

`func (o *GetAccessLists200ResponseInnerClientsInner) GetAddress() GetAccessLists200ResponseInnerClientsInnerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetAddressOk() (*GetAccessLists200ResponseInnerClientsInnerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetAccessLists200ResponseInnerClientsInner) SetAddress(v GetAccessLists200ResponseInnerClientsInnerAddress)`

SetAddress sets Address field to given value.


### GetDirective

`func (o *GetAccessLists200ResponseInnerClientsInner) GetDirective() string`

GetDirective returns the Directive field if non-nil, zero value otherwise.

### GetDirectiveOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetDirectiveOk() (*string, bool)`

GetDirectiveOk returns a tuple with the Directive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirective

`func (o *GetAccessLists200ResponseInnerClientsInner) SetDirective(v string)`

SetDirective sets Directive field to given value.


### GetMeta

`func (o *GetAccessLists200ResponseInnerClientsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAccessLists200ResponseInnerClientsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAccessLists200ResponseInnerClientsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


