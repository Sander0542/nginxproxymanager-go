# PutStreamUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingPort** | Pointer to **int64** |  | [optional] 
**ForwardingHost** | Pointer to [**PutStreamUpdateRequestForwardingHost**](PutStreamUpdateRequestForwardingHost.md) |  | [optional] 
**ForwardingPort** | Pointer to **int64** |  | [optional] 
**TcpForwarding** | Pointer to **boolAsInt** |  | [optional] 
**UdpForwarding** | Pointer to **boolAsInt** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPutStreamUpdateRequest

`func NewPutStreamUpdateRequest() *PutStreamUpdateRequest`

NewPutStreamUpdateRequest instantiates a new PutStreamUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutStreamUpdateRequestWithDefaults

`func NewPutStreamUpdateRequestWithDefaults() *PutStreamUpdateRequest`

NewPutStreamUpdateRequestWithDefaults instantiates a new PutStreamUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingPort

`func (o *PutStreamUpdateRequest) GetIncomingPort() int64`

GetIncomingPort returns the IncomingPort field if non-nil, zero value otherwise.

### GetIncomingPortOk

`func (o *PutStreamUpdateRequest) GetIncomingPortOk() (*int64, bool)`

GetIncomingPortOk returns a tuple with the IncomingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPort

`func (o *PutStreamUpdateRequest) SetIncomingPort(v int64)`

SetIncomingPort sets IncomingPort field to given value.

### HasIncomingPort

`func (o *PutStreamUpdateRequest) HasIncomingPort() bool`

HasIncomingPort returns a boolean if a field has been set.

### GetForwardingHost

`func (o *PutStreamUpdateRequest) GetForwardingHost() PutStreamUpdateRequestForwardingHost`

GetForwardingHost returns the ForwardingHost field if non-nil, zero value otherwise.

### GetForwardingHostOk

`func (o *PutStreamUpdateRequest) GetForwardingHostOk() (*PutStreamUpdateRequestForwardingHost, bool)`

GetForwardingHostOk returns a tuple with the ForwardingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingHost

`func (o *PutStreamUpdateRequest) SetForwardingHost(v PutStreamUpdateRequestForwardingHost)`

SetForwardingHost sets ForwardingHost field to given value.

### HasForwardingHost

`func (o *PutStreamUpdateRequest) HasForwardingHost() bool`

HasForwardingHost returns a boolean if a field has been set.

### GetForwardingPort

`func (o *PutStreamUpdateRequest) GetForwardingPort() int64`

GetForwardingPort returns the ForwardingPort field if non-nil, zero value otherwise.

### GetForwardingPortOk

`func (o *PutStreamUpdateRequest) GetForwardingPortOk() (*int64, bool)`

GetForwardingPortOk returns a tuple with the ForwardingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPort

`func (o *PutStreamUpdateRequest) SetForwardingPort(v int64)`

SetForwardingPort sets ForwardingPort field to given value.

### HasForwardingPort

`func (o *PutStreamUpdateRequest) HasForwardingPort() bool`

HasForwardingPort returns a boolean if a field has been set.

### GetTcpForwarding

`func (o *PutStreamUpdateRequest) GetTcpForwarding() boolAsInt`

GetTcpForwarding returns the TcpForwarding field if non-nil, zero value otherwise.

### GetTcpForwardingOk

`func (o *PutStreamUpdateRequest) GetTcpForwardingOk() (*boolAsInt, bool)`

GetTcpForwardingOk returns a tuple with the TcpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpForwarding

`func (o *PutStreamUpdateRequest) SetTcpForwarding(v boolAsInt)`

SetTcpForwarding sets TcpForwarding field to given value.

### HasTcpForwarding

`func (o *PutStreamUpdateRequest) HasTcpForwarding() bool`

HasTcpForwarding returns a boolean if a field has been set.

### GetUdpForwarding

`func (o *PutStreamUpdateRequest) GetUdpForwarding() boolAsInt`

GetUdpForwarding returns the UdpForwarding field if non-nil, zero value otherwise.

### GetUdpForwardingOk

`func (o *PutStreamUpdateRequest) GetUdpForwardingOk() (*boolAsInt, bool)`

GetUdpForwardingOk returns a tuple with the UdpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpForwarding

`func (o *PutStreamUpdateRequest) SetUdpForwarding(v boolAsInt)`

SetUdpForwarding sets UdpForwarding field to given value.

### HasUdpForwarding

`func (o *PutStreamUpdateRequest) HasUdpForwarding() bool`

HasUdpForwarding returns a boolean if a field has been set.

### GetMeta

`func (o *PutStreamUpdateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PutStreamUpdateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PutStreamUpdateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PutStreamUpdateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


