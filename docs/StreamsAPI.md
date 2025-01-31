# \StreamsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStream**](StreamsAPI.md#CreateStream) | **Post** /nginx/streams | Create a Stream
[**DeleteStream**](StreamsAPI.md#DeleteStream) | **Delete** /nginx/streams/{streamID} | Delete a Stream
[**DisableStream**](StreamsAPI.md#DisableStream) | **Post** /nginx/streams/{streamID}/disable | Disable a Stream
[**EnableStream**](StreamsAPI.md#EnableStream) | **Post** /nginx/streams/{streamID}/enable | Enable a Stream
[**GetStream**](StreamsAPI.md#GetStream) | **Get** /nginx/streams/{streamID} | Get a Stream
[**GetStreams**](StreamsAPI.md#GetStreams) | **Get** /nginx/streams | Get all streams
[**UpdateStream**](StreamsAPI.md#UpdateStream) | **Put** /nginx/streams/{streamID} | Update a Stream



## CreateStream

> CreateStream201Response CreateStream(ctx).CreateStreamRequest(createStreamRequest).Execute()

Create a Stream

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
	createStreamRequest := *openapiclient.NewCreateStreamRequest(int64(123), *openapiclient.NewCreateStreamRequestForwardingHost(), int64(123)) // CreateStreamRequest | Stream Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.CreateStream(context.Background()).CreateStreamRequest(createStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.CreateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStream`: CreateStream201Response
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.CreateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStreamRequest** | [**CreateStreamRequest**](CreateStreamRequest.md) | Stream Payload | 

### Return type

[**CreateStream201Response**](CreateStream201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStream

> bool DeleteStream(ctx, streamID).Execute()

Delete a Stream

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
	streamID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.DeleteStream(context.Background(), streamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStream`: bool
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.DeleteStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamRequest struct via the builder pattern


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


## DisableStream

> bool DisableStream(ctx, streamID).Execute()

Disable a Stream

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
	streamID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.DisableStream(context.Background(), streamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DisableStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableStream`: bool
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.DisableStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableStreamRequest struct via the builder pattern


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


## EnableStream

> bool EnableStream(ctx, streamID).Execute()

Enable a Stream

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
	streamID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.EnableStream(context.Background(), streamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.EnableStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableStream`: bool
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.EnableStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableStreamRequest struct via the builder pattern


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


## GetStream

> CreateStream201Response GetStream(ctx, streamID).Execute()

Get a Stream

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
	streamID := int64(2) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.GetStream(context.Background(), streamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStream`: CreateStream201Response
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateStream201Response**](CreateStream201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreams

> []GetProxyHosts200ResponseInner GetStreams(ctx).Expand(expand).Execute()

Get all streams

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
	resp, r, err := apiClient.StreamsAPI.GetStreams(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreams`: []GetProxyHosts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expansions | 

### Return type

[**[]GetProxyHosts200ResponseInner**](GetProxyHosts200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStream

> CreateStream201Response UpdateStream(ctx, streamID).UpdateProxyHostRequest(updateProxyHostRequest).Execute()

Update a Stream

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
	streamID := int64(2) // int64 | 
	updateProxyHostRequest := *openapiclient.NewUpdateProxyHostRequest() // UpdateProxyHostRequest | Stream Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.UpdateStream(context.Background(), streamID).UpdateProxyHostRequest(updateProxyHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.UpdateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStream`: CreateStream201Response
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.UpdateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProxyHostRequest** | [**UpdateProxyHostRequest**](UpdateProxyHostRequest.md) | Stream Payload | 

### Return type

[**CreateStream201Response**](CreateStream201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

