# UserGetResponse

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

### NewUserGetResponse

`func NewUserGetResponse() *UserGetResponse`

NewUserGetResponse instantiates a new UserGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetResponseWithDefaults

`func NewUserGetResponseWithDefaults() *UserGetResponse`

NewUserGetResponseWithDefaults instantiates a new UserGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGetResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGetResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGetResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *UserGetResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *UserGetResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *UserGetResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *UserGetResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *UserGetResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *UserGetResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *UserGetResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *UserGetResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *UserGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *UserGetResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UserGetResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UserGetResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UserGetResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *UserGetResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserGetResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserGetResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserGetResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *UserGetResponse) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserGetResponse) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserGetResponse) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserGetResponse) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetRoles

`func (o *UserGetResponse) GetRoles() []interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserGetResponse) GetRolesOk() (*[]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserGetResponse) SetRoles(v []interface{})`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserGetResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserGetResponse) GetIsDisabled() boolAsInt`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserGetResponse) GetIsDisabledOk() (*boolAsInt, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserGetResponse) SetIsDisabled(v boolAsInt)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserGetResponse) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


