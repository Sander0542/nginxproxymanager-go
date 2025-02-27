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

// checks if the GetAccessLists200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccessLists200ResponseInner{}

// GetAccessLists200ResponseInner Access List object
type GetAccessLists200ResponseInner struct {
	// Unique identifier
	Id int64 `json:"id"`
	// Date and time of creation
	CreatedOn string `json:"created_on"`
	// Date and time of last update
	ModifiedOn string `json:"modified_on"`
	// User ID
	OwnerUserId int64 `json:"owner_user_id"`
	Name string `json:"name"`
	SatisfyAny bool `json:"satisfy_any"`
	PassAuth bool `json:"pass_auth"`
	Meta map[string]interface{} `json:"meta"`
	Owner *GetAccessLists200ResponseInnerOwner `json:"owner,omitempty"`
	Items []GetAccessLists200ResponseInnerItemsInner `json:"items,omitempty"`
	Clients []GetAccessLists200ResponseInnerClientsInner `json:"clients,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAccessLists200ResponseInner GetAccessLists200ResponseInner

// NewGetAccessLists200ResponseInner instantiates a new GetAccessLists200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccessLists200ResponseInner(id int64, createdOn string, modifiedOn string, ownerUserId int64, name string, satisfyAny bool, passAuth bool, meta map[string]interface{}) *GetAccessLists200ResponseInner {
	this := GetAccessLists200ResponseInner{}
	this.Id = id
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	this.OwnerUserId = ownerUserId
	this.Name = name
	this.SatisfyAny = satisfyAny
	this.PassAuth = passAuth
	this.Meta = meta
	return &this
}

// NewGetAccessLists200ResponseInnerWithDefaults instantiates a new GetAccessLists200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccessLists200ResponseInnerWithDefaults() *GetAccessLists200ResponseInner {
	this := GetAccessLists200ResponseInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetAccessLists200ResponseInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetAccessLists200ResponseInner) SetId(v int64) {
	o.Id = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *GetAccessLists200ResponseInner) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *GetAccessLists200ResponseInner) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *GetAccessLists200ResponseInner) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *GetAccessLists200ResponseInner) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetOwnerUserId returns the OwnerUserId field value
func (o *GetAccessLists200ResponseInner) GetOwnerUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetOwnerUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUserId, true
}

// SetOwnerUserId sets field value
func (o *GetAccessLists200ResponseInner) SetOwnerUserId(v int64) {
	o.OwnerUserId = v
}

// GetName returns the Name field value
func (o *GetAccessLists200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetAccessLists200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetSatisfyAny returns the SatisfyAny field value
func (o *GetAccessLists200ResponseInner) GetSatisfyAny() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SatisfyAny
}

// GetSatisfyAnyOk returns a tuple with the SatisfyAny field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetSatisfyAnyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SatisfyAny, true
}

// SetSatisfyAny sets field value
func (o *GetAccessLists200ResponseInner) SetSatisfyAny(v bool) {
	o.SatisfyAny = v
}

// GetPassAuth returns the PassAuth field value
func (o *GetAccessLists200ResponseInner) GetPassAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PassAuth
}

// GetPassAuthOk returns a tuple with the PassAuth field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetPassAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PassAuth, true
}

// SetPassAuth sets field value
func (o *GetAccessLists200ResponseInner) SetPassAuth(v bool) {
	o.PassAuth = v
}

// GetMeta returns the Meta field value
func (o *GetAccessLists200ResponseInner) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetAccessLists200ResponseInner) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GetAccessLists200ResponseInner) GetOwner() GetAccessLists200ResponseInnerOwner {
	if o == nil || IsNil(o.Owner) {
		var ret GetAccessLists200ResponseInnerOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetOwnerOk() (*GetAccessLists200ResponseInnerOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GetAccessLists200ResponseInner) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given GetAccessLists200ResponseInnerOwner and assigns it to the Owner field.
func (o *GetAccessLists200ResponseInner) SetOwner(v GetAccessLists200ResponseInnerOwner) {
	o.Owner = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetAccessLists200ResponseInner) GetItems() []GetAccessLists200ResponseInnerItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []GetAccessLists200ResponseInnerItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetItemsOk() ([]GetAccessLists200ResponseInnerItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetAccessLists200ResponseInner) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []GetAccessLists200ResponseInnerItemsInner and assigns it to the Items field.
func (o *GetAccessLists200ResponseInner) SetItems(v []GetAccessLists200ResponseInnerItemsInner) {
	o.Items = v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *GetAccessLists200ResponseInner) GetClients() []GetAccessLists200ResponseInnerClientsInner {
	if o == nil || IsNil(o.Clients) {
		var ret []GetAccessLists200ResponseInnerClientsInner
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccessLists200ResponseInner) GetClientsOk() ([]GetAccessLists200ResponseInnerClientsInner, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *GetAccessLists200ResponseInner) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []GetAccessLists200ResponseInnerClientsInner and assigns it to the Clients field.
func (o *GetAccessLists200ResponseInner) SetClients(v []GetAccessLists200ResponseInnerClientsInner) {
	o.Clients = v
}

func (o GetAccessLists200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccessLists200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["owner_user_id"] = o.OwnerUserId
	toSerialize["name"] = o.Name
	toSerialize["satisfy_any"] = o.SatisfyAny
	toSerialize["pass_auth"] = o.PassAuth
	toSerialize["meta"] = o.Meta
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAccessLists200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_on",
		"modified_on",
		"owner_user_id",
		"name",
		"satisfy_any",
		"pass_auth",
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

	varGetAccessLists200ResponseInner := _GetAccessLists200ResponseInner{}

	err = json.Unmarshal(data, &varGetAccessLists200ResponseInner)

	if err != nil {
		return err
	}

	*o = GetAccessLists200ResponseInner(varGetAccessLists200ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "modified_on")
		delete(additionalProperties, "owner_user_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "satisfy_any")
		delete(additionalProperties, "pass_auth")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "items")
		delete(additionalProperties, "clients")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAccessLists200ResponseInner struct {
	value *GetAccessLists200ResponseInner
	isSet bool
}

func (v NullableGetAccessLists200ResponseInner) Get() *GetAccessLists200ResponseInner {
	return v.value
}

func (v *NullableGetAccessLists200ResponseInner) Set(val *GetAccessLists200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccessLists200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccessLists200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccessLists200ResponseInner(val *GetAccessLists200ResponseInner) *NullableGetAccessLists200ResponseInner {
	return &NullableGetAccessLists200ResponseInner{value: val, isSet: true}
}

func (v NullableGetAccessLists200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccessLists200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


