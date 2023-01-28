# \NonUEN2MessagesSubscriptionsCollectionCollectionApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NonUeN2InfoSubscribe**](NonUEN2MessagesSubscriptionsCollectionCollectionApi.md#NonUeN2InfoSubscribe) | **Post** /non-ue-n2-messages/subscriptions | Namf_Communication Non UE N2 Info Subscribe service Operation



## NonUeN2InfoSubscribe

> NonUeN2InfoSubscriptionCreatedData NonUeN2InfoSubscribe(ctx).NonUeN2InfoSubscriptionCreateData(nonUeN2InfoSubscriptionCreateData).Execute()

Namf_Communication Non UE N2 Info Subscribe service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    nonUeN2InfoSubscriptionCreateData := *openapiclient.NewNonUeN2InfoSubscriptionCreateData(*openapiclient.NewN2InformationClass(), "N2NotifyCallbackUri_example") // NonUeN2InfoSubscriptionCreateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonUEN2MessagesSubscriptionsCollectionCollectionApi.NonUeN2InfoSubscribe(context.Background()).NonUeN2InfoSubscriptionCreateData(nonUeN2InfoSubscriptionCreateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonUEN2MessagesSubscriptionsCollectionCollectionApi.NonUeN2InfoSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonUeN2InfoSubscribe`: NonUeN2InfoSubscriptionCreatedData
    fmt.Fprintf(os.Stdout, "Response from `NonUEN2MessagesSubscriptionsCollectionCollectionApi.NonUeN2InfoSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonUeN2InfoSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonUeN2InfoSubscriptionCreateData** | [**NonUeN2InfoSubscriptionCreateData**](NonUeN2InfoSubscriptionCreateData.md) |  | 

### Return type

[**NonUeN2InfoSubscriptionCreatedData**](NonUeN2InfoSubscriptionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

