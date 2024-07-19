# PutUserSetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Current** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 

## Methods

### NewPutUserSetPasswordRequest

`func NewPutUserSetPasswordRequest(type_ string, secret string, ) *PutUserSetPasswordRequest`

NewPutUserSetPasswordRequest instantiates a new PutUserSetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUserSetPasswordRequestWithDefaults

`func NewPutUserSetPasswordRequestWithDefaults() *PutUserSetPasswordRequest`

NewPutUserSetPasswordRequestWithDefaults instantiates a new PutUserSetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PutUserSetPasswordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutUserSetPasswordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutUserSetPasswordRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCurrent

`func (o *PutUserSetPasswordRequest) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *PutUserSetPasswordRequest) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *PutUserSetPasswordRequest) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *PutUserSetPasswordRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetSecret

`func (o *PutUserSetPasswordRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PutUserSetPasswordRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PutUserSetPasswordRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


