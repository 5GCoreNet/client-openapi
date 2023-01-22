# \NotificationSubscriptionCRUDApi

All URIs are relative to *https://example.com/nudsf-dr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAndUpdateNotificationSubscription**](NotificationSubscriptionCRUDApi.md#CreateAndUpdateNotificationSubscription) | **Put** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | NotificationSubscription Create/Update
[**DeleteNotificationSubscription**](NotificationSubscriptionCRUDApi.md#DeleteNotificationSubscription) | **Delete** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | Delete a Notification Subscription of the storage
[**GetNotificationSubscription**](NotificationSubscriptionCRUDApi.md#GetNotificationSubscription) | **Get** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | Notification subscription retrieval
[**UpdateNotificationSubscription**](NotificationSubscriptionCRUDApi.md#UpdateNotificationSubscription) | **Patch** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | NotificationSubscription update



## CreateAndUpdateNotificationSubscription

> NotificationSubscription CreateAndUpdateNotificationSubscription(ctx, realmId, storageId, subscriptionId).NotificationSubscription(notificationSubscription).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).Execute()

NotificationSubscription Create/Update

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
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    subscriptionId := "Subscription01" // string | Identifier of the NotificationSubscription
    notificationSubscription := *openapiclient.NewNotificationSubscription(*openapiclient.NewClientId(), "CallbackReference_example") // NotificationSubscription | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSubscriptionCRUDApi.CreateAndUpdateNotificationSubscription(context.Background(), realmId, storageId, subscriptionId).NotificationSubscription(notificationSubscription).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSubscriptionCRUDApi.CreateAndUpdateNotificationSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAndUpdateNotificationSubscription`: NotificationSubscription
    fmt.Fprintf(os.Stdout, "Response from `NotificationSubscriptionCRUDApi.CreateAndUpdateNotificationSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**subscriptionId** | **string** | Identifier of the NotificationSubscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndUpdateNotificationSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notificationSubscription** | [**NotificationSubscription**](NotificationSubscription.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 

### Return type

[**NotificationSubscription**](NotificationSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationSubscription

> []NotificationSubscription DeleteNotificationSubscription(ctx, realmId, storageId, subscriptionId).ClientId(clientId).GetPrevious(getPrevious).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()

Delete a Notification Subscription of the storage



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
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    subscriptionId := "Subscription01" // string | Identifier of the NotificationSubscription
    clientId := map[string][]openapiclient.ClientId{ ... } // ClientId | Identifies the NF or NFSet
    getPrevious := true // bool | Retrieve the NotificationSubscription before delete (optional) (default to false)
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSubscriptionCRUDApi.DeleteNotificationSubscription(context.Background(), realmId, storageId, subscriptionId).ClientId(clientId).GetPrevious(getPrevious).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSubscriptionCRUDApi.DeleteNotificationSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNotificationSubscription`: []NotificationSubscription
    fmt.Fprintf(os.Stdout, "Response from `NotificationSubscriptionCRUDApi.DeleteNotificationSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**subscriptionId** | **string** | Identifier of the NotificationSubscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | [**ClientId**](ClientId.md) | Identifies the NF or NFSet | 
 **getPrevious** | **bool** | Retrieve the NotificationSubscription before delete | [default to false]
 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**[]NotificationSubscription**](NotificationSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSubscription

> NotificationSubscription GetNotificationSubscription(ctx, realmId, storageId, subscriptionId).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Notification subscription retrieval



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
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    subscriptionId := "Subscription01" // string | Identifier of the NotificationSubscription
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSubscriptionCRUDApi.GetNotificationSubscription(context.Background(), realmId, storageId, subscriptionId).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSubscriptionCRUDApi.GetNotificationSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationSubscription`: NotificationSubscription
    fmt.Fprintf(os.Stdout, "Response from `NotificationSubscriptionCRUDApi.GetNotificationSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**subscriptionId** | **string** | Identifier of the NotificationSubscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | **string** | Features required to be supported by the target NF | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**NotificationSubscription**](NotificationSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationSubscription

> PatchResult UpdateNotificationSubscription(ctx, realmId, storageId, subscriptionId).PatchItem(patchItem).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()

NotificationSubscription update



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
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    subscriptionId := "Subscription01" // string | Identifier of the NotificationSubscription
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | data to patch
    ifMatch := "ifMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSubscriptionCRUDApi.UpdateNotificationSubscription(context.Background(), realmId, storageId, subscriptionId).PatchItem(patchItem).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSubscriptionCRUDApi.UpdateNotificationSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationSubscription`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `NotificationSubscriptionCRUDApi.UpdateNotificationSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**subscriptionId** | **string** | Identifier of the NotificationSubscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchItem** | [**[]PatchItem**](PatchItem.md) | data to patch | 
 **ifMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

