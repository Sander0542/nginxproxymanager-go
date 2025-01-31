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

// checks if the GetSetting200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSetting200Response{}

// GetSetting200Response Setting object
type GetSetting200Response struct {
	// Setting ID
	Id string `json:"id"`
	// Setting Display Name
	Name string `json:"name"`
	// Meaningful description
	Description string `json:"description"`
	Value GetSetting200ResponseValue `json:"value"`
	// Extra metadata
	Meta map[string]interface{} `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _GetSetting200Response GetSetting200Response

// NewGetSetting200Response instantiates a new GetSetting200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSetting200Response(id string, name string, description string, value GetSetting200ResponseValue, meta map[string]interface{}) *GetSetting200Response {
	this := GetSetting200Response{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Value = value
	this.Meta = meta
	return &this
}

// NewGetSetting200ResponseWithDefaults instantiates a new GetSetting200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSetting200ResponseWithDefaults() *GetSetting200Response {
	this := GetSetting200Response{}
	return &this
}

// GetId returns the Id field value
func (o *GetSetting200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSetting200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSetting200Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetSetting200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetSetting200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetSetting200Response) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *GetSetting200Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetSetting200Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetSetting200Response) SetDescription(v string) {
	o.Description = v
}

// GetValue returns the Value field value
func (o *GetSetting200Response) GetValue() GetSetting200ResponseValue {
	if o == nil {
		var ret GetSetting200ResponseValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetSetting200Response) GetValueOk() (*GetSetting200ResponseValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetSetting200Response) SetValue(v GetSetting200ResponseValue) {
	o.Value = v
}

// GetMeta returns the Meta field value
func (o *GetSetting200Response) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetSetting200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetSetting200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o GetSetting200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSetting200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["value"] = o.Value
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSetting200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"value",
		"meta",
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

	varGetSetting200Response := _GetSetting200Response{}

	err = json.Unmarshal(data, &varGetSetting200Response)

	if err != nil {
		return err
	}

	*o = GetSetting200Response(varGetSetting200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "value")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSetting200Response struct {
	value *GetSetting200Response
	isSet bool
}

func (v NullableGetSetting200Response) Get() *GetSetting200Response {
	return v.value
}

func (v *NullableGetSetting200Response) Set(val *GetSetting200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSetting200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSetting200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSetting200Response(val *GetSetting200Response) *NullableGetSetting200Response {
	return &NullableGetSetting200Response{value: val, isSet: true}
}

func (v NullableGetSetting200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSetting200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


