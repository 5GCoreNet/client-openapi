# \ManagePortConfigurationsApi

All URIs are relative to *https://example.com/3gpp-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllManagePortConfigurations**](ManagePortConfigurationsApi.md#FetchAllManagePortConfigurations) | **Get** /{scsAsId}/configurations/{configurationId}/rds-ports | Read all RDS ManagePort Configurations.



## FetchAllManagePortConfigurations

> []ManagePort FetchAllManagePortConfigurations(ctx, scsAsId, configurationId).Execute()

Read all RDS ManagePort Configurations.

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
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagePortConfigurationsApi.FetchAllManagePortConfigurations(context.Background(), scsAsId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagePortConfigurationsApi.FetchAllManagePortConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllManagePortConfigurations`: []ManagePort
    fmt.Fprintf(os.Stdout, "Response from `ManagePortConfigurationsApi.FetchAllManagePortConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllManagePortConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ManagePort**](ManagePort.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

