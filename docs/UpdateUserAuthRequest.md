# UpdateUserAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Current** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 

## Methods

### NewUpdateUserAuthRequest

`func NewUpdateUserAuthRequest(type_ string, secret string, ) *UpdateUserAuthRequest`

NewUpdateUserAuthRequest instantiates a new UpdateUserAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserAuthRequestWithDefaults

`func NewUpdateUserAuthRequestWithDefaults() *UpdateUserAuthRequest`

NewUpdateUserAuthRequestWithDefaults instantiates a new UpdateUserAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateUserAuthRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateUserAuthRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateUserAuthRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCurrent

`func (o *UpdateUserAuthRequest) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *UpdateUserAuthRequest) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *UpdateUserAuthRequest) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *UpdateUserAuthRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateUserAuthRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateUserAuthRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateUserAuthRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


