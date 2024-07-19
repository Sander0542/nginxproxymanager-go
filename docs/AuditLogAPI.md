# \AuditLogAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLog**](AuditLogAPI.md#GetAuditLog) | **Get** /audit-log | Get Audit Log



## GetAuditLog

> GetAuditLog200Response GetAuditLog(ctx).Execute()

Get Audit Log

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
	resp, r, err := apiClient.AuditLogAPI.GetAuditLog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.GetAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLog`: GetAuditLog200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.GetAuditLog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogRequest struct via the builder pattern


### Return type

[**GetAuditLog200Response**](GetAuditLog200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

