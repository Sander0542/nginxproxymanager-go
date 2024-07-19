# CreateAccessList201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**OwnerUserId** | **int32** | User ID | 
**Name** | **string** |  | 
**Directive** | **string** |  | 
**Address** | [**GetAccessLists200ResponseAddress**](GetAccessLists200ResponseAddress.md) |  | 
**SatisfyAny** | **bool** |  | 
**PassAuth** | **bool** |  | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewCreateAccessList201Response

`func NewCreateAccessList201Response(id int32, createdOn string, modifiedOn string, ownerUserId int32, name string, directive string, address GetAccessLists200ResponseAddress, satisfyAny bool, passAuth bool, meta map[string]interface{}, ) *CreateAccessList201Response`

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

`func (o *CreateAccessList201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccessList201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccessList201Response) SetId(v int32)`

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

`func (o *CreateAccessList201Response) GetOwnerUserId() int32`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *CreateAccessList201Response) GetOwnerUserIdOk() (*int32, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *CreateAccessList201Response) SetOwnerUserId(v int32)`

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


### GetDirective

`func (o *CreateAccessList201Response) GetDirective() string`

GetDirective returns the Directive field if non-nil, zero value otherwise.

### GetDirectiveOk

`func (o *CreateAccessList201Response) GetDirectiveOk() (*string, bool)`

GetDirectiveOk returns a tuple with the Directive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirective

`func (o *CreateAccessList201Response) SetDirective(v string)`

SetDirective sets Directive field to given value.


### GetAddress

`func (o *CreateAccessList201Response) GetAddress() GetAccessLists200ResponseAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateAccessList201Response) GetAddressOk() (*GetAccessLists200ResponseAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateAccessList201Response) SetAddress(v GetAccessLists200ResponseAddress)`

SetAddress sets Address field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


