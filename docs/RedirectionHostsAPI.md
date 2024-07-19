# \RedirectionHostsAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRedirectionHostDelete**](RedirectionHostsAPI.md#DeleteRedirectionHostDelete) | **Delete** /nginx/redirection-hosts/{id} | Deletes a existing Redirection Host
[**GetRedirectionHostGet**](RedirectionHostsAPI.md#GetRedirectionHostGet) | **Get** /nginx/redirection-hosts/{id} | Returns a  Redirection Host
[**GetRedirectionHostList**](RedirectionHostsAPI.md#GetRedirectionHostList) | **Get** /nginx/redirection-hosts | Returns a list of Redirection Hosts
[**PostRedirectionHostCreate**](RedirectionHostsAPI.md#PostRedirectionHostCreate) | **Post** /nginx/redirection-hosts | Creates a new Redirection Host
[**PostRedirectionHostDisable**](RedirectionHostsAPI.md#PostRedirectionHostDisable) | **Post** /nginx/redirection-hosts/{id}/disable | Disables a existing Redirection Host
[**PostRedirectionHostEnable**](RedirectionHostsAPI.md#PostRedirectionHostEnable) | **Post** /nginx/redirection-hosts/{id}/enable | Enables a existing Redirection Host
[**PutRedirectionHostUpdate**](RedirectionHostsAPI.md#PutRedirectionHostUpdate) | **Put** /nginx/redirection-hosts/{id} | Updates a existing Redirection Host



## DeleteRedirectionHostDelete

> boolAsInt DeleteRedirectionHostDelete(ctx, id).Execute()

Deletes a existing Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.DeleteRedirectionHostDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.DeleteRedirectionHostDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedirectionHostDelete`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.DeleteRedirectionHostDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedirectionHostDeleteRequest struct via the builder pattern


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


## GetRedirectionHostGet

> RedirectionHostGetResponse GetRedirectionHostGet(ctx, id).Execute()

Returns a  Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.GetRedirectionHostGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.GetRedirectionHostGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedirectionHostGet`: RedirectionHostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.GetRedirectionHostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectionHostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RedirectionHostGetResponse**](RedirectionHostGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedirectionHostList

> []RedirectionHostGetResponse GetRedirectionHostList(ctx).Execute()

Returns a list of Redirection Hosts

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
	resp, r, err := apiClient.RedirectionHostsAPI.GetRedirectionHostList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.GetRedirectionHostList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedirectionHostList`: []RedirectionHostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.GetRedirectionHostList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectionHostListRequest struct via the builder pattern


### Return type

[**[]RedirectionHostGetResponse**](RedirectionHostGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedirectionHostCreate

> RedirectionHostCreateResponse PostRedirectionHostCreate(ctx).PostRedirectionHostCreateRequest(postRedirectionHostCreateRequest).Execute()

Creates a new Redirection Host

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
	postRedirectionHostCreateRequest := *openapiclient.NewPostRedirectionHostCreateRequest([]string{"DomainNames_example"}, int64(302), "HTTPS or $scheme", "jc21.com") // PostRedirectionHostCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectionHostsAPI.PostRedirectionHostCreate(context.Background()).PostRedirectionHostCreateRequest(postRedirectionHostCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.PostRedirectionHostCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedirectionHostCreate`: RedirectionHostCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.PostRedirectionHostCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRedirectionHostCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRedirectionHostCreateRequest** | [**PostRedirectionHostCreateRequest**](PostRedirectionHostCreateRequest.md) |  | 

### Return type

[**RedirectionHostCreateResponse**](RedirectionHostCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedirectionHostDisable

> boolAsInt PostRedirectionHostDisable(ctx, id).Execute()

Disables a existing Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.PostRedirectionHostDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.PostRedirectionHostDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedirectionHostDisable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.PostRedirectionHostDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedirectionHostDisableRequest struct via the builder pattern


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


## PostRedirectionHostEnable

> boolAsInt PostRedirectionHostEnable(ctx, id).Execute()

Enables a existing Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.PostRedirectionHostEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.PostRedirectionHostEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRedirectionHostEnable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.PostRedirectionHostEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedirectionHostEnableRequest struct via the builder pattern


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


## PutRedirectionHostUpdate

> RedirectionHostUpdateResponse PutRedirectionHostUpdate(ctx, id).PutRedirectionHostUpdateRequest(putRedirectionHostUpdateRequest).Execute()

Updates a existing Redirection Host

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
	putRedirectionHostUpdateRequest := *openapiclient.NewPutRedirectionHostUpdateRequest() // PutRedirectionHostUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectionHostsAPI.PutRedirectionHostUpdate(context.Background(), id).PutRedirectionHostUpdateRequest(putRedirectionHostUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.PutRedirectionHostUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRedirectionHostUpdate`: RedirectionHostUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.PutRedirectionHostUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRedirectionHostUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRedirectionHostUpdateRequest** | [**PutRedirectionHostUpdateRequest**](PutRedirectionHostUpdateRequest.md) |  | 

### Return type

[**RedirectionHostUpdateResponse**](RedirectionHostUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

