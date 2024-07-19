# \TokensAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefreshToken**](TokensAPI.md#RefreshToken) | **Get** /tokens | Refresh your access token
[**RequestToken**](TokensAPI.md#RequestToken) | **Post** /tokens | Request a new access token from credentials



## RefreshToken

> RefreshToken200Response RefreshToken(ctx).Execute()

Refresh your access token

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
	resp, r, err := apiClient.TokensAPI.RefreshToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: RefreshToken200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.RefreshToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


### Return type

[**RefreshToken200Response**](RefreshToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestToken

> RefreshToken200Response RequestToken(ctx).RequestTokenRequest(requestTokenRequest).Execute()

Request a new access token from credentials

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
	requestTokenRequest := *openapiclient.NewRequestTokenRequest("Identity_example", "Secret_example") // RequestTokenRequest | Credentials Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.RequestToken(context.Background()).RequestTokenRequest(requestTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.RequestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestToken`: RefreshToken200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.RequestToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestTokenRequest** | [**RequestTokenRequest**](RequestTokenRequest.md) | Credentials Payload | 

### Return type

[**RefreshToken200Response**](RefreshToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

