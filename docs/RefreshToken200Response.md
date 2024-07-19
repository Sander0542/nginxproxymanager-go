# RefreshToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | **string** | Token Expiry ISO Time String | 
**Token** | **string** | JWT Token | 

## Methods

### NewRefreshToken200Response

`func NewRefreshToken200Response(expires string, token string, ) *RefreshToken200Response`

NewRefreshToken200Response instantiates a new RefreshToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshToken200ResponseWithDefaults

`func NewRefreshToken200ResponseWithDefaults() *RefreshToken200Response`

NewRefreshToken200ResponseWithDefaults instantiates a new RefreshToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *RefreshToken200Response) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RefreshToken200Response) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RefreshToken200Response) SetExpires(v string)`

SetExpires sets Expires field to given value.


### GetToken

`func (o *RefreshToken200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RefreshToken200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RefreshToken200Response) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


