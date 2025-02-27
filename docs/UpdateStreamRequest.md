# UpdateStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingPort** | Pointer to **int64** |  | [optional] 
**ForwardingHost** | Pointer to [**GetStreams200ResponseInnerForwardingHost**](GetStreams200ResponseInnerForwardingHost.md) |  | [optional] 
**ForwardingPort** | Pointer to **int64** |  | [optional] 
**TcpForwarding** | Pointer to **bool** |  | [optional] 
**UdpForwarding** | Pointer to **bool** |  | [optional] 
**CertificateId** | Pointer to [**GetProxyHosts200ResponseInnerCertificateId**](GetProxyHosts200ResponseInnerCertificateId.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateStreamRequest

`func NewUpdateStreamRequest() *UpdateStreamRequest`

NewUpdateStreamRequest instantiates a new UpdateStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStreamRequestWithDefaults

`func NewUpdateStreamRequestWithDefaults() *UpdateStreamRequest`

NewUpdateStreamRequestWithDefaults instantiates a new UpdateStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingPort

`func (o *UpdateStreamRequest) GetIncomingPort() int64`

GetIncomingPort returns the IncomingPort field if non-nil, zero value otherwise.

### GetIncomingPortOk

`func (o *UpdateStreamRequest) GetIncomingPortOk() (*int64, bool)`

GetIncomingPortOk returns a tuple with the IncomingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPort

`func (o *UpdateStreamRequest) SetIncomingPort(v int64)`

SetIncomingPort sets IncomingPort field to given value.

### HasIncomingPort

`func (o *UpdateStreamRequest) HasIncomingPort() bool`

HasIncomingPort returns a boolean if a field has been set.

### GetForwardingHost

`func (o *UpdateStreamRequest) GetForwardingHost() GetStreams200ResponseInnerForwardingHost`

GetForwardingHost returns the ForwardingHost field if non-nil, zero value otherwise.

### GetForwardingHostOk

`func (o *UpdateStreamRequest) GetForwardingHostOk() (*GetStreams200ResponseInnerForwardingHost, bool)`

GetForwardingHostOk returns a tuple with the ForwardingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingHost

`func (o *UpdateStreamRequest) SetForwardingHost(v GetStreams200ResponseInnerForwardingHost)`

SetForwardingHost sets ForwardingHost field to given value.

### HasForwardingHost

`func (o *UpdateStreamRequest) HasForwardingHost() bool`

HasForwardingHost returns a boolean if a field has been set.

### GetForwardingPort

`func (o *UpdateStreamRequest) GetForwardingPort() int64`

GetForwardingPort returns the ForwardingPort field if non-nil, zero value otherwise.

### GetForwardingPortOk

`func (o *UpdateStreamRequest) GetForwardingPortOk() (*int64, bool)`

GetForwardingPortOk returns a tuple with the ForwardingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPort

`func (o *UpdateStreamRequest) SetForwardingPort(v int64)`

SetForwardingPort sets ForwardingPort field to given value.

### HasForwardingPort

`func (o *UpdateStreamRequest) HasForwardingPort() bool`

HasForwardingPort returns a boolean if a field has been set.

### GetTcpForwarding

`func (o *UpdateStreamRequest) GetTcpForwarding() bool`

GetTcpForwarding returns the TcpForwarding field if non-nil, zero value otherwise.

### GetTcpForwardingOk

`func (o *UpdateStreamRequest) GetTcpForwardingOk() (*bool, bool)`

GetTcpForwardingOk returns a tuple with the TcpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpForwarding

`func (o *UpdateStreamRequest) SetTcpForwarding(v bool)`

SetTcpForwarding sets TcpForwarding field to given value.

### HasTcpForwarding

`func (o *UpdateStreamRequest) HasTcpForwarding() bool`

HasTcpForwarding returns a boolean if a field has been set.

### GetUdpForwarding

`func (o *UpdateStreamRequest) GetUdpForwarding() bool`

GetUdpForwarding returns the UdpForwarding field if non-nil, zero value otherwise.

### GetUdpForwardingOk

`func (o *UpdateStreamRequest) GetUdpForwardingOk() (*bool, bool)`

GetUdpForwardingOk returns a tuple with the UdpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpForwarding

`func (o *UpdateStreamRequest) SetUdpForwarding(v bool)`

SetUdpForwarding sets UdpForwarding field to given value.

### HasUdpForwarding

`func (o *UpdateStreamRequest) HasUdpForwarding() bool`

HasUdpForwarding returns a boolean if a field has been set.

### GetCertificateId

`func (o *UpdateStreamRequest) GetCertificateId() GetProxyHosts200ResponseInnerCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *UpdateStreamRequest) GetCertificateIdOk() (*GetProxyHosts200ResponseInnerCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *UpdateStreamRequest) SetCertificateId(v GetProxyHosts200ResponseInnerCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *UpdateStreamRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateStreamRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateStreamRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateStreamRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateStreamRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


