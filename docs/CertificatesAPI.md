# \CertificatesAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificateDelete**](CertificatesAPI.md#DeleteCertificateDelete) | **Delete** /nginx/certificates/{id} | Deletes a existing Certificate
[**GetCertificateGet**](CertificatesAPI.md#GetCertificateGet) | **Get** /nginx/certificates/{id} | Returns a  Certificate
[**GetCertificateList**](CertificatesAPI.md#GetCertificateList) | **Get** /nginx/certificates | Returns a list of Certificates
[**GetCertificateTestHTTPChallenge**](CertificatesAPI.md#GetCertificateTestHTTPChallenge) | **Get** /nginx/certificates/{id}/test-http | Tests whether the HTTP challenge should work
[**PostCertificateCreate**](CertificatesAPI.md#PostCertificateCreate) | **Post** /nginx/certificates | Creates a new Certificate



## DeleteCertificateDelete

> boolAsInt DeleteCertificateDelete(ctx, id).Execute()

Deletes a existing Certificate

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
	id := int64(56) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.DeleteCertificateDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.DeleteCertificateDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificateDelete`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.DeleteCertificateDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**boolAsInt**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateGet

> CertificateGetResponse GetCertificateGet(ctx, id).Execute()

Returns a  Certificate

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
	id := int64(56) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.GetCertificateGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificateGet`: CertificateGetResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetCertificateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateGetResponse**](CertificateGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateList

> []CertificateGetResponse GetCertificateList(ctx).Execute()

Returns a list of Certificates

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
	resp, r, err := apiClient.CertificatesAPI.GetCertificateList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificateList`: []CertificateGetResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetCertificateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateListRequest struct via the builder pattern


### Return type

[**[]CertificateGetResponse**](CertificateGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateTestHTTPChallenge

> GetCertificateTestHTTPChallenge(ctx, id).Execute()

Tests whether the HTTP challenge should work

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
	id := int64(56) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificatesAPI.GetCertificateTestHTTPChallenge(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificateTestHTTPChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateTestHTTPChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCertificateCreate

> CertificateCreateResponse PostCertificateCreate(ctx).PostCertificateCreateRequest(postCertificateCreateRequest).Execute()

Creates a new Certificate

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
	postCertificateCreateRequest := *openapiclient.NewPostCertificateCreateRequest("Provider_example") // PostCertificateCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.PostCertificateCreate(context.Background()).PostCertificateCreateRequest(postCertificateCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.PostCertificateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCertificateCreate`: CertificateCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.PostCertificateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCertificateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCertificateCreateRequest** | [**PostCertificateCreateRequest**](PostCertificateCreateRequest.md) |  | 

### Return type

[**CertificateCreateResponse**](CertificateCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

