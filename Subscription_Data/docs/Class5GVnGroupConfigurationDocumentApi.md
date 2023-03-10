# \Class5GVnGroupConfigurationDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create5GVnGroup**](Class5GVnGroupConfigurationDocumentApi.md#Create5GVnGroup) | **Put** /subscription-data/group-data/5g-vn-groups/{externalGroupId} | Create an individual 5G VN Grouop



## Create5GVnGroup

> Model5GVnGroupConfiguration Create5GVnGroup(ctx, externalGroupId).Model5GVnGroupConfiguration(model5GVnGroupConfiguration).Execute()

Create an individual 5G VN Grouop

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    externalGroupId := "externalGroupId_example" // string | 
    model5GVnGroupConfiguration := *openapiclient.NewModel5GVnGroupConfiguration() // Model5GVnGroupConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Class5GVnGroupConfigurationDocumentApi.Create5GVnGroup(context.Background(), externalGroupId).Model5GVnGroupConfiguration(model5GVnGroupConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Class5GVnGroupConfigurationDocumentApi.Create5GVnGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create5GVnGroup`: Model5GVnGroupConfiguration
    fmt.Fprintf(os.Stdout, "Response from `Class5GVnGroupConfigurationDocumentApi.Create5GVnGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreate5GVnGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model5GVnGroupConfiguration** | [**Model5GVnGroupConfiguration**](Model5GVnGroupConfiguration.md) |  | 

### Return type

[**Model5GVnGroupConfiguration**](Model5GVnGroupConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

