# GetAccessLists200ResponseClientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**AccessListId** | **int64** | Access List ID | 
**Address** | [**GetAccessLists200ResponseClientsInnerAddress**](GetAccessLists200ResponseClientsInnerAddress.md) |  | 
**Directive** | **string** |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewGetAccessLists200ResponseClientsInner

`func NewGetAccessLists200ResponseClientsInner(id int64, createdOn string, modifiedOn string, accessListId int64, address GetAccessLists200ResponseClientsInnerAddress, directive string, meta map[string]interface{}, ) *GetAccessLists200ResponseClientsInner`

NewGetAccessLists200ResponseClientsInner instantiates a new GetAccessLists200ResponseClientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessLists200ResponseClientsInnerWithDefaults

`func NewGetAccessLists200ResponseClientsInnerWithDefaults() *GetAccessLists200ResponseClientsInner`

NewGetAccessLists200ResponseClientsInnerWithDefaults instantiates a new GetAccessLists200ResponseClientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessLists200ResponseClientsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessLists200ResponseClientsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessLists200ResponseClientsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAccessLists200ResponseClientsInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAccessLists200ResponseClientsInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAccessLists200ResponseClientsInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAccessLists200ResponseClientsInner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAccessLists200ResponseClientsInner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAccessLists200ResponseClientsInner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetAccessListId

`func (o *GetAccessLists200ResponseClientsInner) GetAccessListId() int64`

GetAccessListId returns the AccessListId field if non-nil, zero value otherwise.

### GetAccessListIdOk

`func (o *GetAccessLists200ResponseClientsInner) GetAccessListIdOk() (*int64, bool)`

GetAccessListIdOk returns a tuple with the AccessListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessListId

`func (o *GetAccessLists200ResponseClientsInner) SetAccessListId(v int64)`

SetAccessListId sets AccessListId field to given value.


### GetAddress

`func (o *GetAccessLists200ResponseClientsInner) GetAddress() GetAccessLists200ResponseClientsInnerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetAccessLists200ResponseClientsInner) GetAddressOk() (*GetAccessLists200ResponseClientsInnerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetAccessLists200ResponseClientsInner) SetAddress(v GetAccessLists200ResponseClientsInnerAddress)`

SetAddress sets Address field to given value.


### GetDirective

`func (o *GetAccessLists200ResponseClientsInner) GetDirective() string`

GetDirective returns the Directive field if non-nil, zero value otherwise.

### GetDirectiveOk

`func (o *GetAccessLists200ResponseClientsInner) GetDirectiveOk() (*string, bool)`

GetDirectiveOk returns a tuple with the Directive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirective

`func (o *GetAccessLists200ResponseClientsInner) SetDirective(v string)`

SetDirective sets Directive field to given value.


### GetMeta

`func (o *GetAccessLists200ResponseClientsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAccessLists200ResponseClientsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAccessLists200ResponseClientsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


