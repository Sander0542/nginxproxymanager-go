# GetAuditLogs200ResponseInnerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | User ID | 
**CreatedOn** | **string** | Created Date | 
**ModifiedOn** | **string** | Modified Date | 
**IsDisabled** | **bool** | Is user Disabled | 
**Email** | **string** | Email | 
**Name** | **string** | Name | 
**Nickname** | **string** | Nickname | 
**Avatar** | **string** | Gravatar URL based on email, without scheme | 
**Roles** | **[]string** | Roles applied | 
**Permissions** | Pointer to [**GetAuditLogs200ResponseInnerUserPermissions**](GetAuditLogs200ResponseInnerUserPermissions.md) |  | [optional] 

## Methods

### NewGetAuditLogs200ResponseInnerUser

`func NewGetAuditLogs200ResponseInnerUser(id int64, createdOn string, modifiedOn string, isDisabled bool, email string, name string, nickname string, avatar string, roles []string, ) *GetAuditLogs200ResponseInnerUser`

NewGetAuditLogs200ResponseInnerUser instantiates a new GetAuditLogs200ResponseInnerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditLogs200ResponseInnerUserWithDefaults

`func NewGetAuditLogs200ResponseInnerUserWithDefaults() *GetAuditLogs200ResponseInnerUser`

NewGetAuditLogs200ResponseInnerUserWithDefaults instantiates a new GetAuditLogs200ResponseInnerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAuditLogs200ResponseInnerUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAuditLogs200ResponseInnerUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAuditLogs200ResponseInnerUser) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAuditLogs200ResponseInnerUser) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAuditLogs200ResponseInnerUser) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAuditLogs200ResponseInnerUser) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAuditLogs200ResponseInnerUser) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAuditLogs200ResponseInnerUser) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAuditLogs200ResponseInnerUser) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetIsDisabled

`func (o *GetAuditLogs200ResponseInnerUser) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *GetAuditLogs200ResponseInnerUser) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *GetAuditLogs200ResponseInnerUser) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetEmail

`func (o *GetAuditLogs200ResponseInnerUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAuditLogs200ResponseInnerUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAuditLogs200ResponseInnerUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *GetAuditLogs200ResponseInnerUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAuditLogs200ResponseInnerUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAuditLogs200ResponseInnerUser) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *GetAuditLogs200ResponseInnerUser) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *GetAuditLogs200ResponseInnerUser) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *GetAuditLogs200ResponseInnerUser) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetAvatar

`func (o *GetAuditLogs200ResponseInnerUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetAuditLogs200ResponseInnerUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetAuditLogs200ResponseInnerUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetRoles

`func (o *GetAuditLogs200ResponseInnerUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetAuditLogs200ResponseInnerUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetAuditLogs200ResponseInnerUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetPermissions

`func (o *GetAuditLogs200ResponseInnerUser) GetPermissions() GetAuditLogs200ResponseInnerUserPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetAuditLogs200ResponseInnerUser) GetPermissionsOk() (*GetAuditLogs200ResponseInnerUserPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetAuditLogs200ResponseInnerUser) SetPermissions(v GetAuditLogs200ResponseInnerUserPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetAuditLogs200ResponseInnerUser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


