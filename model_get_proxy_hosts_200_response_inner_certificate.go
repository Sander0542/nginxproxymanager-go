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
	"gopkg.in/validator.v2"
)

// GetProxyHosts200ResponseInnerCertificate - struct for GetProxyHosts200ResponseInnerCertificate
type GetProxyHosts200ResponseInnerCertificate struct {
	GetCertificates200ResponseInner *GetCertificates200ResponseInner
}

// GetCertificates200ResponseInnerAsGetProxyHosts200ResponseInnerCertificate is a convenience function that returns GetCertificates200ResponseInner wrapped in GetProxyHosts200ResponseInnerCertificate
func GetCertificates200ResponseInnerAsGetProxyHosts200ResponseInnerCertificate(v *GetCertificates200ResponseInner) GetProxyHosts200ResponseInnerCertificate {
	return GetProxyHosts200ResponseInnerCertificate{
		GetCertificates200ResponseInner: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetProxyHosts200ResponseInnerCertificate) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into GetCertificates200ResponseInner
	err = newStrictDecoder(data).Decode(&dst.GetCertificates200ResponseInner)
	if err == nil {
		jsonGetCertificates200ResponseInner, _ := json.Marshal(dst.GetCertificates200ResponseInner)
		if string(jsonGetCertificates200ResponseInner) == "{}" { // empty struct
			dst.GetCertificates200ResponseInner = nil
		} else {
			if err = validator.Validate(dst.GetCertificates200ResponseInner); err != nil {
				dst.GetCertificates200ResponseInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetCertificates200ResponseInner = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetCertificates200ResponseInner = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetProxyHosts200ResponseInnerCertificate)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetProxyHosts200ResponseInnerCertificate)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetProxyHosts200ResponseInnerCertificate) MarshalJSON() ([]byte, error) {
	if src.GetCertificates200ResponseInner != nil {
		return json.Marshal(&src.GetCertificates200ResponseInner)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetProxyHosts200ResponseInnerCertificate) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetCertificates200ResponseInner != nil {
		return obj.GetCertificates200ResponseInner
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetProxyHosts200ResponseInnerCertificate) GetActualInstanceValue() (interface{}) {
	if obj.GetCertificates200ResponseInner != nil {
		return *obj.GetCertificates200ResponseInner
	}

	// all schemas are nil
	return nil
}

type NullableGetProxyHosts200ResponseInnerCertificate struct {
	value *GetProxyHosts200ResponseInnerCertificate
	isSet bool
}

func (v NullableGetProxyHosts200ResponseInnerCertificate) Get() *GetProxyHosts200ResponseInnerCertificate {
	return v.value
}

func (v *NullableGetProxyHosts200ResponseInnerCertificate) Set(val *GetProxyHosts200ResponseInnerCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProxyHosts200ResponseInnerCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProxyHosts200ResponseInnerCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProxyHosts200ResponseInnerCertificate(val *GetProxyHosts200ResponseInnerCertificate) *NullableGetProxyHosts200ResponseInnerCertificate {
	return &NullableGetProxyHosts200ResponseInnerCertificate{value: val, isSet: true}
}

func (v NullableGetProxyHosts200ResponseInnerCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProxyHosts200ResponseInnerCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


