# SettingUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for a Setting | [optional] [readonly] 
**Name** | Pointer to **string** | Description | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Value** | Pointer to **string** | Value | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSettingUpdateResponse

`func NewSettingUpdateResponse() *SettingUpdateResponse`

NewSettingUpdateResponse instantiates a new SettingUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingUpdateResponseWithDefaults

`func NewSettingUpdateResponseWithDefaults() *SettingUpdateResponse`

NewSettingUpdateResponseWithDefaults instantiates a new SettingUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettingUpdateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettingUpdateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettingUpdateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SettingUpdateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SettingUpdateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingUpdateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingUpdateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingUpdateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SettingUpdateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SettingUpdateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SettingUpdateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SettingUpdateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *SettingUpdateResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingUpdateResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingUpdateResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SettingUpdateResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMeta

`func (o *SettingUpdateResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SettingUpdateResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SettingUpdateResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SettingUpdateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


