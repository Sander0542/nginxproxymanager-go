# StreamUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** | Date and time of creation | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** | Date and time of last update | [optional] [readonly] 
**IncomingPort** | Pointer to **int64** |  | [optional] 
**ForwardingHost** | Pointer to [**PutStreamUpdateRequestForwardingHost**](PutStreamUpdateRequestForwardingHost.md) |  | [optional] 
**ForwardingPort** | Pointer to **int64** |  | [optional] 
**TcpForwarding** | Pointer to **boolAsInt** |  | [optional] 
**UdpForwarding** | Pointer to **boolAsInt** |  | [optional] 
**Enabled** | Pointer to **boolAsInt** | Is Enabled | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStreamUpdateResponse

`func NewStreamUpdateResponse() *StreamUpdateResponse`

NewStreamUpdateResponse instantiates a new StreamUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamUpdateResponseWithDefaults

`func NewStreamUpdateResponseWithDefaults() *StreamUpdateResponse`

NewStreamUpdateResponseWithDefaults instantiates a new StreamUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamUpdateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamUpdateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamUpdateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StreamUpdateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *StreamUpdateResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *StreamUpdateResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *StreamUpdateResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *StreamUpdateResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *StreamUpdateResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *StreamUpdateResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *StreamUpdateResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *StreamUpdateResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetIncomingPort

`func (o *StreamUpdateResponse) GetIncomingPort() int64`

GetIncomingPort returns the IncomingPort field if non-nil, zero value otherwise.

### GetIncomingPortOk

`func (o *StreamUpdateResponse) GetIncomingPortOk() (*int64, bool)`

GetIncomingPortOk returns a tuple with the IncomingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPort

`func (o *StreamUpdateResponse) SetIncomingPort(v int64)`

SetIncomingPort sets IncomingPort field to given value.

### HasIncomingPort

`func (o *StreamUpdateResponse) HasIncomingPort() bool`

HasIncomingPort returns a boolean if a field has been set.

### GetForwardingHost

`func (o *StreamUpdateResponse) GetForwardingHost() PutStreamUpdateRequestForwardingHost`

GetForwardingHost returns the ForwardingHost field if non-nil, zero value otherwise.

### GetForwardingHostOk

`func (o *StreamUpdateResponse) GetForwardingHostOk() (*PutStreamUpdateRequestForwardingHost, bool)`

GetForwardingHostOk returns a tuple with the ForwardingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingHost

`func (o *StreamUpdateResponse) SetForwardingHost(v PutStreamUpdateRequestForwardingHost)`

SetForwardingHost sets ForwardingHost field to given value.

### HasForwardingHost

`func (o *StreamUpdateResponse) HasForwardingHost() bool`

HasForwardingHost returns a boolean if a field has been set.

### GetForwardingPort

`func (o *StreamUpdateResponse) GetForwardingPort() int64`

GetForwardingPort returns the ForwardingPort field if non-nil, zero value otherwise.

### GetForwardingPortOk

`func (o *StreamUpdateResponse) GetForwardingPortOk() (*int64, bool)`

GetForwardingPortOk returns a tuple with the ForwardingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPort

`func (o *StreamUpdateResponse) SetForwardingPort(v int64)`

SetForwardingPort sets ForwardingPort field to given value.

### HasForwardingPort

`func (o *StreamUpdateResponse) HasForwardingPort() bool`

HasForwardingPort returns a boolean if a field has been set.

### GetTcpForwarding

`func (o *StreamUpdateResponse) GetTcpForwarding() boolAsInt`

GetTcpForwarding returns the TcpForwarding field if non-nil, zero value otherwise.

### GetTcpForwardingOk

`func (o *StreamUpdateResponse) GetTcpForwardingOk() (*boolAsInt, bool)`

GetTcpForwardingOk returns a tuple with the TcpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpForwarding

`func (o *StreamUpdateResponse) SetTcpForwarding(v boolAsInt)`

SetTcpForwarding sets TcpForwarding field to given value.

### HasTcpForwarding

`func (o *StreamUpdateResponse) HasTcpForwarding() bool`

HasTcpForwarding returns a boolean if a field has been set.

### GetUdpForwarding

`func (o *StreamUpdateResponse) GetUdpForwarding() boolAsInt`

GetUdpForwarding returns the UdpForwarding field if non-nil, zero value otherwise.

### GetUdpForwardingOk

`func (o *StreamUpdateResponse) GetUdpForwardingOk() (*boolAsInt, bool)`

GetUdpForwardingOk returns a tuple with the UdpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpForwarding

`func (o *StreamUpdateResponse) SetUdpForwarding(v boolAsInt)`

SetUdpForwarding sets UdpForwarding field to given value.

### HasUdpForwarding

`func (o *StreamUpdateResponse) HasUdpForwarding() bool`

HasUdpForwarding returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamUpdateResponse) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamUpdateResponse) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamUpdateResponse) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamUpdateResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *StreamUpdateResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StreamUpdateResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StreamUpdateResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StreamUpdateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


