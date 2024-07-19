# GetSetting200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Setting ID | 
**Name** | **string** | Setting Display Name | 
**Description** | **string** | Meaningful description | 
**Value** | [**GetSetting200ResponseValue**](GetSetting200ResponseValue.md) |  | 
**Meta** | **map[string]interface{}** | Extra metadata | 

## Methods

### NewGetSetting200Response

`func NewGetSetting200Response(id string, name string, description string, value GetSetting200ResponseValue, meta map[string]interface{}, ) *GetSetting200Response`

NewGetSetting200Response instantiates a new GetSetting200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSetting200ResponseWithDefaults

`func NewGetSetting200ResponseWithDefaults() *GetSetting200Response`

NewGetSetting200ResponseWithDefaults instantiates a new GetSetting200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSetting200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSetting200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSetting200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetSetting200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSetting200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSetting200Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetSetting200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSetting200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSetting200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetValue

`func (o *GetSetting200Response) GetValue() GetSetting200ResponseValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetSetting200Response) GetValueOk() (*GetSetting200ResponseValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetSetting200Response) SetValue(v GetSetting200ResponseValue)`

SetValue sets Value field to given value.


### GetMeta

`func (o *GetSetting200Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSetting200Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSetting200Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


