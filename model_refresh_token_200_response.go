/*
Nginx Proxy Manager API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.12.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nginxproxymanager

import (
	"encoding/json"
	"fmt"
)

// checks if the RefreshToken200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshToken200Response{}

// RefreshToken200Response Token object
type RefreshToken200Response struct {
	// Token Expiry ISO Time String
	Expires string `json:"expires"`
	// JWT Token
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _RefreshToken200Response RefreshToken200Response

// NewRefreshToken200Response instantiates a new RefreshToken200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshToken200Response(expires string, token string) *RefreshToken200Response {
	this := RefreshToken200Response{}
	this.Expires = expires
	this.Token = token
	return &this
}

// NewRefreshToken200ResponseWithDefaults instantiates a new RefreshToken200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshToken200ResponseWithDefaults() *RefreshToken200Response {
	this := RefreshToken200Response{}
	return &this
}

// GetExpires returns the Expires field value
func (o *RefreshToken200Response) GetExpires() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *RefreshToken200Response) GetExpiresOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *RefreshToken200Response) SetExpires(v string) {
	o.Expires = v
}

// GetToken returns the Token field value
func (o *RefreshToken200Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RefreshToken200Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RefreshToken200Response) SetToken(v string) {
	o.Token = v
}

func (o RefreshToken200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshToken200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expires"] = o.Expires
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefreshToken200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expires",
		"token",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRefreshToken200Response := _RefreshToken200Response{}

	err = json.Unmarshal(data, &varRefreshToken200Response)

	if err != nil {
		return err
	}

	*o = RefreshToken200Response(varRefreshToken200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expires")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefreshToken200Response struct {
	value *RefreshToken200Response
	isSet bool
}

func (v NullableRefreshToken200Response) Get() *RefreshToken200Response {
	return v.value
}

func (v *NullableRefreshToken200Response) Set(val *RefreshToken200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshToken200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshToken200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshToken200Response(val *RefreshToken200Response) *NullableRefreshToken200Response {
	return &NullableRefreshToken200Response{value: val, isSet: true}
}

func (v NullableRefreshToken200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshToken200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


