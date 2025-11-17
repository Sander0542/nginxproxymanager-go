# \CertificatesAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificate**](CertificatesAPI.md#CreateCertificate) | **Post** /nginx/certificates | Create a Certificate
[**DeleteCertificate**](CertificatesAPI.md#DeleteCertificate) | **Delete** /nginx/certificates/{certID} | Delete a Certificate
[**DownloadCertificate**](CertificatesAPI.md#DownloadCertificate) | **Get** /nginx/certificates/{certID}/download | Downloads a Certificate
[**GetCertificate**](CertificatesAPI.md#GetCertificate) | **Get** /nginx/certificates/{certID} | Get a Certificate
[**GetCertificates**](CertificatesAPI.md#GetCertificates) | **Get** /nginx/certificates | Get all certificates
[**GetDNSProviders**](CertificatesAPI.md#GetDNSProviders) | **Get** /nginx/certificates/dns-providers | Get DNS Providers for Certificates
[**RenewCertificate**](CertificatesAPI.md#RenewCertificate) | **Post** /nginx/certificates/{certID}/renew | Renews a Certificate
[**TestHttpReach**](CertificatesAPI.md#TestHttpReach) | **Post** /nginx/certificates/test-http | Test HTTP Reachability
[**UploadCertificate**](CertificatesAPI.md#UploadCertificate) | **Post** /nginx/certificates/{certID}/upload | Uploads a custom Certificate
[**ValidateCertificates**](CertificatesAPI.md#ValidateCertificates) | **Post** /nginx/certificates/validate | Validates given Custom Certificates



## CreateCertificate

> GetCertificates200ResponseInner CreateCertificate(ctx).CreateCertificateRequest(createCertificateRequest).Execute()

Create a Certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	createCertificateRequest := *openapiclient.NewCreateCertificateRequest("letsencrypt") // CreateCertificateRequest | Certificate Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CreateCertificate(context.Background()).CreateCertificateRequest(createCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CreateCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertificate`: GetCertificates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CreateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCertificateRequest** | [**CreateCertificateRequest**](CreateCertificateRequest.md) | Certificate Payload | 

### Return type

[**GetCertificates200ResponseInner**](GetCertificates200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificate

> bool DeleteCertificate(ctx, certID).Execute()

Delete a Certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	certID := int64(2) // int64 | Certificate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.DeleteCertificate(context.Background(), certID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.DeleteCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificate`: bool
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.DeleteCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **int64** | Certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCertificate

> *os.File DownloadCertificate(ctx, certID).Execute()

Downloads a Certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	certID := int64(1) // int64 | Certificate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.DownloadCertificate(context.Background(), certID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.DownloadCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadCertificate`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.DownloadCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **int64** | Certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificate

> GetCertificates200ResponseInner GetCertificate(ctx, certID).Execute()

Get a Certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	certID := int64(1) // int64 | Certificate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.GetCertificate(context.Background(), certID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificate`: GetCertificates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **int64** | Certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCertificates200ResponseInner**](GetCertificates200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificates

> []GetCertificates200ResponseInner GetCertificates(ctx).Expand(expand).Execute()

Get all certificates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	expand := "expand_example" // string | Expansions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.GetCertificates(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificates`: []GetCertificates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expansions | 

### Return type

[**[]GetCertificates200ResponseInner**](GetCertificates200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDNSProviders

> []GetDNSProviders200ResponseInner GetDNSProviders(ctx).Execute()

Get DNS Providers for Certificates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.GetDNSProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetDNSProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDNSProviders`: []GetDNSProviders200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetDNSProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNSProvidersRequest struct via the builder pattern


### Return type

[**[]GetDNSProviders200ResponseInner**](GetDNSProviders200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewCertificate

> GetCertificates200ResponseInner RenewCertificate(ctx, certID).Execute()

Renews a Certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	certID := int64(1) // int64 | Certificate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.RenewCertificate(context.Background(), certID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.RenewCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewCertificate`: GetCertificates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.RenewCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **int64** | Certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCertificates200ResponseInner**](GetCertificates200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestHttpReach

> TestHttpReach(ctx).TestHttpReachRequest(testHttpReachRequest).Execute()

Test HTTP Reachability

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	testHttpReachRequest := *openapiclient.NewTestHttpReachRequest([]string{"Domains_example"}) // TestHttpReachRequest | Test Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificatesAPI.TestHttpReach(context.Background()).TestHttpReachRequest(testHttpReachRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.TestHttpReach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestHttpReachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testHttpReachRequest** | [**TestHttpReachRequest**](TestHttpReachRequest.md) | Test Payload | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCertificate

> UploadCertificate200Response UploadCertificate(ctx, certID).Certificate(certificate).CertificateKey(certificateKey).IntermediateCertificate(intermediateCertificate).Execute()

Uploads a custom Certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	certID := int64(1) // int64 | Certificate ID
	certificate := "certificate_example" // string | 
	certificateKey := "certificateKey_example" // string | 
	intermediateCertificate := "intermediateCertificate_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.UploadCertificate(context.Background(), certID).Certificate(certificate).CertificateKey(certificateKey).IntermediateCertificate(intermediateCertificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.UploadCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCertificate`: UploadCertificate200Response
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.UploadCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certID** | **int64** | Certificate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificate** | **string** |  | 
 **certificateKey** | **string** |  | 
 **intermediateCertificate** | **string** |  | 

### Return type

[**UploadCertificate200Response**](UploadCertificate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCertificates

> ValidateCertificates200Response ValidateCertificates(ctx).Certificate(certificate).CertificateKey(certificateKey).IntermediateCertificate(intermediateCertificate).Execute()

Validates given Custom Certificates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sander0542/nginxproxymanager-go"
)

func main() {
	certificate := "certificate_example" // string | 
	certificateKey := "certificateKey_example" // string | 
	intermediateCertificate := "intermediateCertificate_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.ValidateCertificates(context.Background()).Certificate(certificate).CertificateKey(certificateKey).IntermediateCertificate(intermediateCertificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.ValidateCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateCertificates`: ValidateCertificates200Response
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.ValidateCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificate** | **string** |  | 
 **certificateKey** | **string** |  | 
 **intermediateCertificate** | **string** |  | 

### Return type

[**ValidateCertificates200Response**](ValidateCertificates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

