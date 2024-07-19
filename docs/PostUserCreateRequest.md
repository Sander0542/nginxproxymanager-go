# PostUserCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Nickname** | **string** | Nickname | 
**Email** | **string** | Email Address | 
**Roles** | Pointer to **[]interface{}** | Roles | [optional] 
**IsDisabled** | Pointer to **boolAsInt** | Is Disabled | [optional] 
**Auth** | Pointer to **map[string]interface{}** | Auth Credentials | [optional] 

## Methods

### NewPostUserCreateRequest

`func NewPostUserCreateRequest(name string, nickname string, email string, ) *PostUserCreateRequest`

NewPostUserCreateRequest instantiates a new PostUserCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostUserCreateRequestWithDefaults

`func NewPostUserCreateRequestWithDefaults() *PostUserCreateRequest`

NewPostUserCreateRequestWithDefaults instantiates a new PostUserCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostUserCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostUserCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostUserCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *PostUserCreateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *PostUserCreateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *PostUserCreateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetEmail

`func (o *PostUserCreateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostUserCreateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostUserCreateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoles

`func (o *PostUserCreateRequest) GetRoles() []interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PostUserCreateRequest) GetRolesOk() (*[]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PostUserCreateRequest) SetRoles(v []interface{})`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PostUserCreateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetIsDisabled

`func (o *PostUserCreateRequest) GetIsDisabled() boolAsInt`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *PostUserCreateRequest) GetIsDisabledOk() (*boolAsInt, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *PostUserCreateRequest) SetIsDisabled(v boolAsInt)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *PostUserCreateRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetAuth

`func (o *PostUserCreateRequest) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *PostUserCreateRequest) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *PostUserCreateRequest) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *PostUserCreateRequest) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


