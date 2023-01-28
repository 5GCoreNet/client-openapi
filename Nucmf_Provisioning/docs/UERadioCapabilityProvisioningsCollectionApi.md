# \UERadioCapabilityProvisioningsCollectionApi

All URIs are relative to *https://example.com/nucmf-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProvisioning**](UERadioCapabilityProvisioningsCollectionApi.md#CreateProvisioning) | **Post** /provisionings | Create an Individual UE radio capability provisioning



## CreateProvisioning

> RacsData CreateProvisioning(ctx).RacsData(racsData).Execute()

Create an Individual UE radio capability provisioning

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nucmf_Provisioning"
)

func main() {
    racsData := *openapiclient.NewRacsData(map[string]RacsConfiguration{"key": *openapiclient.NewRacsConfiguration("RacsId_example", []string{"ImeiTacs_example"})}) // RacsData | Create new provisionings for a given SCS/AS.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UERadioCapabilityProvisioningsCollectionApi.CreateProvisioning(context.Background()).RacsData(racsData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UERadioCapabilityProvisioningsCollectionApi.CreateProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProvisioning`: RacsData
    fmt.Fprintf(os.Stdout, "Response from `UERadioCapabilityProvisioningsCollectionApi.CreateProvisioning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **racsData** | [**RacsData**](RacsData.md) | Create new provisionings for a given SCS/AS. | 

### Return type

[**RacsData**](RacsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

