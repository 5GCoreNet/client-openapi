# \IndividualNWDAFMLModelProvisionSubscriptionDocumentApi

All URIs are relative to *https://example.com/nnwdaf-mlmodelprovision/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNWDAFMLModelProvisionSubcription**](IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.md#DeleteNWDAFMLModelProvisionSubcription) | **Delete** /subscriptions/{subscriptionId} | Delete an existing Individual NWDAF ML Model Provision Subscription.
[**UpdateNWDAFMLModelProvisionSubcription**](IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.md#UpdateNWDAFMLModelProvisionSubcription) | **Put** /subscriptions/{subscriptionId} | update an existing Individual NWDAF ML Model Provision Subscription



## DeleteNWDAFMLModelProvisionSubcription

> DeleteNWDAFMLModelProvisionSubcription(ctx, subscriptionId).Execute()

Delete an existing Individual NWDAF ML Model Provision Subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnwdaf_MLModelProvision"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Nnwdaf_MLModelProvision Service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.DeleteNWDAFMLModelProvisionSubcription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.DeleteNWDAFMLModelProvisionSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Nnwdaf_MLModelProvision Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNWDAFMLModelProvisionSubcriptionRequest struct via the builder pattern


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


## UpdateNWDAFMLModelProvisionSubcription

> NwdafMLModelProvSubsc UpdateNWDAFMLModelProvisionSubcription(ctx, subscriptionId).NwdafMLModelProvSubsc(nwdafMLModelProvSubsc).Execute()

update an existing Individual NWDAF ML Model Provision Subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnwdaf_MLModelProvision"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Nnwdaf_MLModelProvision Service.
    nwdafMLModelProvSubsc := *openapiclient.NewNwdafMLModelProvSubsc([]openapiclient.MLEventSubscription{*openapiclient.NewMLEventSubscription(*openapiclient.NewNwdafEvent(), *openapiclient.NewEventFilter())}, "NotifUri_example") // NwdafMLModelProvSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.UpdateNWDAFMLModelProvisionSubcription(context.Background(), subscriptionId).NwdafMLModelProvSubsc(nwdafMLModelProvSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.UpdateNWDAFMLModelProvisionSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNWDAFMLModelProvisionSubcription`: NwdafMLModelProvSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualNWDAFMLModelProvisionSubscriptionDocumentApi.UpdateNWDAFMLModelProvisionSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Nnwdaf_MLModelProvision Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNWDAFMLModelProvisionSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nwdafMLModelProvSubsc** | [**NwdafMLModelProvSubsc**](NwdafMLModelProvSubsc.md) |  | 

### Return type

[**NwdafMLModelProvSubsc**](NwdafMLModelProvSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

