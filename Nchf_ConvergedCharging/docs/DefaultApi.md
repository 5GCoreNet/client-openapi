# \DefaultApi

All URIs are relative to *https://example.com/nchf-convergedcharging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChargingdataChargingDataRefReleasePost**](DefaultApi.md#ChargingdataChargingDataRefReleasePost) | **Post** /chargingdata/{ChargingDataRef}/release | 
[**ChargingdataChargingDataRefUpdatePost**](DefaultApi.md#ChargingdataChargingDataRefUpdatePost) | **Post** /chargingdata/{ChargingDataRef}/update | 
[**ChargingdataPost**](DefaultApi.md#ChargingdataPost) | **Post** /chargingdata | 



## ChargingdataChargingDataRefReleasePost

> ChargingdataChargingDataRefReleasePost(ctx, chargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Nchf_ConvergedCharging"
)

func main() {
    chargingDataRef := "chargingDataRef_example" // string | a unique identifier for a charging data resource in a PLMN
    chargingDataRequest := *openapiclient.NewChargingDataRequest(*openapiclient.NewNFIdentification(*openapiclient.NewNodeFunctionality()), time.Now(), int32(123)) // ChargingDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ChargingdataChargingDataRefReleasePost(context.Background(), chargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChargingdataChargingDataRefReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chargingDataRef** | **string** | a unique identifier for a charging data resource in a PLMN | 

### Other Parameters

Other parameters are passed through a pointer to a apiChargingdataChargingDataRefReleasePostRequest struct via the builder pattern


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


## ChargingdataChargingDataRefUpdatePost

> ChargingDataResponse ChargingdataChargingDataRefUpdatePost(ctx, chargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Nchf_ConvergedCharging"
)

func main() {
    chargingDataRef := "chargingDataRef_example" // string | a unique identifier for a charging data resource in a PLMN
    chargingDataRequest := *openapiclient.NewChargingDataRequest(*openapiclient.NewNFIdentification(*openapiclient.NewNodeFunctionality()), time.Now(), int32(123)) // ChargingDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ChargingdataChargingDataRefUpdatePost(context.Background(), chargingDataRef).ChargingDataRequest(chargingDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChargingdataChargingDataRefUpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChargingdataChargingDataRefUpdatePost`: ChargingDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ChargingdataChargingDataRefUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chargingDataRef** | **string** | a unique identifier for a charging data resource in a PLMN | 

### Other Parameters

Other parameters are passed through a pointer to a apiChargingdataChargingDataRefUpdatePostRequest struct via the builder pattern


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


## ChargingdataPost

> ChargingDataResponse ChargingdataPost(ctx).ChargingDataRequest(chargingDataRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Nchf_ConvergedCharging"
)

func main() {
    chargingDataRequest := *openapiclient.NewChargingDataRequest(*openapiclient.NewNFIdentification(*openapiclient.NewNodeFunctionality()), time.Now(), int32(123)) // ChargingDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ChargingdataPost(context.Background()).ChargingDataRequest(chargingDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChargingdataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChargingdataPost`: ChargingDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ChargingdataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargingdataPostRequest struct via the builder pattern


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

