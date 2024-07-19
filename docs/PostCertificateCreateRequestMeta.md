# PostCertificateCreateRequestMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LetsencryptEmail** | Pointer to **string** |  | [optional] 
**LetsencryptAgree** | Pointer to **boolAsInt** |  | [optional] 
**DnsChallenge** | Pointer to **boolAsInt** |  | [optional] 
**DnsProvider** | Pointer to **string** |  | [optional] 
**DnsProviderCredentials** | Pointer to **string** |  | [optional] 
**PropagationSeconds** | Pointer to **int64** |  | [optional] 

## Methods

### NewPostCertificateCreateRequestMeta

`func NewPostCertificateCreateRequestMeta() *PostCertificateCreateRequestMeta`

NewPostCertificateCreateRequestMeta instantiates a new PostCertificateCreateRequestMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCertificateCreateRequestMetaWithDefaults

`func NewPostCertificateCreateRequestMetaWithDefaults() *PostCertificateCreateRequestMeta`

NewPostCertificateCreateRequestMetaWithDefaults instantiates a new PostCertificateCreateRequestMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLetsencryptEmail

`func (o *PostCertificateCreateRequestMeta) GetLetsencryptEmail() string`

GetLetsencryptEmail returns the LetsencryptEmail field if non-nil, zero value otherwise.

### GetLetsencryptEmailOk

`func (o *PostCertificateCreateRequestMeta) GetLetsencryptEmailOk() (*string, bool)`

GetLetsencryptEmailOk returns a tuple with the LetsencryptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencryptEmail

`func (o *PostCertificateCreateRequestMeta) SetLetsencryptEmail(v string)`

SetLetsencryptEmail sets LetsencryptEmail field to given value.

### HasLetsencryptEmail

`func (o *PostCertificateCreateRequestMeta) HasLetsencryptEmail() bool`

HasLetsencryptEmail returns a boolean if a field has been set.

### GetLetsencryptAgree

`func (o *PostCertificateCreateRequestMeta) GetLetsencryptAgree() boolAsInt`

GetLetsencryptAgree returns the LetsencryptAgree field if non-nil, zero value otherwise.

### GetLetsencryptAgreeOk

`func (o *PostCertificateCreateRequestMeta) GetLetsencryptAgreeOk() (*boolAsInt, bool)`

GetLetsencryptAgreeOk returns a tuple with the LetsencryptAgree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencryptAgree

`func (o *PostCertificateCreateRequestMeta) SetLetsencryptAgree(v boolAsInt)`

SetLetsencryptAgree sets LetsencryptAgree field to given value.

### HasLetsencryptAgree

`func (o *PostCertificateCreateRequestMeta) HasLetsencryptAgree() bool`

HasLetsencryptAgree returns a boolean if a field has been set.

### GetDnsChallenge

`func (o *PostCertificateCreateRequestMeta) GetDnsChallenge() boolAsInt`

GetDnsChallenge returns the DnsChallenge field if non-nil, zero value otherwise.

### GetDnsChallengeOk

`func (o *PostCertificateCreateRequestMeta) GetDnsChallengeOk() (*boolAsInt, bool)`

GetDnsChallengeOk returns a tuple with the DnsChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChallenge

`func (o *PostCertificateCreateRequestMeta) SetDnsChallenge(v boolAsInt)`

SetDnsChallenge sets DnsChallenge field to given value.

### HasDnsChallenge

`func (o *PostCertificateCreateRequestMeta) HasDnsChallenge() bool`

HasDnsChallenge returns a boolean if a field has been set.

### GetDnsProvider

`func (o *PostCertificateCreateRequestMeta) GetDnsProvider() string`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *PostCertificateCreateRequestMeta) GetDnsProviderOk() (*string, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *PostCertificateCreateRequestMeta) SetDnsProvider(v string)`

SetDnsProvider sets DnsProvider field to given value.

### HasDnsProvider

`func (o *PostCertificateCreateRequestMeta) HasDnsProvider() bool`

HasDnsProvider returns a boolean if a field has been set.

### GetDnsProviderCredentials

`func (o *PostCertificateCreateRequestMeta) GetDnsProviderCredentials() string`

GetDnsProviderCredentials returns the DnsProviderCredentials field if non-nil, zero value otherwise.

### GetDnsProviderCredentialsOk

`func (o *PostCertificateCreateRequestMeta) GetDnsProviderCredentialsOk() (*string, bool)`

GetDnsProviderCredentialsOk returns a tuple with the DnsProviderCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProviderCredentials

`func (o *PostCertificateCreateRequestMeta) SetDnsProviderCredentials(v string)`

SetDnsProviderCredentials sets DnsProviderCredentials field to given value.

### HasDnsProviderCredentials

`func (o *PostCertificateCreateRequestMeta) HasDnsProviderCredentials() bool`

HasDnsProviderCredentials returns a boolean if a field has been set.

### GetPropagationSeconds

`func (o *PostCertificateCreateRequestMeta) GetPropagationSeconds() int64`

GetPropagationSeconds returns the PropagationSeconds field if non-nil, zero value otherwise.

### GetPropagationSecondsOk

`func (o *PostCertificateCreateRequestMeta) GetPropagationSecondsOk() (*int64, bool)`

GetPropagationSecondsOk returns a tuple with the PropagationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationSeconds

`func (o *PostCertificateCreateRequestMeta) SetPropagationSeconds(v int64)`

SetPropagationSeconds sets PropagationSeconds field to given value.

### HasPropagationSeconds

`func (o *PostCertificateCreateRequestMeta) HasPropagationSeconds() bool`

HasPropagationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


