# PostStreamCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingPort** | **int64** |  | 
**ForwardingHost** | [**PutStreamUpdateRequestForwardingHost**](PutStreamUpdateRequestForwardingHost.md) |  | 
**ForwardingPort** | **int64** |  | 
**TcpForwarding** | Pointer to **boolAsInt** |  | [optional] 
**UdpForwarding** | Pointer to **boolAsInt** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPostStreamCreateRequest

`func NewPostStreamCreateRequest(incomingPort int64, forwardingHost PutStreamUpdateRequestForwardingHost, forwardingPort int64, ) *PostStreamCreateRequest`

NewPostStreamCreateRequest instantiates a new PostStreamCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostStreamCreateRequestWithDefaults

`func NewPostStreamCreateRequestWithDefaults() *PostStreamCreateRequest`

NewPostStreamCreateRequestWithDefaults instantiates a new PostStreamCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingPort

`func (o *PostStreamCreateRequest) GetIncomingPort() int64`

GetIncomingPort returns the IncomingPort field if non-nil, zero value otherwise.

### GetIncomingPortOk

`func (o *PostStreamCreateRequest) GetIncomingPortOk() (*int64, bool)`

GetIncomingPortOk returns a tuple with the IncomingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPort

`func (o *PostStreamCreateRequest) SetIncomingPort(v int64)`

SetIncomingPort sets IncomingPort field to given value.


### GetForwardingHost

`func (o *PostStreamCreateRequest) GetForwardingHost() PutStreamUpdateRequestForwardingHost`

GetForwardingHost returns the ForwardingHost field if non-nil, zero value otherwise.

### GetForwardingHostOk

`func (o *PostStreamCreateRequest) GetForwardingHostOk() (*PutStreamUpdateRequestForwardingHost, bool)`

GetForwardingHostOk returns a tuple with the ForwardingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingHost

`func (o *PostStreamCreateRequest) SetForwardingHost(v PutStreamUpdateRequestForwardingHost)`

SetForwardingHost sets ForwardingHost field to given value.


### GetForwardingPort

`func (o *PostStreamCreateRequest) GetForwardingPort() int64`

GetForwardingPort returns the ForwardingPort field if non-nil, zero value otherwise.

### GetForwardingPortOk

`func (o *PostStreamCreateRequest) GetForwardingPortOk() (*int64, bool)`

GetForwardingPortOk returns a tuple with the ForwardingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPort

`func (o *PostStreamCreateRequest) SetForwardingPort(v int64)`

SetForwardingPort sets ForwardingPort field to given value.


### GetTcpForwarding

`func (o *PostStreamCreateRequest) GetTcpForwarding() boolAsInt`

GetTcpForwarding returns the TcpForwarding field if non-nil, zero value otherwise.

### GetTcpForwardingOk

`func (o *PostStreamCreateRequest) GetTcpForwardingOk() (*boolAsInt, bool)`

GetTcpForwardingOk returns a tuple with the TcpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpForwarding

`func (o *PostStreamCreateRequest) SetTcpForwarding(v boolAsInt)`

SetTcpForwarding sets TcpForwarding field to given value.

### HasTcpForwarding

`func (o *PostStreamCreateRequest) HasTcpForwarding() bool`

HasTcpForwarding returns a boolean if a field has been set.

### GetUdpForwarding

`func (o *PostStreamCreateRequest) GetUdpForwarding() boolAsInt`

GetUdpForwarding returns the UdpForwarding field if non-nil, zero value otherwise.

### GetUdpForwardingOk

`func (o *PostStreamCreateRequest) GetUdpForwardingOk() (*boolAsInt, bool)`

GetUdpForwardingOk returns a tuple with the UdpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpForwarding

`func (o *PostStreamCreateRequest) SetUdpForwarding(v boolAsInt)`

SetUdpForwarding sets UdpForwarding field to given value.

### HasUdpForwarding

`func (o *PostStreamCreateRequest) HasUdpForwarding() bool`

HasUdpForwarding returns a boolean if a field has been set.

### GetMeta

`func (o *PostStreamCreateRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PostStreamCreateRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PostStreamCreateRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PostStreamCreateRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


