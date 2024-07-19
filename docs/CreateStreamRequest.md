# CreateStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingPort** | **int64** |  | 
**ForwardingHost** | [**GetStreams200ResponseInnerForwardingHost**](GetStreams200ResponseInnerForwardingHost.md) |  | 
**ForwardingPort** | **int64** |  | 
**TcpForwarding** | Pointer to **bool** |  | [optional] 
**UdpForwarding** | Pointer to **bool** |  | [optional] 
**CertificateId** | Pointer to [**GetProxyHosts200ResponseInnerCertificateId**](GetProxyHosts200ResponseInnerCertificateId.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateStreamRequest

`func NewCreateStreamRequest(incomingPort int64, forwardingHost GetStreams200ResponseInnerForwardingHost, forwardingPort int64, ) *CreateStreamRequest`

NewCreateStreamRequest instantiates a new CreateStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStreamRequestWithDefaults

`func NewCreateStreamRequestWithDefaults() *CreateStreamRequest`

NewCreateStreamRequestWithDefaults instantiates a new CreateStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingPort

`func (o *CreateStreamRequest) GetIncomingPort() int64`

GetIncomingPort returns the IncomingPort field if non-nil, zero value otherwise.

### GetIncomingPortOk

`func (o *CreateStreamRequest) GetIncomingPortOk() (*int64, bool)`

GetIncomingPortOk returns a tuple with the IncomingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPort

`func (o *CreateStreamRequest) SetIncomingPort(v int64)`

SetIncomingPort sets IncomingPort field to given value.


### GetForwardingHost

`func (o *CreateStreamRequest) GetForwardingHost() GetStreams200ResponseInnerForwardingHost`

GetForwardingHost returns the ForwardingHost field if non-nil, zero value otherwise.

### GetForwardingHostOk

`func (o *CreateStreamRequest) GetForwardingHostOk() (*GetStreams200ResponseInnerForwardingHost, bool)`

GetForwardingHostOk returns a tuple with the ForwardingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingHost

`func (o *CreateStreamRequest) SetForwardingHost(v GetStreams200ResponseInnerForwardingHost)`

SetForwardingHost sets ForwardingHost field to given value.


### GetForwardingPort

`func (o *CreateStreamRequest) GetForwardingPort() int64`

GetForwardingPort returns the ForwardingPort field if non-nil, zero value otherwise.

### GetForwardingPortOk

`func (o *CreateStreamRequest) GetForwardingPortOk() (*int64, bool)`

GetForwardingPortOk returns a tuple with the ForwardingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPort

`func (o *CreateStreamRequest) SetForwardingPort(v int64)`

SetForwardingPort sets ForwardingPort field to given value.


### GetTcpForwarding

`func (o *CreateStreamRequest) GetTcpForwarding() bool`

GetTcpForwarding returns the TcpForwarding field if non-nil, zero value otherwise.

### GetTcpForwardingOk

`func (o *CreateStreamRequest) GetTcpForwardingOk() (*bool, bool)`

GetTcpForwardingOk returns a tuple with the TcpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpForwarding

`func (o *CreateStreamRequest) SetTcpForwarding(v bool)`

SetTcpForwarding sets TcpForwarding field to given value.

### HasTcpForwarding

`func (o *CreateStreamRequest) HasTcpForwarding() bool`

HasTcpForwarding returns a boolean if a field has been set.

### GetUdpForwarding

`func (o *CreateStreamRequest) GetUdpForwarding() bool`

GetUdpForwarding returns the UdpForwarding field if non-nil, zero value otherwise.

### GetUdpForwardingOk

`func (o *CreateStreamRequest) GetUdpForwardingOk() (*bool, bool)`

GetUdpForwardingOk returns a tuple with the UdpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpForwarding

`func (o *CreateStreamRequest) SetUdpForwarding(v bool)`

SetUdpForwarding sets UdpForwarding field to given value.

### HasUdpForwarding

`func (o *CreateStreamRequest) HasUdpForwarding() bool`

HasUdpForwarding returns a boolean if a field has been set.

### GetCertificateId

`func (o *CreateStreamRequest) GetCertificateId() GetProxyHosts200ResponseInnerCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CreateStreamRequest) GetCertificateIdOk() (*GetProxyHosts200ResponseInnerCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CreateStreamRequest) SetCertificateId(v GetProxyHosts200ResponseInnerCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *CreateStreamRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetMeta

`func (o *CreateStreamRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateStreamRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateStreamRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateStreamRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


