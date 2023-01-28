# \UELocationInformationSubscriptionCreationApi

All URIs are relative to *https://example.com/ngmlc-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationUpdateSubcribe**](UELocationInformationSubscriptionCreationApi.md#LocationUpdateSubcribe) | **Post** /loc-update-subs | subscribe to notifications of UE location information



## LocationUpdateSubcribe

> LocationUpdateSubcribe(ctx).LocUpdateSubs(locUpdateSubs).Execute()

subscribe to notifications of UE location information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ngmlc_Location"
)

func main() {
    locUpdateSubs := *openapiclient.NewLocUpdateSubs("NfInstanceId_example", "NotifURI_example") // LocUpdateSubs | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UELocationInformationSubscriptionCreationApi.LocationUpdateSubcribe(context.Background()).LocUpdateSubs(locUpdateSubs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UELocationInformationSubscriptionCreationApi.LocationUpdateSubcribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationUpdateSubcribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locUpdateSubs** | [**LocUpdateSubs**](LocUpdateSubs.md) |  | 

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

