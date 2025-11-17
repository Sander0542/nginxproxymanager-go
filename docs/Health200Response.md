# Health200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Healthy | 
**Setup** | Pointer to **bool** | Whether the initial setup has been completed | [optional] 
**Version** | [**Health200ResponseVersion**](Health200ResponseVersion.md) |  | 

## Methods

### NewHealth200Response

`func NewHealth200Response(status string, version Health200ResponseVersion, ) *Health200Response`

NewHealth200Response instantiates a new Health200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealth200ResponseWithDefaults

`func NewHealth200ResponseWithDefaults() *Health200Response`

NewHealth200ResponseWithDefaults instantiates a new Health200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Health200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Health200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Health200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSetup

`func (o *Health200Response) GetSetup() bool`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *Health200Response) GetSetupOk() (*bool, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *Health200Response) SetSetup(v bool)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *Health200Response) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetVersion

`func (o *Health200Response) GetVersion() Health200ResponseVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Health200Response) GetVersionOk() (*Health200ResponseVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Health200Response) SetVersion(v Health200ResponseVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


