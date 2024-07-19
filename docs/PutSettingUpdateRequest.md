# PutSettingUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Value | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPutSettingUpdateRequest

`func NewPutSettingUpdateRequest() *PutSettingUpdateRequest`

NewPutSettingUpdateRequest instantiates a new PutSettingUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutSettingUpdateRequestWithDefaults

`func NewPutSettingUpdateRequestWithDefaults() *PutSettingUpdateRequest`

NewPutSettingUpdateRequestWithDefaults instantiates a new PutSettingUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PutSettingUpdateRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PutSettingUpdateRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PutSettingUpdateRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PutSettingUpdateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMeta

`func (o *PutSettingUpdateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PutSettingUpdateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PutSettingUpdateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PutSettingUpdateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


