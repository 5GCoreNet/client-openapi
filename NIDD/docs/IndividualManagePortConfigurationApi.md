# \IndividualManagePortConfigurationApi

All URIs are relative to *https://example.com/3gpp-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndManagePortConfiguration**](IndividualManagePortConfigurationApi.md#DeleteIndManagePortConfiguration) | **Delete** /{scsAsId}/configurations/{configurationId}/rds-ports/{portId} | Delete an Individual ManagePort Configuration resource to release port numbers.
[**FetchIndManagePortConfiguration**](IndividualManagePortConfigurationApi.md#FetchIndManagePortConfiguration) | **Get** /{scsAsId}/configurations/{configurationId}/rds-ports/{portId} | Read an Individual ManagePort Configuration resource to query port numbers.
[**UpdateIndManagePortConfiguration**](IndividualManagePortConfigurationApi.md#UpdateIndManagePortConfiguration) | **Put** /{scsAsId}/configurations/{configurationId}/rds-ports/{portId} | Create a new Individual ManagePort Configuration resource to reserve port numbers.



## DeleteIndManagePortConfiguration

> DeleteIndManagePortConfiguration(ctx, scsAsId, configurationId, portId).Execute()

Delete an Individual ManagePort Configuration resource to release port numbers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDD"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.
    portId := "portId_example" // string | The UE port number.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualManagePortConfigurationApi.DeleteIndManagePortConfiguration(context.Background(), scsAsId, configurationId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualManagePortConfigurationApi.DeleteIndManagePortConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 
**portId** | **string** | The UE port number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndManagePortConfigurationRequest struct via the builder pattern


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


## FetchIndManagePortConfiguration

> ManagePort FetchIndManagePortConfiguration(ctx, scsAsId, configurationId, portId).Execute()

Read an Individual ManagePort Configuration resource to query port numbers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDD"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.
    portId := "portId_example" // string | The UE port number.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualManagePortConfigurationApi.FetchIndManagePortConfiguration(context.Background(), scsAsId, configurationId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualManagePortConfigurationApi.FetchIndManagePortConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndManagePortConfiguration`: ManagePort
    fmt.Fprintf(os.Stdout, "Response from `IndividualManagePortConfigurationApi.FetchIndManagePortConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 
**portId** | **string** | The UE port number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndManagePortConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ManagePort**](ManagePort.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndManagePortConfiguration

> ManagePort UpdateIndManagePortConfiguration(ctx, scsAsId, configurationId, portId).ManagePort(managePort).Execute()

Create a new Individual ManagePort Configuration resource to reserve port numbers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDD"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    configurationId := "configurationId_example" // string | String identifying the individual NIDD configuration resource in the SCEF.
    portId := "portId_example" // string | The UE port number.
    managePort := *openapiclient.NewManagePort("AppId_example") // ManagePort | Contains information to be applied to the individual ManagePort configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualManagePortConfigurationApi.UpdateIndManagePortConfiguration(context.Background(), scsAsId, configurationId, portId).ManagePort(managePort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualManagePortConfigurationApi.UpdateIndManagePortConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndManagePortConfiguration`: ManagePort
    fmt.Fprintf(os.Stdout, "Response from `IndividualManagePortConfigurationApi.UpdateIndManagePortConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 
**portId** | **string** | The UE port number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndManagePortConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **managePort** | [**ManagePort**](ManagePort.md) | Contains information to be applied to the individual ManagePort configuration. | 

### Return type

[**ManagePort**](ManagePort.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

