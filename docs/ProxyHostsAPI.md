# \ProxyHostsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProxyHost**](ProxyHostsAPI.md#CreateProxyHost) | **Post** /nginx/proxy-hosts | Create a Proxy Host
[**DeleteProxyHost**](ProxyHostsAPI.md#DeleteProxyHost) | **Delete** /nginx/proxy-hosts/{hostID} | Delete a Proxy Host
[**DisableProxyHost**](ProxyHostsAPI.md#DisableProxyHost) | **Post** /nginx/proxy-hosts/{hostID}/disable | Disable a Proxy Host
[**EnableProxyHost**](ProxyHostsAPI.md#EnableProxyHost) | **Post** /nginx/proxy-hosts/{hostID}/enable | Enable a Proxy Host
[**GetProxyHost**](ProxyHostsAPI.md#GetProxyHost) | **Get** /nginx/proxy-hosts/{hostID} | Get a Proxy Host
[**GetProxyHosts**](ProxyHostsAPI.md#GetProxyHosts) | **Get** /nginx/proxy-hosts | Get all proxy hosts
[**UpdateProxyHost**](ProxyHostsAPI.md#UpdateProxyHost) | **Put** /nginx/proxy-hosts/{hostID} | Update a Proxy Host



## CreateProxyHost

> GetProxyHosts200ResponseInner CreateProxyHost(ctx).CreateProxyHostRequest(createProxyHostRequest).Execute()

Create a Proxy Host

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
	createProxyHostRequest := *openapiclient.NewCreateProxyHostRequest([]string{"DomainNames_example"}, "http", "127.0.0.1", int64(8080)) // CreateProxyHostRequest | Proxy Host Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.CreateProxyHost(context.Background()).CreateProxyHostRequest(createProxyHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.CreateProxyHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProxyHost`: GetProxyHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.CreateProxyHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProxyHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProxyHostRequest** | [**CreateProxyHostRequest**](CreateProxyHostRequest.md) | Proxy Host Payload | 

### Return type

[**GetProxyHosts200ResponseInner**](GetProxyHosts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProxyHost

> bool DeleteProxyHost(ctx, hostID).Execute()

Delete a Proxy Host

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
	hostID := int64(2) // int64 | The ID of the Proxy Host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.DeleteProxyHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.DeleteProxyHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProxyHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.DeleteProxyHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** | The ID of the Proxy Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProxyHostRequest struct via the builder pattern


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


## DisableProxyHost

> bool DisableProxyHost(ctx, hostID).Execute()

Disable a Proxy Host

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
	hostID := int64(2) // int64 | The ID of the Proxy Host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.DisableProxyHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.DisableProxyHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableProxyHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.DisableProxyHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** | The ID of the Proxy Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableProxyHostRequest struct via the builder pattern


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


## EnableProxyHost

> bool EnableProxyHost(ctx, hostID).Execute()

Enable a Proxy Host

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
	hostID := int64(2) // int64 | The ID of the Proxy Host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.EnableProxyHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.EnableProxyHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableProxyHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.EnableProxyHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** | The ID of the Proxy Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableProxyHostRequest struct via the builder pattern


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


## GetProxyHost

> GetProxyHosts200ResponseInner GetProxyHost(ctx, hostID).Execute()

Get a Proxy Host

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
	hostID := int64(1) // int64 | The ID of the Proxy Host

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.GetProxyHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.GetProxyHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyHost`: GetProxyHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.GetProxyHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** | The ID of the Proxy Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProxyHosts200ResponseInner**](GetProxyHosts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProxyHosts

> []GetProxyHosts200ResponseInner GetProxyHosts(ctx).Expand(expand).Execute()

Get all proxy hosts

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
	resp, r, err := apiClient.ProxyHostsAPI.GetProxyHosts(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.GetProxyHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProxyHosts`: []GetProxyHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.GetProxyHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expansions | 

### Return type

[**[]GetProxyHosts200ResponseInner**](GetProxyHosts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProxyHost

> GetProxyHosts200ResponseInner UpdateProxyHost(ctx, hostID).UpdateProxyHostRequest(updateProxyHostRequest).Execute()

Update a Proxy Host

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
	hostID := int64(2) // int64 | The ID of the Proxy Host
	updateProxyHostRequest := *openapiclient.NewUpdateProxyHostRequest() // UpdateProxyHostRequest | Proxy Host Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyHostsAPI.UpdateProxyHost(context.Background(), hostID).UpdateProxyHostRequest(updateProxyHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyHostsAPI.UpdateProxyHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProxyHost`: GetProxyHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProxyHostsAPI.UpdateProxyHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** | The ID of the Proxy Host | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProxyHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProxyHostRequest** | [**UpdateProxyHostRequest**](UpdateProxyHostRequest.md) | Proxy Host Payload | 

### Return type

[**GetProxyHosts200ResponseInner**](GetProxyHosts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

