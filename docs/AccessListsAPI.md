# \AccessListsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessList**](AccessListsAPI.md#CreateAccessList) | **Post** /nginx/access-lists | Create a Access List
[**DeleteAccessList**](AccessListsAPI.md#DeleteAccessList) | **Delete** /nginx/access-lists/{listID} | Delete a Access List
[**GetAccessList**](AccessListsAPI.md#GetAccessList) | **Get** /nginx/access-lists/{listID} | Get a access List
[**GetAccessLists**](AccessListsAPI.md#GetAccessLists) | **Get** /nginx/access-lists | Get all access lists
[**UpdateAccessList**](AccessListsAPI.md#UpdateAccessList) | **Put** /nginx/access-lists/{listID} | Update a Access List



## CreateAccessList

> GetAccessLists200Response CreateAccessList(ctx).CreateAccessListRequest(createAccessListRequest).Execute()

Create a Access List

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
	createAccessListRequest := *openapiclient.NewCreateAccessListRequest("My Access List") // CreateAccessListRequest | Access List Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessListsAPI.CreateAccessList(context.Background()).CreateAccessListRequest(createAccessListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.CreateAccessList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessList`: GetAccessLists200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.CreateAccessList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessListRequest** | [**CreateAccessListRequest**](CreateAccessListRequest.md) | Access List Payload | 

### Return type

[**GetAccessLists200Response**](GetAccessLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessList

> bool DeleteAccessList(ctx, listID).Execute()

Delete a Access List

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
	listID := int64(2) // int64 | Access List ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessListsAPI.DeleteAccessList(context.Background(), listID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.DeleteAccessList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessList`: bool
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.DeleteAccessList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listID** | **int64** | Access List ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessListRequest struct via the builder pattern


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


## GetAccessList

> GetAccessLists200Response GetAccessList(ctx, listID).Execute()

Get a access List

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
	listID := int64(1) // int64 | Access List ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessListsAPI.GetAccessList(context.Background(), listID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.GetAccessList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessList`: GetAccessLists200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.GetAccessList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listID** | **int64** | Access List ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccessLists200Response**](GetAccessLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLists

> GetAccessLists200Response GetAccessLists(ctx).Expand(expand).Execute()

Get all access lists

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
	resp, r, err := apiClient.AccessListsAPI.GetAccessLists(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.GetAccessLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessLists`: GetAccessLists200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.GetAccessLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expansions | 

### Return type

[**GetAccessLists200Response**](GetAccessLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessList

> GetAccessLists200Response UpdateAccessList(ctx, listID).UpdateAccessListRequest(updateAccessListRequest).Execute()

Update a Access List

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
	listID := int64(2) // int64 | Access List ID
	updateAccessListRequest := *openapiclient.NewUpdateAccessListRequest() // UpdateAccessListRequest | Access List Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessListsAPI.UpdateAccessList(context.Background(), listID).UpdateAccessListRequest(updateAccessListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessListsAPI.UpdateAccessList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessList`: GetAccessLists200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessListsAPI.UpdateAccessList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listID** | **int64** | Access List ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccessListRequest** | [**UpdateAccessListRequest**](UpdateAccessListRequest.md) | Access List Payload | 

### Return type

[**GetAccessLists200Response**](GetAccessLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

