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

// checks if the ValidateCertificates200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateCertificates200Response{}

// ValidateCertificates200Response struct for ValidateCertificates200Response
type ValidateCertificates200Response struct {
	Certificate ValidateCertificates200ResponseCertificate `json:"certificate"`
	CertificateKey bool `json:"certificate_key"`
	AdditionalProperties map[string]interface{}
}

type _ValidateCertificates200Response ValidateCertificates200Response

// NewValidateCertificates200Response instantiates a new ValidateCertificates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateCertificates200Response(certificate ValidateCertificates200ResponseCertificate, certificateKey bool) *ValidateCertificates200Response {
	this := ValidateCertificates200Response{}
	this.Certificate = certificate
	this.CertificateKey = certificateKey
	return &this
}

// NewValidateCertificates200ResponseWithDefaults instantiates a new ValidateCertificates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateCertificates200ResponseWithDefaults() *ValidateCertificates200Response {
	this := ValidateCertificates200Response{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *ValidateCertificates200Response) GetCertificate() ValidateCertificates200ResponseCertificate {
	if o == nil {
		var ret ValidateCertificates200ResponseCertificate
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200Response) GetCertificateOk() (*ValidateCertificates200ResponseCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *ValidateCertificates200Response) SetCertificate(v ValidateCertificates200ResponseCertificate) {
	o.Certificate = v
}

// GetCertificateKey returns the CertificateKey field value
func (o *ValidateCertificates200Response) GetCertificateKey() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CertificateKey
}

// GetCertificateKeyOk returns a tuple with the CertificateKey field value
// and a boolean to check if the value has been set.
func (o *ValidateCertificates200Response) GetCertificateKeyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateKey, true
}

// SetCertificateKey sets field value
func (o *ValidateCertificates200Response) SetCertificateKey(v bool) {
	o.CertificateKey = v
}

func (o ValidateCertificates200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateCertificates200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificate"] = o.Certificate
	toSerialize["certificate_key"] = o.CertificateKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateCertificates200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificate",
		"certificate_key",
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

	varValidateCertificates200Response := _ValidateCertificates200Response{}

	err = json.Unmarshal(data, &varValidateCertificates200Response)

	if err != nil {
		return err
	}

	*o = ValidateCertificates200Response(varValidateCertificates200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "certificate_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateCertificates200Response struct {
	value *ValidateCertificates200Response
	isSet bool
}

func (v NullableValidateCertificates200Response) Get() *ValidateCertificates200Response {
	return v.value
}

func (v *NullableValidateCertificates200Response) Set(val *ValidateCertificates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateCertificates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateCertificates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateCertificates200Response(val *ValidateCertificates200Response) *NullableValidateCertificates200Response {
	return &NullableValidateCertificates200Response{value: val, isSet: true}
}

func (v NullableValidateCertificates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateCertificates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


