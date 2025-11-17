# GetDNSProviders200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the DNS provider, matching the python package | 
**Name** | **string** | Human-readable name of the DNS provider | 
**Credentials** | **string** | Instructions on how to format the credentials for this DNS provider | 

## Methods

### NewGetDNSProviders200ResponseInner

`func NewGetDNSProviders200ResponseInner(id string, name string, credentials string, ) *GetDNSProviders200ResponseInner`

NewGetDNSProviders200ResponseInner instantiates a new GetDNSProviders200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDNSProviders200ResponseInnerWithDefaults

`func NewGetDNSProviders200ResponseInnerWithDefaults() *GetDNSProviders200ResponseInner`

NewGetDNSProviders200ResponseInnerWithDefaults instantiates a new GetDNSProviders200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDNSProviders200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDNSProviders200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDNSProviders200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetDNSProviders200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDNSProviders200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDNSProviders200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetCredentials

`func (o *GetDNSProviders200ResponseInner) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetDNSProviders200ResponseInner) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetDNSProviders200ResponseInner) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


