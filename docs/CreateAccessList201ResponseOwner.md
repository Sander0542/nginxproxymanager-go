# CreateAccessList201ResponseOwner

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
**Permissions** | Pointer to [**CreateAccessList201ResponseOwnerPermissions**](CreateAccessList201ResponseOwnerPermissions.md) |  | [optional] 

## Methods

### NewCreateAccessList201ResponseOwner

`func NewCreateAccessList201ResponseOwner(id int64, createdOn string, modifiedOn string, isDisabled bool, email string, name string, nickname string, avatar string, roles []string, ) *CreateAccessList201ResponseOwner`

NewCreateAccessList201ResponseOwner instantiates a new CreateAccessList201ResponseOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessList201ResponseOwnerWithDefaults

`func NewCreateAccessList201ResponseOwnerWithDefaults() *CreateAccessList201ResponseOwner`

NewCreateAccessList201ResponseOwnerWithDefaults instantiates a new CreateAccessList201ResponseOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAccessList201ResponseOwner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccessList201ResponseOwner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccessList201ResponseOwner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *CreateAccessList201ResponseOwner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *CreateAccessList201ResponseOwner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *CreateAccessList201ResponseOwner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *CreateAccessList201ResponseOwner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *CreateAccessList201ResponseOwner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *CreateAccessList201ResponseOwner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetIsDisabled

`func (o *CreateAccessList201ResponseOwner) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateAccessList201ResponseOwner) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateAccessList201ResponseOwner) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetEmail

`func (o *CreateAccessList201ResponseOwner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAccessList201ResponseOwner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAccessList201ResponseOwner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CreateAccessList201ResponseOwner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccessList201ResponseOwner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccessList201ResponseOwner) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *CreateAccessList201ResponseOwner) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreateAccessList201ResponseOwner) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreateAccessList201ResponseOwner) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetAvatar

`func (o *CreateAccessList201ResponseOwner) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CreateAccessList201ResponseOwner) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CreateAccessList201ResponseOwner) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetRoles

`func (o *CreateAccessList201ResponseOwner) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateAccessList201ResponseOwner) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateAccessList201ResponseOwner) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetPermissions

`func (o *CreateAccessList201ResponseOwner) GetPermissions() CreateAccessList201ResponseOwnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateAccessList201ResponseOwner) GetPermissionsOk() (*CreateAccessList201ResponseOwnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateAccessList201ResponseOwner) SetPermissions(v CreateAccessList201ResponseOwnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateAccessList201ResponseOwner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


