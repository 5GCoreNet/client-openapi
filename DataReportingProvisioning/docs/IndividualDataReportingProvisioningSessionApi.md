# \IndividualDataReportingProvisioningSessionApi

All URIs are relative to *https://example.com/3gpp-data-reporting-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndDataRepProvSession**](IndividualDataReportingProvisioningSessionApi.md#DeleteIndDataRepProvSession) | **Delete** /sessions/{sessionId} | Deletes an already existing Individual Data Reporting Provisioning Session resource.
[**GetIndDataRepProvSession**](IndividualDataReportingProvisioningSessionApi.md#GetIndDataRepProvSession) | **Get** /sessions/{sessionId} | Request the retrieval of an existing Individual Data Reporting Provisioning Session resource.



## DeleteIndDataRepProvSession

> DeleteIndDataRepProvSession(ctx, sessionId).Execute()

Deletes an already existing Individual Data Reporting Provisioning Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReportingProvisioning"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingProvisioningSessionApi.DeleteIndDataRepProvSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingProvisioningSessionApi.DeleteIndDataRepProvSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndDataRepProvSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndDataRepProvSession

> DataReportingProvisioningSession GetIndDataRepProvSession(ctx, sessionId).Execute()

Request the retrieval of an existing Individual Data Reporting Provisioning Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DataReportingProvisioning"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Data Reporting Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDataReportingProvisioningSessionApi.GetIndDataRepProvSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDataReportingProvisioningSessionApi.GetIndDataRepProvSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndDataRepProvSession`: DataReportingProvisioningSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualDataReportingProvisioningSessionApi.GetIndDataRepProvSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Data Reporting Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndDataRepProvSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataReportingProvisioningSession**](DataReportingProvisioningSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

