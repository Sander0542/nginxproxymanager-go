# LoginAsUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | **float32** | Token Expiry Unix Time | 
**Token** | **string** | JWT Token | 
**User** | [**GetAccessLists200ResponseInnerOwner**](GetAccessLists200ResponseInnerOwner.md) |  | 

## Methods

### NewLoginAsUser200Response

`func NewLoginAsUser200Response(expires float32, token string, user GetAccessLists200ResponseInnerOwner, ) *LoginAsUser200Response`

NewLoginAsUser200Response instantiates a new LoginAsUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginAsUser200ResponseWithDefaults

`func NewLoginAsUser200ResponseWithDefaults() *LoginAsUser200Response`

NewLoginAsUser200ResponseWithDefaults instantiates a new LoginAsUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *LoginAsUser200Response) GetExpires() float32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *LoginAsUser200Response) GetExpiresOk() (*float32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *LoginAsUser200Response) SetExpires(v float32)`

SetExpires sets Expires field to given value.


### GetToken

`func (o *LoginAsUser200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoginAsUser200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoginAsUser200Response) SetToken(v string)`

SetToken sets Token field to given value.


### GetUser

`func (o *LoginAsUser200Response) GetUser() GetAccessLists200ResponseInnerOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoginAsUser200Response) GetUserOk() (*GetAccessLists200ResponseInnerOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoginAsUser200Response) SetUser(v GetAccessLists200ResponseInnerOwner)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


