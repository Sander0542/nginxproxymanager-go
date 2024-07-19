# SettingGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for a Setting | [optional] [readonly] 
**Name** | Pointer to **string** | Description | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Value** | Pointer to **string** | Value | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSettingGetResponse

`func NewSettingGetResponse() *SettingGetResponse`

NewSettingGetResponse instantiates a new SettingGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingGetResponseWithDefaults

`func NewSettingGetResponseWithDefaults() *SettingGetResponse`

NewSettingGetResponseWithDefaults instantiates a new SettingGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettingGetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettingGetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettingGetResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SettingGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SettingGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingGetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingGetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SettingGetResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SettingGetResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SettingGetResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SettingGetResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *SettingGetResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingGetResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingGetResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SettingGetResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMeta

`func (o *SettingGetResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SettingGetResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SettingGetResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SettingGetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


