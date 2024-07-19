# TokenRefreshResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | JWT | [optional] 
**Expires** | Pointer to **time.Time** | Token expiry time | [optional] 
**Scope** | Pointer to **string** | Scope of the Token, defaults to &#39;user&#39; | [optional] 

## Methods

### NewTokenRefreshResponse

`func NewTokenRefreshResponse() *TokenRefreshResponse`

NewTokenRefreshResponse instantiates a new TokenRefreshResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRefreshResponseWithDefaults

`func NewTokenRefreshResponseWithDefaults() *TokenRefreshResponse`

NewTokenRefreshResponseWithDefaults instantiates a new TokenRefreshResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenRefreshResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenRefreshResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenRefreshResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenRefreshResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpires

`func (o *TokenRefreshResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenRefreshResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenRefreshResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenRefreshResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetScope

`func (o *TokenRefreshResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenRefreshResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenRefreshResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TokenRefreshResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


