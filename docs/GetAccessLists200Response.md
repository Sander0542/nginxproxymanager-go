# GetAccessLists200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**OwnerUserId** | **int64** | User ID | 
**Name** | **string** |  | 
**Meta** | **map[string]interface{}** |  | 
**SatisfyAny** | **bool** |  | 
**PassAuth** | **bool** |  | 
**ProxyHostCount** | **int64** |  | 

## Methods

### NewGetAccessLists200Response

`func NewGetAccessLists200Response(id int64, createdOn string, modifiedOn string, ownerUserId int64, name string, meta map[string]interface{}, satisfyAny bool, passAuth bool, proxyHostCount int64, ) *GetAccessLists200Response`

NewGetAccessLists200Response instantiates a new GetAccessLists200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessLists200ResponseWithDefaults

`func NewGetAccessLists200ResponseWithDefaults() *GetAccessLists200Response`

NewGetAccessLists200ResponseWithDefaults instantiates a new GetAccessLists200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessLists200Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessLists200Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessLists200Response) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAccessLists200Response) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAccessLists200Response) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAccessLists200Response) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAccessLists200Response) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAccessLists200Response) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAccessLists200Response) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *GetAccessLists200Response) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *GetAccessLists200Response) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *GetAccessLists200Response) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetName

`func (o *GetAccessLists200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAccessLists200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAccessLists200Response) SetName(v string)`

SetName sets Name field to given value.


### GetMeta

`func (o *GetAccessLists200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAccessLists200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAccessLists200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetSatisfyAny

`func (o *GetAccessLists200Response) GetSatisfyAny() bool`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *GetAccessLists200Response) GetSatisfyAnyOk() (*bool, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *GetAccessLists200Response) SetSatisfyAny(v bool)`

SetSatisfyAny sets SatisfyAny field to given value.


### GetPassAuth

`func (o *GetAccessLists200Response) GetPassAuth() bool`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *GetAccessLists200Response) GetPassAuthOk() (*bool, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *GetAccessLists200Response) SetPassAuth(v bool)`

SetPassAuth sets PassAuth field to given value.


### GetProxyHostCount

`func (o *GetAccessLists200Response) GetProxyHostCount() int64`

GetProxyHostCount returns the ProxyHostCount field if non-nil, zero value otherwise.

### GetProxyHostCountOk

`func (o *GetAccessLists200Response) GetProxyHostCountOk() (*int64, bool)`

GetProxyHostCountOk returns a tuple with the ProxyHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHostCount

`func (o *GetAccessLists200Response) SetProxyHostCount(v int64)`

SetProxyHostCount sets ProxyHostCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


