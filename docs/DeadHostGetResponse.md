# DeadHostGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** | Date and time of creation | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** | Date and time of last update | [optional] [readonly] 
**DomainNames** | Pointer to **[]string** | Domain Names separated by a comma | [optional] 
**CertificateId** | Pointer to [**PutProxyHostUpdateRequestCertificateId**](PutProxyHostUpdateRequestCertificateId.md) |  | [optional] 
**SslForced** | Pointer to **boolAsInt** | Is SSL Forced | [optional] 
**HstsEnabled** | Pointer to **boolAsInt** | Is HSTS Enabled | [optional] 
**HstsSubdomains** | Pointer to **boolAsInt** | Is HSTS applicable to all subdomains | [optional] 
**Http2Support** | Pointer to **boolAsInt** | HTTP2 Protocol Support | [optional] 
**AdvancedConfig** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **boolAsInt** | Is Enabled | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeadHostGetResponse

`func NewDeadHostGetResponse() *DeadHostGetResponse`

NewDeadHostGetResponse instantiates a new DeadHostGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeadHostGetResponseWithDefaults

`func NewDeadHostGetResponseWithDefaults() *DeadHostGetResponse`

NewDeadHostGetResponseWithDefaults instantiates a new DeadHostGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeadHostGetResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeadHostGetResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeadHostGetResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeadHostGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *DeadHostGetResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *DeadHostGetResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *DeadHostGetResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *DeadHostGetResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *DeadHostGetResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *DeadHostGetResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *DeadHostGetResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *DeadHostGetResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetDomainNames

`func (o *DeadHostGetResponse) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *DeadHostGetResponse) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *DeadHostGetResponse) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *DeadHostGetResponse) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetCertificateId

`func (o *DeadHostGetResponse) GetCertificateId() PutProxyHostUpdateRequestCertificateId`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *DeadHostGetResponse) GetCertificateIdOk() (*PutProxyHostUpdateRequestCertificateId, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *DeadHostGetResponse) SetCertificateId(v PutProxyHostUpdateRequestCertificateId)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *DeadHostGetResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSslForced

`func (o *DeadHostGetResponse) GetSslForced() boolAsInt`

GetSslForced returns the SslForced field if non-nil, zero value otherwise.

### GetSslForcedOk

`func (o *DeadHostGetResponse) GetSslForcedOk() (*boolAsInt, bool)`

GetSslForcedOk returns a tuple with the SslForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslForced

`func (o *DeadHostGetResponse) SetSslForced(v boolAsInt)`

SetSslForced sets SslForced field to given value.

### HasSslForced

`func (o *DeadHostGetResponse) HasSslForced() bool`

HasSslForced returns a boolean if a field has been set.

### GetHstsEnabled

`func (o *DeadHostGetResponse) GetHstsEnabled() boolAsInt`

GetHstsEnabled returns the HstsEnabled field if non-nil, zero value otherwise.

### GetHstsEnabledOk

`func (o *DeadHostGetResponse) GetHstsEnabledOk() (*boolAsInt, bool)`

GetHstsEnabledOk returns a tuple with the HstsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsEnabled

`func (o *DeadHostGetResponse) SetHstsEnabled(v boolAsInt)`

SetHstsEnabled sets HstsEnabled field to given value.

### HasHstsEnabled

`func (o *DeadHostGetResponse) HasHstsEnabled() bool`

HasHstsEnabled returns a boolean if a field has been set.

### GetHstsSubdomains

`func (o *DeadHostGetResponse) GetHstsSubdomains() boolAsInt`

GetHstsSubdomains returns the HstsSubdomains field if non-nil, zero value otherwise.

### GetHstsSubdomainsOk

`func (o *DeadHostGetResponse) GetHstsSubdomainsOk() (*boolAsInt, bool)`

GetHstsSubdomainsOk returns a tuple with the HstsSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomains

`func (o *DeadHostGetResponse) SetHstsSubdomains(v boolAsInt)`

SetHstsSubdomains sets HstsSubdomains field to given value.

### HasHstsSubdomains

`func (o *DeadHostGetResponse) HasHstsSubdomains() bool`

HasHstsSubdomains returns a boolean if a field has been set.

### GetHttp2Support

`func (o *DeadHostGetResponse) GetHttp2Support() boolAsInt`

GetHttp2Support returns the Http2Support field if non-nil, zero value otherwise.

### GetHttp2SupportOk

`func (o *DeadHostGetResponse) GetHttp2SupportOk() (*boolAsInt, bool)`

GetHttp2SupportOk returns a tuple with the Http2Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Support

`func (o *DeadHostGetResponse) SetHttp2Support(v boolAsInt)`

SetHttp2Support sets Http2Support field to given value.

### HasHttp2Support

`func (o *DeadHostGetResponse) HasHttp2Support() bool`

HasHttp2Support returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *DeadHostGetResponse) GetAdvancedConfig() string`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *DeadHostGetResponse) GetAdvancedConfigOk() (*string, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *DeadHostGetResponse) SetAdvancedConfig(v string)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *DeadHostGetResponse) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *DeadHostGetResponse) GetEnabled() boolAsInt`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeadHostGetResponse) GetEnabledOk() (*boolAsInt, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeadHostGetResponse) SetEnabled(v boolAsInt)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeadHostGetResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMeta

`func (o *DeadHostGetResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DeadHostGetResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DeadHostGetResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DeadHostGetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


