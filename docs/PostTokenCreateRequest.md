# PostTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | Email Address or other 3rd party providers identifier | 
**Secret** | **string** | A password or key | 
**Scope** | Pointer to **string** | Scope of the Token, defaults to &#39;user&#39; | [optional] 

## Methods

### NewPostTokenCreateRequest

`func NewPostTokenCreateRequest(identity string, secret string, ) *PostTokenCreateRequest`

NewPostTokenCreateRequest instantiates a new PostTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTokenCreateRequestWithDefaults

`func NewPostTokenCreateRequestWithDefaults() *PostTokenCreateRequest`

NewPostTokenCreateRequestWithDefaults instantiates a new PostTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *PostTokenCreateRequest) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *PostTokenCreateRequest) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *PostTokenCreateRequest) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### GetSecret

`func (o *PostTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PostTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PostTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetScope

`func (o *PostTokenCreateRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PostTokenCreateRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PostTokenCreateRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PostTokenCreateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


