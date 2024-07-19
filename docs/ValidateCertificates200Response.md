# ValidateCertificates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**ValidateCertificates200ResponseCertificate**](ValidateCertificates200ResponseCertificate.md) |  | 
**CertificateKey** | **bool** |  | 

## Methods

### NewValidateCertificates200Response

`func NewValidateCertificates200Response(certificate ValidateCertificates200ResponseCertificate, certificateKey bool, ) *ValidateCertificates200Response`

NewValidateCertificates200Response instantiates a new ValidateCertificates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateCertificates200ResponseWithDefaults

`func NewValidateCertificates200ResponseWithDefaults() *ValidateCertificates200Response`

NewValidateCertificates200ResponseWithDefaults instantiates a new ValidateCertificates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ValidateCertificates200Response) GetCertificate() ValidateCertificates200ResponseCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ValidateCertificates200Response) GetCertificateOk() (*ValidateCertificates200ResponseCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ValidateCertificates200Response) SetCertificate(v ValidateCertificates200ResponseCertificate)`

SetCertificate sets Certificate field to given value.


### GetCertificateKey

`func (o *ValidateCertificates200Response) GetCertificateKey() bool`

GetCertificateKey returns the CertificateKey field if non-nil, zero value otherwise.

### GetCertificateKeyOk

`func (o *ValidateCertificates200Response) GetCertificateKeyOk() (*bool, bool)`

GetCertificateKeyOk returns a tuple with the CertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKey

`func (o *ValidateCertificates200Response) SetCertificateKey(v bool)`

SetCertificateKey sets CertificateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


