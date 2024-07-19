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

// checks if the AccessListUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessListUpdateResponse{}

// AccessListUpdateResponse struct for AccessListUpdateResponse
type AccessListUpdateResponse struct {
	// Unique identifier
	Id *int64 `json:"id,omitempty"`
	// Date and time of creation
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// Date and time of last update
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	// Name of the Access List
	Name *string `json:"name,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// NewAccessListUpdateResponse instantiates a new AccessListUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessListUpdateResponse() *AccessListUpdateResponse {
	this := AccessListUpdateResponse{}
	return &this
}

// NewAccessListUpdateResponseWithDefaults instantiates a new AccessListUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessListUpdateResponseWithDefaults() *AccessListUpdateResponse {
	this := AccessListUpdateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessListUpdateResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListUpdateResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessListUpdateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AccessListUpdateResponse) SetId(v int64) {
	o.Id = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *AccessListUpdateResponse) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListUpdateResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *AccessListUpdateResponse) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *AccessListUpdateResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *AccessListUpdateResponse) GetModifiedOn() time.Time {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListUpdateResponse) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *AccessListUpdateResponse) HasModifiedOn() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *AccessListUpdateResponse) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessListUpdateResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListUpdateResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessListUpdateResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessListUpdateResponse) SetName(v string) {
	o.Name = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AccessListUpdateResponse) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListUpdateResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AccessListUpdateResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *AccessListUpdateResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o AccessListUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessListUpdateResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableAccessListUpdateResponse struct {
	value *AccessListUpdateResponse
	isSet bool
}

func (v NullableAccessListUpdateResponse) Get() *AccessListUpdateResponse {
	return v.value
}

func (v *NullableAccessListUpdateResponse) Set(val *AccessListUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessListUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessListUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessListUpdateResponse(val *AccessListUpdateResponse) *NullableAccessListUpdateResponse {
	return &NullableAccessListUpdateResponse{value: val, isSet: true}
}

func (v NullableAccessListUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessListUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


