# \Class404HostsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create404Host**](Class404HostsAPI.md#Create404Host) | **Post** /nginx/dead-hosts | Create a 404 Host
[**DeleteDeadHost**](Class404HostsAPI.md#DeleteDeadHost) | **Delete** /nginx/dead-hosts/{hostID} | Delete a 404 Host
[**DisableDeadHost**](Class404HostsAPI.md#DisableDeadHost) | **Post** /nginx/dead-hosts/{hostID}/disable | Disable a 404 Host
[**EnableDeadHost**](Class404HostsAPI.md#EnableDeadHost) | **Post** /nginx/dead-hosts/{hostID}/enable | Enable a 404 Host
[**GetDeadHost**](Class404HostsAPI.md#GetDeadHost) | **Get** /nginx/dead-hosts/{hostID} | Get a 404 Host
[**GetDeadHosts**](Class404HostsAPI.md#GetDeadHosts) | **Get** /nginx/dead-hosts | Get all 404 hosts
[**UpdateDeadHost**](Class404HostsAPI.md#UpdateDeadHost) | **Put** /nginx/dead-hosts/{hostID} | Update a 404 Host



## Create404Host

> GetDeadHosts200ResponseInner Create404Host(ctx).Create404HostRequest(create404HostRequest).Execute()

Create a 404 Host

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
	create404HostRequest := *openapiclient.NewCreate404HostRequest([]string{"DomainNames_example"}) // Create404HostRequest | 404 Host Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class404HostsAPI.Create404Host(context.Background()).Create404HostRequest(create404HostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.Create404Host``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create404Host`: GetDeadHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.Create404Host`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate404HostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create404HostRequest** | [**Create404HostRequest**](Create404HostRequest.md) | 404 Host Payload | 

### Return type

[**GetDeadHosts200ResponseInner**](GetDeadHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeadHost

> bool DeleteDeadHost(ctx, hostID).Execute()

Delete a 404 Host

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
	hostID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class404HostsAPI.DeleteDeadHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.DeleteDeadHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeadHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.DeleteDeadHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeadHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableDeadHost

> bool DisableDeadHost(ctx, hostID).Execute()

Disable a 404 Host

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
	hostID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class404HostsAPI.DisableDeadHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.DisableDeadHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableDeadHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.DisableDeadHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableDeadHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableDeadHost

> bool EnableDeadHost(ctx, hostID).Execute()

Enable a 404 Host

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
	hostID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class404HostsAPI.EnableDeadHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.EnableDeadHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableDeadHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.EnableDeadHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableDeadHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeadHost

> GetDeadHosts200ResponseInner GetDeadHost(ctx, hostID).Expand(expand).Execute()

Get a 404 Host

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
	hostID := int64(1) // int64 | 
	expand := "expand_example" // string | Expansions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class404HostsAPI.GetDeadHost(context.Background(), hostID).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.GetDeadHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeadHost`: GetDeadHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.GetDeadHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeadHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expansions | 

### Return type

[**GetDeadHosts200ResponseInner**](GetDeadHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeadHosts

> []GetDeadHosts200ResponseInner GetDeadHosts(ctx).Expand(expand).Execute()

Get all 404 hosts

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
	resp, r, err := apiClient.Class404HostsAPI.GetDeadHosts(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.GetDeadHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeadHosts`: []GetDeadHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.GetDeadHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeadHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expansions | 

### Return type

[**[]GetDeadHosts200ResponseInner**](GetDeadHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeadHost

> GetDeadHosts200ResponseInner UpdateDeadHost(ctx, hostID).UpdateDeadHostRequest(updateDeadHostRequest).Execute()

Update a 404 Host

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
	hostID := int64(2) // int64 | 
	updateDeadHostRequest := *openapiclient.NewUpdateDeadHostRequest() // UpdateDeadHostRequest | 404 Host Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Class404HostsAPI.UpdateDeadHost(context.Background(), hostID).UpdateDeadHostRequest(updateDeadHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Class404HostsAPI.UpdateDeadHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeadHost`: GetDeadHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `Class404HostsAPI.UpdateDeadHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeadHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeadHostRequest** | [**UpdateDeadHostRequest**](UpdateDeadHostRequest.md) | 404 Host Payload | 

### Return type

[**GetDeadHosts200ResponseInner**](GetDeadHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

