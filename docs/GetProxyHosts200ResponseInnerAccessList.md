# GetProxyHosts200ResponseInnerAccessList

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

### NewGetProxyHosts200ResponseInnerAccessList

`func NewGetProxyHosts200ResponseInnerAccessList(id int64, createdOn string, modifiedOn string, ownerUserId int64, name string, meta map[string]interface{}, satisfyAny bool, passAuth bool, proxyHostCount int64, ) *GetProxyHosts200ResponseInnerAccessList`

NewGetProxyHosts200ResponseInnerAccessList instantiates a new GetProxyHosts200ResponseInnerAccessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProxyHosts200ResponseInnerAccessListWithDefaults

`func NewGetProxyHosts200ResponseInnerAccessListWithDefaults() *GetProxyHosts200ResponseInnerAccessList`

NewGetProxyHosts200ResponseInnerAccessListWithDefaults instantiates a new GetProxyHosts200ResponseInnerAccessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetProxyHosts200ResponseInnerAccessList) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProxyHosts200ResponseInnerAccessList) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetProxyHosts200ResponseInnerAccessList) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetProxyHosts200ResponseInnerAccessList) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetProxyHosts200ResponseInnerAccessList) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetProxyHosts200ResponseInnerAccessList) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *GetProxyHosts200ResponseInnerAccessList) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *GetProxyHosts200ResponseInnerAccessList) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetName

`func (o *GetProxyHosts200ResponseInnerAccessList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetProxyHosts200ResponseInnerAccessList) SetName(v string)`

SetName sets Name field to given value.


### GetMeta

`func (o *GetProxyHosts200ResponseInnerAccessList) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProxyHosts200ResponseInnerAccessList) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetSatisfyAny

`func (o *GetProxyHosts200ResponseInnerAccessList) GetSatisfyAny() bool`

GetSatisfyAny returns the SatisfyAny field if non-nil, zero value otherwise.

### GetSatisfyAnyOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetSatisfyAnyOk() (*bool, bool)`

GetSatisfyAnyOk returns a tuple with the SatisfyAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfyAny

`func (o *GetProxyHosts200ResponseInnerAccessList) SetSatisfyAny(v bool)`

SetSatisfyAny sets SatisfyAny field to given value.


### GetPassAuth

`func (o *GetProxyHosts200ResponseInnerAccessList) GetPassAuth() bool`

GetPassAuth returns the PassAuth field if non-nil, zero value otherwise.

### GetPassAuthOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetPassAuthOk() (*bool, bool)`

GetPassAuthOk returns a tuple with the PassAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassAuth

`func (o *GetProxyHosts200ResponseInnerAccessList) SetPassAuth(v bool)`

SetPassAuth sets PassAuth field to given value.


### GetProxyHostCount

`func (o *GetProxyHosts200ResponseInnerAccessList) GetProxyHostCount() int64`

GetProxyHostCount returns the ProxyHostCount field if non-nil, zero value otherwise.

### GetProxyHostCountOk

`func (o *GetProxyHosts200ResponseInnerAccessList) GetProxyHostCountOk() (*int64, bool)`

GetProxyHostCountOk returns a tuple with the ProxyHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHostCount

`func (o *GetProxyHosts200ResponseInnerAccessList) SetProxyHostCount(v int64)`

SetProxyHostCount sets ProxyHostCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


