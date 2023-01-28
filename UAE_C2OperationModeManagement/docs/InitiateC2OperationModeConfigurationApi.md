# \InitiateC2OperationModeConfigurationApi

All URIs are relative to *https://example.com/uae-c2opmode-mngt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitiateC2OpModeConfig**](InitiateC2OperationModeConfigurationApi.md#InitiateC2OpModeConfig) | **Post** /initiate | Request the provisioning of C2 Operation Mode configuration information for a UAS (i.e. pair of UAV and UAV-C).



## InitiateC2OpModeConfig

> C2Result InitiateC2OpModeConfig(ctx).ConfigureData(configureData).Execute()

Request the provisioning of C2 Operation Mode configuration information for a UAS (i.e. pair of UAV and UAV-C).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/UAE_C2OperationModeManagement"
)

func main() {
    configureData := *openapiclient.NewConfigureData("UassId_example", openapiclient.UasId{Interface{}: new(interface{})}, []openapiclient.C2CommMode{*openapiclient.NewC2CommMode()}, []openapiclient.C2CommModeSwitching{*openapiclient.NewC2CommModeSwitching()}, "NotificationUri_example", *openapiclient.NewC2CommMode(), *openapiclient.NewC2SwitchPolicies()) // ConfigureData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InitiateC2OperationModeConfigurationApi.InitiateC2OpModeConfig(context.Background()).ConfigureData(configureData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InitiateC2OperationModeConfigurationApi.InitiateC2OpModeConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateC2OpModeConfig`: C2Result
    fmt.Fprintf(os.Stdout, "Response from `InitiateC2OperationModeConfigurationApi.InitiateC2OpModeConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateC2OpModeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureData** | [**ConfigureData**](ConfigureData.md) |  | 

### Return type

[**C2Result**](C2Result.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

