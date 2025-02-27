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

// checks if the GetAuditLog200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAuditLog200Response{}

// GetAuditLog200Response Audit Log object
type GetAuditLog200Response struct {
	// Unique identifier
	Id int64 `json:"id"`
	// Date and time of creation
	CreatedOn string `json:"created_on"`
	// Date and time of last update
	ModifiedOn string `json:"modified_on"`
	// User ID
	UserId int64 `json:"user_id"`
	ObjectType string `json:"object_type"`
	// Unique identifier
	ObjectId int64 `json:"object_id"`
	Action string `json:"action"`
	Meta map[string]interface{} `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _GetAuditLog200Response GetAuditLog200Response

// NewGetAuditLog200Response instantiates a new GetAuditLog200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAuditLog200Response(id int64, createdOn string, modifiedOn string, userId int64, objectType string, objectId int64, action string, meta map[string]interface{}) *GetAuditLog200Response {
	this := GetAuditLog200Response{}
	this.Id = id
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	this.UserId = userId
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Action = action
	this.Meta = meta
	return &this
}

// NewGetAuditLog200ResponseWithDefaults instantiates a new GetAuditLog200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAuditLog200ResponseWithDefaults() *GetAuditLog200Response {
	this := GetAuditLog200Response{}
	return &this
}

// GetId returns the Id field value
func (o *GetAuditLog200Response) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetAuditLog200Response) SetId(v int64) {
	o.Id = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *GetAuditLog200Response) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *GetAuditLog200Response) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *GetAuditLog200Response) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *GetAuditLog200Response) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetUserId returns the UserId field value
func (o *GetAuditLog200Response) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GetAuditLog200Response) SetUserId(v int64) {
	o.UserId = v
}

// GetObjectType returns the ObjectType field value
func (o *GetAuditLog200Response) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *GetAuditLog200Response) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *GetAuditLog200Response) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *GetAuditLog200Response) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetAction returns the Action field value
func (o *GetAuditLog200Response) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GetAuditLog200Response) SetAction(v string) {
	o.Action = v
}

// GetMeta returns the Meta field value
func (o *GetAuditLog200Response) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *GetAuditLog200Response) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// SetMeta sets field value
func (o *GetAuditLog200Response) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o GetAuditLog200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAuditLog200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["user_id"] = o.UserId
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["action"] = o.Action
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAuditLog200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_on",
		"modified_on",
		"user_id",
		"object_type",
		"object_id",
		"action",
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

	varGetAuditLog200Response := _GetAuditLog200Response{}

	err = json.Unmarshal(data, &varGetAuditLog200Response)

	if err != nil {
		return err
	}

	*o = GetAuditLog200Response(varGetAuditLog200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "modified_on")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "action")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAuditLog200Response struct {
	value *GetAuditLog200Response
	isSet bool
}

func (v NullableGetAuditLog200Response) Get() *GetAuditLog200Response {
	return v.value
}

func (v *NullableGetAuditLog200Response) Set(val *GetAuditLog200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAuditLog200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAuditLog200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAuditLog200Response(val *GetAuditLog200Response) *NullableGetAuditLog200Response {
	return &NullableGetAuditLog200Response{value: val, isSet: true}
}

func (v NullableGetAuditLog200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAuditLog200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


