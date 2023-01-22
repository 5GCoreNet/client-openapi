# \IndividualNIDDDownlinkDataDeliveryApi

All URIs are relative to *https://example.com/3gpp-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndDownlinkDataDelivery**](IndividualNIDDDownlinkDataDeliveryApi.md#DeleteIndDownlinkDataDelivery) | **Delete** /{scsAsId}/configurations/{configurationId}/downlink-data-deliveries/{downlinkDataDeliveryId} | Delete an NIDD downlink data delivery resource.
[**FetchIndDownlinkDataDelivery**](IndividualNIDDDownlinkDataDeliveryApi.md#FetchIndDownlinkDataDelivery) | **Get** /{scsAsId}/configurations/{configurationId}/downlink-data-deliveries/{downlinkDataDeliveryId} | Read pending NIDD downlink data delivery resource.
[**ModifyIndDownlinkDataDelivery**](IndividualNIDDDownlinkDataDeliveryApi.md#ModifyIndDownlinkDataDelivery) | **Patch** /{scsAsId}/configurations/{configurationId}/downlink-data-deliveries/{downlinkDataDeliveryId} | Modify an existing Individual NIDD downlink data delivery resource.
[**UpdateIndDownlinkDataDelivery**](IndividualNIDDDownlinkDataDeliveryApi.md#UpdateIndDownlinkDataDelivery) | **Put** /{scsAsId}/configurations/{configurationId}/downlink-data-deliveries/{downlinkDataDeliveryId} | Replace an NIDD downlink data delivery resource.



## DeleteIndDownlinkDataDelivery

> DeleteIndDownlinkDataDelivery(ctx, scsAsId, configurationId, downlinkDataDeliveryId).Execute()

Delete an NIDD downlink data delivery resource.

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
    downlinkDataDeliveryId := "downlinkDataDeliveryId_example" // string | String identifying the individual NIDD downlink data delivery in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDDownlinkDataDeliveryApi.DeleteIndDownlinkDataDelivery(context.Background(), scsAsId, configurationId, downlinkDataDeliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDDownlinkDataDeliveryApi.DeleteIndDownlinkDataDelivery``: %v\n", err)
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
**downlinkDataDeliveryId** | **string** | String identifying the individual NIDD downlink data delivery in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndDownlinkDataDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndDownlinkDataDelivery

> NiddDownlinkDataTransfer FetchIndDownlinkDataDelivery(ctx, scsAsId, configurationId, downlinkDataDeliveryId).Execute()

Read pending NIDD downlink data delivery resource.

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
    downlinkDataDeliveryId := "downlinkDataDeliveryId_example" // string | String identifying the individual NIDD downlink data delivery in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDDownlinkDataDeliveryApi.FetchIndDownlinkDataDelivery(context.Background(), scsAsId, configurationId, downlinkDataDeliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDDownlinkDataDeliveryApi.FetchIndDownlinkDataDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndDownlinkDataDelivery`: NiddDownlinkDataTransfer
    fmt.Fprintf(os.Stdout, "Response from `IndividualNIDDDownlinkDataDeliveryApi.FetchIndDownlinkDataDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 
**downlinkDataDeliveryId** | **string** | String identifying the individual NIDD downlink data delivery in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndDownlinkDataDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NiddDownlinkDataTransfer**](NiddDownlinkDataTransfer.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndDownlinkDataDelivery

> NiddDownlinkDataTransfer ModifyIndDownlinkDataDelivery(ctx, scsAsId, configurationId, downlinkDataDeliveryId).NiddDownlinkDataTransferPatch(niddDownlinkDataTransferPatch).Execute()

Modify an existing Individual NIDD downlink data delivery resource.

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
    downlinkDataDeliveryId := "downlinkDataDeliveryId_example" // string | String identifying the individual NIDD downlink data delivery in the SCEF.
    niddDownlinkDataTransferPatch := *openapiclient.NewNiddDownlinkDataTransferPatch() // NiddDownlinkDataTransferPatch | Contains the parameters to update an individual NIDD downlink data delivery resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDDownlinkDataDeliveryApi.ModifyIndDownlinkDataDelivery(context.Background(), scsAsId, configurationId, downlinkDataDeliveryId).NiddDownlinkDataTransferPatch(niddDownlinkDataTransferPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDDownlinkDataDeliveryApi.ModifyIndDownlinkDataDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndDownlinkDataDelivery`: NiddDownlinkDataTransfer
    fmt.Fprintf(os.Stdout, "Response from `IndividualNIDDDownlinkDataDeliveryApi.ModifyIndDownlinkDataDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 
**downlinkDataDeliveryId** | **string** | String identifying the individual NIDD downlink data delivery in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndDownlinkDataDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **niddDownlinkDataTransferPatch** | [**NiddDownlinkDataTransferPatch**](NiddDownlinkDataTransferPatch.md) | Contains the parameters to update an individual NIDD downlink data delivery resource. | 

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


## UpdateIndDownlinkDataDelivery

> NiddDownlinkDataTransfer UpdateIndDownlinkDataDelivery(ctx, scsAsId, configurationId, downlinkDataDeliveryId).NiddDownlinkDataTransfer(niddDownlinkDataTransfer).Execute()

Replace an NIDD downlink data delivery resource.

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
    downlinkDataDeliveryId := "downlinkDataDeliveryId_example" // string | String identifying the individual NIDD downlink data delivery in the SCEF.
    niddDownlinkDataTransfer := openapiclient.NiddDownlinkDataTransfer{Interface{}: new(interface{})} // NiddDownlinkDataTransfer | Contains information to be applied to the individual NIDD downlink data delivery.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNIDDDownlinkDataDeliveryApi.UpdateIndDownlinkDataDelivery(context.Background(), scsAsId, configurationId, downlinkDataDeliveryId).NiddDownlinkDataTransfer(niddDownlinkDataTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNIDDDownlinkDataDeliveryApi.UpdateIndDownlinkDataDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndDownlinkDataDelivery`: NiddDownlinkDataTransfer
    fmt.Fprintf(os.Stdout, "Response from `IndividualNIDDDownlinkDataDeliveryApi.UpdateIndDownlinkDataDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**configurationId** | **string** | String identifying the individual NIDD configuration resource in the SCEF. | 
**downlinkDataDeliveryId** | **string** | String identifying the individual NIDD downlink data delivery in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndDownlinkDataDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **niddDownlinkDataTransfer** | [**NiddDownlinkDataTransfer**](NiddDownlinkDataTransfer.md) | Contains information to be applied to the individual NIDD downlink data delivery. | 

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

