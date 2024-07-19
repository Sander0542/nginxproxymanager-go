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


// CreateStreamRequestForwardingHost struct for CreateStreamRequestForwardingHost
type CreateStreamRequestForwardingHost struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateStreamRequestForwardingHost) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreateStreamRequestForwardingHost)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateStreamRequestForwardingHost) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableCreateStreamRequestForwardingHost struct {
	value *CreateStreamRequestForwardingHost
	isSet bool
}

func (v NullableCreateStreamRequestForwardingHost) Get() *CreateStreamRequestForwardingHost {
	return v.value
}

func (v *NullableCreateStreamRequestForwardingHost) Set(val *CreateStreamRequestForwardingHost) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStreamRequestForwardingHost) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStreamRequestForwardingHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStreamRequestForwardingHost(val *CreateStreamRequestForwardingHost) *NullableCreateStreamRequestForwardingHost {
	return &NullableCreateStreamRequestForwardingHost{value: val, isSet: true}
}

func (v NullableCreateStreamRequestForwardingHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStreamRequestForwardingHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


