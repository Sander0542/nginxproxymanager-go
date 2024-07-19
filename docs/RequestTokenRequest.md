# RequestTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 

## Methods

### NewRequestTokenRequest

`func NewRequestTokenRequest(identity string, secret string, ) *RequestTokenRequest`

NewRequestTokenRequest instantiates a new RequestTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTokenRequestWithDefaults

`func NewRequestTokenRequestWithDefaults() *RequestTokenRequest`

NewRequestTokenRequestWithDefaults instantiates a new RequestTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *RequestTokenRequest) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *RequestTokenRequest) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *RequestTokenRequest) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### GetScope

`func (o *RequestTokenRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RequestTokenRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RequestTokenRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RequestTokenRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSecret

`func (o *RequestTokenRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RequestTokenRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RequestTokenRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


