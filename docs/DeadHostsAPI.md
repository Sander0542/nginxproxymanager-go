# \DeadHostsAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeadHostDelete**](DeadHostsAPI.md#DeleteDeadHostDelete) | **Delete** /nginx/dead-hosts/{id} | Deletes a existing 404 Host
[**GetDeadHostGet**](DeadHostsAPI.md#GetDeadHostGet) | **Get** /nginx/dead-hosts/{id} | Returns a  404 Host
[**GetDeadHostList**](DeadHostsAPI.md#GetDeadHostList) | **Get** /nginx/dead-hosts | Returns a list of 404 Hosts
[**PostDeadHostCreate**](DeadHostsAPI.md#PostDeadHostCreate) | **Post** /nginx/dead-hosts | Creates a new 404 Host
[**PostDeadHostDisable**](DeadHostsAPI.md#PostDeadHostDisable) | **Post** /nginx/dead-hosts/{id}/disable | Disables a existing 404 Host
[**PostDeadHostEnable**](DeadHostsAPI.md#PostDeadHostEnable) | **Post** /nginx/dead-hosts/{id}/enable | Enables a existing 404 Host
[**PutDeadHostUpdate**](DeadHostsAPI.md#PutDeadHostUpdate) | **Put** /nginx/dead-hosts/{id} | Updates a existing 404 Host



## DeleteDeadHostDelete

> boolAsInt DeleteDeadHostDelete(ctx, id).Execute()

Deletes a existing 404 Host

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
	resp, r, err := apiClient.DeadHostsAPI.DeleteDeadHostDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.DeleteDeadHostDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeadHostDelete`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.DeleteDeadHostDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeadHostDeleteRequest struct via the builder pattern


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


## GetDeadHostGet

> DeadHostGetResponse GetDeadHostGet(ctx, id).Execute()

Returns a  404 Host

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
	resp, r, err := apiClient.DeadHostsAPI.GetDeadHostGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.GetDeadHostGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeadHostGet`: DeadHostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.GetDeadHostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeadHostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeadHostGetResponse**](DeadHostGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeadHostList

> []DeadHostGetResponse GetDeadHostList(ctx).Execute()

Returns a list of 404 Hosts

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
	resp, r, err := apiClient.DeadHostsAPI.GetDeadHostList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.GetDeadHostList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeadHostList`: []DeadHostGetResponse
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.GetDeadHostList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeadHostListRequest struct via the builder pattern


### Return type

[**[]DeadHostGetResponse**](DeadHostGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeadHostCreate

> DeadHostCreateResponse PostDeadHostCreate(ctx).PostDeadHostCreateRequest(postDeadHostCreateRequest).Execute()

Creates a new 404 Host

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
	postDeadHostCreateRequest := *openapiclient.NewPostDeadHostCreateRequest([]string{"DomainNames_example"}) // PostDeadHostCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeadHostsAPI.PostDeadHostCreate(context.Background()).PostDeadHostCreateRequest(postDeadHostCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.PostDeadHostCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeadHostCreate`: DeadHostCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.PostDeadHostCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDeadHostCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postDeadHostCreateRequest** | [**PostDeadHostCreateRequest**](PostDeadHostCreateRequest.md) |  | 

### Return type

[**DeadHostCreateResponse**](DeadHostCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeadHostDisable

> boolAsInt PostDeadHostDisable(ctx, id).Execute()

Disables a existing 404 Host

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
	resp, r, err := apiClient.DeadHostsAPI.PostDeadHostDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.PostDeadHostDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeadHostDisable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.PostDeadHostDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeadHostDisableRequest struct via the builder pattern


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


## PostDeadHostEnable

> boolAsInt PostDeadHostEnable(ctx, id).Execute()

Enables a existing 404 Host

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
	resp, r, err := apiClient.DeadHostsAPI.PostDeadHostEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.PostDeadHostEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeadHostEnable`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.PostDeadHostEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeadHostEnableRequest struct via the builder pattern


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


## PutDeadHostUpdate

> DeadHostUpdateResponse PutDeadHostUpdate(ctx, id).PutDeadHostUpdateRequest(putDeadHostUpdateRequest).Execute()

Updates a existing 404 Host

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
	putDeadHostUpdateRequest := *openapiclient.NewPutDeadHostUpdateRequest() // PutDeadHostUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeadHostsAPI.PutDeadHostUpdate(context.Background(), id).PutDeadHostUpdateRequest(putDeadHostUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeadHostsAPI.PutDeadHostUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDeadHostUpdate`: DeadHostUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DeadHostsAPI.PutDeadHostUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeadHostUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putDeadHostUpdateRequest** | [**PutDeadHostUpdateRequest**](PutDeadHostUpdateRequest.md) |  | 

### Return type

[**DeadHostUpdateResponse**](DeadHostUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

