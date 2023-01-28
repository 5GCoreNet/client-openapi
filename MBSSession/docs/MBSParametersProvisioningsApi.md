# \MBSParametersProvisioningsApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMBSParamsProvisionings**](MBSParametersProvisioningsApi.md#GetMBSParamsProvisionings) | **Get** /mbs-pp | Request to retrieve all the active MBS Parameters Provisioning resources at the NEF.



## GetMBSParamsProvisionings

> []MbsPpData GetMBSParamsProvisionings(ctx).Execute()

Request to retrieve all the active MBS Parameters Provisioning resources at the NEF.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSParametersProvisioningsApi.GetMBSParamsProvisionings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSParametersProvisioningsApi.GetMBSParamsProvisionings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMBSParamsProvisionings`: []MbsPpData
    fmt.Fprintf(os.Stdout, "Response from `MBSParametersProvisioningsApi.GetMBSParamsProvisionings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMBSParamsProvisioningsRequest struct via the builder pattern


### Return type

[**[]MbsPpData**](MbsPpData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

