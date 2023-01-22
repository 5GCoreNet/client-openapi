# \NIDDDownlinkDataDeliveriesApi

All URIs are relative to *https://example.com/3gpp-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownlinkDataDelivery**](NIDDDownlinkDataDeliveriesApi.md#CreateDownlinkDataDelivery) | **Post** /{scsAsId}/configurations/{configurationId}/downlink-data-deliveries | Create an NIDD downlink data delivery resource related to a particular NIDD configuration resource.
[**FetchAllDownlinkDataDeliveries**](NIDDDownlinkDataDeliveriesApi.md#FetchAllDownlinkDataDeliveries) | **Get** /{scsAsId}/configurations/{configurationId}/downlink-data-deliveries | Read all pending NIDD downlink data delivery resources related to a particular NIDD configuration resource.



## CreateDownlinkDataDelivery

> NiddDownlinkDataTransfer CreateDownlinkDataDelivery(ctx, scsAsId, configurationId).NiddDownlinkDataTransfer(niddDownlinkDataTransfer).Execute()

Create an NIDD downlink data delivery resource related to a particular NIDD configuration resource.

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
    niddDownlinkDataTransfer := openapiclient.NiddDownlinkDataTransfer{Interface{}: new(interface{})} // NiddDownlinkDataTransfer | Contains the data to create a NIDD downlink data delivery.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NIDDDownlinkDataDeliveriesApi.CreateDownlinkDataDelivery(context.Background(), scsAsId, configurationId).NiddDownlinkDataTransfer(niddDownlinkDataTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDDownlinkDataDeliveriesApi.CreateDownlinkDataDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDownlinkDataDelivery`: NiddDownlinkDataTransfer
    fmt.Fprintf(os.Stdout, "Response from `NIDDDownlinkDataDeliveriesApi.CreateDownlinkDataDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownlinkDataDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **niddDownlinkDataTransfer** | [**NiddDownlinkDataTransfer**](NiddDownlinkDataTransfer.md) | Contains the data to create a NIDD downlink data delivery. | 

### Return type

[**NiddDownlinkDataTransfer**](NiddDownlinkDataTransfer.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllDownlinkDataDeliveries

> []NiddDownlinkDataTransfer FetchAllDownlinkDataDeliveries(ctx, scsAsId, configurationId).Execute()

Read all pending NIDD downlink data delivery resources related to a particular NIDD configuration resource.

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
    resp, r, err := apiClient.NIDDDownlinkDataDeliveriesApi.FetchAllDownlinkDataDeliveries(context.Background(), scsAsId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NIDDDownlinkDataDeliveriesApi.FetchAllDownlinkDataDeliveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllDownlinkDataDeliveries`: []NiddDownlinkDataTransfer
    fmt.Fprintf(os.Stdout, "Response from `NIDDDownlinkDataDeliveriesApi.FetchAllDownlinkDataDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllDownlinkDataDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]NiddDownlinkDataTransfer**](NiddDownlinkDataTransfer.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

