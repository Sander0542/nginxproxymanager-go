# \StreamsAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStreamDelete**](StreamsAPI.md#DeleteStreamDelete) | **Delete** /nginx/streams/{id} | Deletes a existing Stream
[**GetStreamGet**](StreamsAPI.md#GetStreamGet) | **Get** /nginx/streams/{id} | Returns a  Steam
[**GetStreamList**](StreamsAPI.md#GetStreamList) | **Get** /nginx/streams | Returns a list of Steams
[**PostStreamCreate**](StreamsAPI.md#PostStreamCreate) | **Post** /nginx/streams | Creates a new Stream
[**PostStreamDisable**](StreamsAPI.md#PostStreamDisable) | **Post** /nginx/streams/{id}/disable | Disables a existing Stream
[**PostStreamEnable**](StreamsAPI.md#PostStreamEnable) | **Post** /nginx/streams/{id}/enable | Enables a existing Stream
[**PutStreamUpdate**](StreamsAPI.md#PutStreamUpdate) | **Put** /nginx/streams/{id} | Updates a existing Stream



## DeleteStreamDelete

> boolAsInt DeleteStreamDelete(ctx, id).Execute()

Deletes a existing Stream

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
	resp, r, err := apiClient.StreamsAPI.DeleteStreamDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.DeleteStreamDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStreamDelete`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.DeleteStreamDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamDeleteRequest struct via the builder pattern


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


## GetStreamGet

> StreamGetResponse GetStreamGet(ctx, id).Execute()

Returns a  Steam

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
	resp, r, err := apiClient.StreamsAPI.GetStreamGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamGet`: StreamGetResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamGetResponse**](StreamGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamList

> []StreamGetResponse GetStreamList(ctx).Execute()

Returns a list of Steams

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
	resp, r, err := apiClient.StreamsAPI.GetStreamList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.GetStreamList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamList`: []StreamGetResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.GetStreamList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamListRequest struct via the builder pattern


### Return type

[**[]StreamGetResponse**](StreamGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStreamCreate

> StreamCreateResponse PostStreamCreate(ctx).PostStreamCreateRequest(postStreamCreateRequest).Execute()

Creates a new Stream

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
	postStreamCreateRequest := *openapiclient.NewPostStreamCreateRequest(int64(123), *openapiclient.NewPutStreamUpdateRequestForwardingHost(), int64(123)) // PostStreamCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.PostStreamCreate(context.Background()).PostStreamCreateRequest(postStreamCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.PostStreamCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStreamCreate`: StreamCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.PostStreamCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStreamCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postStreamCreateRequest** | [**PostStreamCreateRequest**](PostStreamCreateRequest.md) |  | 

### Return type

[**StreamCreateResponse**](StreamCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStreamDisable

> boolAsInt PostStreamDisable(ctx, id).Execute()

Disables a existing Stream

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
	resp, r, err := apiClient.StreamsAPI.PostStreamDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.PostStreamDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStreamDisable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.PostStreamDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStreamDisableRequest struct via the builder pattern


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


## PostStreamEnable

> boolAsInt PostStreamEnable(ctx, id).Execute()

Enables a existing Stream

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
	resp, r, err := apiClient.StreamsAPI.PostStreamEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.PostStreamEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStreamEnable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.PostStreamEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStreamEnableRequest struct via the builder pattern


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


## PutStreamUpdate

> StreamUpdateResponse PutStreamUpdate(ctx, id).PutStreamUpdateRequest(putStreamUpdateRequest).Execute()

Updates a existing Stream

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
	putStreamUpdateRequest := *openapiclient.NewPutStreamUpdateRequest() // PutStreamUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.PutStreamUpdate(context.Background(), id).PutStreamUpdateRequest(putStreamUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.PutStreamUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutStreamUpdate`: StreamUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.PutStreamUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStreamUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putStreamUpdateRequest** | [**PutStreamUpdateRequest**](PutStreamUpdateRequest.md) |  | 

### Return type

[**StreamUpdateResponse**](StreamUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

