# \AccessListsAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccessListDelete**](AccessListsAPI.md#DeleteAccessListDelete) | **Delete** /nginx/access-lists/{id} | Deletes a existing Access List
[**GetAccessListGet**](AccessListsAPI.md#GetAccessListGet) | **Get** /nginx/access-lists/{id} | Returns a  Access List
[**GetAccessListList**](AccessListsAPI.md#GetAccessListList) | **Get** /nginx/access-lists | Returns a list of Access Lists
[**PostAccessListCreate**](AccessListsAPI.md#PostAccessListCreate) | **Post** /nginx/access-lists | Creates a new Access List
[**PutAccessListUpdate**](AccessListsAPI.md#PutAccessListUpdate) | **Put** /nginx/access-lists/{id} | Updates a existing Access List



## DeleteAccessListDelete

> boolAsInt DeleteAccessListDelete(ctx, id).Execute()

Deletes a existing Access List

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
	resp, r, err := apiClient.AccessListsAPI.DeleteAccessListDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.DeleteAccessListDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessListDelete`: boolAsInt
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.DeleteAccessListDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessListDeleteRequest struct via the builder pattern


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


## GetAccessListGet

> AccessListGetResponse GetAccessListGet(ctx, id).Execute()

Returns a  Access List

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
	resp, r, err := apiClient.AccessListsAPI.GetAccessListGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.GetAccessListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessListGet`: AccessListGetResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.GetAccessListGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessListGetResponse**](AccessListGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessListList

> []AccessListGetResponse GetAccessListList(ctx).Execute()

Returns a list of Access Lists

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
	resp, r, err := apiClient.AccessListsAPI.GetAccessListList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.GetAccessListList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessListList`: []AccessListGetResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.GetAccessListList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessListListRequest struct via the builder pattern


### Return type

[**[]AccessListGetResponse**](AccessListGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccessListCreate

> AccessListCreateResponse PostAccessListCreate(ctx).PostAccessListCreateRequest(postAccessListCreateRequest).Execute()

Creates a new Access List

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
	postAccessListCreateRequest := *openapiclient.NewPostAccessListCreateRequest("Name_example") // PostAccessListCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessListsAPI.PostAccessListCreate(context.Background()).PostAccessListCreateRequest(postAccessListCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.PostAccessListCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccessListCreate`: AccessListCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.PostAccessListCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessListCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postAccessListCreateRequest** | [**PostAccessListCreateRequest**](PostAccessListCreateRequest.md) |  | 

### Return type

[**AccessListCreateResponse**](AccessListCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccessListUpdate

> AccessListUpdateResponse PutAccessListUpdate(ctx, id).PutAccessListUpdateRequest(putAccessListUpdateRequest).Execute()

Updates a existing Access List

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
	putAccessListUpdateRequest := *openapiclient.NewPutAccessListUpdateRequest() // PutAccessListUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessListsAPI.PutAccessListUpdate(context.Background(), id).PutAccessListUpdateRequest(putAccessListUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.PutAccessListUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAccessListUpdate`: AccessListUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.PutAccessListUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccessListUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putAccessListUpdateRequest** | [**PutAccessListUpdateRequest**](PutAccessListUpdateRequest.md) |  | 

### Return type

[**AccessListUpdateResponse**](AccessListUpdateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

