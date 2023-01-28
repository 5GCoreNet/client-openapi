# \ApplicationRequirementsCollectionDocumentApi

All URIs are relative to *https://example.com/vae-dynamic-group/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupConfiguration**](ApplicationRequirementsCollectionDocumentApi.md#CreateGroupConfiguration) | **Post** /group-configurations | VAE_Dynamice_Group resource create service Operation



## CreateGroupConfiguration

> GroupConfigurationData CreateGroupConfiguration(ctx).GroupConfigurationData(groupConfigurationData).Execute()

VAE_Dynamice_Group resource create service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_DynamicGroup"
)

func main() {
    groupConfigurationData := *openapiclient.NewGroupConfigurationData("GroupId_example", "Definition_example", "LeaderId_example", "NotifUri_example") // GroupConfigurationData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRequirementsCollectionDocumentApi.CreateGroupConfiguration(context.Background()).GroupConfigurationData(groupConfigurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRequirementsCollectionDocumentApi.CreateGroupConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupConfiguration`: GroupConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRequirementsCollectionDocumentApi.CreateGroupConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupConfigurationData** | [**GroupConfigurationData**](GroupConfigurationData.md) |  | 

### Return type

[**GroupConfigurationData**](GroupConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

