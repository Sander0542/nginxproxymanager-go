# GetCertificates200ResponseInnerMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** |  | [optional] 
**CertificateKey** | Pointer to **string** |  | [optional] 
**DnsChallenge** | Pointer to **bool** |  | [optional] 
**DnsProvider** | Pointer to **string** |  | [optional] 
**DnsProviderCredentials** | Pointer to **string** |  | [optional] 
**LetsencryptAgree** | Pointer to **bool** |  | [optional] 
**LetsencryptCertificate** | Pointer to **map[string]interface{}** |  | [optional] 
**LetsencryptEmail** | Pointer to **string** |  | [optional] 
**PropagationSeconds** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetCertificates200ResponseInnerMeta

`func NewGetCertificates200ResponseInnerMeta() *GetCertificates200ResponseInnerMeta`

NewGetCertificates200ResponseInnerMeta instantiates a new GetCertificates200ResponseInnerMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertificates200ResponseInnerMetaWithDefaults

`func NewGetCertificates200ResponseInnerMetaWithDefaults() *GetCertificates200ResponseInnerMeta`

NewGetCertificates200ResponseInnerMetaWithDefaults instantiates a new GetCertificates200ResponseInnerMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *GetCertificates200ResponseInnerMeta) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *GetCertificates200ResponseInnerMeta) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *GetCertificates200ResponseInnerMeta) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *GetCertificates200ResponseInnerMeta) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateKey

`func (o *GetCertificates200ResponseInnerMeta) GetCertificateKey() string`

GetCertificateKey returns the CertificateKey field if non-nil, zero value otherwise.

### GetCertificateKeyOk

`func (o *GetCertificates200ResponseInnerMeta) GetCertificateKeyOk() (*string, bool)`

GetCertificateKeyOk returns a tuple with the CertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKey

`func (o *GetCertificates200ResponseInnerMeta) SetCertificateKey(v string)`

SetCertificateKey sets CertificateKey field to given value.

### HasCertificateKey

`func (o *GetCertificates200ResponseInnerMeta) HasCertificateKey() bool`

HasCertificateKey returns a boolean if a field has been set.

### GetDnsChallenge

`func (o *GetCertificates200ResponseInnerMeta) GetDnsChallenge() bool`

GetDnsChallenge returns the DnsChallenge field if non-nil, zero value otherwise.

### GetDnsChallengeOk

`func (o *GetCertificates200ResponseInnerMeta) GetDnsChallengeOk() (*bool, bool)`

GetDnsChallengeOk returns a tuple with the DnsChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChallenge

`func (o *GetCertificates200ResponseInnerMeta) SetDnsChallenge(v bool)`

SetDnsChallenge sets DnsChallenge field to given value.

### HasDnsChallenge

`func (o *GetCertificates200ResponseInnerMeta) HasDnsChallenge() bool`

HasDnsChallenge returns a boolean if a field has been set.

### GetDnsProvider

`func (o *GetCertificates200ResponseInnerMeta) GetDnsProvider() string`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *GetCertificates200ResponseInnerMeta) GetDnsProviderOk() (*string, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *GetCertificates200ResponseInnerMeta) SetDnsProvider(v string)`

SetDnsProvider sets DnsProvider field to given value.

### HasDnsProvider

`func (o *GetCertificates200ResponseInnerMeta) HasDnsProvider() bool`

HasDnsProvider returns a boolean if a field has been set.

### GetDnsProviderCredentials

`func (o *GetCertificates200ResponseInnerMeta) GetDnsProviderCredentials() string`

GetDnsProviderCredentials returns the DnsProviderCredentials field if non-nil, zero value otherwise.

### GetDnsProviderCredentialsOk

`func (o *GetCertificates200ResponseInnerMeta) GetDnsProviderCredentialsOk() (*string, bool)`

GetDnsProviderCredentialsOk returns a tuple with the DnsProviderCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProviderCredentials

`func (o *GetCertificates200ResponseInnerMeta) SetDnsProviderCredentials(v string)`

SetDnsProviderCredentials sets DnsProviderCredentials field to given value.

### HasDnsProviderCredentials

`func (o *GetCertificates200ResponseInnerMeta) HasDnsProviderCredentials() bool`

HasDnsProviderCredentials returns a boolean if a field has been set.

### GetLetsencryptAgree

`func (o *GetCertificates200ResponseInnerMeta) GetLetsencryptAgree() bool`

GetLetsencryptAgree returns the LetsencryptAgree field if non-nil, zero value otherwise.

### GetLetsencryptAgreeOk

`func (o *GetCertificates200ResponseInnerMeta) GetLetsencryptAgreeOk() (*bool, bool)`

GetLetsencryptAgreeOk returns a tuple with the LetsencryptAgree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencryptAgree

`func (o *GetCertificates200ResponseInnerMeta) SetLetsencryptAgree(v bool)`

SetLetsencryptAgree sets LetsencryptAgree field to given value.

### HasLetsencryptAgree

`func (o *GetCertificates200ResponseInnerMeta) HasLetsencryptAgree() bool`

HasLetsencryptAgree returns a boolean if a field has been set.

### GetLetsencryptCertificate

`func (o *GetCertificates200ResponseInnerMeta) GetLetsencryptCertificate() map[string]interface{}`

GetLetsencryptCertificate returns the LetsencryptCertificate field if non-nil, zero value otherwise.

### GetLetsencryptCertificateOk

`func (o *GetCertificates200ResponseInnerMeta) GetLetsencryptCertificateOk() (*map[string]interface{}, bool)`

GetLetsencryptCertificateOk returns a tuple with the LetsencryptCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencryptCertificate

`func (o *GetCertificates200ResponseInnerMeta) SetLetsencryptCertificate(v map[string]interface{})`

SetLetsencryptCertificate sets LetsencryptCertificate field to given value.

### HasLetsencryptCertificate

`func (o *GetCertificates200ResponseInnerMeta) HasLetsencryptCertificate() bool`

HasLetsencryptCertificate returns a boolean if a field has been set.

### GetLetsencryptEmail

`func (o *GetCertificates200ResponseInnerMeta) GetLetsencryptEmail() string`

GetLetsencryptEmail returns the LetsencryptEmail field if non-nil, zero value otherwise.

### GetLetsencryptEmailOk

`func (o *GetCertificates200ResponseInnerMeta) GetLetsencryptEmailOk() (*string, bool)`

GetLetsencryptEmailOk returns a tuple with the LetsencryptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsencryptEmail

`func (o *GetCertificates200ResponseInnerMeta) SetLetsencryptEmail(v string)`

SetLetsencryptEmail sets LetsencryptEmail field to given value.

### HasLetsencryptEmail

`func (o *GetCertificates200ResponseInnerMeta) HasLetsencryptEmail() bool`

HasLetsencryptEmail returns a boolean if a field has been set.

### GetPropagationSeconds

`func (o *GetCertificates200ResponseInnerMeta) GetPropagationSeconds() int64`

GetPropagationSeconds returns the PropagationSeconds field if non-nil, zero value otherwise.

### GetPropagationSecondsOk

`func (o *GetCertificates200ResponseInnerMeta) GetPropagationSecondsOk() (*int64, bool)`

GetPropagationSecondsOk returns a tuple with the PropagationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationSeconds

`func (o *GetCertificates200ResponseInnerMeta) SetPropagationSeconds(v int64)`

SetPropagationSeconds sets PropagationSeconds field to given value.

### HasPropagationSeconds

`func (o *GetCertificates200ResponseInnerMeta) HasPropagationSeconds() bool`

HasPropagationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


