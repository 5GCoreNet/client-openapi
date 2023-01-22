# \DefaultApi

All URIs are relative to *https://example.com/nchf-offlineonlycharging/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfflinechargingdataOfflineChargingDataRefReleasePost**](DefaultApi.md#OfflinechargingdataOfflineChargingDataRefReleasePost) | **Post** /offlinechargingdata/{OfflineChargingDataRef}/release | 
[**OfflinechargingdataOfflineChargingDataRefUpdatePost**](DefaultApi.md#OfflinechargingdataOfflineChargingDataRefUpdatePost) | **Post** /offlinechargingdata/{OfflineChargingDataRef}/update | 
[**OfflinechargingdataPost**](DefaultApi.md#OfflinechargingdataPost) | **Post** /offlinechargingdata | 



## OfflinechargingdataOfflineChargingDataRefReleasePost

> OfflinechargingdataOfflineChargingDataRefReleasePost(ctx, offlineChargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    offlineChargingDataRef := "offlineChargingDataRef_example" // string | a unique identifier for a charging data resource in a PLMN
    chargingDataRequest := *openapiclient.NewChargingDataRequest(*openapiclient.NewNFIdentification(*openapiclient.NewNodeFunctionality()), time.Now(), int32(123)) // ChargingDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OfflinechargingdataOfflineChargingDataRefReleasePost(context.Background(), offlineChargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OfflinechargingdataOfflineChargingDataRefReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**offlineChargingDataRef** | **string** | a unique identifier for a charging data resource in a PLMN | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfflinechargingdataOfflineChargingDataRefReleasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargingDataRequest** | [**ChargingDataRequest**](ChargingDataRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfflinechargingdataOfflineChargingDataRefUpdatePost

> ChargingDataResponse OfflinechargingdataOfflineChargingDataRefUpdatePost(ctx, offlineChargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    offlineChargingDataRef := "offlineChargingDataRef_example" // string | a unique identifier for a charging data resource in a PLMN
    chargingDataRequest := *openapiclient.NewChargingDataRequest(*openapiclient.NewNFIdentification(*openapiclient.NewNodeFunctionality()), time.Now(), int32(123)) // ChargingDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OfflinechargingdataOfflineChargingDataRefUpdatePost(context.Background(), offlineChargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OfflinechargingdataOfflineChargingDataRefUpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfflinechargingdataOfflineChargingDataRefUpdatePost`: ChargingDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OfflinechargingdataOfflineChargingDataRefUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**offlineChargingDataRef** | **string** | a unique identifier for a charging data resource in a PLMN | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfflinechargingdataOfflineChargingDataRefUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargingDataRequest** | [**ChargingDataRequest**](ChargingDataRequest.md) |  | 

### Return type

[**ChargingDataResponse**](ChargingDataResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfflinechargingdataPost

> ChargingDataResponse OfflinechargingdataPost(ctx).ChargingDataRequest(chargingDataRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    chargingDataRequest := *openapiclient.NewChargingDataRequest(*openapiclient.NewNFIdentification(*openapiclient.NewNodeFunctionality()), time.Now(), int32(123)) // ChargingDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OfflinechargingdataPost(context.Background()).ChargingDataRequest(chargingDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OfflinechargingdataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfflinechargingdataPost`: ChargingDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OfflinechargingdataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfflinechargingdataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargingDataRequest** | [**ChargingDataRequest**](ChargingDataRequest.md) |  | 

### Return type

[**ChargingDataResponse**](ChargingDataResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

