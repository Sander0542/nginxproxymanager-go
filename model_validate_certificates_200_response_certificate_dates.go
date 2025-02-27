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

// checks if the ValidateCertificates200ResponseCertificateDates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateCertificates200ResponseCertificateDates{}

// ValidateCertificates200ResponseCertificateDates struct for ValidateCertificates200ResponseCertificateDates
type ValidateCertificates200ResponseCertificateDates struct {
	From int64 `json:"from"`
	To int64 `json:"to"`
	AdditionalProperties map[string]interface{}
}

type _ValidateCertificates200ResponseCertificateDates ValidateCertificates200ResponseCertificateDates

// NewValidateCertificates200ResponseCertificateDates instantiates a new ValidateCertificates200ResponseCertificateDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateCertificates200ResponseCertificateDates(from int64, to int64) *ValidateCertificates200ResponseCertificateDates {
	this := ValidateCertificates200ResponseCertificateDates{}
	this.From = from
	this.To = to
	return &this
}

// NewValidateCertificates200ResponseCertificateDatesWithDefaults instantiates a new ValidateCertificates200ResponseCertificateDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateCertificates200ResponseCertificateDatesWithDefaults() *ValidateCertificates200ResponseCertificateDates {
	this := ValidateCertificates200ResponseCertificateDates{}
	return &this
}

// GetFrom returns the From field value
func (o *ValidateCertificates200ResponseCertificateDates) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200ResponseCertificateDates) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *ValidateCertificates200ResponseCertificateDates) SetFrom(v int64) {
	o.From = v
}

// GetTo returns the To field value
func (o *ValidateCertificates200ResponseCertificateDates) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200ResponseCertificateDates) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *ValidateCertificates200ResponseCertificateDates) SetTo(v int64) {
	o.To = v
}

func (o ValidateCertificates200ResponseCertificateDates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateCertificates200ResponseCertificateDates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateCertificates200ResponseCertificateDates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"to",
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

	varValidateCertificates200ResponseCertificateDates := _ValidateCertificates200ResponseCertificateDates{}

	err = json.Unmarshal(data, &varValidateCertificates200ResponseCertificateDates)

	if err != nil {
		return err
	}

	*o = ValidateCertificates200ResponseCertificateDates(varValidateCertificates200ResponseCertificateDates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateCertificates200ResponseCertificateDates struct {
	value *ValidateCertificates200ResponseCertificateDates
	isSet bool
}

func (v NullableValidateCertificates200ResponseCertificateDates) Get() *ValidateCertificates200ResponseCertificateDates {
	return v.value
}

func (v *NullableValidateCertificates200ResponseCertificateDates) Set(val *ValidateCertificates200ResponseCertificateDates) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateCertificates200ResponseCertificateDates) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateCertificates200ResponseCertificateDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateCertificates200ResponseCertificateDates(val *ValidateCertificates200ResponseCertificateDates) *NullableValidateCertificates200ResponseCertificateDates {
	return &NullableValidateCertificates200ResponseCertificateDates{value: val, isSet: true}
}

func (v NullableValidateCertificates200ResponseCertificateDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateCertificates200ResponseCertificateDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


