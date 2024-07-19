# TokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | JWT | [optional] 
**Expires** | Pointer to **time.Time** | Token expiry time | [optional] 

## Methods

### NewTokenCreateResponse

`func NewTokenCreateResponse() *TokenCreateResponse`

NewTokenCreateResponse instantiates a new TokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateResponseWithDefaults

`func NewTokenCreateResponseWithDefaults() *TokenCreateResponse`

NewTokenCreateResponseWithDefaults instantiates a new TokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenCreateResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenCreateResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenCreateResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenCreateResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpires

`func (o *TokenCreateResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenCreateResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenCreateResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenCreateResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


