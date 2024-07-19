# UpdateUserPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | Pointer to **string** | Visibility Type | [optional] 
**AccessLists** | Pointer to **string** | Access Lists Permissions | [optional] 
**DeadHosts** | Pointer to **string** | 404 Hosts Permissions | [optional] 
**ProxyHosts** | Pointer to **string** | Proxy Hosts Permissions | [optional] 
**RedirectionHosts** | Pointer to **string** | Redirection Permissions | [optional] 
**Streams** | Pointer to **string** | Streams Permissions | [optional] 
**Certificates** | Pointer to **string** | Certificates Permissions | [optional] 

## Methods

### NewUpdateUserPermissionsRequest

`func NewUpdateUserPermissionsRequest() *UpdateUserPermissionsRequest`

NewUpdateUserPermissionsRequest instantiates a new UpdateUserPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPermissionsRequestWithDefaults

`func NewUpdateUserPermissionsRequestWithDefaults() *UpdateUserPermissionsRequest`

NewUpdateUserPermissionsRequestWithDefaults instantiates a new UpdateUserPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *UpdateUserPermissionsRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateUserPermissionsRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateUserPermissionsRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateUserPermissionsRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccessLists

`func (o *UpdateUserPermissionsRequest) GetAccessLists() string`

GetAccessLists returns the AccessLists field if non-nil, zero value otherwise.

### GetAccessListsOk

`func (o *UpdateUserPermissionsRequest) GetAccessListsOk() (*string, bool)`

GetAccessListsOk returns a tuple with the AccessLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLists

`func (o *UpdateUserPermissionsRequest) SetAccessLists(v string)`

SetAccessLists sets AccessLists field to given value.

### HasAccessLists

`func (o *UpdateUserPermissionsRequest) HasAccessLists() bool`

HasAccessLists returns a boolean if a field has been set.

### GetDeadHosts

`func (o *UpdateUserPermissionsRequest) GetDeadHosts() string`

GetDeadHosts returns the DeadHosts field if non-nil, zero value otherwise.

### GetDeadHostsOk

`func (o *UpdateUserPermissionsRequest) GetDeadHostsOk() (*string, bool)`

GetDeadHostsOk returns a tuple with the DeadHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadHosts

`func (o *UpdateUserPermissionsRequest) SetDeadHosts(v string)`

SetDeadHosts sets DeadHosts field to given value.

### HasDeadHosts

`func (o *UpdateUserPermissionsRequest) HasDeadHosts() bool`

HasDeadHosts returns a boolean if a field has been set.

### GetProxyHosts

`func (o *UpdateUserPermissionsRequest) GetProxyHosts() string`

GetProxyHosts returns the ProxyHosts field if non-nil, zero value otherwise.

### GetProxyHostsOk

`func (o *UpdateUserPermissionsRequest) GetProxyHostsOk() (*string, bool)`

GetProxyHostsOk returns a tuple with the ProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHosts

`func (o *UpdateUserPermissionsRequest) SetProxyHosts(v string)`

SetProxyHosts sets ProxyHosts field to given value.

### HasProxyHosts

`func (o *UpdateUserPermissionsRequest) HasProxyHosts() bool`

HasProxyHosts returns a boolean if a field has been set.

### GetRedirectionHosts

`func (o *UpdateUserPermissionsRequest) GetRedirectionHosts() string`

GetRedirectionHosts returns the RedirectionHosts field if non-nil, zero value otherwise.

### GetRedirectionHostsOk

`func (o *UpdateUserPermissionsRequest) GetRedirectionHostsOk() (*string, bool)`

GetRedirectionHostsOk returns a tuple with the RedirectionHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionHosts

`func (o *UpdateUserPermissionsRequest) SetRedirectionHosts(v string)`

SetRedirectionHosts sets RedirectionHosts field to given value.

### HasRedirectionHosts

`func (o *UpdateUserPermissionsRequest) HasRedirectionHosts() bool`

HasRedirectionHosts returns a boolean if a field has been set.

### GetStreams

`func (o *UpdateUserPermissionsRequest) GetStreams() string`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *UpdateUserPermissionsRequest) GetStreamsOk() (*string, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *UpdateUserPermissionsRequest) SetStreams(v string)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *UpdateUserPermissionsRequest) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetCertificates

`func (o *UpdateUserPermissionsRequest) GetCertificates() string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *UpdateUserPermissionsRequest) GetCertificatesOk() (*string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *UpdateUserPermissionsRequest) SetCertificates(v string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *UpdateUserPermissionsRequest) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


