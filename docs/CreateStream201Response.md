# CreateStream201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier | [readonly] 
**CreatedOn** | **string** | Date and time of creation | [readonly] 
**ModifiedOn** | **string** | Date and time of last update | [readonly] 
**OwnerUserId** | **int64** | User ID | 
**IncomingPort** | **int64** |  | 
**ForwardingHost** | [**CreateStreamRequestForwardingHost**](CreateStreamRequestForwardingHost.md) |  | 
**ForwardingPort** | **int64** |  | 
**TcpForwarding** | **bool** |  | 
**UdpForwarding** | **bool** |  | 
**Enabled** | **bool** | Is Enabled | 
**Meta** | **map[string]interface{}** |  | 

## Methods

### NewCreateStream201Response

`func NewCreateStream201Response(id int64, createdOn string, modifiedOn string, ownerUserId int64, incomingPort int64, forwardingHost CreateStreamRequestForwardingHost, forwardingPort int64, tcpForwarding bool, udpForwarding bool, enabled bool, meta map[string]interface{}, ) *CreateStream201Response`

NewCreateStream201Response instantiates a new CreateStream201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStream201ResponseWithDefaults

`func NewCreateStream201ResponseWithDefaults() *CreateStream201Response`

NewCreateStream201ResponseWithDefaults instantiates a new CreateStream201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateStream201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStream201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStream201Response) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedOn

`func (o *CreateStream201Response) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *CreateStream201Response) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *CreateStream201Response) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *CreateStream201Response) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *CreateStream201Response) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *CreateStream201Response) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetOwnerUserId

`func (o *CreateStream201Response) GetOwnerUserId() int64`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *CreateStream201Response) GetOwnerUserIdOk() (*int64, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *CreateStream201Response) SetOwnerUserId(v int64)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetIncomingPort

`func (o *CreateStream201Response) GetIncomingPort() int64`

GetIncomingPort returns the IncomingPort field if non-nil, zero value otherwise.

### GetIncomingPortOk

`func (o *CreateStream201Response) GetIncomingPortOk() (*int64, bool)`

GetIncomingPortOk returns a tuple with the IncomingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPort

`func (o *CreateStream201Response) SetIncomingPort(v int64)`

SetIncomingPort sets IncomingPort field to given value.


### GetForwardingHost

`func (o *CreateStream201Response) GetForwardingHost() CreateStreamRequestForwardingHost`

GetForwardingHost returns the ForwardingHost field if non-nil, zero value otherwise.

### GetForwardingHostOk

`func (o *CreateStream201Response) GetForwardingHostOk() (*CreateStreamRequestForwardingHost, bool)`

GetForwardingHostOk returns a tuple with the ForwardingHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingHost

`func (o *CreateStream201Response) SetForwardingHost(v CreateStreamRequestForwardingHost)`

SetForwardingHost sets ForwardingHost field to given value.


### GetForwardingPort

`func (o *CreateStream201Response) GetForwardingPort() int64`

GetForwardingPort returns the ForwardingPort field if non-nil, zero value otherwise.

### GetForwardingPortOk

`func (o *CreateStream201Response) GetForwardingPortOk() (*int64, bool)`

GetForwardingPortOk returns a tuple with the ForwardingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingPort

`func (o *CreateStream201Response) SetForwardingPort(v int64)`

SetForwardingPort sets ForwardingPort field to given value.


### GetTcpForwarding

`func (o *CreateStream201Response) GetTcpForwarding() bool`

GetTcpForwarding returns the TcpForwarding field if non-nil, zero value otherwise.

### GetTcpForwardingOk

`func (o *CreateStream201Response) GetTcpForwardingOk() (*bool, bool)`

GetTcpForwardingOk returns a tuple with the TcpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpForwarding

`func (o *CreateStream201Response) SetTcpForwarding(v bool)`

SetTcpForwarding sets TcpForwarding field to given value.


### GetUdpForwarding

`func (o *CreateStream201Response) GetUdpForwarding() bool`

GetUdpForwarding returns the UdpForwarding field if non-nil, zero value otherwise.

### GetUdpForwardingOk

`func (o *CreateStream201Response) GetUdpForwardingOk() (*bool, bool)`

GetUdpForwardingOk returns a tuple with the UdpForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpForwarding

`func (o *CreateStream201Response) SetUdpForwarding(v bool)`

SetUdpForwarding sets UdpForwarding field to given value.


### GetEnabled

`func (o *CreateStream201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateStream201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateStream201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *CreateStream201Response) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateStream201Response) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateStream201Response) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


