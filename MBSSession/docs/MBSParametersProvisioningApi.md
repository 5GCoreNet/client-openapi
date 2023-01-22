# \MBSParametersProvisioningApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSParamsProvisioning**](MBSParametersProvisioningApi.md#CreateMBSParamsProvisioning) | **Post** /mbs-pp | Request the creation of a new MBS Parameters Provisioning.



## CreateMBSParamsProvisioning

> MbsPpData CreateMBSParamsProvisioning(ctx).MbsPpData(mbsPpData).Execute()

Request the creation of a new MBS Parameters Provisioning.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mbsPpData := *openapiclient.NewMbsPpData("AfId_example") // MbsPpData | Representation of the new MBS Parameters Provisioning to be created at the NEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSParametersProvisioningApi.CreateMBSParamsProvisioning(context.Background()).MbsPpData(mbsPpData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSParametersProvisioningApi.CreateMBSParamsProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSParamsProvisioning`: MbsPpData
    fmt.Fprintf(os.Stdout, "Response from `MBSParametersProvisioningApi.CreateMBSParamsProvisioning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSParamsProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mbsPpData** | [**MbsPpData**](MbsPpData.md) | Representation of the new MBS Parameters Provisioning to be created at the NEF. | 

### Return type

[**MbsPpData**](MbsPpData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

