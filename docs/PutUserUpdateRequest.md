# PutUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Nickname** | Pointer to **string** | Nickname | [optional] 
**Email** | Pointer to **string** | Email Address | [optional] 
**Roles** | Pointer to **[]interface{}** | Roles | [optional] 
**IsDisabled** | Pointer to **boolAsInt** | Is Disabled | [optional] 

## Methods

### NewPutUserUpdateRequest

`func NewPutUserUpdateRequest() *PutUserUpdateRequest`

NewPutUserUpdateRequest instantiates a new PutUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUserUpdateRequestWithDefaults

`func NewPutUserUpdateRequestWithDefaults() *PutUserUpdateRequest`

NewPutUserUpdateRequestWithDefaults instantiates a new PutUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutUserUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutUserUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutUserUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutUserUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *PutUserUpdateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *PutUserUpdateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *PutUserUpdateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *PutUserUpdateRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *PutUserUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PutUserUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PutUserUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PutUserUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRoles

`func (o *PutUserUpdateRequest) GetRoles() []interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PutUserUpdateRequest) GetRolesOk() (*[]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PutUserUpdateRequest) SetRoles(v []interface{})`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PutUserUpdateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetIsDisabled

`func (o *PutUserUpdateRequest) GetIsDisabled() boolAsInt`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *PutUserUpdateRequest) GetIsDisabledOk() (*boolAsInt, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *PutUserUpdateRequest) SetIsDisabled(v boolAsInt)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *PutUserUpdateRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


