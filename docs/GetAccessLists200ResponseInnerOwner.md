# GetAccessLists200ResponseInnerOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**IsDisabled** | **bool** | Is user Disabled | 
**Email** | **string** | Email | 
**Name** | **string** | Name | 
**Nickname** | **string** | Nickname | 
**Avatar** | **string** | Gravatar URL based on email, without scheme | 
**Roles** | **[]string** | Roles applied | 
**Permissions** | Pointer to [**GetAccessLists200ResponseInnerOwnerPermissions**](GetAccessLists200ResponseInnerOwnerPermissions.md) |  | [optional] 

## Methods

### NewGetAccessLists200ResponseInnerOwner

`func NewGetAccessLists200ResponseInnerOwner(id int64, createdOn string, modifiedOn string, isDisabled bool, email string, name string, nickname string, avatar string, roles []string, ) *GetAccessLists200ResponseInnerOwner`

NewGetAccessLists200ResponseInnerOwner instantiates a new GetAccessLists200ResponseInnerOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessLists200ResponseInnerOwnerWithDefaults

`func NewGetAccessLists200ResponseInnerOwnerWithDefaults() *GetAccessLists200ResponseInnerOwner`

NewGetAccessLists200ResponseInnerOwnerWithDefaults instantiates a new GetAccessLists200ResponseInnerOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessLists200ResponseInnerOwner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessLists200ResponseInnerOwner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessLists200ResponseInnerOwner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *GetAccessLists200ResponseInnerOwner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GetAccessLists200ResponseInnerOwner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GetAccessLists200ResponseInnerOwner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *GetAccessLists200ResponseInnerOwner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GetAccessLists200ResponseInnerOwner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GetAccessLists200ResponseInnerOwner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetIsDisabled

`func (o *GetAccessLists200ResponseInnerOwner) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *GetAccessLists200ResponseInnerOwner) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *GetAccessLists200ResponseInnerOwner) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetEmail

`func (o *GetAccessLists200ResponseInnerOwner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAccessLists200ResponseInnerOwner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAccessLists200ResponseInnerOwner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *GetAccessLists200ResponseInnerOwner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAccessLists200ResponseInnerOwner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAccessLists200ResponseInnerOwner) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *GetAccessLists200ResponseInnerOwner) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *GetAccessLists200ResponseInnerOwner) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *GetAccessLists200ResponseInnerOwner) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetAvatar

`func (o *GetAccessLists200ResponseInnerOwner) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetAccessLists200ResponseInnerOwner) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetAccessLists200ResponseInnerOwner) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetRoles

`func (o *GetAccessLists200ResponseInnerOwner) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetAccessLists200ResponseInnerOwner) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetAccessLists200ResponseInnerOwner) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetPermissions

`func (o *GetAccessLists200ResponseInnerOwner) GetPermissions() GetAccessLists200ResponseInnerOwnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetAccessLists200ResponseInnerOwner) GetPermissionsOk() (*GetAccessLists200ResponseInnerOwnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetAccessLists200ResponseInnerOwner) SetPermissions(v GetAccessLists200ResponseInnerOwnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetAccessLists200ResponseInnerOwner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


