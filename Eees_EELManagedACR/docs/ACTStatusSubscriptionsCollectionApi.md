# \ACTStatusSubscriptionsCollectionApi

All URIs are relative to *https://example.com/eees-eel-acr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateACTStatusSubsc**](ACTStatusSubscriptionsCollectionApi.md#CreateACTStatusSubsc) | **Post** /subscriptions | Request the creation of a subscription to ACT status reporting.
[**GetACTStatusSubscriptions**](ACTStatusSubscriptionsCollectionApi.md#GetACTStatusSubscriptions) | **Get** /subscriptions | Retrieve all the active ACT Status Subscriptions managed by the EES.



## CreateACTStatusSubsc

> ACTStatusSubsc CreateACTStatusSubsc(ctx).ACTStatusSubsc(aCTStatusSubsc).Execute()

Request the creation of a subscription to ACT status reporting.

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
    aCTStatusSubsc := *openapiclient.NewACTStatusSubsc("EasId_example", "NotificationUri_example") // ACTStatusSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACTStatusSubscriptionsCollectionApi.CreateACTStatusSubsc(context.Background()).ACTStatusSubsc(aCTStatusSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACTStatusSubscriptionsCollectionApi.CreateACTStatusSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateACTStatusSubsc`: ACTStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `ACTStatusSubscriptionsCollectionApi.CreateACTStatusSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateACTStatusSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCTStatusSubsc** | [**ACTStatusSubsc**](ACTStatusSubsc.md) |  | 

### Return type

[**ACTStatusSubsc**](ACTStatusSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetACTStatusSubscriptions

> []ACTStatusSubsc GetACTStatusSubscriptions(ctx).Execute()

Retrieve all the active ACT Status Subscriptions managed by the EES.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACTStatusSubscriptionsCollectionApi.GetACTStatusSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACTStatusSubscriptionsCollectionApi.GetACTStatusSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACTStatusSubscriptions`: []ACTStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `ACTStatusSubscriptionsCollectionApi.GetACTStatusSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetACTStatusSubscriptionsRequest struct via the builder pattern


### Return type

[**[]ACTStatusSubsc**](ACTStatusSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

