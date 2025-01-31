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

// checks if the ReportsHosts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsHosts200Response{}

// ReportsHosts200Response struct for ReportsHosts200Response
type ReportsHosts200Response struct {
	// Proxy Hosts Count
	Proxy *int64 `json:"proxy,omitempty"`
	// Redirection Hosts Count
	Redirection *int64 `json:"redirection,omitempty"`
	// Streams Count
	Stream *int64 `json:"stream,omitempty"`
	// 404 Hosts Count
	Dead *int64 `json:"dead,omitempty"`
}

// NewReportsHosts200Response instantiates a new ReportsHosts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsHosts200Response() *ReportsHosts200Response {
	this := ReportsHosts200Response{}
	return &this
}

// NewReportsHosts200ResponseWithDefaults instantiates a new ReportsHosts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsHosts200ResponseWithDefaults() *ReportsHosts200Response {
	this := ReportsHosts200Response{}
	return &this
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *ReportsHosts200Response) GetProxy() int64 {
	if o == nil || IsNil(o.Proxy) {
		var ret int64
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsHosts200Response) GetProxyOk() (*int64, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *ReportsHosts200Response) HasProxy() bool {
	if o != nil && !IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given int64 and assigns it to the Proxy field.
func (o *ReportsHosts200Response) SetProxy(v int64) {
	o.Proxy = &v
}

// GetRedirection returns the Redirection field value if set, zero value otherwise.
func (o *ReportsHosts200Response) GetRedirection() int64 {
	if o == nil || IsNil(o.Redirection) {
		var ret int64
		return ret
	}
	return *o.Redirection
}

// GetRedirectionOk returns a tuple with the Redirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsHosts200Response) GetRedirectionOk() (*int64, bool) {
	if o == nil || IsNil(o.Redirection) {
		return nil, false
	}
	return o.Redirection, true
}

// HasRedirection returns a boolean if a field has been set.
func (o *ReportsHosts200Response) HasRedirection() bool {
	if o != nil && !IsNil(o.Redirection) {
		return true
	}

	return false
}

// SetRedirection gets a reference to the given int64 and assigns it to the Redirection field.
func (o *ReportsHosts200Response) SetRedirection(v int64) {
	o.Redirection = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ReportsHosts200Response) GetStream() int64 {
	if o == nil || IsNil(o.Stream) {
		var ret int64
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsHosts200Response) GetStreamOk() (*int64, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ReportsHosts200Response) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given int64 and assigns it to the Stream field.
func (o *ReportsHosts200Response) SetStream(v int64) {
	o.Stream = &v
}

// GetDead returns the Dead field value if set, zero value otherwise.
func (o *ReportsHosts200Response) GetDead() int64 {
	if o == nil || IsNil(o.Dead) {
		var ret int64
		return ret
	}
	return *o.Dead
}

// GetDeadOk returns a tuple with the Dead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsHosts200Response) GetDeadOk() (*int64, bool) {
	if o == nil || IsNil(o.Dead) {
		return nil, false
	}
	return o.Dead, true
}

// HasDead returns a boolean if a field has been set.
func (o *ReportsHosts200Response) HasDead() bool {
	if o != nil && !IsNil(o.Dead) {
		return true
	}

	return false
}

// SetDead gets a reference to the given int64 and assigns it to the Dead field.
func (o *ReportsHosts200Response) SetDead(v int64) {
	o.Dead = &v
}

func (o ReportsHosts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsHosts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !IsNil(o.Redirection) {
		toSerialize["redirection"] = o.Redirection
	}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	if !IsNil(o.Dead) {
		toSerialize["dead"] = o.Dead
	}
	return toSerialize, nil
}

type NullableReportsHosts200Response struct {
	value *ReportsHosts200Response
	isSet bool
}

func (v NullableReportsHosts200Response) Get() *ReportsHosts200Response {
	return v.value
}

func (v *NullableReportsHosts200Response) Set(val *ReportsHosts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsHosts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsHosts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsHosts200Response(val *ReportsHosts200Response) *NullableReportsHosts200Response {
	return &NullableReportsHosts200Response{value: val, isSet: true}
}

func (v NullableReportsHosts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsHosts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


