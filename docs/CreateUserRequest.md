# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Nickname** | **string** | Nickname | 
**Email** | **string** | Email | 
**Roles** | Pointer to **[]string** | Roles applied | [optional] 
**IsDisabled** | Pointer to **bool** | Is user Disabled | [optional] 
**Auth** | Pointer to **map[string]interface{}** | Auth Credentials | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(name string, nickname string, email string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *CreateUserRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreateUserRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreateUserRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetEmail

`func (o *CreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoles

`func (o *CreateUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetIsDisabled

`func (o *CreateUserRequest) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateUserRequest) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateUserRequest) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *CreateUserRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAuth

`func (o *CreateUserRequest) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CreateUserRequest) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CreateUserRequest) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *CreateUserRequest) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


