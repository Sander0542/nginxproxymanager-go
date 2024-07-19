# UpdateSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**UpdateSettingRequestMeta**](UpdateSettingRequestMeta.md) |  | [optional] 

## Methods

### NewUpdateSettingRequest

`func NewUpdateSettingRequest() *UpdateSettingRequest`

NewUpdateSettingRequest instantiates a new UpdateSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSettingRequestWithDefaults

`func NewUpdateSettingRequestWithDefaults() *UpdateSettingRequest`

NewUpdateSettingRequestWithDefaults instantiates a new UpdateSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateSettingRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSettingRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSettingRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateSettingRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateSettingRequest) GetMeta() UpdateSettingRequestMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateSettingRequest) GetMetaOk() (*UpdateSettingRequestMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateSettingRequest) SetMeta(v UpdateSettingRequestMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateSettingRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


