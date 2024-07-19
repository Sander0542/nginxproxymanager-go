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

// checks if the Health200ResponseVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Health200ResponseVersion{}

// Health200ResponseVersion The version object
type Health200ResponseVersion struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Revision int64 `json:"revision"`
	AdditionalProperties map[string]interface{}
}

type _Health200ResponseVersion Health200ResponseVersion

// NewHealth200ResponseVersion instantiates a new Health200ResponseVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealth200ResponseVersion(major int64, minor int64, revision int64) *Health200ResponseVersion {
	this := Health200ResponseVersion{}
	this.Major = major
	this.Minor = minor
	this.Revision = revision
	return &this
}

// NewHealth200ResponseVersionWithDefaults instantiates a new Health200ResponseVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealth200ResponseVersionWithDefaults() *Health200ResponseVersion {
	this := Health200ResponseVersion{}
	return &this
}

// GetMajor returns the Major field value
func (o *Health200ResponseVersion) GetMajor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Major
}

// GetMajorOk returns a tuple with the Major field value
// and a boolean to check if the value has been set.
func (o *Health200ResponseVersion) GetMajorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Major, true
}

// SetMajor sets field value
func (o *Health200ResponseVersion) SetMajor(v int64) {
	o.Major = v
}

// GetMinor returns the Minor field value
func (o *Health200ResponseVersion) GetMinor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Minor
}

// GetMinorOk returns a tuple with the Minor field value
// and a boolean to check if the value has been set.
func (o *Health200ResponseVersion) GetMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minor, true
}

// SetMinor sets field value
func (o *Health200ResponseVersion) SetMinor(v int64) {
	o.Minor = v
}

// GetRevision returns the Revision field value
func (o *Health200ResponseVersion) GetRevision() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *Health200ResponseVersion) GetRevisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *Health200ResponseVersion) SetRevision(v int64) {
	o.Revision = v
}

func (o Health200ResponseVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Health200ResponseVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["major"] = o.Major
	toSerialize["minor"] = o.Minor
	toSerialize["revision"] = o.Revision

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Health200ResponseVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"major",
		"minor",
		"revision",
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

	varHealth200ResponseVersion := _Health200ResponseVersion{}

	err = json.Unmarshal(data, &varHealth200ResponseVersion)

	if err != nil {
		return err
	}

	*o = Health200ResponseVersion(varHealth200ResponseVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "major")
		delete(additionalProperties, "minor")
		delete(additionalProperties, "revision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHealth200ResponseVersion struct {
	value *Health200ResponseVersion
	isSet bool
}

func (v NullableHealth200ResponseVersion) Get() *Health200ResponseVersion {
	return v.value
}

func (v *NullableHealth200ResponseVersion) Set(val *Health200ResponseVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableHealth200ResponseVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableHealth200ResponseVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealth200ResponseVersion(val *Health200ResponseVersion) *NullableHealth200ResponseVersion {
	return &NullableHealth200ResponseVersion{value: val, isSet: true}
}

func (v NullableHealth200ResponseVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealth200ResponseVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


