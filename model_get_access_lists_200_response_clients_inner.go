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

// checks if the GetAccessLists200ResponseClientsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccessLists200ResponseClientsInner{}

// GetAccessLists200ResponseClientsInner struct for GetAccessLists200ResponseClientsInner
type GetAccessLists200ResponseClientsInner struct {
	// Unique identifier
	Id int64 `json:"id"`
	// Date and time of creation
	CreatedOn string `json:"created_on"`
	// Date and time of last update
	ModifiedOn string `json:"modified_on"`
	// Access List ID
	AccessListId int64 `json:"access_list_id"`
	Address GetAccessLists200ResponseClientsInnerAddress `json:"address"`
	Directive string `json:"directive"`
	Meta map[string]interface{} `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _GetAccessLists200ResponseClientsInner GetAccessLists200ResponseClientsInner

// NewGetAccessLists200ResponseClientsInner instantiates a new GetAccessLists200ResponseClientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccessLists200ResponseClientsInner(id int64, createdOn string, modifiedOn string, accessListId int64, address GetAccessLists200ResponseClientsInnerAddress, directive string, meta map[string]interface{}) *GetAccessLists200ResponseClientsInner {
	this := GetAccessLists200ResponseClientsInner{}
	this.Id = id
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	this.AccessListId = accessListId
	this.Address = address
	this.Directive = directive
	this.Meta = meta
	return &this
}

// NewGetAccessLists200ResponseClientsInnerWithDefaults instantiates a new GetAccessLists200ResponseClientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccessLists200ResponseClientsInnerWithDefaults() *GetAccessLists200ResponseClientsInner {
	this := GetAccessLists200ResponseClientsInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetAccessLists200ResponseClientsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetAccessLists200ResponseClientsInner) SetId(v int64) {
	o.Id = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *GetAccessLists200ResponseClientsInner) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *GetAccessLists200ResponseClientsInner) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *GetAccessLists200ResponseClientsInner) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *GetAccessLists200ResponseClientsInner) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetAccessListId returns the AccessListId field value
func (o *GetAccessLists200ResponseClientsInner) GetAccessListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccessListId
}

// GetAccessListIdOk returns a tuple with the AccessListId field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetAccessListIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessListId, true
}

// SetAccessListId sets field value
func (o *GetAccessLists200ResponseClientsInner) SetAccessListId(v int64) {
	o.AccessListId = v
}

// GetAddress returns the Address field value
func (o *GetAccessLists200ResponseClientsInner) GetAddress() GetAccessLists200ResponseClientsInnerAddress {
	if o == nil {
		var ret GetAccessLists200ResponseClientsInnerAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetAddressOk() (*GetAccessLists200ResponseClientsInnerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *GetAccessLists200ResponseClientsInner) SetAddress(v GetAccessLists200ResponseClientsInnerAddress) {
	o.Address = v
}

// GetDirective returns the Directive field value
func (o *GetAccessLists200ResponseClientsInner) GetDirective() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Directive
}

// GetDirectiveOk returns a tuple with the Directive field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetDirectiveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Directive, true
}

// SetDirective sets field value
func (o *GetAccessLists200ResponseClientsInner) SetDirective(v string) {
	o.Directive = v
}

// GetMeta returns the Meta field value
func (o *GetAccessLists200ResponseClientsInner) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseClientsInner) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetAccessLists200ResponseClientsInner) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o GetAccessLists200ResponseClientsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccessLists200ResponseClientsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["access_list_id"] = o.AccessListId
	toSerialize["address"] = o.Address
	toSerialize["directive"] = o.Directive
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAccessLists200ResponseClientsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_on",
		"modified_on",
		"access_list_id",
		"address",
		"directive",
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

	varGetAccessLists200ResponseClientsInner := _GetAccessLists200ResponseClientsInner{}

	err = json.Unmarshal(data, &varGetAccessLists200ResponseClientsInner)

	if err != nil {
		return err
	}

	*o = GetAccessLists200ResponseClientsInner(varGetAccessLists200ResponseClientsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "modified_on")
		delete(additionalProperties, "access_list_id")
		delete(additionalProperties, "address")
		delete(additionalProperties, "directive")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAccessLists200ResponseClientsInner struct {
	value *GetAccessLists200ResponseClientsInner
	isSet bool
}

func (v NullableGetAccessLists200ResponseClientsInner) Get() *GetAccessLists200ResponseClientsInner {
	return v.value
}

func (v *NullableGetAccessLists200ResponseClientsInner) Set(val *GetAccessLists200ResponseClientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccessLists200ResponseClientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccessLists200ResponseClientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccessLists200ResponseClientsInner(val *GetAccessLists200ResponseClientsInner) *NullableGetAccessLists200ResponseClientsInner {
	return &NullableGetAccessLists200ResponseClientsInner{value: val, isSet: true}
}

func (v NullableGetAccessLists200ResponseClientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccessLists200ResponseClientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


