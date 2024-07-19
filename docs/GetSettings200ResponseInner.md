# GetSettings200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Setting ID | 
**Name** | **string** | Setting Display Name | 
**Description** | **string** | Meaningful description | 
**Value** | [**GetSettings200ResponseInnerValue**](GetSettings200ResponseInnerValue.md) |  | 
**Meta** | **map[string]interface{}** | Extra metadata | 

## Methods

### NewGetSettings200ResponseInner

`func NewGetSettings200ResponseInner(id string, name string, description string, value GetSettings200ResponseInnerValue, meta map[string]interface{}, ) *GetSettings200ResponseInner`

NewGetSettings200ResponseInner instantiates a new GetSettings200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSettings200ResponseInnerWithDefaults

`func NewGetSettings200ResponseInnerWithDefaults() *GetSettings200ResponseInner`

NewGetSettings200ResponseInnerWithDefaults instantiates a new GetSettings200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSettings200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSettings200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSettings200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetSettings200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSettings200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSettings200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetSettings200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSettings200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSettings200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetValue

`func (o *GetSettings200ResponseInner) GetValue() GetSettings200ResponseInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetSettings200ResponseInner) GetValueOk() (*GetSettings200ResponseInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetSettings200ResponseInner) SetValue(v GetSettings200ResponseInnerValue)`

SetValue sets Value field to given value.


### GetMeta

`func (o *GetSettings200ResponseInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSettings200ResponseInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSettings200ResponseInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


