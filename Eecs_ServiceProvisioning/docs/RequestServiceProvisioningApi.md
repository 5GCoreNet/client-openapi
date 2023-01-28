# \RequestServiceProvisioningApi

All URIs are relative to *https://example.com/eecs-serviceprovisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestServProv**](RequestServiceProvisioningApi.md#RequestServProv) | **Post** /request | Request service provisioning information.



## RequestServProv

> ECSServProvResp RequestServProv(ctx).ECSServProvReq(eCSServProvReq).Execute()

Request service provisioning information.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eecs_ServiceProvisioning"
)

func main() {
    eCSServProvReq := *openapiclient.NewECSServProvReq("EecId_example") // ECSServProvReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestServiceProvisioningApi.RequestServProv(context.Background()).ECSServProvReq(eCSServProvReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestServiceProvisioningApi.RequestServProv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestServProv`: ECSServProvResp
    fmt.Fprintf(os.Stdout, "Response from `RequestServiceProvisioningApi.RequestServProv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestServProvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eCSServProvReq** | [**ECSServProvReq**](ECSServProvReq.md) |  | 

### Return type

[**ECSServProvResp**](ECSServProvResp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

