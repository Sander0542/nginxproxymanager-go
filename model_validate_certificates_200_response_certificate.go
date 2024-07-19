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

// checks if the ValidateCertificates200ResponseCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateCertificates200ResponseCertificate{}

// ValidateCertificates200ResponseCertificate struct for ValidateCertificates200ResponseCertificate
type ValidateCertificates200ResponseCertificate struct {
	Cn string `json:"cn"`
	Issuer string `json:"issuer"`
	Dates ValidateCertificates200ResponseCertificateDates `json:"dates"`
	AdditionalProperties map[string]interface{}
}

type _ValidateCertificates200ResponseCertificate ValidateCertificates200ResponseCertificate

// NewValidateCertificates200ResponseCertificate instantiates a new ValidateCertificates200ResponseCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateCertificates200ResponseCertificate(cn string, issuer string, dates ValidateCertificates200ResponseCertificateDates) *ValidateCertificates200ResponseCertificate {
	this := ValidateCertificates200ResponseCertificate{}
	this.Cn = cn
	this.Issuer = issuer
	this.Dates = dates
	return &this
}

// NewValidateCertificates200ResponseCertificateWithDefaults instantiates a new ValidateCertificates200ResponseCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateCertificates200ResponseCertificateWithDefaults() *ValidateCertificates200ResponseCertificate {
	this := ValidateCertificates200ResponseCertificate{}
	return &this
}

// GetCn returns the Cn field value
func (o *ValidateCertificates200ResponseCertificate) GetCn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cn
}

// GetCnOk returns a tuple with the Cn field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200ResponseCertificate) GetCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cn, true
}

// SetCn sets field value
func (o *ValidateCertificates200ResponseCertificate) SetCn(v string) {
	o.Cn = v
}

// GetIssuer returns the Issuer field value
func (o *ValidateCertificates200ResponseCertificate) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200ResponseCertificate) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *ValidateCertificates200ResponseCertificate) SetIssuer(v string) {
	o.Issuer = v
}

// GetDates returns the Dates field value
func (o *ValidateCertificates200ResponseCertificate) GetDates() ValidateCertificates200ResponseCertificateDates {
	if o == nil {
		var ret ValidateCertificates200ResponseCertificateDates
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200ResponseCertificate) GetDatesOk() (*ValidateCertificates200ResponseCertificateDates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dates, true
}

// SetDates sets field value
func (o *ValidateCertificates200ResponseCertificate) SetDates(v ValidateCertificates200ResponseCertificateDates) {
	o.Dates = v
}

func (o ValidateCertificates200ResponseCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateCertificates200ResponseCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cn"] = o.Cn
	toSerialize["issuer"] = o.Issuer
	toSerialize["dates"] = o.Dates

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateCertificates200ResponseCertificate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cn",
		"issuer",
		"dates",
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

	varValidateCertificates200ResponseCertificate := _ValidateCertificates200ResponseCertificate{}

	err = json.Unmarshal(data, &varValidateCertificates200ResponseCertificate)

	if err != nil {
		return err
	}

	*o = ValidateCertificates200ResponseCertificate(varValidateCertificates200ResponseCertificate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cn")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "dates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateCertificates200ResponseCertificate struct {
	value *ValidateCertificates200ResponseCertificate
	isSet bool
}

func (v NullableValidateCertificates200ResponseCertificate) Get() *ValidateCertificates200ResponseCertificate {
	return v.value
}

func (v *NullableValidateCertificates200ResponseCertificate) Set(val *ValidateCertificates200ResponseCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateCertificates200ResponseCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateCertificates200ResponseCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateCertificates200ResponseCertificate(val *ValidateCertificates200ResponseCertificate) *NullableValidateCertificates200ResponseCertificate {
	return &NullableValidateCertificates200ResponseCertificate{value: val, isSet: true}
}

func (v NullableValidateCertificates200ResponseCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateCertificates200ResponseCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


