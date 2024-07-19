/*
Nginx Proxy Manager API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.12.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nginxproxymanager

import (
	"encoding/json"
)

// checks if the CreateCertificate400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCertificate400Response{}

// CreateCertificate400Response Error
type CreateCertificate400Response struct {
	Error *CreateCertificate400ResponseError `json:"error,omitempty"`
}

// NewCreateCertificate400Response instantiates a new CreateCertificate400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCertificate400Response() *CreateCertificate400Response {
	this := CreateCertificate400Response{}
	return &this
}

// NewCreateCertificate400ResponseWithDefaults instantiates a new CreateCertificate400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCertificate400ResponseWithDefaults() *CreateCertificate400Response {
	this := CreateCertificate400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateCertificate400Response) GetError() CreateCertificate400ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret CreateCertificate400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificate400Response) GetErrorOk() (*CreateCertificate400ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateCertificate400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CreateCertificate400ResponseError and assigns it to the Error field.
func (o *CreateCertificate400Response) SetError(v CreateCertificate400ResponseError) {
	o.Error = &v
}

func (o CreateCertificate400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCertificate400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableCreateCertificate400Response struct {
	value *CreateCertificate400Response
	isSet bool
}

func (v NullableCreateCertificate400Response) Get() *CreateCertificate400Response {
	return v.value
}

func (v *NullableCreateCertificate400Response) Set(val *CreateCertificate400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCertificate400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCertificate400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCertificate400Response(val *CreateCertificate400Response) *NullableCreateCertificate400Response {
	return &NullableCreateCertificate400Response{value: val, isSet: true}
}

func (v NullableCreateCertificate400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCertificate400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


