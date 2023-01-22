# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnwdaf-mlmodelprovision/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNWDAFMLModelProvisionSubcription**](SubscriptionsCollectionApi.md#CreateNWDAFMLModelProvisionSubcription) | **Post** /subscriptions | Create a new Individual NWDAF ML Model Provision Subscription resource.



## CreateNWDAFMLModelProvisionSubcription

> NwdafMLModelProvSubsc CreateNWDAFMLModelProvisionSubcription(ctx).NwdafMLModelProvSubsc(nwdafMLModelProvSubsc).Execute()

Create a new Individual NWDAF ML Model Provision Subscription resource.

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
    nwdafMLModelProvSubsc := *openapiclient.NewNwdafMLModelProvSubsc([]openapiclient.MLEventSubscription{*openapiclient.NewMLEventSubscription(*openapiclient.NewNwdafEvent(), *openapiclient.NewEventFilter())}, "NotifUri_example") // NwdafMLModelProvSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionApi.CreateNWDAFMLModelProvisionSubcription(context.Background()).NwdafMLModelProvSubsc(nwdafMLModelProvSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionApi.CreateNWDAFMLModelProvisionSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNWDAFMLModelProvisionSubcription`: NwdafMLModelProvSubsc
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionApi.CreateNWDAFMLModelProvisionSubcription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNWDAFMLModelProvisionSubcriptionRequest struct via the builder pattern


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

