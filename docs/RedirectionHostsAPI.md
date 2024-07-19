# \RedirectionHostsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRedirectionHost**](RedirectionHostsAPI.md#CreateRedirectionHost) | **Post** /nginx/redirection-hosts | Create a Redirection Host
[**DeleteRedirectionHost**](RedirectionHostsAPI.md#DeleteRedirectionHost) | **Delete** /nginx/redirection-hosts/{hostID} | Delete a Redirection Host
[**DisableRedirectionHost**](RedirectionHostsAPI.md#DisableRedirectionHost) | **Post** /nginx/redirection-hosts/{hostID}/disable | Disable a Redirection Host
[**EnableRedirectionHost**](RedirectionHostsAPI.md#EnableRedirectionHost) | **Post** /nginx/redirection-hosts/{hostID}/enable | Enable a Redirection Host
[**GetRedirectionHost**](RedirectionHostsAPI.md#GetRedirectionHost) | **Get** /nginx/redirection-hosts/{hostID} | Get a Redirection Host
[**GetRedirectionHosts**](RedirectionHostsAPI.md#GetRedirectionHosts) | **Get** /nginx/redirection-hosts | Get all Redirection hosts
[**UpdateRedirectionHost**](RedirectionHostsAPI.md#UpdateRedirectionHost) | **Put** /nginx/redirection-hosts/{hostID} | Update a Redirection Host



## CreateRedirectionHost

> GetRedirectionHosts200ResponseInner CreateRedirectionHost(ctx).CreateRedirectionHostRequest(createRedirectionHostRequest).Execute()

Create a Redirection Host

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
	createRedirectionHostRequest := *openapiclient.NewCreateRedirectionHostRequest([]string{"DomainNames_example"}, int64(302), "ForwardScheme_example", "jc21.com") // CreateRedirectionHostRequest | Redirection Host Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectionHostsAPI.CreateRedirectionHost(context.Background()).CreateRedirectionHostRequest(createRedirectionHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.CreateRedirectionHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRedirectionHost`: GetRedirectionHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.CreateRedirectionHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRedirectionHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRedirectionHostRequest** | [**CreateRedirectionHostRequest**](CreateRedirectionHostRequest.md) | Redirection Host Payload | 

### Return type

[**GetRedirectionHosts200ResponseInner**](GetRedirectionHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRedirectionHost

> bool DeleteRedirectionHost(ctx, hostID).Execute()

Delete a Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.DeleteRedirectionHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.DeleteRedirectionHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRedirectionHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.DeleteRedirectionHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRedirectionHostRequest struct via the builder pattern


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


## DisableRedirectionHost

> bool DisableRedirectionHost(ctx, hostID).Execute()

Disable a Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.DisableRedirectionHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.DisableRedirectionHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableRedirectionHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.DisableRedirectionHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRedirectionHostRequest struct via the builder pattern


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


## EnableRedirectionHost

> bool EnableRedirectionHost(ctx, hostID).Execute()

Enable a Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.EnableRedirectionHost(context.Background(), hostID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.EnableRedirectionHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableRedirectionHost`: bool
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.EnableRedirectionHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRedirectionHostRequest struct via the builder pattern


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


## GetRedirectionHost

> GetRedirectionHosts200ResponseInner GetRedirectionHost(ctx, hostID).Expand(expand).Execute()

Get a Redirection Host

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
	resp, r, err := apiClient.RedirectionHostsAPI.GetRedirectionHost(context.Background(), hostID).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.GetRedirectionHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedirectionHost`: GetRedirectionHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.GetRedirectionHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectionHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expansions | 

### Return type

[**GetRedirectionHosts200ResponseInner**](GetRedirectionHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedirectionHosts

> []GetRedirectionHosts200ResponseInner GetRedirectionHosts(ctx).Expand(expand).Execute()

Get all Redirection hosts

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
	resp, r, err := apiClient.RedirectionHostsAPI.GetRedirectionHosts(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.GetRedirectionHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedirectionHosts`: []GetRedirectionHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.GetRedirectionHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedirectionHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expansions | 

### Return type

[**[]GetRedirectionHosts200ResponseInner**](GetRedirectionHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRedirectionHost

> GetRedirectionHosts200ResponseInner UpdateRedirectionHost(ctx, hostID).UpdateRedirectionHostRequest(updateRedirectionHostRequest).Execute()

Update a Redirection Host

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
	updateRedirectionHostRequest := *openapiclient.NewUpdateRedirectionHostRequest() // UpdateRedirectionHostRequest | Redirection Host       Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedirectionHostsAPI.UpdateRedirectionHost(context.Background(), hostID).UpdateRedirectionHostRequest(updateRedirectionHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionHostsAPI.UpdateRedirectionHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRedirectionHost`: GetRedirectionHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RedirectionHostsAPI.UpdateRedirectionHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRedirectionHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRedirectionHostRequest** | [**UpdateRedirectionHostRequest**](UpdateRedirectionHostRequest.md) | Redirection Host       Payload | 

### Return type

[**GetRedirectionHosts200ResponseInner**](GetRedirectionHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

