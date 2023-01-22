# \UEIPReachabilitySubscriptionCreationApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UeReachIpSubscribe**](UEIPReachabilitySubscriptionCreationApi.md#UeReachIpSubscribe) | **Post** /{imsUeId}/access-data/ps-domain/ue-reach-subscriptions | subscribe to notifications of UE reachability



## UeReachIpSubscribe

> CreatedUeReachabilitySubscription UeReachIpSubscribe(ctx, imsUeId).UeReachabilitySubscription(ueReachabilitySubscription).PrivateIdentity(privateIdentity).Execute()

subscribe to notifications of UE reachability

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    ueReachabilitySubscription := *openapiclient.NewUeReachabilitySubscription(time.Now(), "CallbackReference_example") // UeReachabilitySubscription | 
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEIPReachabilitySubscriptionCreationApi.UeReachIpSubscribe(context.Background(), imsUeId).UeReachabilitySubscription(ueReachabilitySubscription).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEIPReachabilitySubscriptionCreationApi.UeReachIpSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UeReachIpSubscribe`: CreatedUeReachabilitySubscription
    fmt.Fprintf(os.Stdout, "Response from `UEIPReachabilitySubscriptionCreationApi.UeReachIpSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUeReachIpSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueReachabilitySubscription** | [**UeReachabilitySubscription**](UeReachabilitySubscription.md) |  | 
 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**CreatedUeReachabilitySubscription**](CreatedUeReachabilitySubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

