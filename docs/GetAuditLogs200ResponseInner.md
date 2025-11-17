# GetAuditLogs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**UserId** | **int64** | User ID | 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** | Unique identifier | [readonly] 
**Action** | **string** |  | 
**Meta** | **map[string]interface{}** |  | 
**User** | Pointer to [**GetAuditLogs200ResponseInnerUser**](GetAuditLogs200ResponseInnerUser.md) |  | [optional] 

## Methods

### NewGetAuditLogs200ResponseInner

`func NewGetAuditLogs200ResponseInner(id int64, createdOn string, modifiedOn string, userId int64, objectType string, objectId int64, action string, meta map[string]interface{}, ) *GetAuditLogs200ResponseInner`

NewGetAuditLogs200ResponseInner instantiates a new GetAuditLogs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditLogs200ResponseInnerWithDefaults

`func NewGetAuditLogs200ResponseInnerWithDefaults() *GetAuditLogs200ResponseInner`

NewGetAuditLogs200ResponseInnerWithDefaults instantiates a new GetAuditLogs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAuditLogs200ResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAuditLogs200ResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAuditLogs200ResponseInner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAuditLogs200ResponseInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAuditLogs200ResponseInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAuditLogs200ResponseInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAuditLogs200ResponseInner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAuditLogs200ResponseInner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAuditLogs200ResponseInner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetUserId

`func (o *GetAuditLogs200ResponseInner) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetAuditLogs200ResponseInner) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetAuditLogs200ResponseInner) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetObjectType

`func (o *GetAuditLogs200ResponseInner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GetAuditLogs200ResponseInner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GetAuditLogs200ResponseInner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *GetAuditLogs200ResponseInner) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *GetAuditLogs200ResponseInner) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *GetAuditLogs200ResponseInner) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetAction

`func (o *GetAuditLogs200ResponseInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAuditLogs200ResponseInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAuditLogs200ResponseInner) SetAction(v string)`

SetAction sets Action field to given value.


### GetMeta

`func (o *GetAuditLogs200ResponseInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAuditLogs200ResponseInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAuditLogs200ResponseInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetUser

`func (o *GetAuditLogs200ResponseInner) GetUser() GetAuditLogs200ResponseInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetAuditLogs200ResponseInner) GetUserOk() (*GetAuditLogs200ResponseInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetAuditLogs200ResponseInner) SetUser(v GetAuditLogs200ResponseInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetAuditLogs200ResponseInner) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


