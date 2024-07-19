# \ProxyHostsAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProxyHostDelete**](ProxyHostsAPI.md#DeleteProxyHostDelete) | **Delete** /nginx/proxy-hosts/{id} | Deletes a existing Proxy Host
[**GetProxyHostGet**](ProxyHostsAPI.md#GetProxyHostGet) | **Get** /nginx/proxy-hosts/{id} | Returns a  Proxy Host
[**GetProxyHostList**](ProxyHostsAPI.md#GetProxyHostList) | **Get** /nginx/proxy-hosts | Returns a list of Proxy Hosts
[**PostProxyHostCreate**](ProxyHostsAPI.md#PostProxyHostCreate) | **Post** /nginx/proxy-hosts | Creates a new Proxy Host
[**PostProxyHostDisable**](ProxyHostsAPI.md#PostProxyHostDisable) | **Post** /nginx/proxy-hosts/{id}/disable | Disables a existing Proxy Host
[**PostProxyHostEnable**](ProxyHostsAPI.md#PostProxyHostEnable) | **Post** /nginx/proxy-hosts/{id}/enable | Enables a existing Proxy Host
[**PutProxyHostUpdate**](ProxyHostsAPI.md#PutProxyHostUpdate) | **Put** /nginx/proxy-hosts/{id} | Updates a existing Proxy Host



## DeleteProxyHostDelete

> boolAsInt DeleteProxyHostDelete(ctx, id).Execute()

Deletes a existing Proxy Host

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
	resp, r, err := apiClient.ProxyHostsAPI.DeleteProxyHostDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.DeleteProxyHostDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProxyHostDelete`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.DeleteProxyHostDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProxyHostDeleteRequest struct via the builder pattern


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


## GetProxyHostGet

> ProxyHostGetResponse GetProxyHostGet(ctx, id).Execute()

Returns a  Proxy Host

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
	resp, r, err := apiClient.ProxyHostsAPI.GetProxyHostGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.GetProxyHostGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyHostGet`: ProxyHostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.GetProxyHostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyHostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyHostGetResponse**](ProxyHostGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxyHostList

> []ProxyHostGetResponse GetProxyHostList(ctx).Execute()

Returns a list of Proxy Hosts

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
	resp, r, err := apiClient.ProxyHostsAPI.GetProxyHostList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.GetProxyHostList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyHostList`: []ProxyHostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.GetProxyHostList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyHostListRequest struct via the builder pattern


### Return type

[**[]ProxyHostGetResponse**](ProxyHostGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProxyHostCreate

> ProxyHostCreateResponse PostProxyHostCreate(ctx).PostProxyHostCreateRequest(postProxyHostCreateRequest).Execute()

Creates a new Proxy Host

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
	postProxyHostCreateRequest := *openapiclient.NewPostProxyHostCreateRequest([]string{"DomainNames_example"}, "ForwardScheme_example", "ForwardHost_example", int64(123)) // PostProxyHostCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.PostProxyHostCreate(context.Background()).PostProxyHostCreateRequest(postProxyHostCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.PostProxyHostCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProxyHostCreate`: ProxyHostCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.PostProxyHostCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProxyHostCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postProxyHostCreateRequest** | [**PostProxyHostCreateRequest**](PostProxyHostCreateRequest.md) |  | 

### Return type

[**ProxyHostCreateResponse**](ProxyHostCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProxyHostDisable

> boolAsInt PostProxyHostDisable(ctx, id).Execute()

Disables a existing Proxy Host

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
	resp, r, err := apiClient.ProxyHostsAPI.PostProxyHostDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.PostProxyHostDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProxyHostDisable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.PostProxyHostDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProxyHostDisableRequest struct via the builder pattern


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


## PostProxyHostEnable

> boolAsInt PostProxyHostEnable(ctx, id).Execute()

Enables a existing Proxy Host

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
	resp, r, err := apiClient.ProxyHostsAPI.PostProxyHostEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.PostProxyHostEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProxyHostEnable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.PostProxyHostEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProxyHostEnableRequest struct via the builder pattern


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


## PutProxyHostUpdate

> ProxyHostUpdateResponse PutProxyHostUpdate(ctx, id).PutProxyHostUpdateRequest(putProxyHostUpdateRequest).Execute()

Updates a existing Proxy Host

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
	putProxyHostUpdateRequest := *openapiclient.NewPutProxyHostUpdateRequest() // PutProxyHostUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.PutProxyHostUpdate(context.Background(), id).PutProxyHostUpdateRequest(putProxyHostUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.PutProxyHostUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProxyHostUpdate`: ProxyHostUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.PutProxyHostUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProxyHostUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putProxyHostUpdateRequest** | [**PutProxyHostUpdateRequest**](PutProxyHostUpdateRequest.md) |  | 

### Return type

[**ProxyHostUpdateResponse**](ProxyHostUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

