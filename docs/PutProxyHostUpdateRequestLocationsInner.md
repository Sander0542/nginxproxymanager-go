# PutProxyHostUpdateRequestLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Path** | **string** |  | 
**ForwardScheme** | **string** |  | 
**ForwardHost** | **string** |  | 
**ForwardPort** | **int64** |  | 
**ForwardPath** | Pointer to **string** |  | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 

## Methods

### NewPutProxyHostUpdateRequestLocationsInner

`func NewPutProxyHostUpdateRequestLocationsInner(path string, forwardScheme string, forwardHost string, forwardPort int64, ) *PutProxyHostUpdateRequestLocationsInner`

NewPutProxyHostUpdateRequestLocationsInner instantiates a new PutProxyHostUpdateRequestLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutProxyHostUpdateRequestLocationsInnerWithDefaults

`func NewPutProxyHostUpdateRequestLocationsInnerWithDefaults() *PutProxyHostUpdateRequestLocationsInner`

NewPutProxyHostUpdateRequestLocationsInnerWithDefaults instantiates a new PutProxyHostUpdateRequestLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutProxyHostUpdateRequestLocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutProxyHostUpdateRequestLocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutProxyHostUpdateRequestLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PutProxyHostUpdateRequestLocationsInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PutProxyHostUpdateRequestLocationsInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *PutProxyHostUpdateRequestLocationsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PutProxyHostUpdateRequestLocationsInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetForwardScheme

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardScheme() string`

GetForwardScheme returns the ForwardScheme field if non-nil, zero value otherwise.

### GetForwardSchemeOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardSchemeOk() (*string, bool)`

GetForwardSchemeOk returns a tuple with the ForwardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardScheme

`func (o *PutProxyHostUpdateRequestLocationsInner) SetForwardScheme(v string)`

SetForwardScheme sets ForwardScheme field to given value.


### GetForwardHost

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardHost() string`

GetForwardHost returns the ForwardHost field if non-nil, zero value otherwise.

### GetForwardHostOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardHostOk() (*string, bool)`

GetForwardHostOk returns a tuple with the ForwardHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardHost

`func (o *PutProxyHostUpdateRequestLocationsInner) SetForwardHost(v string)`

SetForwardHost sets ForwardHost field to given value.


### GetForwardPort

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardPort() int64`

GetForwardPort returns the ForwardPort field if non-nil, zero value otherwise.

### GetForwardPortOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardPortOk() (*int64, bool)`

GetForwardPortOk returns a tuple with the ForwardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPort

`func (o *PutProxyHostUpdateRequestLocationsInner) SetForwardPort(v int64)`

SetForwardPort sets ForwardPort field to given value.


### GetForwardPath

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardPath() string`

GetForwardPath returns the ForwardPath field if non-nil, zero value otherwise.

### GetForwardPathOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetForwardPathOk() (*string, bool)`

GetForwardPathOk returns a tuple with the ForwardPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardPath

`func (o *PutProxyHostUpdateRequestLocationsInner) SetForwardPath(v string)`

SetForwardPath sets ForwardPath field to given value.

### HasForwardPath

`func (o *PutProxyHostUpdateRequestLocationsInner) HasForwardPath() bool`

HasForwardPath returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *PutProxyHostUpdateRequestLocationsInner) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *PutProxyHostUpdateRequestLocationsInner) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *PutProxyHostUpdateRequestLocationsInner) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *PutProxyHostUpdateRequestLocationsInner) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


