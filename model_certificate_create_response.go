/*
Nginx Proxy Manager API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.x.x
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nginxproxymanager

import (
	"encoding/json"
	"time"
)

// checks if the CertificateCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateCreateResponse{}

// CertificateCreateResponse struct for CertificateCreateResponse
type CertificateCreateResponse struct {
	// Unique identifier
	Id *int64 `json:"id,omitempty"`
	// Date and time of creation
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// Date and time of last update
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	Provider *string `json:"provider,omitempty" validate:"regexp=^(letsencrypt|other)$"`
	// Nice Name for the custom certificate
	NiceName *string `json:"nice_name,omitempty"`
	// Domain Names separated by a comma
	DomainNames []string `json:"domain_names,omitempty"`
	// Date and time of expiration
	ExpiresOn *time.Time `json:"expires_on,omitempty"`
	Meta *PostCertificateCreateRequestMeta `json:"meta,omitempty"`
}

// NewCertificateCreateResponse instantiates a new CertificateCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateCreateResponse() *CertificateCreateResponse {
	this := CertificateCreateResponse{}
	return &this
}

// NewCertificateCreateResponseWithDefaults instantiates a new CertificateCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateCreateResponseWithDefaults() *CertificateCreateResponse {
	this := CertificateCreateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CertificateCreateResponse) SetId(v int64) {
	o.Id = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *CertificateCreateResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetModifiedOn() time.Time {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasModifiedOn() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *CertificateCreateResponse) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CertificateCreateResponse) SetProvider(v string) {
	o.Provider = &v
}

// GetNiceName returns the NiceName field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetNiceName() string {
	if o == nil || IsNil(o.NiceName) {
		var ret string
		return ret
	}
	return *o.NiceName
}

// GetNiceNameOk returns a tuple with the NiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetNiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.NiceName) {
		return nil, false
	}
	return o.NiceName, true
}

// HasNiceName returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasNiceName() bool {
	if o != nil && !IsNil(o.NiceName) {
		return true
	}

	return false
}

// SetNiceName gets a reference to the given string and assigns it to the NiceName field.
func (o *CertificateCreateResponse) SetNiceName(v string) {
	o.NiceName = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *CertificateCreateResponse) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetExpiresOn() time.Time {
	if o == nil || IsNil(o.ExpiresOn) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresOn) {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasExpiresOn() bool {
	if o != nil && !IsNil(o.ExpiresOn) {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given time.Time and assigns it to the ExpiresOn field.
func (o *CertificateCreateResponse) SetExpiresOn(v time.Time) {
	o.ExpiresOn = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CertificateCreateResponse) GetMeta() PostCertificateCreateRequestMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PostCertificateCreateRequestMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreateResponse) GetMetaOk() (*PostCertificateCreateRequestMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CertificateCreateResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PostCertificateCreateRequestMeta and assigns it to the Meta field.
func (o *CertificateCreateResponse) SetMeta(v PostCertificateCreateRequestMeta) {
	o.Meta = &v
}

func (o CertificateCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["created_on"] = o.CreatedOn
	}
	if !IsNil(o.ModifiedOn) {
		toSerialize["modified_on"] = o.ModifiedOn
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.NiceName) {
		toSerialize["nice_name"] = o.NiceName
	}
	if !IsNil(o.DomainNames) {
		toSerialize["domain_names"] = o.DomainNames
	}
	if !IsNil(o.ExpiresOn) {
		toSerialize["expires_on"] = o.ExpiresOn
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableCertificateCreateResponse struct {
	value *CertificateCreateResponse
	isSet bool
}

func (v NullableCertificateCreateResponse) Get() *CertificateCreateResponse {
	return v.value
}

func (v *NullableCertificateCreateResponse) Set(val *CertificateCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateCreateResponse(val *CertificateCreateResponse) *NullableCertificateCreateResponse {
	return &NullableCertificateCreateResponse{value: val, isSet: true}
}

func (v NullableCertificateCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


