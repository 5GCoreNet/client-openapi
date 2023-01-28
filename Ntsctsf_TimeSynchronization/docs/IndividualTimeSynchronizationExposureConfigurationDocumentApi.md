# \IndividualTimeSynchronizationExposureConfigurationDocumentApi

All URIs are relative to *https://example.com/ntsctsf-time-sync/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualTimeSynchronizationExposureConfiguration**](IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#CreateIndividualTimeSynchronizationExposureConfiguration) | **Post** /subscriptions/{subscriptionId}/configurations | Craete a new Individual Time Synchronization Exposure Configuration
[**DeleteIndividualTimeSynchronizationExposureConfiguration**](IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#DeleteIndividualTimeSynchronizationExposureConfiguration) | **Delete** /subscriptions/{subscriptionId}/configurations/{configurationId} | Delete an Individual TimeSynchronization Exposure Configuration
[**GetIndividualTimeSynchronizationExposureConfiguration**](IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#GetIndividualTimeSynchronizationExposureConfiguration) | **Get** /subscriptions/{subscriptionId}/configurations/{configurationId} | Reads an existing Individual Time Synchronization Exposure Configuration
[**ReplaceIndividualTimeSynchronizationExposureConfiguration**](IndividualTimeSynchronizationExposureConfigurationDocumentApi.md#ReplaceIndividualTimeSynchronizationExposureConfiguration) | **Put** /subscriptions/{subscriptionId}/configurations/{configurationId} | Replace an individual Time Synchronization Exposure Configuration



## CreateIndividualTimeSynchronizationExposureConfiguration

> TimeSyncExposureConfig CreateIndividualTimeSynchronizationExposureConfiguration(ctx, subscriptionId).TimeSyncExposureConfig1(timeSyncExposureConfig1).Execute()

Craete a new Individual Time Synchronization Exposure Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_TimeSynchronization"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription.
    timeSyncExposureConfig1 := *openapiclient.NewTimeSyncExposureConfig1(int32(123), *openapiclient.NewPtpInstance(*openapiclient.NewInstanceType(), *openapiclient.NewProtocol(), "PtpProfile_example"), int32(123), "ConfigNotifId_example", "ConfigNotifUri_example") // TimeSyncExposureConfig1 | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.CreateIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId).TimeSyncExposureConfig1(timeSyncExposureConfig1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureConfigurationDocumentApi.CreateIndividualTimeSynchronizationExposureConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualTimeSynchronizationExposureConfiguration`: TimeSyncExposureConfig
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureConfigurationDocumentApi.CreateIndividualTimeSynchronizationExposureConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualTimeSynchronizationExposureConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeSyncExposureConfig1** | [**TimeSyncExposureConfig1**](TimeSyncExposureConfig1.md) | Contains the information for the creation the resource. | 

### Return type

[**TimeSyncExposureConfig**](TimeSyncExposureConfig.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualTimeSynchronizationExposureConfiguration

> DeleteIndividualTimeSynchronizationExposureConfiguration(ctx, subscriptionId, configurationId).Execute()

Delete an Individual TimeSynchronization Exposure Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_TimeSynchronization"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription.
    configurationId := "configurationId_example" // string | String identifying an Individual Time Synchronization Exposure Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.DeleteIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureConfigurationDocumentApi.DeleteIndividualTimeSynchronizationExposureConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription. | 
**configurationId** | **string** | String identifying an Individual Time Synchronization Exposure Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualTimeSynchronizationExposureConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualTimeSynchronizationExposureConfiguration

> TimeSyncExposureConfig GetIndividualTimeSynchronizationExposureConfiguration(ctx, subscriptionId, configurationId).Execute()

Reads an existing Individual Time Synchronization Exposure Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_TimeSynchronization"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription.
    configurationId := "configurationId_example" // string | String identifying an Individual Time Synchronization Exposure Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.GetIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureConfigurationDocumentApi.GetIndividualTimeSynchronizationExposureConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualTimeSynchronizationExposureConfiguration`: TimeSyncExposureConfig
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureConfigurationDocumentApi.GetIndividualTimeSynchronizationExposureConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription. | 
**configurationId** | **string** | String identifying an Individual Time Synchronization Exposure Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualTimeSynchronizationExposureConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimeSyncExposureConfig**](TimeSyncExposureConfig.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIndividualTimeSynchronizationExposureConfiguration

> TimeSyncExposureConfig ReplaceIndividualTimeSynchronizationExposureConfiguration(ctx, subscriptionId, configurationId).TimeSyncExposureConfig1(timeSyncExposureConfig1).Execute()

Replace an individual Time Synchronization Exposure Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_TimeSynchronization"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying an Individual Time Synchronization Exposure Subscription.
    configurationId := "configurationId_example" // string | String identifying an Individual Time Synchronization Exposure Configuration.
    timeSyncExposureConfig1 := *openapiclient.NewTimeSyncExposureConfig1(int32(123), *openapiclient.NewPtpInstance(*openapiclient.NewInstanceType(), *openapiclient.NewProtocol(), "PtpProfile_example"), int32(123), "ConfigNotifId_example", "ConfigNotifUri_example") // TimeSyncExposureConfig1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.ReplaceIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId, configurationId).TimeSyncExposureConfig1(timeSyncExposureConfig1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureConfigurationDocumentApi.ReplaceIndividualTimeSynchronizationExposureConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualTimeSynchronizationExposureConfiguration`: TimeSyncExposureConfig
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureConfigurationDocumentApi.ReplaceIndividualTimeSynchronizationExposureConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an Individual Time Synchronization Exposure Subscription. | 
**configurationId** | **string** | String identifying an Individual Time Synchronization Exposure Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualTimeSynchronizationExposureConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeSyncExposureConfig1** | [**TimeSyncExposureConfig1**](TimeSyncExposureConfig1.md) |  | 

### Return type

[**TimeSyncExposureConfig**](TimeSyncExposureConfig.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

