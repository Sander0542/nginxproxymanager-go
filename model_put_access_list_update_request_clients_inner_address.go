/*
Nginx Proxy Manager API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.x.x
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nginxproxymanager

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// PutAccessListUpdateRequestClientsInnerAddress - struct for PutAccessListUpdateRequestClientsInnerAddress
type PutAccessListUpdateRequestClientsInnerAddress struct {
	String *string
}

// stringAsPutAccessListUpdateRequestClientsInnerAddress is a convenience function that returns string wrapped in PutAccessListUpdateRequestClientsInnerAddress
func StringAsPutAccessListUpdateRequestClientsInnerAddress(v *string) PutAccessListUpdateRequestClientsInnerAddress {
	return PutAccessListUpdateRequestClientsInnerAddress{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PutAccessListUpdateRequestClientsInnerAddress) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PutAccessListUpdateRequestClientsInnerAddress)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PutAccessListUpdateRequestClientsInnerAddress)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PutAccessListUpdateRequestClientsInnerAddress) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PutAccessListUpdateRequestClientsInnerAddress) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePutAccessListUpdateRequestClientsInnerAddress struct {
	value *PutAccessListUpdateRequestClientsInnerAddress
	isSet bool
}

func (v NullablePutAccessListUpdateRequestClientsInnerAddress) Get() *PutAccessListUpdateRequestClientsInnerAddress {
	return v.value
}

func (v *NullablePutAccessListUpdateRequestClientsInnerAddress) Set(val *PutAccessListUpdateRequestClientsInnerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAccessListUpdateRequestClientsInnerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAccessListUpdateRequestClientsInnerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAccessListUpdateRequestClientsInnerAddress(val *PutAccessListUpdateRequestClientsInnerAddress) *NullablePutAccessListUpdateRequestClientsInnerAddress {
	return &NullablePutAccessListUpdateRequestClientsInnerAddress{value: val, isSet: true}
}

func (v NullablePutAccessListUpdateRequestClientsInnerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAccessListUpdateRequestClientsInnerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


