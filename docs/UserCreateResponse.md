# UserCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** | Date and time of creation | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** | Date and time of last update | [optional] [readonly] 
**Name** | Pointer to **string** | Name | [optional] 
**Nickname** | Pointer to **string** | Nickname | [optional] 
**Email** | Pointer to **string** | Email Address | [optional] 
**Avatar** | Pointer to **string** | Avatar | [optional] [readonly] 
**Roles** | Pointer to **[]interface{}** | Roles | [optional] 
**IsDisabled** | Pointer to **boolAsInt** | Is Disabled | [optional] 

## Methods

### NewUserCreateResponse

`func NewUserCreateResponse() *UserCreateResponse`

NewUserCreateResponse instantiates a new UserCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateResponseWithDefaults

`func NewUserCreateResponseWithDefaults() *UserCreateResponse`

NewUserCreateResponseWithDefaults instantiates a new UserCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserCreateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCreateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCreateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *UserCreateResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *UserCreateResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *UserCreateResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *UserCreateResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *UserCreateResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *UserCreateResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *UserCreateResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *UserCreateResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *UserCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *UserCreateResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UserCreateResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UserCreateResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UserCreateResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *UserCreateResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreateResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *UserCreateResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserCreateResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserCreateResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserCreateResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetRoles

`func (o *UserCreateResponse) GetRoles() []interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreateResponse) GetRolesOk() (*[]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreateResponse) SetRoles(v []interface{})`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserCreateResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserCreateResponse) GetIsDisabled() boolAsInt`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserCreateResponse) GetIsDisabledOk() (*boolAsInt, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserCreateResponse) SetIsDisabled(v boolAsInt)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserCreateResponse) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


