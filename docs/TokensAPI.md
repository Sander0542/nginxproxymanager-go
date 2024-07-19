# \TokensAPI

All URIs are relative to *http://localhost:81/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenRefresh**](TokensAPI.md#GetTokenRefresh) | **Get** /tokens | Returns a new token.
[**PostTokenCreate**](TokensAPI.md#PostTokenCreate) | **Post** /tokens | Creates a new token.



## GetTokenRefresh

> TokenRefreshResponse GetTokenRefresh(ctx).Execute()

Returns a new token.

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
	resp, r, err := apiClient.TokensAPI.GetTokenRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokenRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenRefresh`: TokenRefreshResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokenRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRefreshRequest struct via the builder pattern


### Return type

[**TokenRefreshResponse**](TokenRefreshResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTokenCreate

> TokenCreateResponse PostTokenCreate(ctx).PostTokenCreateRequest(postTokenCreateRequest).Execute()

Creates a new token.

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
	postTokenCreateRequest := *openapiclient.NewPostTokenCreateRequest("john@example.com", "correct horse battery staple") // PostTokenCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostTokenCreate(context.Background()).PostTokenCreateRequest(postTokenCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTokenCreate`: TokenCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTokenCreateRequest** | [**PostTokenCreateRequest**](PostTokenCreateRequest.md) |  | 

### Return type

[**TokenCreateResponse**](TokenCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

